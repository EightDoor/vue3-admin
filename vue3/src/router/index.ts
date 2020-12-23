import { createRouter, createWebHashHistory, RouteRecordRaw } from 'vue-router'
import { canUserAccess } from '@/authority'

import Layout from '@/layout/layout/layout.vue'
import Login from '@/views/login/login.vue'
import LayoutChildren from '@/views/layout-children.vue';

import Home from '@/views/home/home.vue'
import Test from '@/views/Test.vue'
import Demo from '@/views/Demo.vue'

import Depart from '@/views/sys/depart.vue'
import SysUser from "@/views/sys/user.vue"
import SysRole from '@/views/sys/role.vue';
import SysMenu from '@/views/sys/menu.vue';

const routes: RouteRecordRaw[] = [
  {
    path: '/',
    name: 'layout',
    component: Layout,
    redirect: '/home',
    children: [
      {
        name: 'home',
        path: 'home',
        component: Home,
        meta: {
          title: '首页',
          icon: '23',
        },
      },
      {
        name: "sys",
        path: "sys",
        component: LayoutChildren,
        meta: {
          title: "系统管理",
          icon: "1"
        },
        children: [
          {
            name: 'user',
            path: 'user',
            component: SysUser,
            meta: {
              title: '用户管理',
              icon: '1',
            },
          },
          {
            name: 'menu',
            path: 'menu',
            component: SysMenu,
            meta: {
              title: '菜单管理',
              icon: '1',
            },
          },
          {
            name: 'role',
            path: 'role',
            component: SysRole,
            meta: {
              title: '角色管理',
              icon: '1',
            },
          },
          {
            name: 'depart',
            path: 'depart',
            component: Depart,
            meta: {
              title: '部门管理',
              icon: '1',
            },
          },
        ]
      },
      {
        name: 'test',
        path: 'test',
        component: Test,
        meta: {
          title: '测试',
          icon: '23',
        },
        children: [
          {
            name: 'demo',
            path: 'demo',
            component: Demo,
            meta: {
              title: 'demo',
              icon: '23',
            },
          },
        ],
      },
    ],
  },
  {
    path: '/login',
    name: 'login',
    component: Login,
  },
]
const app = createRouter({
  history: createWebHashHistory(),
  routes,
})

// 全局路由前置钩子
app.beforeEach(async () => {
  await canUserAccess()
})

export default app
