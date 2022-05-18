<template>
  <a-sub-menu :key="menuInfo.key" v-bind="$attrs">
    <template #title>
      <span>
        <!--图标-->
        {{ menuInfo.title }}</span
      >
    </template>
    <template v-for="item in menuInfo.children">
      <template v-if="!item.children">
        <a-menu-item :key="item.key" @click="jumpTo(item)">
          <!--图标-->
          <span>{{ item.title }}</span>
        </a-menu-item>
      </template>
      <template v-else>
        <sub-menu :key="item.key" :menu-info="item" />
      </template>
    </template>
  </a-sub-menu>
</template>

<script lang="ts">
import { defineComponent, toRaw } from 'vue';
import { useRouter } from 'vue-router';
import { useStore } from 'vuex';
import { MenuItem } from '@/types/layout/menu';
import { CURRENT_MENU, STORELETMENUPATH } from '@/utils/constant';
import localStorefrom from '@/utils/store';
import { SETCRUMBSLIST } from '@/store/mutation-types';
import { MenuFormatBrumb } from './menu-common';
import log from '@/utils/log';
import localStore from '@/utils/store';
import { formatArr } from '@/utils';

export default defineComponent({
  name: 'SubMenu',
  props: {
    menuInfo: {
      type: Object,
      default: () => ({}),
    },
  },
  setup() {
    const router = useRouter();
    const store = useStore();
    // methods
    async function jumpTo(item: MenuItem) {
      if (item.path) {
        store.commit(SETCRUMBSLIST, toRaw(item.crumbs));
        await localStorefrom.set(CURRENT_MENU, toRaw(item));
        localStorefrom.get(STORELETMENUPATH).then((res) => {
          log.i(res, '点击二级菜单-获取的存储值');
          let data: MenuItem[] = [];
          if (res) {
            data = [...res, toRaw(item)];
          } else {
            data = [toRaw(item)];
          }
          localStore.set(STORELETMENUPATH, formatArr(data)).then(() => {
            router.push({
              path: item.path || '',
            });
            MenuFormatBrumb(item);
          });
        });
      }
    }
    return {
      // methods
      jumpTo,
    };
  },
});
</script>
