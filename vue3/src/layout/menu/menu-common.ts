import { PanesType } from '@/store/sys/sys-crumbs'
import store from '@/store/index'
import { MENUTABS } from '@/store/mutation-types'
import { toRaw } from 'vue'

const MenuFormatBrumb = (val: PanesType): void => {
  store.commit(MENUTABS, toRaw(val))
}

export { MenuFormatBrumb }
