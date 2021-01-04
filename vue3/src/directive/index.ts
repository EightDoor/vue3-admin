import { App } from 'vue'
import PermissionButton from './permission-button'

const Directive = (app: App<Element>) => {
  // 按钮权限
  // eslint-disable-next-line @typescript-eslint/ban-ts-comment
  // @ts-ignore
  app.directive('bt-auth', PermissionButton)
}

export default Directive
