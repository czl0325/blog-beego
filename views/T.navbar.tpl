{{ define "navbar" }}
    <div class="navbar navbar-default navbar-static-top">
        <div class="container">
            <a class="navbar-brand" href="/">我的博客</a>
            <ul class="nav navbar-nav">
                <li {{if compare .HomeIndex 0}}class="active"{{end}}><a href="/">首页</a></li>
                <li {{if compare .HomeIndex 1}}class="active"{{end}}><a href="/category">分类</a></li>
                <li {{if compare .HomeIndex 2}}class="active"{{end}}><a href="/topic">文章</a></li>
            </ul>

            <div class="pull-right">
                <ul class="nav navbar-nav">
                    <li><a id="uname"></a></li>
                </ul>
            </div>
        </div>
    </div>

    <script type="text/javascript">
        $(document).ready(function () {
            let name = getCookie("name")
            if (name) {
                $("#uname").html(name + "   退出")
                $("#uname").attr("href", "/login?exit=true")
            } else {
                $("#uname").html("管理员登录")
                $("#uname").attr("href", "/login")
            }
        })
    </script>
{{ end }}