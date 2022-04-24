<template>
  <base-container>
    <fs-crud ref="crudRef" v-bind="crudBinding" />
  </base-container>
</template>

<script lang="ts">
import { defineComponent, ref, onMounted } from 'vue';
import { useExpose, useCrud } from '@fast-crud/fast-crud';
import BaseContainer from '@/components/BaseContainer/index.vue';
import createCrudOptions from './crud';

export default defineComponent({
  name: 'DemoAvue', // 实际开发中可以修改一下name
  components: {
    BaseContainer,
  },
  setup() {
    // crud组件的ref
    const crudRef = ref<any>();
    // crud 配置的ref
    const crudBinding = ref<any>();
    // 暴露的方法
    const { expose } = useExpose({ crudRef, crudBinding });
    // 你的crud配置
    const { crudOptions } = createCrudOptions({ expose });
    // 初始化crud配置
    // eslint-disable-next-line @typescript-eslint/no-unused-vars,no-unused-vars
    const { resetCrudOptions } = useCrud({ expose, crudOptions });
    // 你可以调用此方法，重新初始化crud配置
    // resetCrudOptions(options)
    // 页面打开后获取列表数据
    onMounted(() => {
      expose.doRefresh();
    });
    return {
      crudBinding,
      crudRef,
    };
  },
});
</script>
