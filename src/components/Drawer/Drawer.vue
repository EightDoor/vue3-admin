<template>
  <a-drawer
    :title="title"
    :placement="placement"
    v-model:visible="commdrawerData.visible"
    :width="width"
    @close="onClose"
  >
    <div class="drawerContainer">
      <slot></slot>
    </div>
    <div class="drawerBottom" v-if="okText || cancelText">
      <a-button
        :loading="loading"
        v-if="cancelText"
        style="margin-right: 20px"
        @click="onCancel"
      >
        {{ cancelText }}
      </a-button>
      <a-button :loading="loading" v-if="okText" type="primary" @click="onOk">
        {{ okText }}
      </a-button>
    </div>
  </a-drawer>
</template>
<script lang="ts">
export interface DrawerProps {
  visible: boolean
  title: string
  placement?: string
  loading?: boolean
}
import { defineComponent, reactive, watch } from 'vue'

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
  setup(props, { emit }) {
    const commdrawerData = reactive({
      visible: false,
    })
    function onCancel() {
      emit('on-close')
      commdrawerData.visible = false
    }
    function onOk() {
      emit('on-ok')
    }
    function onClose() {
      emit('on-close')
      commdrawerData.visible = false
    }
    watch(
      () => props.visible,
      (data) => {
        commdrawerData.visible = data
      }
    )
    return {
      commdrawerData,
      onCancel,
      onOk,
      onClose,
    }
  },
})

export default CommonDrawer
</script>
<style lang="less" scoped>
@import './Drawer.less';
</style>
