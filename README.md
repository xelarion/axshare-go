Dockerize axshare-vue

    docker build -t ervincheung/axshare_go .
    docker run -d -p 10524:10524 --rm --name dockerize-axshare_go ervincheung/axshare_go
***

**api/**
（Service 应用相关目录）一般用来放着 OpenAPI/Swagger 的 spec、JSON 的 schema 文件或者 protocol 的定义。

**build/**
打包（packaging）和 CI 相关文件。比如 Docker，OS（deb，rpm，pkg）相关的配置和脚本文件可放在 build/package 目录下，而 CI （travis，drone 等）相关文件可放置 build/ci 目录下。

**cmd/** 
入口文件

**configs/**
配置文件或者模版文件。

**deployments/**
IaaS，PaaS 或者容器编排系统的配置和模版文件。

**docs/**
设计或者用户文档。

**examples/**
项目（应用或者库）相关的示例代码。

**githooks/**
放置 Git hooks。

**init/**
系统初始化（如 systemd，upstart，sysv）和进程管理（如 runit，supervisord）相关工具的配置。

**internal/**
私有的 application 或者库代码（不希望 package 的接口被扩散到同层目录以外的空间中）

**pkg/**
用来放置库代码，可被项目内部或外部引用。

**scripts/**
构建，安装，分析等相关操作的脚本。

**test/**
额外的测试应用和测试数据，如 test/data。

**third_party/**
外部的第三方工具、代码或其他组件。

**tools/**
项目相关的一些 tool，其代码可引用 pkg/ 和 internal/ 目录下的 package。



