import { ColumnProps } from 'ant-design-vue/lib/table/interface'

export interface CommonResponse<T> extends PaginType {
  list: T[]
}

// 请求接口分页
export interface PaginType {
  page?: number
  page_size?: number
  total?: number
}
// 表格默认的分页
export interface TablePaginType {
  current: number
  pageSize: number
  total: number
}

export interface CommonResponseSing<T> extends PaginType {
  list: T
}

// 表格
export interface TableDataType<T> extends PaginType {
  data: T[]
  columns: AntColumn[]
  loading: boolean
}
interface AntColumn extends ColumnProps, AntColumnSlot {}
interface AntColumnSlot {
  slots?: {
    customRender: string
  }
}

// 通用的业务类型字段
export interface CommonTableList {
  id?: string
  created_at?: string
  updated_at?: string
}

// 树形tree-select
export interface CommonTreeSelect {
  title?: string
  value?: string
  key?: string
}
// 树形结构选中selectKey
export interface CommonTreeSelectKeys {
  checked: string[]
  halfChecked: string[]
}
