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
    <template #status="{ text }">
      {{ formatStatus(text) }}
    </template>
    <template #depart="{ record }">
      {{ record.depart_info.name }}
    </template>
    <template #action="{ record }">
      <a-button
        v-bt-auth:power
        type="primary"
        style="margin-right: 15px"
        @click="PowerAllocation(record)"
      />

      <a-button
        v-bt-auth:edit
        type="primary"
        style="margin-right: 15px"
        @click="Editor(record)"
      />
      <a-popconfirm
        title="确定删除吗?"
        ok-text="删除"
        cancel-text="取消"
        @confirm="Del(record)"
      >
        <a-button v-bt-auth:del type="danger" />
      </a-popconfirm>
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
    <a-form :label-col="{ span: 4 }" :wrapper-col="{ span: 20 }">
      <a-form-item label="角色名称" v-bind="validateInfos.role_name">
        <a-input v-model:value="modelRef.role_name"></a-input>
      </a-form-item>
      <a-form-item label="描述">
        <a-input v-model:value="modelRef.remark"></a-input>
      </a-form-item>
    </a-form>
  </common-drawer>
  <common-tree
    :visible="allocationTree.visible"
    :data="allocationTree.data"
    :loading="allocationTree.loading"
    @on-close="Close"
    @on-submit="SubmitOk"
  />
</template>
<script lang="ts">
import { defineComponent, onMounted, reactive, toRaw } from "vue";
import { useForm } from "@ant-design-vue/use";
import {
  CommonTreeSelectKeys,
  TableDataType,
  TablePaginType,
} from "@/types/type";
import { MenuType, RoleType } from "@/types/sys";
import { http } from "@/utils/request";
import CommonDrawer, { DrawerProps } from "@/components/Drawer/Drawer.vue";
import CommonButton from "@/components/Button/Button.vue";
import CommonTree from "@/views/common/tree.vue";
import { message } from "ant-design-vue";
import { Method } from "axios";

export interface AllocateType {
  visible: boolean;
  loading: boolean;
  data?: string[];
  allocateId: string;
}
const SysRole = defineComponent({
  name: "SysRole",
  components: {
    CommonButton,
    CommonDrawer,
    CommonTree,
  },
  setup() {
    const modelRef = reactive<RoleType>({
      remark: "",
      role_name: "",
    });
    const editId = reactive<{ id: string | undefined }>({ id: "" });
    const commonDrawerData = reactive<DrawerProps>({
      title: "添加",
      loading: false,
      visible: false,
    });
    const allocationTree = reactive<AllocateType>({
      visible: false,
      loading: false,
      data: [],
      allocateId: "",
    });
    const rulesRef = reactive({
      role_name: [
        {
          required: true,
          message: "请输入角色名称",
        },
      ],
    });
    const tableData = reactive<TableDataType<RoleType>>({
      data: [],
      page: 1,
      page_size: 10,
      total: 0,
      loading: false,
      columns: [
        {
          title: "描述",
          key: "remark",
          dataIndex: "remark",
        },
        {
          title: "角色名称",
          key: "role_name",
          dataIndex: "role_name",
        },
        {
          title: "操作",
          key: "action",
          slots: {
            customRender: "action",
          },
        },
      ],
    });

    function getList() {
      tableData.loading = true;
      http<RoleType>({
        url: "role",
        method: "GET",
        params: {
          page: tableData.page,
          page_size: tableData.page_size,
        },
      }).then((res) => {
        tableData.loading = false;
        tableData.page = res.page;
        tableData.page_size = res.page_size;
        tableData.total = res.total;
        tableData.data = res.list;
      });
    }
    onMounted(() => {
      getList();
    });

    const { resetFields, validate, validateInfos } = useForm(
      modelRef,
      rulesRef
    );
    function Submit() {
      validate().then(() => {
        let url = "role";
        let method: Method = "POST";
        const body = toRaw(modelRef);
        commonDrawerData.loading = true;
        if (editId.id) {
          url = `role/${editId.id}`;
          method = "PUT";
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
    function ChangAdd() {
      resetFields();
      commonDrawerData.visible = true;
      editId.id = "";
    }

    function Editor(record: RoleType) {
      if (record.id) {
        editId.id = record.id;
        commonDrawerData.title = "修改";
        commonDrawerData.visible = true;
        modelRef.remark = record.remark;
        modelRef.role_name = record.role_name;
      }
    }
    function Del(record: RoleType) {
      http({ url: "role/" + record.id, method: "delete" }).then(() => {
        message.success("删除成功");
        getList();
      });
    }
    function Change(pagin: TablePaginType) {
      tableData.page = pagin.current;
      getList();
    }
    function PowerAllocation(record: RoleType) {
      allocationTree.loading = true;
      allocationTree.visible = true;
      if (record.id != null) {
        allocationTree.allocateId = record.id;
      }
      http<MenuType>({
        url: "/role/permissions/" + record.id,
        method: "get",
      }).then((res) => {
        const list: string[] = [];
        res.list.forEach((item) => {
          return list.push(item.id || "");
        });
        allocationTree.data = list;
        allocationTree.loading = false;
      });
    }
    function Close() {
      allocationTree.visible = false;
    }
    function SubmitOk(val: CommonTreeSelectKeys) {
      const data = {
        role_id: allocationTree.allocateId,
        menu_id: val.checked.join(","),
      };
      allocationTree.loading = true;
      http<MenuType>({
        url: "/role/permissions",
        method: "post",
        body: data,
      })
        .then(() => {
          message.success("更新成功");
          allocationTree.loading = false;
          allocationTree.visible = false;
        })
        .catch(() => {
          allocationTree.loading = false;
        });
    }
    return {
      //data
      tableData,
      commonDrawerData,
      modelRef,
      allocationTree,

      // methods
      ChangAdd,
      Submit,
      Editor,
      Del,
      Change,
      PowerAllocation,
      Close,
      SubmitOk,

      // form
      validateInfos,
    };
  },
});

export default SysRole;
</script>
