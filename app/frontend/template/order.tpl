{{ define "order" }}
{{ template "header" . }}
<div class="row">
    <div class="card border-0" style="width: 100%;">
        <div class="card-body row">
            {{ range $.orders }}
            <div class="card">
                <div class="card-body">
                    <h6 class="card-subtitle mb-2 text-muted">{{.CreatedDate}} Order ID: {{.OrderId}}</h6>
                    <ul class="list-group col-lg-12 col-sm-15">
                        {{ range .Items }}
                        <li class="list-group-item border-0">
                            <div class="card border-0">
                                <div class="card-body row">
                                    <div class="col-3">
                                        <img src="{{ .Picture }}" style="max-width: 100px;max-height: 50px" alt="">
                                    </div>
                                    <div class="col-3">
                                        <div class="mt-1">{{ .ProductName }}</div>
                                    </div>
                                    <div class="col-2">
                                        <div class="mt-1">x {{ .Qty }}</div>
                                    </div>
                                    <div class="col-4">
                                        <div class="mt-1">Cost: {{ .Cost }}</div>
                                    </div>
                                </div>
                            </div>
                        </li>
                        {{ end}}
                    </ul>
                </div>
            </div>
            <p>
                {{ end}}
        </div>
    </div>
</div>
</div>
{{ template "footer" . }}
{{ end }}