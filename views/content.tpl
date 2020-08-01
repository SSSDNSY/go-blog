{{define "content"}}
    <!-- Main Content -->
    <div class="container">
        <div class="row">
            <div class="col-lg-8 col-md-10 mx-auto">
                {{range .Topics}}
                    <div class="post-preview">
                        <a href="/topic/blog/{{.Id}}">
                            <h2 class="post-title">
                                {{.Title}}
                            </h2>
                            <div class="post-subtitle">
                                {{ substr .Content 0 50}}
                            </div>
                        </a>
                        <div class="post-meta">
                            <span class="fas fa-user">&nbsp;</span> pengzh&nbsp;&nbsp;&nbsp;
                            <span class="fas fa-calendar-alt"></span>&nbsp;{{ date .Updated "Y-m-d H:i:s"}}
                        </div>
                    </div>
                    <hr>
                {{end}}
                <!-- Pager -->
                <div class="" style="display:{{.leftDis}}">
                    <a class="btn btn-primary float-left" href="?lPageNumber={{.lPageNumber }}">&larr;Pre </a>
                </div>
                <div class="clearfix" style="display: {{.rightDis}}">
                    <a class="btn btn-primary float-right" href="?rPageNumber={{.rPageNumber}}">Next &rarr;</a>
                </div>
            </div>
        </div>
    </div>

    <hr>

{{end}}