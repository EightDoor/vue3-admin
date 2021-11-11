export interface MenuItem {
  key: number;
  title: string;
  icon?: string;
  path?: string;
  children?: MenuItem[];
  id: number;
  parentId: number;
  crumbs?: string;
  isHome?: boolean;
}
export interface MenusInfo {
  selectedKeys: string[];
  list: MenuItem[];
  openKeys: string[];
}
