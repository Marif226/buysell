<!doctype html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport"
          content="width=device-width, user-scalable=no, initial-scale=1.0, maximum-scale=1.0, minimum-scale=1.0">
    <meta http-equiv="X-UA-Compatible" content="ie=edge">
    <title>BUYSELL</title>
</head>
<body>
    <h1>Hello BUYSELL!</h1>
    <h4>Kazakhstan</h4>
    <form action="/" method="get">
        Search by ad name: <input type="text" name="title"><br>
        <input type="submit" value="Search">
    </form>
    <#list products as product>
        <div>
            <p><b>${product.title}</b> ${product.price} | <a href="/product/${product.id}">Details</a></p>
        </div>
        <#else>
        <h3>No products</h3>
    </#list>
    <hr>
    <h3>Create new product</h3>
    <form action="/product/create" method="post" enctype="multipart/form-data">
        Ad's name: <input type="text" name="title"></br>
        Description: <input type="text" name="description"></br>
        Price: <input type="number" name="price"></br>
        City: <input type="text" name="city"></br>
        Author: <input type="text" name="author"></br>
        1st photo: <input type="file" name="file1"></br>
        2nd photo: <input type="file" name="file2"></br>
        3rd photo: <input type="file" name="file3"></br>
        <input type="submit" value="Submit">
    </form>
</body>
</html>