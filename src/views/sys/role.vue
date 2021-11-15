<template>
  <div>
    <common-button v-bt-auth:add icon-name="add" @change="ChangAdd" />
    <a-table
      style="margin-top: 15px"
      :columns="tableData.columns"
      :dataSource="tableData.data"
      :loading="tableData.loading"
      row-key="id"
      :pagination="{
        total: tableData.total,
      }"
      @change="Change"
    >
      <template #bodyCell="{column, record}">
        <template v-if="column.key === 'action'">
          <a-button
              v-bt-auth:power
              type="primary"
              style="margin-right: 15px"
              @click="PowerAllocation(record)"
          />

          <a-button v-bt-auth:edit type="primary" style="margin-right: 15px" @click="Editor(record)" />
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
        <a-form-item label="角色名称" name="roleName">
          <a-input v-model:value="formData.roleName"></a-input>
        </a-form-item>
        <a-form-item label="描述">
          <a-input v-model:value="formData.remark"></a-input>
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
  </div>
</template>
<script lang="ts">
import {
  defineComponent, onMounted, reactive, ref, toRaw,
} from 'vue';
import { message } from 'ant-design-vue';
import { Method } from 'axios';
import type { UnwrapRef } from 'vue';
import {
  CommonTreeSelectKeys,
  TableDataType,
  TablePaginType,
} from '@/types/type';
import { MenuType, RoleType } from '@/types/sys';
import http from '@/utils/request';
import CommonDrawer, { DrawerProps } from '@/components/Drawer/Drawer.vue';
import CommonButton from '@/components/Button/Button.vue';
import CommonTree from '@/views/common/tree.vue';
import { searchParam } from '@/utils/search_param';

export interface AllocateType {
  visible: boolean;
  loading: boolean;
  data?: number[];
  allocateId: string;
}
const SysRole = defineComponent({
  name: 'SysRole',
  isRouter: true,
  components: {
    CommonButton,
    CommonDrawer,
    CommonTree,
  },
  setup() {
    const formData: UnwrapRef<RoleType> = reactive({
      remark: '',
      roleName: '',
    });
    const formRef = ref();
    const editId = reactive<{ id: number | undefined }>({ id: 0 });
    const commonDrawerData = reactive<DrawerProps>({
      title: '添加',
      loading: false,
      visible: false,
    });
    const allocationTree = reactive<AllocateType>({
      visible: false,
      loading: false,
      data: [],
      allocateId: '',
    });
    const rules = {
      roleName: [
        {
          required: true,
          message: '请输入角色名称',
        },
      ],
    };
    const tableData = reactive<TableDataType<RoleType>>({
      data: [],
      page: 1,
      pageSize: 10,
      total: 0,
      loading: false,
      columns: [
        {
          title: '角色名称',
          key: 'roleName',
          dataIndex: 'roleName',
        },
        {
          title: '描述',
          key: 'remark',
          dataIndex: 'remark',
        },
        {
          title: '操作',
          key: 'action',
        },
      ],
    });

    function getList() {
      tableData.loading = true;
      http<RoleType>({
        url: `role${searchParam({
          page: tableData.page,
          limit: tableData.pageSize,
        })}`,
        method: 'GET',
      }).then((res) => {
        tableData.loading = false;
        tableData.page = res.list?.page;
        tableData.total = res.list?.total;
        tableData.data = res.list?.data || [];
        console.log(tableData, 'data');
      });
    }
    onMounted(() => {
      getList();
    });

    function Submit() {
      formRef.value.validate().then(() => {
        let url = 'role';
        let method: Method = 'POST';
        const body = toRaw(formData);
        commonDrawerData.loading = true;
        if (editId.id) {
          url = `role/${editId.id}`;
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

    function Editor(record: RoleType) {
      if (record.id) {
        editId.id = record.id;
        commonDrawerData.title = '修改';
        commonDrawerData.visible = true;
        formData.remark = record.remark;
        formData.roleName = record.roleName;
      }
    }
    function Del(record: RoleType) {
      http({ url: `role/${record.id}`, method: 'delete' }).then(() => {
        message.success('删除成功');
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
        allocationTree.allocateId = String(record.id);
      }
      http<MenuType>({
        url: `/role/menus/${record.id}`,
        method: 'get',
      }).then((res) => {
        if (res && res.data) {
          const menus = res.data.menuId?.split(',').map((item) => Number(item));
          if (menus) {
            allocationTree.data = menus;
          }
        } else {
          allocationTree.data = [];
        }
        allocationTree.loading = false;
      });
    }
    function Close() {
      allocationTree.visible = false;
    }
    function SubmitOk(val: CommonTreeSelectKeys) {
      const data = {
        roleId: allocationTree.allocateId,
        menuId: val.checked.join(','),
      };
      allocationTree.loading = true;
      http<MenuType>({
        url: '/role/relationAndMenu',
        method: 'post',
        body: data,
      })
        .then(() => {
          message.success('更新成功');
          allocationTree.loading = false;
          allocationTree.visible = false;
        })
        .catch(() => {
          allocationTree.loading = false;
        });
    }
    return {
      // data
      tableData,
      commonDrawerData,
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
      formData,
      formRef,
      rules,
    };
  },
});

export default SysRole;
</script>
