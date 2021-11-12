<template>
  <common-button v-bt-auth:add icon-name="add" @change="ChangAdd" />
  <a-table
    style="margin-top: 15px"
    :columns="tableData.columns"
    :data-source="tableData.data"
    :loading="tableData.loading"
    row-key="id"
    :pagination="{
      total: tableData.total,
    }"
    @change="Change"
  >
    <template #bodyCell="{column,text, record}">
      <template v-if="column.key === 'status'">
        {{ formatStatus(text) }}
      </template>
      <template v-if="column.key === 'action'">
        <a-button
            v-bt-auth:power
            type="primary"
            style="margin-right: 15px"
            @click="PowerAllocation()"
        />

        <a-button v-bt-auth:edit type="primary" style="margin-right: 15px" @click="Editor(record)" />
        <a-button
            v-bt-auth:setting
            type="primary"
            style="margin-right: 15px"
            @click="Setting(record)"
        />
        <a-popconfirm title="确定删除吗?" ok-text="删除" cancel-text="取消" @confirm="Del(record)">
          <a-button v-bt-auth:del danger />
        </a-popconfirm>
      </template>
    </template>
  </a-table>

  <common-drawer
    :title="commonDrawerData.title"
    :visible="commonDrawerData.visible"
    :loading="commonDrawerData.loading"
    ok-text="确定"
    cancel-text="取消"
    @on-ok="Submit()"
    @on-close="commonDrawerData.visible = false"
  >
    <a-form :model="formData" :rules="rules" ref="formRef" :label-col="{ span: 4 }" :wrapper-col="{ span: 20 }">
      <a-form-item label="字典名称" name="name">
        <a-input v-model:value="formData.name"></a-input>
      </a-form-item>
      <a-form-item label="字典编号" name="serialNumber">
        <a-input v-model:value="formData.serialNumber"></a-input>
      </a-form-item>
      <a-form-item label="描述" name="describe">
        <a-input v-model:value="formData.describe"></a-input>
      </a-form-item>
    </a-form>
  </common-drawer>
  <dict-drawer ref="RefDictDrawer" />
</template>
<script lang="ts">
import {
  defineComponent, onMounted, reactive, toRaw, ref,
} from 'vue';
import { useForm } from '@ant-design-vue/use';
import { message } from 'ant-design-vue';
import { Method } from 'axios';
import { TableDataType, TablePaginType } from '@/types/type';
import { SysDict } from '@/types/sys/dict';
import http from '@/utils/request';
import CommonDrawer, { DrawerProps } from '@/components/Drawer/Drawer.vue';
import CommonButton from '@/components/Button/Button.vue';
import DictDrawer from './dict-drawer.vue';
import { searchParam } from '@/utils/search_param';

export interface AllocateType {
  visible: boolean;
  loading: boolean;
  data?: string[];
  allocateId: string;
}
const SysDictView = defineComponent({
  name: 'SysDict',
  isRouter: true,
  components: {
    CommonButton,
    CommonDrawer,
    DictDrawer,
  },
  setup() {
    const formRef = ref();
    const formData = reactive<SysDict>({
      name: '',
      serialNumber: '',
      describe: '',
    });
    const RefDictDrawer = ref();
    const editId = reactive<{ id: number | undefined }>({ id: 0 });
    const commonDrawerData = reactive<DrawerProps>({
      title: '添加',
      loading: false,
      visible: false,
    });
    const rules = {
      name: [
        {
          required: true,
          message: '请输入名称',
        },
      ],
    };
    const tableData = reactive<TableDataType<SysDict>>({
      data: [],
      page: 1,
      pageSize: 10,
      total: 0,
      loading: false,
      columns: [
        {
          title: '字典名称',
          key: 'name',
          dataIndex: 'name',
        },
        {
          title: '字典编号',
          key: 'serial_number',
          dataIndex: 'serial_number',
        },
        {
          title: '描述',
          key: 'describe',
          dataIndex: 'describe',
        },
        {
          title: '操作',
          key: 'action',
        },
      ],
    });

    function getList() {
      tableData.loading = true;
      http<SysDict>({
        url: `dict${searchParam({
          page: tableData.page,
          limit: tableData.pageSize,
        })}`,
        method: 'GET',
      }).then((res) => {
        tableData.loading = false;
        tableData.page = res.list?.page;
        tableData.pageSize = res.list?.pageSize;
        tableData.total = res.list?.total;
        tableData.data = res.list?.data || [];
      });
    }
    onMounted(() => {
      getList();
    });

    function Submit() {
      formRef.value.validate().then(() => {
        let url = 'dict';
        let method: Method = 'POST';
        const body = toRaw(formData);
        commonDrawerData.loading = true;
        if (editId.id) {
          url = `dict/${editId.id}`;
          method = 'PUT';
        }
        http({
          url,
          method,
          body,
        }).then(() => {
          message.success(`${commonDrawerData.title}成功`);
          commonDrawerData.loading = false;
          commonDrawerData.visible = false;
          getList();
        });
      });
    }

    function resetFields() {
      if (formRef.value) {
        formRef.value.resetFields();
      }
    }

    function ChangAdd() {
      resetFields();
      commonDrawerData.visible = true;
      editId.id = 0;
    }

    function Editor(record: SysDict) {
      if (record.id) {
        editId.id = record.id;
        commonDrawerData.title = '修改';
        commonDrawerData.visible = true;
        formData.name = record.name;
        formData.serialNumber = record.serialNumber;
        formData.describe = record.describe;
      }
    }
    function Del(record: SysDict) {
      http({ url: `dict/${record.id}`, method: 'delete' }).then(() => {
        message.success('删除成功');
        getList();
      });
    }
    function Change(pagin: TablePaginType) {
      tableData.page = pagin.current;
      getList();
    }

    const Setting = (record: SysDict) => {
      RefDictDrawer.value.IsShow(record);
    };

    function PowerAllocation() {
      //
    }

    function formatStatus(text: string) {
      return text;
    }

    return {
      // data
      tableData,
      commonDrawerData,
      RefDictDrawer,
      // methods
      ChangAdd,
      Submit,
      Editor,
      Del,
      Change,
      Setting,
      PowerAllocation,
      formatStatus,

      // form
      formData,
      formRef,
      rules,
    };
  },
});

export default SysDictView;
</script>
