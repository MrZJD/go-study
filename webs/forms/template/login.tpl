<html>
    <head>
        <title></title>
    </head>
    <body>
        <form action="/login" method="post">
            用户名：<input type="text" name="username">
            密码：<input type="password" name="password">
            <input type="submit" value="登录">
            <input type="hidden" name="token" value="{{.}}">
        </form>
        哈哈哈
    </body>
</html>