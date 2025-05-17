{{define "sign-up"}}
{{template "header" .}}

<div class="row justify-content-center">
    <div class="col-4">
        <form method="post" action="/auth/register">
            <div class="mb-3">
                <label for="email" class="form-label">Email{{template "required"}}</label>
                <input type="email" class="form-control" id="email" name="email" value="test@test.com">
            </div>
            <div class="mb-3">
                <label for="password" class="form-label">Password {{template "required"}}</label>
                <input type="password" class="form-control" id="password" name="password" value="123">
            </div>
            <div class="mb-3">
                <label for="password_confirm" class="form-label">Password Confirm {{template "required"}}</label>
                <!-- for 和 id对应，当点击标签是，焦点自动转移到输入框上 -->
                <input type="password" class="form-control" id="password_confirm" name="password_confirm" value="123">
                <!-- name作为form提交时的键名 -->
            </div>
            <div class="mb-3">
                Already have account, click here to <a href="/sign-in">Sign in</a>
            </div>
            <button type="submit" class="btn btn-primary">Sign In</button>
        </form>
    </div>
</div>
{{template "footer" .}}
{{end}}