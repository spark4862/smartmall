{{define "home"}}

{{template "header" .}}
<!-- .是对模板传入参数 -->

<div class="row">
    {{range .items}}
    <div class="card col-xl-3 col-lg-4 col-md-6 col-sm-12 p-4 border-0">
        <a href="/product?id={{ .Id }}" class="btn">
            <img src="{{.Picture}}" class="card-img-top" alt="...">
            <!-- 注意外面的引号 -->
            <div class="card-body">
                <p class="card-text">{{.Name}}</p>
                <h5 class="card-title">{{.Price}}</h5>
            </div>
        </a>
    </div>
    {{end}}
</div>

{{template "footer" .}}

{{end}}