import { createApp } from 'vue';
import AntDesign from 'ant-design-vue';
import App from './App.vue';

// 路由
import router from './router/index';
// ant-design-vue
import 'ant-design-vue/dist/antd.css'; // or 'ant-design-vue/dist/antd.less'
// vuex
import store from '@/store/index';
// 公用style
import '@/assets/style/common.less';
// 自定义指令
import Directive from '@/directive/index';

const app = createApp(App);
Directive(app);
app.use(store).use(router).use(AntDesign).mount('#app');
