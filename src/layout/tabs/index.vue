<template>
  <div class="container">
    <a-tabs
      v-model:activeKey="activeKey"
      hide-add
      type="editable-card"
      @change="OnChange"
      @edit="OnEdit"
    >
      <a-tab-pane
        v-for="(pane, index) in panes"
        :key="index"
        :tab="pane.title"
        :closable="pane.closable"
      >
        <slot />
      </a-tab-pane>
    </a-tabs>
  </div>
</template>
<script lang="ts">
import {
  computed,
  defineComponent, ref, toRaw, watch,
} from 'vue';
import { useStore } from 'vuex';
import { useRouter } from 'vue-router';
import _ from 'lodash-es';
import { PanesType } from '@/store/sys/sys-crumbs';
import localStore from '@/utils/store';
import { CURRENT_MENU, STORELETMENUPATH } from '@/utils/constant';
import { SysTab, SysTabDel } from '@/types/sys/tab';
import { DELETETABS, DELETETABSACTION } from '@/store/mutation-types';
import log from '@/utils/log';

export default defineComponent({
  name: 'TabsHome',
  setup() {
    const store = useStore();
    const router = useRouter();
    const activeKey = ref(0);

    const panes = ref<{
      title?: string;
      closable?: boolean;
    }[]>([]);

    const pansList = computed(() => store.state.crumbs.panes);
    watch(
      pansList,
      (newVal) => {
        log.i(newVal, '更新的值');
        localStore.get<SysTab>(CURRENT_MENU).then((res) => {
          if (newVal && newVal.length > 0) {
            log.i(res, '获取到的菜单');
            log.i(newVal, 'newVal');
            if (res) {
              const result = newVal.findIndex((item) => item.id === res.id);
              // 设置当前被选择项
              if (result !== -1) {
                activeKey.value = result;
              }
            }
            panes.value = newVal;
          }
        });
      },
      {
        deep: true,
        immediate: true,
      },
    );

    const FormatData = () => {
      const data = toRaw(store.state.crumbs.panes);
      log.i(data, 'FormatData - data');
      const list: PanesType[] = [];
      data.forEach((item: PanesType) => {
        list.push(toRaw(item));
      });
      return list;
    };
    const OnChange = (val: number) => {
      const result = FormatData()[val];
      activeKey.value = val;
      router.push(result.path);
      store.commit(DELETETABS, { selectData: toRaw(result) });
    };

    const OnEdit = (val: number) => {
      const value = _.cloneDeep(toRaw(activeKey.value));
      const result: SysTabDel = {
        delData: FormatData()[val],
        selectData: '',
      };
      if (FormatData().length > 1) {
        const index = value >= val ? value - 1 : value;

        result.selectData = FormatData()[index];
      }
      store.dispatch(DELETETABSACTION, toRaw(result));
    };
    return {
      activeKey,
      OnChange,
      panes,
      OnEdit,
    };
  },
});
</script>
<style lang="less" scoped>
@import "index.less";
</style>
