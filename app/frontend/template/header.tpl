{{define "header"}}
<!doctype html>
<html lang="en">

<head>
    <meta charset="utf-8">
    <meta name="viewport" content="width=device-width, initial-scale=1">
    <!-- width=device-width：将视口的宽度设置为设备的屏幕宽度。
    initial-scale=1：设置页面的初始缩放比例为 1，即页面在加载时不进行缩放。 -->
    {{ if $.redirect }}
    <meta http-equiv="refresh" content="5; url={{$.redirect}}">
    <!-- http-equiv="refresh"：指示浏览器在指定时间后刷新页面或跳转到新的 URL。
    content="5; url=/checkout/result"：设置刷新或跳转的时间间隔和目标地址。 -->
    {{ end }}
    <title>{{$.title}}</title>
    <!-- 访问顶级作用域对象 -->
    <link href="/static/css/bootstrap.min.css" rel="stylesheet" crossorigin="anonymous">
    <!-- 如果crossorigin="use-credentials"，浏览器在请求 styles.css 时会附带用户的凭据，如 Cookies 和认证头部 -->
    <!--html头中 Access-Control-Allow-Credentials: true：明确表示服务器允许浏览器发送凭据 -->
    <link rel="stylesheet" href="https://cdnjs.cloudflare.com/ajax/libs/font-awesome/6.7.1/css/all.min.css">
    <!--  <link> 标签是 HTML 中用于在文档与外部资源之间建立关系的元素 -->
    <!-- rel 属性用于指定当前文档与被链接资源之间的关系，其值是一个或多个由空格分隔的关键字。不同的 rel 值表示不同的关系 -->
</head>

<body class="min-vh-100">
    <header>
        <nav class="navbar navbar-expand-lg bg-body-tertiary">
            <div class="container-fluid">
                <!-- container-fluid占满窗口 -->
                <img class="navbar-brand" src="/static/img/logo.png" href="#" alt="logo" style="height: 3em;" />
                <button class="navbar-toggler" type="button" data-bs-toggle="collapse"
                    data-bs-target="#navbarSupportedContent" aria-controls="navbarSupportedContent"
                    aria-expanded="false" aria-label="Toggle navigation">
                    <span class="navbar-toggler-icon"></span>
                </button>
                <div class="collapse navbar-collapse" id="navbarSupportedContent">
                    <ul class="navbar-nav me-auto mb-2 mb-lg-0">
                        <li class="nav-item">
                            <a class="nav-link active" aria-current="page" href="/">Home</a>
                        </li>

                        <li class="nav-item dropdown">
                            <a class="nav-link dropdown-toggle" href="#" role="button" data-bs-toggle="dropdown"
                                aria-expanded="false">
                                Categories
                            </a>
                            <ul class="dropdown-menu">
                                <li><a class="dropdown-item" href="/category/meat">Meat</a></li>
                                <li><a class="dropdown-item" href="/category/Vegetarian">Vegetarian</a></li>
                            </ul>
                        </li>
                        <li class="nav-item">
                            <a class="nav-link" aria-disabled="true" href="/about">About</a>
                        </li>
                    </ul>
                    <form class="d-flex ms-auto" role="search" action="/search" method="get">
                        <input class="form-control me-2" type="search" name="q" placeholder="Search" aria-label="Search"
                            value="{{ .q }}">
                        <button class="btn btn-outline-success" type="submit">Search</button>
                    </form>

                    {{ template "cart-num" . }}

                    {{if .user_id}}
                    <div class="dropdown">
                        <div class="ms-3 dropdown-toggle" data-bs-toggle="dropdown">
                            <i class="fa-solid fa-user fa-xl"></i>
                            <span>hello</span>
                            <!--  p标签：用于定义段落，是组织文本内容的基本单位。span标签：用于对文本中的一小部分进行样式或脚本操作，适用于不需要换行的内容。 -->
                        </div>
                        <!-- data-bs-toggle 是 Bootstrap 5 中用于启用特定组件行为的 HTML 数据属性。通过将此属性添加到元素上，可以无需编写 JavaScript 代码，直接实现如折叠（collapse）、模态框（modal）、标签页（tab）等交互功能 -->
                        <ul class="dropdown-menu dropdown-menu-end mt-3">
                            <li><a class="dropdown-item" href="/order">Order Center</a></li>
                            <li><a class="dropdown-item" href="#">Order Center</a></li>
                            <li>
                                <form action="/auth/logout" method="post">
                                    <button class="dropdown-item" type="submit">Logout</button>
                                </form>
                            </li>
                        </ul>
                    </div>

                    {{else}}
                    <div class="ms-3">
                        <a href="/sign-in" class="btn btn-primary">Sign in</a>
                    </div>
                    {{end}}
                </div>
            </div>
        </nav>
    </header>

    <main style="min-height: calc(83vh);">
        <div class="container-fluid py-3">
            {{end}}