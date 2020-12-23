import { CommonTableList, CommonTreeSelect } from '@/types/type'

export interface DepartType extends CommonTreeSelect {
  parent_id: string
  name: string
  order_num: number
  id: string
  children?: Array<any>
}

export interface UserType extends CommonTableList {
  // 用户登录密码
  pass_word: string
  // 用户登录账号
  account: string
  // 用户显示昵称
  nick_name: string
  // 邮箱地址
  email?: string
  // 所属状态是否有效  1是有效 0是失效
  status: 0 | 1
  // 头像地址
  avatar?: string
  // 部门id
  dept_id: string
  // 用户手机号码
  phone_num?: string
}

export interface RoleType extends CommonTableList {
  // 描述
  remark?: string;
  // 角色名称
  role_name: string
}

export interface MenuType extends CommonTableList {
  id: string;
  // 父级id
  parent_id: string;
  // 菜单名称
  name: string;
  // 菜单类型：1. 目录 2. 菜单  3. 按钮
  type: number;
  // 排序
  order_num: number;
  // 权限标识，接口标识
  perms?: string
  // 菜单标识，前端路由name
  code: string;
  children?: Array<any>
  title?: string;
  value?:string;
}
