import { ColumnProps } from 'ant-design-vue/lib/table/interface';

// 请求接口分页
export interface PaginType {
    page?: number;
    page_size?: number;
    total?: number;
}

export interface CommonResponse<T> extends PaginType {
  list: T[];
  data: T | null
}

// 表格默认的分页
export interface TablePaginType {
  current: number;
  pageSize: number;
  total: number;
}

export interface CommonResponseSing<T> extends PaginType {
  list: T;
}

interface AntColumnSlot {
    slots?: {
        customRender: string;
    };
    title: string;
    dataIndex?: string;
    key?: string;
    fixed?: string | boolean;
    width?: number;
}

interface AntColumn extends ColumnProps, AntColumnSlot {}

// 表格
export interface TableDataType<T> extends PaginType {
  data: T[];
  columns: AntColumn[];
  loading: boolean;
}

// 通用的业务类型字段
export interface CommonTableList {
  id?: number;
  created_at?: string;
  updated_at?: string;
}

// 树形tree-select
export interface CommonTreeSelect {
  title?: string;
  value?: string;
  key?: string;
}
// 树形结构选中selectKey
export interface CommonTreeSelectKeys {
  checked: string[];
  halfChecked: string[];
}
