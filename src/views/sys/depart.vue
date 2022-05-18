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
    :columns="tableCont.columns"
    row-key="id"
    :data-source="tableCont.data"
    :pagination="{
      total: tableCont.total,
    }"
    :loading="tableCont.loading"
    @change="Change"
  >
    <template #bodyCell="{column, record}">
      <template v-if="column.key === 'action'">
        <a-button v-bt-auth:edit type="primary" style="margin-right: 15px" @click="Editor(record)" />
        <a-popconfirm title="确定删除吗?" ok-text="删除" cancel-text="取消" @confirm="Del(record)">
          <a-button v-bt-auth:del danger />
        </a-popconfirm>
      </template>
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
    <a-form :model="formData" :rules="rules" ref="formRef" :label-col="{ span: 4 }" :wrapper-col="{ span: 20 }">
      <a-form-item label="父级id" name="parentId">
        <a-tree-select
          v-model:value="formData.parentId"
          style="width: 100%"
          :dropdown-style="{ maxHeight: '400px', overflow: 'auto' }"
          :tree-data="treeOptions.options"
          placeholder="请选择"
          tree-default-expand-all
        ></a-tree-select>
      </a-form-item>
      <a-form-item label="名称" name="name">
        <a-input v-model:value="formData.name"></a-input>
      </a-form-item>
      <a-form-item label="排序" name="orderNum">
        <a-input-number v-model:value="formData.orderNum"></a-input-number>
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
import { cloneDeep } from 'lodash-es';
import CommonButton from '@/components/Button/Button.vue';
import CommonDrawer, { DrawerProps } from '@/components/Drawer/Drawer.vue';
import http from '@/utils/request';
import { DepartType } from '@/types/sys';
import { TableDataType, TablePaginType } from '@/types/type';
import { ListObjCompare, ListToTree } from '@/utils';
import { searchParam } from '@/utils/search_param';

const SysDepart = defineComponent({
  name: 'SysDepart',
  isRouter: true,
  components: { CommonButton, CommonDrawer },
  setup() {
    const tableCont = reactive<TableDataType<DepartType>>({
      page: 1,
      pageSize: 10,
      total: 0,
      loading: false,
      columns: [
        {
          title: '名称',
          dataIndex: 'name',
        },
        {
          title: '父级节点',
          dataIndex: 'parentId',
        },
        {
          title: '排序',
          dataIndex: 'orderNum',
        },
        {
          title: '操作',
          key: 'action',
        },
      ],
      data: [],
    });
    const treeOptions = reactive<{ options: DepartType[] }>({ options: [] });
    const drawerData = reactive<DrawerProps>({
      title: '添加',
      visible: false,
      loading: false,
    });
    const formRef = ref();
    const formData = reactive<DepartType>({
      parentId: '',
      name: '',
      orderNum: 0,
      id: '',
    });
    const rules = {
      parent_id: [
        {
          required: true,
          message: '请选择父级',
        },
      ],
      name: [
        {
          required: true,
          message: '请输入名称',
        },
      ],
      order_num: [
        {
          required: true,
          message: '请输入排序',
        },
      ],
    };
    function getList() {
      tableCont.loading = true;
      http<DepartType>({
        url: `/dept${searchParam({
          page: tableCont.page,
          limit: tableCont.pageSize,
        })}`,
        method: 'GET',
      }).then((res) => {
        const list = res.list?.data.sort(ListObjCompare('order_num')) || [];
        tableCont.loading = false;
        tableCont.data = ListToTree(list);
        tableCont.total = res.list?.total;
        list.forEach((item) => {
          item.title = item.name;
          item.value = item.id;
        });
        treeOptions.options = ListToTree(list);
      });
    }
    onMounted(() => {
      getList();
    });
    const onSubmit = () => {
      formRef.value.validate()
        .then(() => {
          drawerData.loading = true;
          const data = cloneDeep(toRaw(formData));
          // @ts-ignore
          delete data.id;
          let method: Method = 'POST';
          if (formData.id) {
            method = 'PUT';
          }
          http<DepartType>({
            url: formData.id ? `dept/${formData.id}` : 'dept',
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

    function resetFields() {
      if (formRef.value) {
        formRef.value.resetFields();
      }
    }

    function ChangeClick() {
      drawerData.title = '添加';
      drawerData.visible = true;
      resetFields();
      formData.parentId = '0';
    }

    function Editor(record: DepartType) {
      const data = toRaw(record);
      drawerData.title = '修改';
      drawerData.visible = true;
      formData.name = data.name;
      formData.orderNum = data.orderNum;
      formData.parentId = data.parentId;
      formData.id = data.id;
    }
    function Del(record: DepartType) {
      const data = toRaw(record);
      http({ url: `dept/${data.id}`, method: 'delete' }).then(() => {
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

      // table
      Editor,
      Del,

      // form
      resetFields,
      rules,
      formData,
      formRef,
      onSubmit,
    };
  },
});
export default SysDepart;
</script>
<style scoped lang="less">
@import "./depart.less";
</style>
