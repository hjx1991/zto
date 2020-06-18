<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>register</title>
</head>
<body>
    <form method="post" action="/register">
    {{.xsrfdata}}
        用户名:<input name="username" type="text"/>
        密 码:<input name="password" type="text"/>
        <input type="submit" value="提交">
    </form>
</body>
</html>