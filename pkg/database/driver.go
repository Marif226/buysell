package database

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"sync"
)

const Version = "1.0"

// Driver is a database driver struct
type Driver struct {
	mutex   sync.Mutex
	mutexes map[string]*sync.Mutex
	Dir     string
}

func NewDriver(dir string) (*Driver, error) {
	// clean directory path
	dir = filepath.Clean(dir)

	// create a new driver
	driver := Driver{
		Dir:     dir,
		mutexes: make(map[string]*sync.Mutex),
	}

	// check if database directory exists
	if _, err := stat(dir); err == nil {
		log.Printf("using '%s' (database already exists)\n", dir)
		return &driver, nil
	}

	log.Printf("creating a new database at '%s' ...\n", dir)
	// make a database directory dir
	return &driver, os.MkdirAll(dir, 0755)
}

func (d *Driver) getOrCreateMutex(collection string) *sync.Mutex {
	m, ok := d.mutexes[collection]
	d.mutex.Lock()
	defer d.mutex.Unlock()

	if !ok {
		m = &sync.Mutex{}
		d.mutexes[collection] = m
	}

	return m
}

// Write writes a record v to a collection with a name resource
func (d *Driver) Write(collection, resource string, v interface{}) error {
	if collection == "" {
		return fmt.Errorf("missing collection - no place to save record")
	}

	if resource == "" {
		return fmt.Errorf("missing resource - unable to save record (no name)")
	}

	mutex := d.getOrCreateMutex(collection)
	mutex.Lock()
	defer mutex.Unlock()

	// dir/collection/resource.json
	dir := filepath.Join(d.Dir, collection)
	finalPath := filepath.Join(dir, resource+".json")
	temporaryPath := finalPath + ".tmp"

	// make a directory for collection inside the database directory
	if err := os.MkdirAll(dir, 0755); err != nil {
		return err
	}

	// encode struct v to json
	b, err := json.MarshalIndent(v, "", "\t")
	if err != nil {
		return err
	}
	b = append(b, byte('\n'))

	if err := os.WriteFile(temporaryPath, b, 0644); err != nil {
		return err
	}

	return os.Rename(temporaryPath, finalPath)
}

// Read reads a record in a collection with a name matching to given resource
func (d *Driver) Read(collection, resource string, v interface{}) error {
	if collection == "" {
		return fmt.Errorf("missing collection - unable to read")
	}

	if resource == "" {
		return fmt.Errorf("missing resource - unable to read record (no name)")
	}

	// path to a record dir/collection/resource
	record := filepath.Join(d.Dir, collection, resource)

	// check if path exists
	if _, err := stat(record); err != nil {
		return err
	}

	b, err := os.ReadFile(record + ".json")
	if err != nil {
		return err
	}

	return json.Unmarshal(b, &v)
}

// ReadAll reads all records in a collection and returns them as a slice of strings
func (d *Driver) ReadAll(collection string) ([]string, error) {
	if collection == "" {
		return nil, fmt.Errorf("missing collection - unable to read")
	}

	// collection directory with the records
	dir := filepath.Join(d.Dir, collection)

	if _, err := stat(dir); err != nil {
		return nil, err
	}

	files, _ := os.ReadDir(dir)

	var records []string

	for _, file := range files {
		// read file by its name
		b, err := os.ReadFile(filepath.Join(dir, file.Name()))
		if err != nil {
			return nil, err
		}

		// append the file read to records
		records = append(records, string(b))
	}

	return records, nil
}

// Delete deletes given collection or file in the collection
func (d *Driver) Delete(collection, resource string) error {
	if collection == "" {
		return fmt.Errorf("missing collection - unable to delete")
	}

	// path to the document to be deleted
	path := filepath.Join(collection, resource)

	mutex := d.getOrCreateMutex(collection)
	mutex.Lock()
	defer mutex.Unlock()

	dir := filepath.Join(d.Dir, path)

	fi, err := stat(dir)
	switch {
	case fi == nil, err != nil:
		return fmt.Errorf("unable to find file or directory named '%v'", path)
	case fi.Mode().IsDir():
		// when resource is empty, delete the whole collection
		return os.RemoveAll(dir)
	case fi.Mode().IsRegular():
		// delete concrete resource document
		return os.RemoveAll(dir + ".json")
	}

	return nil
}

// stat checks if given path exists and returns FileInfo
func stat(path string) (os.FileInfo, error) {
	fi, err := os.Stat(path)
	if os.IsNotExist(err) {
		fi, err = os.Stat(path + ".json")
	}

	return fi, err
}
