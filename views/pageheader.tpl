{{define "pageheader"}}

    <!-- Page Header -->
    <header class="masthead" style="background-image: url('{{.image}}')">
        <div class="overlay"></div>
        <div class="container">
            <div class="row">
                <div class="col-lg-8 col-md-10 mx-auto">
                    <div class="site-heading">
                        <h1>{{.title}}</h1>
                        <span id="pageHeader" class="subheading">{{.subtitle}}</span>
                    </div>
                </div>
            </div>
        </div>
    </header>
{{end}}