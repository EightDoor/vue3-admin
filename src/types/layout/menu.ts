export interface MenuItem {
  key: number;
  title: string;
  icon?: string;
  path?: string;
  children?: MenuItem[];
  id: number;
  parentId: number;
  crumbs?: string;
  closable?: boolean;
  isHome?: number
  name?: string;
}
export interface MenusInfo {
  selectedKeys: string[];
  list: MenuItem[];
  openKeys: number[];
}
