<template>
  <common-button v-bt-auth:add="{ title: true }" title="添加" icon-name="add" @change="ChangAdd" />
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
    <template #bodyCell="{column, text, record}">
      <template v-if="column.key === 'status'">
        <div :style="text === 0 ? { color: 'red' } : { color: 'green' }">{{ formatStatus(text) }}</div>
      </template>
      <template v-if="column.key === 'depart'">{{ record.deptId }}</template>
      <template v-if="column.key === 'avatar'">
        <a-tag v-if="!record.avatar" color="red">暂无头像</a-tag>
        <img v-else class="avatar" :src="record.avatar" :alt="record.nickName">
      </template>

      <template v-if="column.key === 'action'">
        <a-button
            v-bt-auth:power
            type="primary"
            style="margin-right: 15px"
            @click="Allocate(record)"
        />
        <a-button v-bt-auth:edit type="primary" style="margin-right: 15px" @click="Editor(record)" />
        <a-button v-bt-auth:updatePassword danger style="margin-right: 15px" @click="updatePassword(record)" />
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
    <a-form
        :model="formData"
        ref="formRef"
        :rules="rules"
        :label-col="{ span: 4 }"
        :wrapper-col="{ span: 20 }"
        name="avatar"
    >
      <a-form-item label="头像">
        <image-upload v-model:list="formData.avatar"/>
      </a-form-item>
      <a-form-item label="账号"  name="modelRef">
        <a-input v-model:value="formData.account"></a-input>
      </a-form-item>
      <a-form-item label="昵称"  name="nickName">
        <a-input v-model:value="formData.nickName"></a-input>
      </a-form-item>
      <a-form-item label="部门"  name="deptId">
        <a-tree-select
          v-model:value="formData.deptId"
          style="width: 100%"
          :dropdown-style="{ maxHeight: '400px', overflow: 'auto' }"
          :tree-data="treeOptions.options"
          placeholder="请选择部门"
          tree-default-expand-all
        ></a-tree-select>
      </a-form-item>
      <a-form-item label="状态"  name="status">
        <a-radio-group v-model:value="formData.status" name="radioGroup">
          <a-radio :value="0">失效</a-radio>
          <a-radio :value="1">有效</a-radio>
        </a-radio-group>
      </a-form-item>
      <a-form-item label="邮箱" name="email">
        <a-input v-model:value="formData.email"></a-input>
      </a-form-item>
      <a-form-item label="手机号码" name="phoneNum">
        <a-input v-model:value="formData.phoneNum"></a-input>
      </a-form-item>
      <a-form-item v-if="!editId.id" label="密码" name="passWord">
        <a-input v-model:value="formData.passWord"></a-input>
      </a-form-item>
    </a-form>
  </common-drawer>
  <common-drawer
    title="角色列表"
    ok-text="确定"
    cancel-text="取消"
    :visible="allocationTree.visible"
    :loading="allocationTree.loading"
    @on-close="Close"
    @on-ok="SubmitOk"
  >
    <a-spin :spinning="allocationTree.loading">
      <a-checkbox-group
        v-if="roleList.length > 0"
        v-model:value="selectkeys"
        class="checkoutContainer"
      >
        <div v-for="(item, index) in roleList" :key="index">
          <a-checkbox :value="item.id">{{ item.remark }}</a-checkbox>
        </div>
      </a-checkbox-group>
      <a-empty v-else />
    </a-spin>
  </common-drawer>

  <!-- 修改密码 -->
  <a-modal v-model:visible="updatePasswdVis" title="修改密码" @ok="updatePasswdHandleOk">
    <a-input v-model:value="password" placeholder="请输入新密码" />
  </a-modal>
</template>
<script lang="ts">
import {
  defineComponent, onMounted, reactive, toRaw, ref,
} from 'vue';
import { PlusOutlined, LoadingOutlined } from '@ant-design/icons-vue';
import { message } from 'ant-design-vue';
import { Method } from 'axios';
import { TableDataType, TablePaginType } from '@/types/type';
import { DepartType, RoleType, UserType } from '@/types/sys';
import http from '@/utils/request';
import CommonDrawer, { DrawerProps } from '@/components/Drawer/Drawer.vue';
import CommonButton from '@/components/Button/Button.vue';
import { ListObjCompare, ListToTree } from '@/utils';
import { AllocateType } from '@/views/sys/role.vue';
import { searchParam } from '@/utils/search_param';
import log from '@/utils/log';
import ImageUpload from '@/components/ImageUpload/index.vue';
import BusinessUtils from '@/utils/business';

interface SysUserRole {
  userId: number;
  roleId: number;
  id?: number;
}
interface UserAndRole {
  userId: number;
  data: SysUserRole[],
}

const SysUser = defineComponent({
  name: 'SysUser',
  isRouter: true,
  components: {
    CommonButton,
    CommonDrawer,
    ImageUpload,
  },
  setup() {
    const selectkeys = ref<number[]>([]);
    const roleList = ref<RoleType[]>([]);
    const allocationTree = reactive<AllocateType>({
      visible: false,
      loading: false,
      allocateId: '',
    });
    const formRef = ref();
    const formData = reactive<UserType>({
      account: '',
      nickName: '',
      email: '',
      status: 1,
      avatar: [],
      deptId: 0,
      phoneNum: '',
    });
    const treeOptions = reactive<{ options: DepartType[] }>({ options: [] });
    const editId = reactive<{ id: number | undefined }>({ id: 0 });
    const commonDrawerData = reactive<DrawerProps>({
      title: '添加',
      loading: false,
      visible: false,
    });
    const uploadImageData = reactive({
      fileList: [],
      loading: false,
      imageUrl: '',
      action: 'test',
    });
    const rules = {
      account: [
        {
          required: true,
          message: '请输入账号',
        },
      ],
      nickName: [
        {
          required: true,
          message: '请输入昵称',
        },
      ],
      deptId: [
        {
          required: true,
          message: '请选择部门',
        },
      ],
    };
    const tableData = reactive<TableDataType<UserType>>({
      data: [],
      page: 1,
      pageSize: 10,
      total: 0,
      loading: false,
      columns: [
        {
          title: '账号',
          key: 'account',
          dataIndex: 'account',
        },
        {
          title: '名称',
          key: 'nickName',
          dataIndex: 'nickName',
        },
        {
          title: '头像',
          key: 'avatar',
          dataIndex: 'avatar',
        },
        {
          title: '部门',
          key: 'deptId',
          dataIndex: 'deptId',
        },
        {
          title: '状态',
          key: 'status',
          dataIndex: 'status',
        },
        {
          title: '邮箱',
          key: 'email',
          dataIndex: 'email',
        },
        {
          title: '操作',
          key: 'action',
          dataIndex: 'action',
        },
      ],
    });

    function getList() {
      tableData.loading = true;
      http<UserType>({
        url: `user${searchParam({
          limit: 10,
          page: 1,
        })}`,
        method: 'GET',
      }).then((res) => {
        tableData.loading = false;
        tableData.page = res.list?.page;
        tableData.pageSize = res.list?.pageSize;
        tableData.total = res.list?.total;
        tableData.data = res.list?.data || [];
        log.i(tableData, 'table');
      });
    }
    function getDepartList() {
      http<DepartType>({
        url: `/dept${searchParam({
          page: 1,
          limit: 1000,
        })}`,
        method: 'GET',
      }).then((res) => {
        const list = res.list?.data.sort(ListObjCompare('orderNum'));
        if (list) {
          list.forEach((item) => {
            item.title = item.name;
            item.value = item.id;
            item.key = item.id;
          });
          treeOptions.options = ListToTree(list);
        }
      });
    }
    function getRoleList() {
      http<RoleType>({
        url: `/role${searchParam({
          page: 1,
          limit: 1000,
        })}`,
        method: 'get',
      }).then((res) => {
        roleList.value = res.list?.data || [];
      });
    }
    onMounted(() => {
      getList();
      getDepartList();
      getRoleList();
    });

    function formatStatus(text: number): string {
      let result = '';
      switch (text) {
        case 0:
          result = '失效';
          break;
        case 1:
          result = '有效';
          break;
        default:
          result = '未知';
      }
      return result;
    }

    function Submit() {
      formRef.value.validate().then(() => {
        let url = 'user';
        let method: Method = 'POST';
        const data = toRaw(formData);
        commonDrawerData.loading = true;
        if (editId.id) {
          url = `user/${editId.id}`;
          method = 'PUT';
        }
        if (data.avatar) {
          data.avatar = BusinessUtils.formatUploadImg(data.avatar);
        }
        http({
          url,
          method,
          body: data,
        }).then(() => {
          message.success(`${commonDrawerData.title}成功`);
          commonDrawerData.loading = false;
          commonDrawerData.visible = false;
          getList();
        });
      });
    }
    function ChangAdd() {
      if (formRef.value) {
        formRef.value.resetFields();
      }
      commonDrawerData.visible = true;
      editId.id = 0;
    }

    function getBase64(img, callback) {
      const reader = new FileReader();
      reader.addEventListener('load', () => callback(reader.result));
      reader.readAsDataURL(img);
    }

    function handleChange(info) {
      if (info.file.status === 'uploading') {
        uploadImageData.loading = true;
        return;
      }
      if (info.file.status === 'done') {
        // Get this url from response in real world.
        getBase64(info.file.originFileObj, (imageUrl: string) => {
          uploadImageData.imageUrl = imageUrl;
          uploadImageData.loading = false;
        });
      }
      if (info.file.status === 'error') {
        uploadImageData.loading = false;
      }
    }
    function beforeUpload(file) {
      const isJpgOrPng = file.type === 'image/jpeg' || file.type === 'image/png';
      if (!isJpgOrPng) {
        message.error('请上传jpg格式图片');
      }
      const isLt2M = file.size / 1024 / 1024 < 2;
      if (!isLt2M) {
        message.error('图片上传最大为2M');
      }
      return isJpgOrPng && isLt2M;
    }

    function Editor(record: UserType) {
      if (record.id) {
        editId.id = record.id;
        commonDrawerData.title = '修改';
        commonDrawerData.visible = true;
        formData.account = record.account;
        formData.nickName = record.nickName;
        formData.email = record.email;
        formData.status = record.status;
        formData.avatar = BusinessUtils.formatUploadShow(record.avatar);
        formData.deptId = Number(record.deptId);
        formData.phoneNum = record.phoneNum;
      }
    }
    function Del(record: UserType) {
      http({ url: `user/${record.id}`, method: 'delete' }).then(() => {
        message.success('删除成功');
        getList();
      });
    }
    function Change(pagin: TablePaginType) {
      tableData.page = pagin.current;
      getList();
    }
    function Allocate(record: UserType) {
      allocationTree.loading = true;
      allocationTree.visible = true;
      if (record.id != null) {
        allocationTree.allocateId = String(record.id);
      }
      http<SysUserRole>({
        url: `/user/roleList/${record.id}${searchParam({
          page: 1,
          limit: 1000,
        })}`,
        method: 'get',
      }).then((res) => {
        const list: number[] = [];
        res.list?.data.forEach((item) => {
          if (item.id != null) {
            list.push(Number(item.roleId));
          }
        });
        selectkeys.value = list;
        allocationTree.loading = false;
      });
    }
    function Close() {
      allocationTree.visible = false;
    }
    function SubmitOk() {
      const userId = Number(allocationTree.allocateId);
      const result = selectkeys.value.map((item) => ({
        userId,
        roleId: Number(item),
      }));
      const data: UserAndRole = {
        userId,
        data: result,
      };
      allocationTree.loading = true;
      http<UserType>({
        url: '/user/userRole',
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

    // 修改密码
    const updatePasswdVis = ref(false);
    const updateData = ref();
    const password = ref();
    function updatePasswdHandleOk() {
      if (password.value) {
        http<boolean>({
          url: '/user/updatePasswd',
          method: 'post',
          body: {
            id: updateData.value.id,
            password: password.value,
          },
        }).then((res) => {
          console.log(res.data, 'res');
          if (res.data) {
            message.success('修改成功');
            updatePasswdVis.value = false;
          } else {
            message.error('修改失败');
          }
        });
      } else {
        message.info('请输入新密码');
      }
    }
    function updatePassword(row) {
      updatePasswdVis.value = true;
      updateData.value = row;
      console.log(row, 'row');
    }
    return {
      // data
      tableData,
      commonDrawerData,
      formRef,
      treeOptions,
      uploadImageData,
      allocationTree,
      selectkeys,
      roleList,
      editId,
      updatePasswdVis,
      updatePassword,
      password,

      // methods
      ChangAdd,
      formatStatus,
      Submit,
      handleChange,
      beforeUpload,
      Editor,
      Del,
      Change,
      Allocate,
      Close,
      SubmitOk,
      updatePasswdHandleOk,
      // form
      rules,
      formData,
    };
  },
});

export default SysUser;
</script>
<style scoped lang="less">
.checkoutContainer {
  display: flex;
  flex: 1;
  flex-direction: row;
  flex-wrap: wrap;
}
.checkoutContainer {
  display: flex;
  flex: 1;
  flex-direction: row;
  flex-wrap: wrap;
}

.avatar {
  width:50px;
  height: 50px;
  border-radius: 50%;
}
</style>
