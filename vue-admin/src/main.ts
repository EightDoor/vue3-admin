import { createApp } from 'vue'
import App from './App.vue'
import './index.css'

// 路由
import router from './router/index';
// ant-design-vue
import AntDesign from 'ant-design-vue';
import 'ant-design-vue/dist/antd.css'; // or 'ant-design-vue/dist/antd.less'

createApp(App).use(router).use(AntDesign).mount('#app')
