<template>
  <common-drawer
    title="字典配置"
    cancel-text="取消"
    :visible="visible"
    @on-close="Close()"
  >
    <a-button type="primary" @click="Add()">添加</a-button>
    <a-table
      :columns="tableData.columns"
      :pagination="{
        total: tableData.total,
      }"
      :data-source="tableData.data"
      :loading="tableData.loading"
    >
      <template #bodyCell="{column, record}">
        <template v-if="column.dataIndex === 'action'" >
          <a-button type="primary" class="button-right" @click="Edit(record)">
            编辑
          </a-button>
          <a-popconfirm
              title="确定删除吗?"
              ok-text="删除"
              cancel-text="取消"
              @confirm="Del(record)"
          >
            <a-button type="primary" class="button-right" danger> 删除 </a-button>
          </a-popconfirm>
        </template>
      </template>

    </a-table>
    <a-modal
      v-model:visible="modalForm.visible"
      :title="modalForm.title"
      :confirm-loading="modalForm.loading"
      @ok="handleOk"
    >
      <a-form :model="formData" :rules="rules" ref="formRef" :label-col="labelCol" :wrapper-col="wrapperCol">
        <a-form-item label="字典名称" name="label">
          <a-input v-model:value="formData.label" />
        </a-form-item>
        <a-form-item label="字典值" name="value">
          <a-input v-model:value="formData.value" />
        </a-form-item>
        <a-form-item label="描述" name="describe">
          <a-textarea v-model:value="formData.describe" allow-clear />
        </a-form-item>
      </a-form>
    </a-modal>
  </common-drawer>
</template>
<script lang="ts">
import {
  defineComponent, ref, reactive, toRaw,
} from 'vue';
import { message } from 'ant-design-vue';
import { Method } from 'axios';
import CommonDrawer from '@/components/Drawer/Drawer.vue';
import { SysDict, SysDictItem } from '@/types/sys/dict';
import { TableDataType } from '@/types/type';
import http from '@/utils/request';
import { searchParam } from '@/utils/search_param';

const DictDrawer = defineComponent({
  components: {
    CommonDrawer,
  },
  setup() {
    const editParentId = ref(-1);
    const editId = ref(-1);
    const modalForm = reactive({
      title: '添加',
      visible: false,
      loading: false,
    });
    const formRef = ref();
    const formData = reactive<SysDictItem>({
      value: '',
      label: '',
      describe: '',
      dictId: -1,
    });
    const rules = {
      value: [
        {
          required: true,
          message: '请输入字典项',
        },
      ],
      label: [
        {
          required: true,
          message: '请输入名称',
        },
      ],
    };
    const tableData = reactive<TableDataType<SysDictItem>>({
      data: [],
      page: 1,
      pageSize: 10,
      total: 0,
      loading: false,
      columns: [
        {
          title: '名称',
          dataIndex: 'label',
        },
        {
          title: '数据值',
          dataIndex: 'value',
        },
        {
          title: '描述',
          dataIndex: 'describe',
        },
        {
          title: '操作',
          dataIndex: 'action',
        },
      ],
    });
    const visible = ref(false);
    const GetList = () => {
      tableData.loading = true;
      const search = `/dict-item${searchParam({
        page: tableData.page,
        limit: tableData.pageSize,
        search: {
          $and: [{
            dictId: editParentId.value,
          }],
        },
      })}`;
      http<SysDictItem>({
        url: search,
        method: 'GET',
      }).then((res) => {
        tableData.data = res.list?.data || [];
        tableData.total = res.list?.total;
        tableData.loading = false;
      });
    };
    const Close = () => {
      visible.value = false;
    };
    const IsShow = (record: SysDict) => {
      editParentId.value = record.id || -1;
      visible.value = true;
      GetList();
    };
    const Edit = (record: SysDictItem) => {
      const editData = toRaw(record);
      delete editData.dictId;
      formData.value = editData.value;
      formData.label = editData.label;
      formData.describe = editData.describe;
      console.log(
        '🚀 ~ file: dict-drawer.vue ~ line 139 ~ Edit ~ record',
        record,
      );
      editId.value = editData.id || -1;
      modalForm.visible = true;
    };

    function resetFields() {
      if (formRef.value) {
        formRef.value.resetFields();
      }
    }

    const Add = () => {
      modalForm.visible = true;
      resetFields();
    };
    const handleOk = () => {
      modalForm.loading = true;
      formRef.value.validate()
        .then(() => {
          const data = toRaw(formData);
          data.dictId = editParentId.value;
          let method: Method = 'POST';
          let url = 'dict-item';
          if (editId.value !== -1 && editId.value) {
            method = 'PUT';
            url = `${url}/${editId.value}`;
          }
          http({
            url,
            method,
            body: data,
          }).then((res) => {
            console.log(res);
            message.success(`${modalForm.title}成功`);
            modalForm.loading = false;
            modalForm.visible = false;
            GetList();
          });
        })
        .catch((err) => {
          console.error(err);
          modalForm.loading = false;
        });
    };
    const Del = (record: SysDictItem) => {
      http({
        url: `dict-item/${record.id}`,
        method: 'DELETE',
      }).then((res) => {
        console.log(res);
        message.success('删除成功');
        GetList();
      });
    };
    return {
      // data
      visible,
      tableData,
      modalForm,
      labelCol: { span: 4 },
      wrapperCol: { span: 14 },

      // methods
      Close,
      IsShow,
      Edit,
      Add,
      Del,
      resetFields,
      handleOk,

      // form
      formData,
      formRef,
      rules,
    };
  },
});
export default DictDrawer;
</script>
