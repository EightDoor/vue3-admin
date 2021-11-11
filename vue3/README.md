# 文件结构
> public 静态资源不打包

> src
>
> > assets 静态资源(打包)
> >
> > components 公共组件
> >
> > layout 布局
> >
> > router 路由
> >
> > store vuex状态管理
> >
> > types typescript类型文件定义
> >
> > utils 工具类
> >
> > views 页面
>
> App router-view渲染
>
> authority 权限校验
>
> main 主入口
>
> shims-vue.d 全局vue类型定义

> .env 公共配置
>
> .env.development 开发配置
>
> .env.production 生产配置
>
> .gitignore git提交忽略文件或者文件夹
>
> .prettierrc prettierrc格式化配置
>
> vue.config vue-cli自定义配置

# 技术栈
- typescript + vue3 + composition-api + ant-design-vue2.x + vue-router@4.x + vuex@4.x

# 完成功能

- 用户管理  <> 关联角色
- 角色管理  <> 关联菜单
- 菜单管理  
- 部门管理
- 登录、权限校验
- 动态路由生成 - 按钮权限
- 菜单栏切换面包屑
- tab展示历史导航路由
- 字典管理

# 待办

- 首页样式调整美化

  # 运行

- 安装依赖 `npm install or yarn install`

- 运行 `npm run dev`

# 资料
- [eslint-vue-plugin](https://eslint.vuejs.org/rules/valid-v-model.html)
# 使用

## 按钮权限使用

- ```vue
  // 使用的是按钮权限的名字
  <a-button type="danger" v-bt-auth:del></a-button>
  // 或者
  // 使用的是自定义的名称  '删除 - 1'
  <a-button type="danger" v-bt-auth:del="{title:true}">删除 - 1</a-button>
  ```

