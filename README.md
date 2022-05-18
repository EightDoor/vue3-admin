# 技术栈
- typescript + vue3 + composition-api + ant-design-vue2.x + vue-router@4.x + vuex@4.x

# 完成功能

- 用户管理 <> 关联角色
- 角色管理 <> 关联菜单
- 菜单管理
- 部门管理
- 字典管理
- 登录、权限校验
- 动态路由生成 - 按钮权限
- 菜单栏切换面包屑
- tab展示历史导航路由
# 运行
- 安装依赖 `pnpm install`
- 运行 `npm run dev`

## 待办
- [ ] 首页样式调整美化
- [ ]使用fast-crud开发其他的功能列表
  - [x]  完成基础demo展示
  - [ ]  增加复杂select自定义表格数据
- [ ] 个人用户修改密码
- [ ] 用户登陆日志记录
- [ ] 接口请求记录
- [x] 使用 @nestjsx/crud-request  定制查询条件
# 问题
- [x] 刷新 顶部面包屑显示不对问题
- [x] 刷新页面  默认选中当前选择的路由
- [ ] 刷新页面直接跳转404 问题
- [ ] 表单无法重置，重置失效，以后更新版本看看是否可行

# 其他
- 配套后台 nodejs
  - [vue3-admin-nest](https://github.com/EightDoor/vue3-admin-nest)
- 注意事项
  - 创建组件的时候，如果为动态路由，需要添加 isRouter: true (后续更改方式)
  - 动态路由方案参考 https://juejin.cn/post/6951557699079569422

## 按钮权限使用
- ```vue
  // 使用的是按钮权限的名字
  <a-button type="danger" v-bt-auth:del></a-button>
  // 或者
  // 使用的是自定义的名称  '删除 - 1'
  <a-button type="danger" v-bt-auth:del="{title:true}">删除 - 1</a-button>
  ```
