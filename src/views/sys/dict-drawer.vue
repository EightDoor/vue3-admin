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
      <template #action="{ record }">
        <a-button type="primary" class="button-right" @click="Edit(record)">
          编辑
        </a-button>
        <a-popconfirm
          title="确定删除吗?"
          ok-text="删除"
          cancel-text="取消"
          @confirm="Del(record)"
        >
          <a-button type="primary" class="button-right"> 删除 </a-button>
        </a-popconfirm>
      </template>
    </a-table>
    <a-modal
      v-model:visible="modalForm.visible"
      :title="modalForm.title"
      :confirm-loading="modalForm.loading"
      @ok="handleOk"
    >
      <a-form :label-col="labelCol" :wrapper-col="wrapperCol">
        <a-form-item label="字典名称" v-bind="validateInfos.label">
          <a-input v-model:value="modalRef.label" />
        </a-form-item>
        <a-form-item label="字典值" v-bind="validateInfos.value">
          <a-input v-model:value="modalRef.value" />
        </a-form-item>
        <a-form-item label="描述">
          <a-textarea v-model:value="modalRef.describe" allow-clear />
        </a-form-item>
      </a-form>
    </a-modal>
  </common-drawer>
</template>
<script lang="ts">
import CommonDrawer from "@/components/Drawer/Drawer.vue";
import { SysDict, SysDictItem } from "@/types/sys/dict";
import { TableDataType } from "@/types/type";
import { http } from "@/utils/request";
import { defineComponent, ref, reactive, toRaw } from "vue";
import { useForm } from "@ant-design-vue/use";
import { message } from "ant-design-vue";
import { Method } from "axios";

const DictDrawer = defineComponent({
  components: {
    CommonDrawer,
  },
  setup() {
    const editParentId = ref("");
    const editId = ref("");
    const modalForm = reactive({
      title: "添加",
      visible: false,
      loading: false,
    });
    const modalRef = reactive<SysDictItem>({
      value: "",
      label: "",
      describe: "",
      dict_id: "",
    });
    const ruleRef = reactive({
      value: [
        {
          required: true,
          message: "请输入字典项",
        },
      ],
      label: [
        {
          required: true,
          message: "请输入名称",
        },
      ],
    });
    const { resetFields, validate, validateInfos } = useForm(modalRef, ruleRef);
    const tableData = reactive<TableDataType<SysDictItem>>({
      data: [],
      page: 1,
      page_size: 10,
      total: 0,
      loading: false,
      columns: [
        {
          title: "名称",
          dataIndex: "label",
        },
        {
          title: "数据值",
          dataIndex: "value",
        },
        {
          title: "描述",
          dataIndex: "describe",
        },
        {
          title: "操作",
          dataIndex: "action",
          slots: { customRender: "action" },
        },
      ],
    });
    const visible = ref(false);
    const GetList = () => {
      tableData.loading = true;
      http<SysDictItem>({
        url: "/dict-item/" + editParentId.value,
        method: "GET",
      }).then((res) => {
        tableData.data = res.list;
        tableData.total = res.total;
        console.log("🚀 ~ file: dict-drawer.vue ~ line 42 ~ http ~ res", res);
        tableData.loading = false;
      });
    };
    const Close = () => {
      visible.value = false;
    };
    const IsShow = (record: SysDict) => {
      console.log(
        "🚀 ~ file: dict-drawer.vue ~ line 26 ~ IsShow ~ record",
        record
      );
      editParentId.value = toRaw(record.id) || "";
      visible.value = true;
      GetList();
    };
    const Edit = (record: SysDictItem) => {
      const editData = toRaw(record);
      delete editData.dict_id;
      modalRef.value = editData.value;
      modalRef.label = editData.label;
      modalRef.describe = editData.describe;
      console.log(
        "🚀 ~ file: dict-drawer.vue ~ line 139 ~ Edit ~ record",
        record
      );
      editId.value = editData.id || "";
      modalForm.visible = true;
    };
    const Add = () => {
      modalForm.visible = true;
      resetFields();
    };
    const handleOk = () => {
      modalForm.loading = true;
      validate()
        .then(() => {
          const data = toRaw(modalRef);
          data.dict_id = editParentId.value;
          let method: Method = "POST";
          let url = "dict-item";
          if (editId.value) {
            method = "PUT";
            url = `${url}/${editId.value}`;
          }
          http({
            url,
            method,
            body: data,
          }).then((res) => {
            console.log(res);
            message.success(modalForm.title + "成功");
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
        url: "dict-item/" + record.id,
        method: "DELETE",
      }).then((res) => {
        console.log(res);
        message.success("删除成功");
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
      modalRef,

      // methods
      Close,
      IsShow,
      Edit,
      Add,
      Del,
      resetFields,
      validateInfos,
      handleOk,
    };
  },
});
export default DictDrawer;
</script>
