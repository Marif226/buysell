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
<h4>Product's details</h4>
<#list images as img>
    <img src="/images/${img.id}" height="60px"><br><br>
</#list>
<b>Product's name</b>${product.title}<br>
<b>Product's description</b>${product.description}<br>
<b>Price</b>${product.price}<br>
<b>City</b>${product.city}<br>
<b>Author</b>${product.author}<br>
<hr>
<form action="/product/delete/${product.id}" method="post">
    <input type="submit" value="Delete product">
</form>
</body>
</html>