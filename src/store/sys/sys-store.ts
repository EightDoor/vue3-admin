import { Commit } from 'vuex';
import { RouteRecordRaw } from 'vue-router';
import { Key } from 'ant-design-vue/es/_util/type';
import { message } from 'ant-design-vue';
import { cloneDeep } from 'lodash';

import {
  COLLAPSED,
  LOGIN,
  LOGINRESET,
  RESET,
  SET_MENUS_MUTATION,
  SET_SYS,
  SETUSERINFO,
  USERINFOMENUS,
} from '@/store/mutation-types';
import {
  LoginType, MenuType, UserInformation, UserType,
} from '@/types/sys';
import http from '@/utils/request';
import { PERMISSIONBUTTONS, TOKEN } from '@/utils/constant';
import { MenuItem } from '@/types/layout/menu';

import { ListObjCompare, ListToTree } from '@/utils';
import store from '../index';
import log from '@/utils/log';

interface GetToken {
    accessToken: string;
    expiresIn: string;
    msg: string;
}
export interface SysStoreType {
  menus: MenuType[];
  userInfo: Partial<UserType>;
  userInfoMenus: MenuType[];
  permissionButtons: MenuType[];
  collapsed: boolean;
}
export type CustomMenus = RouteRecordRaw & { id: number; parentId: number, children?: any[] };

const getUserInfo = (): Promise<UserInformation | null> => new Promise((resolve, reject) => {
  http<UserInformation>({
    url: 'auth/userInfo',
    method: 'get',
  })
    .then((res) => {
      resolve(res.data);
    })
    .catch((err) => {
      reject(err);
    });
});

async function findModName(modules: any, keyName?: string) {
  const list: any[] = [];
  Object.keys(modules).forEach((key) => {
    const file = modules[key].default;
    if (keyName && file.name === keyName) {
      list.push(file);
      return;
    }
    list.push(file);
  });
  return list;
}

async function baseLayout(): Promise<RouteRecordRaw[]> {
  const modules = await import.meta.globEager('../../layout/layout/**.vue');
  // 加载基础页面结构 layout
  const list:CustomMenus[] = [];
  const result: any[] = await findModName(modules, 'LayoutHome');
  if (result.length > 0) {
    list.push(
      {
        path: '/',
        name: 'LayoutHome',
        component: result[0],
        id: Date.now(),
        parentId: 0,
        children: [],
      },
    );
  } else {
    message.error('layout文件读取失败，请检查');
  }
  return list;
}

// 查询是否存在上级
function queryWhetherParent(parentId: number, item: MenuType[]) {
  let rPath = '';
  const data = item.find((v) => v.id === parentId);
  if (data) {
    if (data.parentId !== 0) {
      queryWhetherParent(data.parentId, item);
    }
    rPath += `/${data.name}`;
    return rPath;
  }
  return '';
}

const FormatMenuTree = async (menuData: MenuType[]): Promise<RouteRecordRaw[]> => {
  // {
  //   name: 'home',
  //     path: 'home',
  //   component: () => import('@/views/home/home.vue'),
  //   meta: {
  //   title: '首页',
  //     icon: '23',
  // },
  //   redirect: '/home',
  // },
  const result: CustomMenus[] = [];
  const modules = await import.meta.globEager('../../views/**/*.vue');
  const childModules = await import.meta.globEager('../../views/**.vue');

  // TODO 待验证是否必须要,     if (!file.isRouter) return;
  menuData.forEach((menuItem) => {
    const fileKey = Object.keys(modules).find((key) => modules[key].default && modules[key].default.name === menuItem.name);
    if (fileKey) {
      const allPath = queryWhetherParent(menuItem.parentId, menuData);
      const file = modules[fileKey].default;
      const obj:CustomMenus = {
        path: `${allPath}/${file.name}`,
        name: menuItem.name,
        component: file,
        redirect: menuItem.redirect,
        meta: {
          title: menuItem.title,
          icon: menuItem.icon,
        },
        id: menuItem.id,
        parentId: menuItem.parentId,
      };
      if (menuItem.name === file.name) {
        result.push(obj);
      } else if (menuItem.type === 1) {
        // 目录类型
        findModName(childModules).then((r) => {
          if (r && r.length > 0) {
            [obj.component] = r;
          }
          result.push(obj);
        });
      }
    }
  });

  const l = await baseLayout();
  const [layout] = l;
  layout.children = ListToTree(result);
  log.d(layout, '加载的路由');

  return [layout];
};
function ListToTreeMenus(jsonData, id = 'id', pid = 'parentId'): MenuItem[] {
  const result: MenuItem[] = [];
  const temp = {};
  for (let i = 0; i < jsonData.length; i += 1) {
    temp[jsonData[i][id]] = jsonData[i]; // 以id作为索引存储元素，可以无需遍历直接定位元素
  }
  for (let j = 0; j < jsonData.length; j += 1) {
    const currentElement = jsonData[j];
    const tempCurrentElementParent = temp[currentElement[pid]]; // 临时变量里面的当前元素的父元素
    if (tempCurrentElementParent) {
      // 如果存在父元素
      if (!tempCurrentElementParent.children) {
        // 如果父元素没有chindren键
        tempCurrentElementParent.children = []; // 设上父元素的children键
      }
      // 组合路由跳转地址
      currentElement.path = `/${tempCurrentElementParent.path}/${currentElement.path}`;
      // 组合面包屑
      currentElement.crumbs = `${tempCurrentElementParent.title},${currentElement.title}`;
      tempCurrentElementParent.children.push(currentElement); // 给父元素加上当前元素作为子元素
    } else {
      // 不存在父元素，意味着当前元素是一级元素
      result.push(currentElement);
    }
  }
  return result;
}
function formatMenus(menus: MenuType[], status = false) {
  return menus.filter((item) => (status ? item.type === 3 : item.type !== 3));
}
export default {
  namespace: true,
  state: {
    menus: [],
    userInfoMenus: [],
    userInfo: null,
    collapsed: false,
  },
  mutations: {
    [COLLAPSED](state: SysStoreType): void {
      state.collapsed = !state.collapsed;
    },
    [SET_MENUS_MUTATION](state: SysStoreType, payload: UserInformation): void {
      const menus = formatMenus(payload.menus);
      let result: MenuItem[] = [];
      const list = cloneDeep(menus.sort(ListObjCompare('order_num')));
      list.forEach((item) => {
        if (!item.hidden) {
          result.push({
            key: item.id || 0,
            title: item.title,
            path: item.name,
            icon: item.icon,
            id: item.id,
            parentId: item.parentId,
            crumbs: `${item.title}`,
          });
        }
      });
      result = ListToTreeMenus(result);
      state.menus = result;
      store.dispatch(RESET, payload.menus);
    },
    [SETUSERINFO](state: SysStoreType, payload: UserInformation): void {
      state.userInfo = payload.userInfo;
    },
    [USERINFOMENUS](state: SysStoreType, payload: UserInformation): void {
      state.userInfoMenus = payload.menus;
    },
    [PERMISSIONBUTTONS](state: SysStoreType, payload: UserInformation): void {
      const menus = cloneDeep(payload.menus);
      const data = formatMenus(menus, true);
      data.forEach((item) => {
        const r = menus.find((v) => v.id === item.parentId);
        if (r) {
          item.name = r.id;
        }
      });
      state.permissionButtons = data;
    },
    [LOGINRESET](state: SysStoreType): void {
      state.permissionButtons = [];
      state.userInfoMenus = [];
      state.menus = [];
      state.userInfo = {
        nickName: '',
        status: 0,
        account: '',
        deptId: '',
      };
    },
  },
  actions: {
    [SET_SYS]({
      commit,
    }: {
      commit: Commit;
    }): Promise<{ userInfo: UserType | null; menus: RouteRecordRaw[] }> {
      return new Promise((resolve, reject) => {
        // 获取数据然后直接存储
        getUserInfo()
          .then(async (res) => {
            if (res) {
              commit(PERMISSIONBUTTONS, res);
              commit(USERINFOMENUS, res);
              commit(SETUSERINFO, res);
              commit(SET_MENUS_MUTATION, res);
              const menus = await FormatMenuTree(formatMenus(res.menus));
              resolve({
                userInfo: res.userInfo,
                menus,
              });
            }
          })
          .catch((err) => {
            reject(err);
          });
      });
    },
    [LOGIN](s: unknown, payload: LoginType): Promise<Key> {
      return new Promise<string>((resolve, reject) => {
        const body = {
          username: payload.account,
          password: payload.pass_word,
        };
        http<GetToken>({
          url: 'auth/login',
          method: 'post',
          body,
        })
          .then(async (res) => {
            localStorage.setItem(TOKEN, res.data?.accessToken ?? '');
            resolve(res.data?.accessToken ?? '');
          })
          .catch((err) => {
            reject(err);
          });
      });
    },
  },
};
