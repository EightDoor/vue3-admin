import { PanesType } from '@/store/sys/sys-crumbs'
import store from '@/store/index'
import { MENUTABS } from '@/store/mutation-types'
import { toRaw } from 'vue'
import { MenuItem } from '@/types/layout/menu'

const MenuFormatBrumb = (val: MenuItem): void => {
  const data: PanesType = {
    id: val.id || '',
    title: val.title,
    parent_id: val.parent_id || '',
    path: val.path || '',
  }
  store.commit(MENUTABS, toRaw(data))
}

export { MenuFormatBrumb }
