import { CommonTableList, CommonTreeSelect } from '@/types/type';

export interface MenuType extends CommonTableList {
    crumbs?: string;
    key?: number;
    id: number;
    createdAt?: any;
    updatedAt?: any;
    deletedAt?: any;
    parentId: number;
    title: string;
    type?: number;
    orderNum?: number;
    perms?: any;
    name?: string;
    path?: string;
    component?: string;
    redirect?: any;
    icon?: any;
    isHome?: boolean;
    children?: Array<MenuType>;
    value?: string;
    hidden?: number;
}

export interface DepartType extends CommonTreeSelect {
  parent_id: string;
  name: string;
  order_num: number;
  id?: string;
  children?: Array<MenuType>;
}

export interface UserType extends CommonTableList {
  // 用户登录密码
  pass_word: string;
  // 用户登录账号
  account: string;
  // 用户显示昵称
  nick_name: string;
  // 邮箱地址
  email?: string;
  // 所属状态是否有效  1是有效 0是失效
  status: 0 | 1;
  // 头像地址
  avatar?: string;
  // 部门id
  dept_id: string;
  // 用户手机号码
  phone_num?: string;
  depart_info?: DepartType;
}

export interface RoleType extends CommonTableList {
  // 描述
  remark?: string;
  // 角色名称
  role_name: string;
  id?: number;
}

export interface LoginType {
  account: string;
  pass_word: string;
}
export interface UserInfoType {
  menu: MenuType[];
  user_info: UserType;
}

// 新定义的获取用户信息
interface Role {
    id: number;
    createdAt?: any;
    updatedAt?: any;
    deletedAt?: any;
    remark: string;
    roleName: string;
}

export interface UserInfo {
    id: number;
    createdAt?: any;
    updatedAt?: any;
    deletedAt?: any;
    account: string;
    nickName: string;
    email?: any;
    status: number;
    avatar?: any;
    deptId: string;
    phoneNum?: any;
}

export interface UserInformation {
    userInfo: UserInfo;
    menus: MenuType[];
    roles: Role[];
}
