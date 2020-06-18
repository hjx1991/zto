<!DOCTYPE html>
<html lang="zh-CN">
<head>
	<title>Demo</title>
	<meta charset="utf-8">
</head>
<body>
     <form action="/register" method="post">
     {{.xsrfdata}}
       <input type="text" name="username"/>
       <input type="text" name="password"/>
       <input  type="submit" value="Post"/>
     </form>

</body>
</html>