{{define "navbar"}}
    <nav class="navbar navbar-expand-lg navbar-light bg-light">
        <a class="navbar-brand" href="/">我的博客</a>
        <button class="navbar-toggler" data-toggle="collapse" data-target="#navbar-collapse-1"
                type="button">
            <span class="icon-bar navbar-left"></span>
            <span class="icon-bar navbar-left"></span>
            <span class="icon-bar navbar-left"></span>
        </button>

        <div class="collapse navbar-collapse" id="navbar-collapse-1">
            <ul class="navbar-nav mr-auto">
                <li class="nav-item {{if .IsHome}}active{{end}}"><a href="/" class="nav-link ">首页</a></li>
                <li class="nav-item {{if .IsCategory}}active{{end}}"><a href="/category" class="nav-link "> 分类</a>
                </li>
                <li class="nav-item {{if .IsTopic}}active{{end}}"><a href="/topic" class="nav-link"> 文章</a></li>
                <li class="nav-item {{if .IsFile}}active{{end}}"><a href="/file" class="nav-link"> 文件</a></li>
                {{if .IsLogin}}
                    <li class="nav-item">
                        <a class="nav-link" href="/login?exit=true">退出</a>
                    </li>
                {{else}}
                    <li class="nav-item">
                        <a class="nav-link" href="/login">管理员登录</a>
                    </li>
                {{end}}
            </ul>
            <form class="form-inline my-2 my-lg-0">
                <input class="form-control " type="search" placeholder="搜索内容" aria-label="搜索内容">
            </form>
        </div>
    </nav>

{{end}}