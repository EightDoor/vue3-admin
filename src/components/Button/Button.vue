<template>
  <a-button :type="type" @click="CommonClick()">
    <template v-if="iconName" #icon>
      <PlusOutlined v-if="iconName === 'add'" />
    </template>
    {{ title }}
  </a-button>
</template>
<script lang="ts">

import { defineComponent, PropType } from 'vue';
import { PlusOutlined } from '@ant-design/icons-vue';

type buttonType = 'default' | 'primary' | 'ghost' | 'dashed' | 'link' | 'text';

const CommonButton = defineComponent({
  name: 'CommonButton',
  components: {
    PlusOutlined,
  },
  props: {
    title: {
      type: String,
      default: '默认',
    },
    iconName: {
      type: String,
      default: '',
    },
    type: {
      type: String as PropType<buttonType>,
      default: 'primary',
    },
  },
  emits: ['change'],
  setup(props, { emit }) {
    function CommonClick() {
      emit('change');
    }
    return {
      CommonClick,
    };
  },
});

export default CommonButton;
</script>
