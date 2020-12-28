<template>
  <div class="login">
    <div class="container">
      <h1>后台登录</h1>
      <a-form :labe-col="labelCol" :wrapper-col="wrapperCol" class="form">
        <a-form-item label="账户" v-bind="validateInfos.name">
          <a-input v-model:value="modelRef.name" />
        </a-form-item>
        <a-form-item label="密码" v-bind="validateInfos.password">
          <a-input type="password" v-model:value="modelRef.password" />
        </a-form-item>
        <a-form-item :wrapper-col="{ span: 18, offset: 6 }">
          <a-button
            type="primary"
            :loading="submitData.loading"
            @click="onSubmit"
            >登录</a-button
          >
        </a-form-item>
      </a-form>
    </div>
  </div>
</template>
<script lang="ts">
import { defineComponent, reactive, toRaw } from 'vue'
import { useForm } from '@ant-design-vue/use'
import { message } from 'ant-design-vue'
import { RouteRecordRaw, useRouter } from 'vue-router'
import { useStore } from 'vuex'
import { LOGIN } from '@/store/mutation-types'
import Layout from '@/layout/layout/layout.vue'
import Home from '@/views/home/home.vue'
import LayoutChildren from '@/views/layout-children.vue'
import SysUser from '@/views/sys/user.vue'
import SysMenu from '@/views/sys/menu.vue'
import SysRole from '@/views/sys/role.vue'
import Depart from '@/views/sys/depart.vue'

interface LoginType {
  name: string
  password: string
}
const Login = defineComponent({
  name: 'login',
  setup() {
    const router = useRouter()
    const submitData = reactive({ loading: false })
    const wrapperCol = { span: 14 }
    const labelCol = { span: 4 }
    const store = useStore()
    const modelRef = reactive<LoginType>({
      name: '',
      password: '',
    })
    const rulesRef = reactive({
      name: [
        {
          required: true,
          message: '请输入姓名',
        },
      ],
      password: [
        {
          required: true,
          message: '请输入密码',
        },
      ],
    })

    const { validateInfos, validate } = useForm(modelRef, rulesRef)

    const onSubmit = (e: { preventDefault: () => void }) => {
      e.preventDefault()
      validate()
        .then(() => {
          const data = toRaw(modelRef)
          submitData.loading = true
          store
            .dispatch(LOGIN, {
              account: data.name,
              pass_word: data.password,
            })
            .then(() => {
              submitData.loading = false
              message.success('登录成功!')
              const dynamicRoutes: RouteRecordRaw[] = [
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
                      name: 'sys',
                      path: 'sys',
                      component: LayoutChildren,
                      meta: {
                        title: '系统管理',
                        icon: '1',
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
                      ],
                    },
                  ],
                },
              ]
              dynamicRoutes.forEach((item) => {
                router.addRoute(item)
              })
              router.push('/home')
            })
        })
        .catch((err) => {
          console.log('error', err)
        })
    }
    return {
      modelRef,
      wrapperCol,
      labelCol,
      validateInfos,
      submitData,
      onSubmit,
    }
  },
})
export default Login
</script>
<style lang="less" scoped>
@import 'login.less';
</style>
