# 目录结构

> conf 配置
>
> configure 读取配置文件
>
> controller 控制器
>
> db 数据库
>
> > index 数据库连接配置
> >
> > nest.sql 数据文件
>
> docs swagger文档
>
> log zap日志
>
> middleware 中间件
>
> model 类型定义
>
> router 路由定义
>
> > v1 v1版本路由
> >
> > router 路由入口
> >
> > whitelistRouter 不需要登录鉴权的路由
>
> service 服务
>
> tmp air运行生成目录
>
> utils 工具类
>
> > db 数据库相关的
> >
> > R   请求统一封装的返回
>
> .gitignore git忽略文件或者文件夹
>
> air.toml air配置
>
> main 主入口

## 完成

- 用户管理	<> 关联角色
- 角色管理    <> 关联菜单
- 菜单管理
- 部门管理
- 字典管理
- 登录、权限校验
- swagger 文档生成
- 日志zap集成
- 热重载 air

## 运行

- 自己创建数据库，从db找到nest.sql 导入到数据库中，修改conf下面的DB配置为自己的
- air 运行 `air -c air.toml -d or air` 启动端口为8081
- swagger访问路径 http://localhost:8081/swagger/index.html
- 交叉编译工具 gox
    - `打包 gox -output="bin" -osarch="linux/amd64"`