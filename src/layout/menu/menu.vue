<template>
  <a-menu
    v-model:openKeys="menusInfo.openKeys"
    v-model:selectedKeys="menusInfo.selectedKeys"
    mode="inline"
    theme="dark"
    :inline-collapsed="collapsed"
    :collapsed="collapsed"
  >
    <template v-for="item in getMenus" >
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
  </a-menu>
</template>
<script lang="ts">
import {
  defineComponent,
  reactive,
  onMounted,
  computed,
  toRaw,
  watch,
} from 'vue';
import { useRouter } from 'vue-router';
import { useStore } from 'vuex';
import SubMenu from './menu-item.vue';
import { MenuItem, MenusInfo } from '@/types/layout/menu';
import { STORELETMENUPATH } from '@/utils/constant';
import { MenuType } from '@/types/sys';
import localStore from '@/utils/store';
import { SETCRUMBSLIST } from '@/store/mutation-types';
import { MenuFormatBrumb } from './menu-common';
import { PanesType } from '@/store/sys/sys-crumbs';

interface InitTopTabs extends MenuItem {
  crumbs: string;
}
export default defineComponent({
  name: 'CommonMenu',
  components: {
    SubMenu,
  },
  setup() {
    const router = useRouter();
    const store = useStore();
    const menusInfo = reactive<MenusInfo>({
      selectedKeys: [],
      list: [],
      openKeys: [],
    });
    const getUserInfoMenus = computed(() => store.state.sys.userInfoMenus);

    const FormatSelectKey = (res) => {
      menusInfo.selectedKeys = [res.key || res.id || ''];
      const { parent_id } = res;
      const data: MenuType[] = toRaw(getUserInfoMenus.value);
      const r = data.find((item) => item.id === parent_id);
      if (r) {
        menusInfo.openKeys = [String(r.id) || ''];
      }
    };

    onMounted(() => {
      menusInfo.list = [];
      localStore.get<InitTopTabs>(STORELETMENUPATH).then((res) => {
        if (res) {
          // 初始化顶部面包屑
          store.commit(SETCRUMBSLIST, toRaw(res.crumbs));
          // 初始化顶部tabs
          MenuFormatBrumb(res);
          FormatSelectKey(res);
        }
      });
    });

    // methods
    function jumpTo(item: MenuItem) {
      if (item.path) {
        store.commit(SETCRUMBSLIST, toRaw(item.crumbs));
        localStore.set(STORELETMENUPATH, toRaw(item)).then(() => {
          router.push({
            path: item.path || '',
          });
          MenuFormatBrumb(item);
        });
      }
    }
    watch(
      () => store.state.crumbs.selectPane,
      (newValue: PanesType) => {
        const data = toRaw(newValue);
        FormatSelectKey(data);
      },
    );
    return {
      // data
      menusInfo,
      getMenus: computed(() => store.state.sys.menus),
      collapsed: computed(() => store.state.sys.collapsed),
      // methods
      jumpTo,
    };
  },
});
</script>
<style scoped lang="less">
@import "./menu.less";
</style>
