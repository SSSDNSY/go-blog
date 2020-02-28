{{define "navbar"}}
    <nav class="navbar  navbar-default ">
        <div class="container">
            <div class="navbar-header">
                <button class="navbar-toggle collapsed" data-toggle="collapse" data-target="#navbar-collapse-1" type="button">
                    <span class="icon-bar navbar-left"></span>
                    <span class="icon-bar navbar-left"></span>
                    <span class="icon-bar navbar-left"></span>
                </button>
                <a class="navbar-brand" href="/">我的博客</a>
            </div>
            <ul class="nav navbar-nav collapse navbar-collapse" id="navbar-collapse-1">
                <li {{if .IsHome}}class="active" {{end}}><a href="/">首页</a></li>
                <li {{if .IsCategory}}class="active" {{end}} class=""><a href="/category"> 分类</a></li>
                <li {{if .IsTopic}}class="active" {{end}} class=""><a href="/topic"> 文章</a></li>
                <li {{if .IsFile}}class="active" {{end}} class=""><a href="/file"> 文件</a></li>
            </ul>
            <ul class="nav navbar-nav navbar-right collapse navbar-collapse">
                {{if .IsLogin}}
                    <li><a class="" href="/login?exit=true">退出</a></li>
                {{else}}
                    <li><a class="" href="/login">管理员登录</a></li>
                {{end}}
            </ul>
        </div>
    </nav>
{{end}}