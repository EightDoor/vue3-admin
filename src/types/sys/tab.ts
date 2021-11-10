import { PanesType } from '@/store/sys/sys-crumbs'

// tabs
export interface SysTab {
  crumbs: string
  icon: string
  id: string
  key: string
  parent_id: string
  path: string
  title: string
}

export interface SysTabDel {
  delData: PanesType
  selectData: PanesType | string
}
