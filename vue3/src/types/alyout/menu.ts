export interface MenuItem {
  key: string
  title: string
  path?: string
  children?: MenuItem[]
}
export interface MenusInfo {
  selectedKeys: string[]
  collapsed: boolean
  list: MenuItem[]
}
