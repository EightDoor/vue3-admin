// 按钮权限自定义指令
import store from "@/store/index";
import localStore from "@/utils/store";
import { STORELETMENUPATH } from "@/utils/constant";
import { MenuItem } from "@/types/layout/menu";
import { DirectiveBinding } from "vue";

type ButtonPermissionType<T> = "add" | "editor" | "del" | T;

const ButtonPermissionType = {
  beforeMount: (el: HTMLElement): void => {
    el.style.display = "none";
  },
  mounted: (el: HTMLElement, binding: DirectiveBinding): void => {
    const arg = binding.arg;
    const value = binding.value;
    const permissions = store.state.sys.permissionButtons;
    localStore.get<MenuItem>(STORELETMENUPATH).then((res) => {
      const data = permissions.filter((item) => item.name === res?.id);
      if (data.length > 0) {
        const r = data.find((item) => item.perms === arg);
        if (r && r.id) {
          el.style.display = "inline-block";
          if (!value?.title) {
            el.textContent = r.title;
          }
          return;
        } else {
          el.style.display = "none";
          return;
        }
      } else {
        el.style.display = "none";
        return;
      }
    });
  },
};
export default ButtonPermissionType;
