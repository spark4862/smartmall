{{define "sign-in"}}
<!-- 该定义可以用于取消代码对tpl引用的后缀 -->
{{template "header" .}}

<div class="row justify-content-center">
    <div class="col-4">
        <form method="post" action="/auth/login?next={{ .Referer }}">
            <div class="mb-3">
                <label for="email" class="form-label">Email{{template "required"}}</label>
                <input type="email" class="form-control" id="email" name="email" value="test@test.com">
            </div>
            <div class="mb-3">
                <label for="password" class="form-label">Password {{template "required"}}</label>
                <input type="password" class="form-control" id="password" name="password" value="123">
            </div>
            <div class="mb-3">
                Don't have account, click here to <a href="/sign-up">Sign up</a>
            </div>
            <button type="submit" class="btn btn-primary">Sign In</button>
        </form>
    </div>
</div>
{{template "footer" .}}
{{end}}