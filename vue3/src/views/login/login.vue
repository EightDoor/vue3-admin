<template>
  <div class="login">
    <div class="container">
      <h1>后台登录</h1>
      <a-form :labe-col="labelCol" :wrapper-col="wrapperCol" class="form">
        <a-form-item label="账户" v-bind="validateInfos.name">
          <a-input v-model:value="modelRef.name" />
        </a-form-item>
        <a-form-item label="密码" v-bind="validateInfos.password">
          <a-input-password
            @pressEnter="onSubmit"
            v-model:value="modelRef.password"
          />
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
import { useRouter } from 'vue-router'
import { useStore } from 'vuex'
import { LOGIN, LOGINRESET } from '@/store/mutation-types'
import { localForage } from '@/utils/localforage'
import { STORELETMENUPATH } from '@/utils/constant'

interface LoginType {
  name: string
  password: string
}
const Login = defineComponent({
  name: 'Login',
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
              localForage.setItem(STORELETMENUPATH, {})
              store.commit(`${LOGINRESET}`)
              router.push('/home')
            })
            .catch(() => {
              submitData.loading = false
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
