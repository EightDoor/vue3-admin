<template>
  <div class="login">
    <div class="container">
      <h1>后台登录</h1>
      <a-form
        ref="formRef"
        :model="formState"
        :labe-col="labelCol"
        :wrapper-col="wrapperCol"
        :rules="rules"
        class="form"
      >
        <a-form-item label="账户" name="name">
          <a-input v-model:value="formState.name" />
        </a-form-item>
        <a-form-item label="密码" name="password">
          <a-input-password
            v-model:value="formState.password"
            @pressEnter="onSubmit"
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
import {
  defineComponent, reactive, ref, toRaw,
} from 'vue';
import { message } from 'ant-design-vue';
import { useRouter } from 'vue-router';
import { useStore } from 'vuex';
import { LOGIN, LOGINRESET } from '@/store/mutation-types';
import localStore from '@/utils/store';
import { STORELETMENUPATH } from '@/utils/constant';

interface LoginType {
  name: string;
  password: string;
}
const Login = defineComponent({
  name: 'ViewsLogin',
  setup() {
    const router = useRouter();
    const submitData = reactive({ loading: false });
    const wrapperCol = { span: 14 };
    const labelCol = { span: 4 };
    const store = useStore();

    const formRef = ref();
    const formState = reactive<LoginType>({
      name: '',
      password: '',
    });
    const rules = {
      name: [
        {
          required: true,
          message: '请输入姓名',
          trigger: 'blur',
        },
      ],
      password: [
        {
          required: true,
          message: '请输入密码',
          trigger: 'blur',
        },
      ],
    };

    const onSubmit = (e: { preventDefault: () => void }) => {
      e.preventDefault();
      formRef.value
        .validate()
        .then(() => {
          const data = toRaw(formState);
          submitData.loading = true;
          store
            .dispatch(LOGIN, {
              account: data.name,
              pass_word: data.password,
            })
            .then(() => {
              submitData.loading = false;
              message.success('登录成功!');
              store.commit(`${LOGINRESET}`);
              router.push('/');
            })
            .catch(() => {
              submitData.loading = false;
            });
        })
        .catch((err) => {
          console.log('error', err);
        });
    };
    return {
      formState,
      wrapperCol,
      labelCol,
      rules,
      submitData,
      onSubmit,
      formRef,
    };
  },
});
export default Login;
</script>
<style lang="less" scoped>
@import "login.less";
</style>
