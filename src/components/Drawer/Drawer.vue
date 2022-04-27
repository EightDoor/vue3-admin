<template>
  <a-drawer
    v-model:visible="commdrawerData.visible"
    :title="title"
    :placement="placement"
    :width="width"
    @close="onClose"
    destroyOnClose
  >
    <div class="drawerContainer">
      <slot></slot>
    </div>
    <div v-if="okText || cancelText" class="drawerBottom">
      <a-button
        v-if="cancelText"
        :loading="commdrawerData.loading"
        style="margin-right: 20px"
        @click="onCancel"
      >
        {{ cancelText }}
      </a-button>
      <a-button
        v-if="okText"
        :loading="commdrawerData.loading"
        type="primary"
        @click="onOk"
      >
        {{ okText }}
      </a-button>
    </div>
  </a-drawer>
</template>
<script lang="ts">
import { defineComponent, reactive, watch } from 'vue';

export interface DrawerProps {
  visible: boolean;
  title: string;
  placement?: string;
  loading?: boolean;
}

const CommonDrawer = defineComponent({
  name: 'ComponentDrawer',
  props: {
    title: {
      type: String,
      required: true,
    },
    placement: {
      type: String,
      default: 'right',
    },
    okText: {
      type: String,
      default: '',
    },
    cancelText: {
      type: String,
      default: '',
    },
    width: {
      type: String,
      default: '45%',
    },
    visible: {
      type: Boolean,
      default: false,
    },
    loading: {
      type: Boolean,
      default: false,
    },
  },
  emits: ['on-close', 'on-ok'],
  setup(props, { emit }) {
    const commdrawerData = reactive({
      visible: false,
      loading: false,
    });
    function onCancel() {
      emit('on-close');
      commdrawerData.visible = false;
    }
    function onOk() {
      emit('on-ok');

      // 超时自定关闭loading
      setTimeout(() => {
        commdrawerData.loading = false;
      }, 5000);
    }
    function onClose() {
      emit('on-close');
      commdrawerData.visible = false;
    }
    watch(
      () => props.visible,
      (data) => {
        commdrawerData.visible = data;
        commdrawerData.loading = false;
      },
    );
    return {
      commdrawerData,
      onCancel,
      onOk,
      onClose,
    };
  },
});

export default CommonDrawer;
</script>
<style lang="less" scoped>
@import './Drawer.less';
</style>
