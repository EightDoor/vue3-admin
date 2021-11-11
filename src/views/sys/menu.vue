<template>
  <div class="space-margin-bottom">
    <common-button
      v-bt-auth:add="{ title: true }"
      title="添加"
      icon-name="add"
      @change="ChangeClick()"
    />
  </div>
  <a-table
    :scroll="{ x: 500 }"
    :columns="tableCont.columns"
    row-key="id"
    :data-source="tableCont.data"
    :pagination="{
      total: tableCont.total,
    }"
    :loading="tableCont.loading"
    @change="Change"
  >
    <template #action="{ record }">
      <a-button
        v-bt-auth:edit
        type="primary"
        style="margin-right: 15px"
        @click="Editor(record)"
      ></a-button>
      <a-popconfirm
        title="确定删除吗?"
        ok-text="删除"
        cancel-text="取消"
        @confirm="Del(record)"
      >
        <a-button v-bt-auth:del type="danger"></a-button>
      </a-popconfirm>
    </template>
  </a-table>
  <common-drawer
    :title="drawerData.title"
    :visible="drawerData.visible"
    cancel-text="取消"
    ok-text="确定"
    :loading="drawerData.loading"
    @onClose="drawerData.visible = false"
    @onOk="onSubmit"
  >
    <a-form :rules="rules" ref="formRef" :label-col="{ span: 4 }" :wrapper-col="{ span: 20 }">
      <a-form-item label="父级id" >
        <a-tree-select
          v-model:value="validateInfos.parentId"
          style="width: 100%"
          :dropdown-style="{ maxHeight: '400px', overflow: 'auto' }"
          :tree-data="treeOptions.options"
          placeholder="请选择"
          tree-default-expand-all
        >
        </a-tree-select>
      </a-form-item>
      <a-form-item label="名称" >
        <a-input v-model:value="validateInfos.title"></a-input>
      </a-form-item>
      <a-form-item label="菜单类型" >
        <a-select v-model:value="validateInfos.type">
          <a-select-option
            v-for="(item, index) in optionsType"
            :key="index"
            :value="item.value"
          >
            {{ item.label }}
          </a-select-option>
        </a-select>
      </a-form-item>
      <a-form-item v-if="validateInfos.type === 3" label="权限标识">
        <a-input v-model:value="validateInfos.perms"></a-input>
      </a-form-item>
      <div v-if="validateInfos.type !== 3">
        <a-form-item label="菜单name">
          <a-input v-model:value="validateInfos.name"></a-input>
        </a-form-item>
        <a-form-item label="是否首页">
          <a-radio-group v-model:value="validateInfos.isHome">
            <a-radio :value="true"> 是 </a-radio>
            <a-radio :value="false"> 否 </a-radio>
          </a-radio-group>
        </a-form-item>
        <a-form-item label="重定向地址">
          <a-input v-model:value="validateInfos.redirect"></a-input>
        </a-form-item>
        <a-form-item label="图标">
          <a-input v-model:value="validateInfos.icon"></a-input>
        </a-form-item>
        <a-form-item label="是否隐藏">
          <a-radio-group v-model:value="validateInfos.hidden" name="radioGroup">
            <a-radio :value="0"> 否 </a-radio>
            <a-radio :value="1"> 是 </a-radio>
          </a-radio-group>
        </a-form-item>
      </div>
      <a-form-item label="排序" >
        <a-input-number v-model:value="validateInfos.orderNum"></a-input-number>
      </a-form-item>
    </a-form>
  </common-drawer>
</template>
<script lang="ts">
import {
  defineComponent, reactive, onMounted, toRaw, ref,
} from 'vue';
import { useForm } from '@ant-design-vue/use';
import { message } from 'ant-design-vue';
import { Method } from 'axios';
import { cloneDeep } from 'lodash';
import CommonButton from '@/components/Button/Button.vue';
import CommonDrawer, { DrawerProps } from '@/components/Drawer/Drawer.vue';
import http from '@/utils/request';
import { MenuType } from '@/types/sys';
import { TableDataType, TablePaginType } from '@/types/type';
import { ListObjCompare, ListToTree } from '@/utils';
import { searchParam } from '@/utils/search_param';

const SysMenu = defineComponent({
  name: 'SysMenu',
  isRouter: true,
  components: { CommonButton, CommonDrawer },
  setup() {
    const optionsType = ref([
      {
        value: 1,
        label: '目录',
      },
      {
        value: 2,
        label: '菜单',
      },
      {
        value: 3,
        label: '按钮',
      },
    ]);
    const width = 150;
    const tableCont = reactive<TableDataType<MenuType>>({
      page: 1,
      pageSize: 1000,
      total: 0,
      loading: false,
      columns: [
        {
          title: '名称',
          dataIndex: 'title',
          width: width * 2,
          fixed: true,
        },
        {
          title: '类型',
          dataIndex: 'type',
          width: width / 2,
        },
        {
          title: '重定向地址',
          dataIndex: 'redirect',
          width,
        },
        {
          title: '图标',
          dataIndex: 'icon',
          width,
        },
        {
          title: '父级节点',
          dataIndex: 'parentId',
          width,
        },
        {
          title: '权限标识',
          dataIndex: 'perms',
          width: width / 2,
        },
        {
          title: '菜单标识',
          dataIndex: 'name',
          width: width / 2,
        },
        {
          title: '排序',
          dataIndex: 'orderNum',
          width: width / 2,
        },
        {
          title: '操作',
          key: 'action',
          fixed: 'right',
          width: 200,
          slots: {
            customRender: 'action',
          },
        },
      ],
      data: [],
    });
    const treeOptions = reactive<{ options: MenuType[] }>({ options: [] });
    const drawerData = reactive<DrawerProps>({
      title: '添加',
      visible: false,
      loading: false,
    });
    const formRef = ref();
    let validateInfos = reactive<MenuType>({
      parentId: 0,
      path: '',
      component: '',
      redirect: '',
      icon: '',
      title: '',
      perms: '',
      name: '',
      type: 1,
      orderNum: 1,
      id: 0,
      hidden: 0,
      isHome: false,
    });
    const rules = reactive({
      parent_id: [
        {
          required: true,
          message: '请选择父级',
        },
      ],
      title: [
        {
          required: true,
          message: '请输入名称',
        },
      ],
      type: [
        {
          required: true,
          message: '请选择类型',
        },
      ],
      order_num: [
        {
          required: true,
          message: '请输入排序',
        },
      ],
    });
    function getList() {
      tableCont.loading = true;
      http<MenuType>({
        url: `/menu${searchParam({
          page: tableCont.page,
          limit: tableCont.pageSize,
        })}`,
        method: 'GET',
      }).then((res) => {
        const list = cloneDeep(res.list?.data.sort(ListObjCompare('order_num'))) || [];
        tableCont.loading = false;
        tableCont.data = ListToTree(list);
        tableCont.total = res.list?.total;
        const treeMenus = cloneDeep(res.list?.data || []);
        treeMenus.forEach((item) => {
          item.value = String(item.id);
        });
        treeOptions.options = ListToTree(treeMenus);
      });
    }
    onMounted(() => {
      getList();
    });
    const onSubmit = () => {
      formRef.value.validate()
        .then(() => {
          drawerData.loading = true;
          const data = cloneDeep(toRaw(validateInfos));

          // @ts-ignore
          delete data.id;
          let method: Method = 'POST';
          if (validateInfos.id) {
            method = 'PUT';
          }
          http<MenuType>({
            url: validateInfos.id ? `menu/${validateInfos.id}` : 'menu',
            method,
            body: data,
          }).then(() => {
            message.success(`${drawerData.title}成功`);
            drawerData.loading = false;
            drawerData.visible = false;
            getList();
          });
        })
        .catch((err) => {
          console.log('error', err);
        });
    };
    function ChangeClick() {
      drawerData.title = '添加';
      drawerData.visible = true;
      formRef.value.resetFields();
      validateInfos.parentId = 0;
    }
    function Editor(record: MenuType) {
      drawerData.title = '修改';
      drawerData.visible = true;
      validateInfos = Object.assign(validateInfos, record);
    }
    function Del(record: MenuType) {
      const data = toRaw(record);
      http({ url: `menu/${data.id}`, method: 'delete' }).then(() => {
        message.success('删除成功');
        getList();
      });
    }
    function Change(pagination: TablePaginType) {
      tableCont.page = pagination.current;
      getList();
    }

    return {
      tableCont,
      ChangeClick,
      drawerData,
      treeOptions,
      Change,
      optionsType,

      // table
      Editor,
      Del,

      // form
      validateInfos,
      formRef,
      rules,
      onSubmit,
    };
  },
});
export default SysMenu;
</script>
<style scoped lang="less">
@import "./depart.less";
</style>
