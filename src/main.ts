import { createApp } from 'vue';
import AntDesign from 'ant-design-vue';
import ui from '@fast-crud/ui-antdv';
import { FastCrud } from '@fast-crud/fast-crud';
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

// 引入fast-crud
import '@fast-crud/fast-crud/dist/style.css';

const app = createApp(App);
Directive(app);

// 先安装ui
app.use(ui);
// 然后安装FastCrud
app.use(FastCrud, {
  // 此处配置公共的dictRequest（字典请求）
  // async dictRequest({ dict }) {
  //     return await request({ url: dict.url }); //根据dict的url，异步返回一个字典数组
  // },
  // 公共crud配置
  commonOptions() {
    return {
      request: {
        // 接口请求配置
        // 你项目后台接口大概率与fast-crud所需要的返回结构不一致，所以需要配置此项
        // 请参考文档http://fast-crud.docmirror.cn/api/crud-options/request.html
        transformQuery: ({ page, form, sort }) =>
        // 转换为你pageRequest所需要的请求参数结构
          ({ page, form, sort }),
        transformRes: ({ res }) =>
        // 将pageRequest的返回数据，转换为fast-crud所需要的格式
        // return {records,currentPage,pageSize,total};
          ({
            records: res.list.data,
            currentPage: res.list.page,
            pageSize: 20,
            total: res.list.total,
          }),

      },
    };
  },
});
app.use(store).use(router).use(AntDesign).mount('#app');
