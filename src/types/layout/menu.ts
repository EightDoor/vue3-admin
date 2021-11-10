export interface MenuItem {
  key: string
  title: string
  icon?: string
  path?: string
  children?: MenuItem[]
  id?: string
  parent_id?: string
  crumbs?: string
  is_home?: boolean
}
export interface MenusInfo {
  selectedKeys: string[]
  list: MenuItem[]
  openKeys: string[]
}
