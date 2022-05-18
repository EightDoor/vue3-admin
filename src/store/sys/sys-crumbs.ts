import { Commit } from 'vuex';
import {
  SETCRUMBSLIST,
  MENUTABS,
  DELETETABS,
  DELETETABSACTION,
  RESET,
  RESETMU, SET_MENU_LIST, RESETMU_CURRENT_CRUMB, RESETMU_ACTION, CLEAR_CRUMBS,
} from '@/store/mutation-types';
import { MenuItem } from '@/types/layout/menu';
import { SysTabDel } from '@/types/sys/tab';
import { CURRENT_MENU, STORELETMENUPATH } from '@/utils/constant';
import localStore from '@/utils/store';
import localStorefrom from '@/utils/store';
import log from '@/utils/log';

export interface PanesType {
  id: number;
  title: string;
  path: string;
  parentId: number;
  closable?: boolean;
  timestamp?: number;
}
export interface CrumbsStoreType {
  list: string[];
  panes: PanesType[];
  delPane: PanesType;
  selectPane: PanesType | string;
}
interface ResetType {
  result: PanesType;
  title: string;
}
export default {
  namespace: true,
  state: {
    list: '',
    panes: [],
    selectPane: '',
  },
  mutations: {
    [SETCRUMBSLIST](state: CrumbsStoreType, payload: string[]): void {
      state.list = payload;
    },
    [MENUTABS](state: CrumbsStoreType, payload: PanesType): void {
      const data: PanesType[] = state.panes;
      const doesIsExist = data.findIndex((item) => item.id === payload.id);
      if (doesIsExist === -1) {
        data.push(payload);
        state.panes = data.filter((item) => item.title);
      } else {
        data[doesIsExist].timestamp = Date.now();
        state.panes = data;
      }
      log.i(state.panes, 'MENUTABS - pans');
    },
    [DELETETABS](state: CrumbsStoreType, payload: SysTabDel): void {
      state.selectPane = payload.selectData;
      if (payload.delData) {
        state.panes = state.panes.filter(
          (item: PanesType) => item.id !== payload.delData.id,
        );
      }
    },
    /**
     * 初始化设置面包屑所有值
     * @param state
     * @param payload
     */
    [RESETMU](state: CrumbsStoreType, payload: PanesType[]) {
      // 读取面包屑所有值
      state.panes = payload;
    },
    /**
     * 初始化设置面包屑 当前选中值
     * @param state
     * @param payload
     */
    [RESETMU_CURRENT_CRUMB](state: CrumbsStoreType, payload: PanesType): void {
      // 读取面包屑所有值
      state.selectPane = payload;
      state.list = [payload.title];
    },
    [SET_MENU_LIST](state: CrumbsStoreType, payload: PanesType[]): void {
      state.panes = payload;
    },
    [CLEAR_CRUMBS](state: CrumbsStoreType): void {
      state.panes = [];
      state.list = [];
      state.selectPane = '';
    },
  },
  actions: {
    [DELETETABSACTION](
      { commit }: { commit: Commit },
      payload: SysTabDel,
    ): void {
      // 设置菜单选中项
      localStore.set(STORELETMENUPATH, payload.selectData).then(() => {
        commit(DELETETABS, payload);
      });
    },
    // 设置当前的登录默认显示的选中项
    // 默认首页只能是一级的
    // TODO 待完善  刷新页面直接选择对应的选中项
    async [RESET]({ commit }: { commit: Commit }, payload: MenuItem[]) {
      console.log(payload, 'payload');
      const data = payload.filter((item) => item.isHome);
      if (data.length > 0) {
        const r = data[0];
        const result = {
          id: r.id || '',
          title: r.title,
          path: `/${r.name}`,
          parent_id: r.parentId || '',
          closable: false,
        };

        // 读取面包屑所有值
        const allCrumbs = await localStorefrom.get(STORELETMENUPATH);
        if (allCrumbs) {
          commit(RESETMU, allCrumbs);
        } else {
          await localStorefrom.set(STORELETMENUPATH, [result]);
          commit(RESETMU, [result]);
        }
        // 读取当前选择的面包屑
        const currentCrumb = await localStorefrom.get(CURRENT_MENU);
        if (currentCrumb) {
          commit(RESETMU_CURRENT_CRUMB, currentCrumb);
        } else {
          commit(RESETMU_CURRENT_CRUMB, r.title);
        }
      }
    },
  },
};
