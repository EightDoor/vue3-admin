import { createApp } from 'vue'
import App from './App.vue'
import './index.css'

// 路由
import router from './router/index'
// ant-design-vue
import AntDesign from 'ant-design-vue'
import 'ant-design-vue/dist/antd.css' // or 'ant-design-vue/dist/antd.less'
// vuex
import store from '@/store/index'
// 公用style
import '@/assets/style/common.less'

const app = createApp(App)
app.use(store).use(router).use(AntDesign).mount('#app')
