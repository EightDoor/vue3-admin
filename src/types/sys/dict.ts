import { CommonTableList } from '../type'

// 字典管理
export interface SysDict extends CommonTableList {
  name: string
  serial_number?: string
  describe?: string
}
// 字典项
export interface SysDictItem extends CommonTableList {
  // 名称
  label: string
  // 数据值
  value: string
  // 描述
  describe: string
  // 字典值id
  dict_id?: string
}
