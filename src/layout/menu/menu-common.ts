import { markRaw } from 'vue';
import { PanesType } from '@/store/sys/sys-crumbs';
import store from '@/store/index';
import { MENUTABS } from '@/store/mutation-types';
import { MenuItem } from '@/types/layout/menu';

const MenuFormatBrumb = (val: MenuItem): void => {
  const data: PanesType = {
    id: val.id || 0,
    title: val.title,
    parentId: val.parentId || 0,
    path: val.path || '',
  };
  store.commit(MENUTABS, markRaw(data));
};

export { MenuFormatBrumb };
