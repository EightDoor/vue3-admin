
export interface CommonResponse<T> extends PaginType {
  list: T[]
}

// 分页
export interface PaginType {
  page?: number
  page_size?: number,
  total?: number
}

export interface CommonResponseSing<T> extends PaginType {
  list: T
}

// 表格
export interface TableDataType<T> extends PaginType{
  data: T[]
  columns: any,
  loading: boolean
}

// 通用的业务类型字段
export interface CommonTableList {
  id?: string;
  created_at?: string;
  updated_at?: string;
}

// 树形tree-select
export interface CommonTreeSelect {
  title?: string
  value?: string
  key?: string;
}
