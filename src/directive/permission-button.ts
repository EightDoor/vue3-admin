// 按钮权限自定义指令
import { DirectiveBinding, markRaw, unref } from 'vue';
import store from '@/store/index';
import localStore from '@/utils/store';
import { CURRENT_MENU } from '@/utils/constant';
import { MenuItem } from '@/types/layout/menu';
import log from '@/utils/log';

const ButtonPermissionType = {
  beforeMount: (el: HTMLElement): void => {
    el.style.display = 'none';
  },
  mounted: (el: HTMLElement, binding: DirectiveBinding): void => {
    const { arg } = binding;
    const { value } = binding;
    const permissions = unref(store.state.sys.permissionButtons);
    localStore.get<MenuItem>(CURRENT_MENU).then((res) => {
      const data = permissions.filter((item) => item.name === res?.id);
      if (data.length > 0) {
        const r = data.find((item) => item.perms === arg);
        if (r && r.id) {
          el.style.display = 'inline-block';
          if (!value?.title) {
            el.textContent = r.title;
          }
        } else {
          el.style.display = 'none';
        }
      } else {
        el.style.display = 'none';
      }
    });
  },
};
export default ButtonPermissionType;
