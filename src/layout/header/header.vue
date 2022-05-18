<template>
  <a-layout-header class="container">
    <div class="content">
      <div class="button-icon-container">
        <div @click="ToggleCollapsed()">
          <MenuUnfoldOutlined v-if="collapsed" class="button-icon" />
          <MenuFoldOutlined v-else class="button-icon" />
        </div>
        <a-breadcrumb>
          <a-breadcrumb-item
            v-for="(item, index) in crumbs"
            :key="index"
            class="button-breadcrumb"
          >
            {{ item }} {{ index > 0 && index + 1 !== crumbs.length ? '/' : '' }}
          </a-breadcrumb-item>
        </a-breadcrumb>
      </div>
      <a-dropdown>
        <a class="ant-dropdown-link" @click="(e) => e.preventDefault()">
          <a-avatar :size="50">
            <template #icon><UserOutlined /></template>
          </a-avatar>
        </a>
        <template #overlay>
          <a-menu>
            <a-menu-item v-for="(item, index) in data" :key="index">
              <a @click="GoTo(item)">{{ item }}</a>
            </a-menu-item>
          </a-menu>
        </template>
      </a-dropdown>
    </div>
  </a-layout-header>
  <CommonDrawer :visible="visible" title="个人中心" @on-close="visible = false">
    <a-form>
      <a-form-item label="修改密码" v-bind="validateInfos.password">
        <a-input v-model:value="modelRef.password" />
      </a-form-item>
      <a-form-item>
        <a-button type="primary" @click.prevent="onSubmit()">更改</a-button>
      </a-form-item>
    </a-form>
  </CommonDrawer>
</template>
<script lang="ts">
import {
  UserOutlined,
  MenuUnfoldOutlined,
  MenuFoldOutlined,
} from '@ant-design/icons-vue';
import { defineComponent, ref, computed, reactive, toRaw } from 'vue';
import { useRouter } from 'vue-router';
import { useStore } from 'vuex';
import { Form } from 'ant-design-vue';
import { CLEAR_CRUMBS, COLLAPSED } from '@/store/mutation-types';
import log from '@/utils/log';
import CommonDrawer from '@/components/Drawer/Drawer.vue';
import utilsLocal from '@/utils/store';

export default defineComponent({
  name: 'CommonHeader',
  components: {
    UserOutlined,
    MenuUnfoldOutlined,
    MenuFoldOutlined,
    CommonDrawer,
  },
  setup() {
    const data = ref<string[]>(['个人中心', '退出']);
    const router = useRouter();
    const store = useStore();

    // 个人中心
    const visible = ref(false);
    const { useForm } = Form;
    const modelRef = reactive({
      password: '',
      icon: '',
    });
    const ruleRef = reactive({
      password: [
        {
          required: true,
          message: '请输入密码',
        },
      ],
    });

    const { resetFields, validate, validateInfos } = useForm(modelRef, ruleRef);

    function onSubmit() {
      validate()
        .then(() => {
          log.d(toRaw(modelRef), '表单值');
        })
        .catch((err) => {
          log.e(err, '表单验证错误');
        });
    }

    function GoTo(val: string) {
      log.i(val, '选择的值');
      switch (val) {
        case '退出':
          localStorage.clear();
          utilsLocal.clear();
          store.commit(CLEAR_CRUMBS);
          router.replace('/login');
          break;
        case '个人中心':
          visible.value = true;
          break;
        default:
      }
    }
    function ToggleCollapsed() {
      store.commit(COLLAPSED);
    }
    const crumbs = computed({
      get: () => {
        let r = [];
        try {
          r = store.state.crumbs.list.split(',');
        } catch (err) {
          //          console.log('err: ', err);
        }
        return r;
      },
      set: () => {
        // do
      },
    });

    return {
      // data
      data,
      collapsed: computed(() => store.state.sys.collapsed),
      crumbs,
      // methods
      GoTo,
      ToggleCollapsed,

      visible,
      validateInfos,
      resetFields,
      modelRef,
      onSubmit,
    };
  },
});
</script>
<style scoped lang="less">
@import 'header.less';
</style>
