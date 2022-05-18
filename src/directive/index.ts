import { App } from 'vue';
import PermissionButton from './permission-button';

const Directive = (app: App<Element>): void => {
  // 按钮权限
  app.directive('bt-auth', PermissionButton);
};

export default Directive;
