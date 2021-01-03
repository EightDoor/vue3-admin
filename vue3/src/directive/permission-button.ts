// 按钮权限自定义指令
import store from '@/store/index'
import { DirectiveBinding } from '@vue/runtime-core'
import { localForage } from '@/utils/localforage'
import { STORELETMENUPATH } from '@/utils/constant'
import { MenuItem } from '@/types/layout/menu'

type ButtonPermissionType<T> = 'add' | 'editor' | 'del' | T

function ButtonPermissionType(
  el: HTMLElement,
  binding: DirectiveBinding<ButtonPermissionType<any>>
) {
  const arg = binding.arg
  const value = binding.value
  const permissions = store.state.sys.permissionButtons
  localForage.getItem<MenuItem>(STORELETMENUPATH).then((res) => {
    const data = permissions.filter((item) => item.name === res?.id)
    if (data.length > 0) {
      const r = data.find((item) => item.perms === arg)
      if (r && r.id) {
        el.style.display = 'inline-block'
        if (!value?.title) {
          el.textContent = r.title
        }
        return
      } else {
        el.style.display = 'none'
        return
      }
    } else {
      el.style.display = 'none'
      return
    }
    el.style.display = 'none'
    return
  })
}

export default ButtonPermissionType
