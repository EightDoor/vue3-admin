import {createRouter, createWebHashHistory, RouteRecordRaw} from 'vue-router';
import {canUserAccess} from "../authority";

import Home from '@/pages/home/home.vue';
import Login from '@/pages/login/login.vue';

const routes:RouteRecordRaw[] = [
    {
        path: '/',
        name: 'home',
        component: Home,
    },
    {
        path: '/login',
        name: 'login',
        component: Login
    }
]
const app = createRouter({
    history: createWebHashHistory(),
    routes
})

// 全局路由前置钩子
app.beforeEach(async(to, form)=>{
    await canUserAccess();
})

export default app;
