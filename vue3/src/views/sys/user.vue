<template>
  <common-button title="添加" icon-name="add" @change="ChangAdd" />
  <a-table
    style="margin-top: 15px"
    :columns="tableData.columns"
    :data-source="tableData.data"
    :loading="tableData.loading"
    rowKey="id"
    :pagination="{
      total: tableData.total,
    }"
    @change="Change"
  >
    <template #status="{ text }">
      <div :style="text === 0 ? { color: 'red' } : { color: 'green' }">
        {{ formatStatus(text) }}
      </div>
    </template>
    <template #depart="{ record }">
      {{ record.depart_info.name }}
    </template>
    <template #action="{ record }">
      <a-button
        type="primary"
        style="margin-right: 15px"
        @click="Allocate(record)"
        >角色分配</a-button
      >
      <a-button
        type="primary"
        style="margin-right: 15px"
        @click="Editor(record)"
        >编辑</a-button
      >
      <a-popconfirm
        title="确定删除吗?"
        ok-text="删除"
        cancel-text="取消"
        @confirm="Del(record)"
      >
        <a-button type="danger">删除</a-button>
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
      <a-form-item label="头像">
        <a-upload
          v-model:fileList="modelRef.avatar"
          name="avatar"
          list-type="picture-card"
          class="avatar-uploader"
          :show-upload-list="false"
          :action="uploadImageData.action"
          :before-upload="beforeUpload"
          @change="handleChange"
        >
          <img
            v-if="uploadImageData.imageUrl"
            :src="uploadImageData.imageUrl"
            alt="avatar"
          />
          <div v-else>
            <loading-outlined v-if="uploadImageData.loading" />
            <plus-outlined v-else />
            <div class="ant-upload-text">上传</div>
          </div>
        </a-upload>
      </a-form-item>
      <a-form-item label="账号" v-bind="validateInfos.account">
        <a-input v-model:value="modelRef.account"></a-input>
      </a-form-item>
      <a-form-item label="昵称" v-bind="validateInfos.nick_name">
        <a-input v-model:value="modelRef.nick_name"></a-input>
      </a-form-item>
      <a-form-item label="部门" v-bind="validateInfos.dept_id">
        <a-tree-select
          v-model:value="modelRef.dept_id"
          style="width: 100%"
          :dropdown-style="{ maxHeight: '400px', overflow: 'auto' }"
          :tree-data="treeOptions.options"
          placeholder="请选择部门"
          tree-default-expand-all
        >
        </a-tree-select>
      </a-form-item>
      <a-form-item label="状态" v-bind="validateInfos.status">
        <a-radio-group name="radioGroup" v-model:value="modelRef.status">
          <a-radio :value="0"> 失效 </a-radio>
          <a-radio :value="1"> 有效 </a-radio>
        </a-radio-group>
      </a-form-item>
      <a-form-item label="邮箱">
        <a-input v-model:value="modelRef.email"></a-input>
      </a-form-item>
      <a-form-item label="手机号码">
        <a-input v-model:value="modelRef.phone_num"></a-input>
      </a-form-item>
      <a-form-item label="密码">
        <a-input v-model:value="modelRef.pass_word"></a-input>
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
        class="checkoutContainer"
        v-model:value="selectkeys"
        v-if="roleList.length > 0"
      >
        <div v-for="(item, index) in roleList" :key="index">
          <a-checkbox :value="item.id"> {{ item.remark }} </a-checkbox>
        </div>
      </a-checkbox-group>
      <a-empty v-else />
    </a-spin>
  </common-drawer>
</template>
<script lang="ts">
import { defineComponent, onMounted, reactive, toRaw, ref } from 'vue'
import { useForm } from '@ant-design-vue/use'
import { TableDataType, TablePaginType } from '@/types/type'
import { DepartType, RoleType, UserType } from '@/types/sys'
import { http } from '@/utils/request'
import CommonDrawer, { DrawerProps } from '@/components/Drawer/Drawer.vue'
import CommonButton from '@/components/Button/Button.vue'
import { ListObjCompare, ListToTree } from '@/utils'
import { PlusOutlined, LoadingOutlined } from '@ant-design/icons-vue'
import { message } from 'ant-design-vue'
import { Method } from 'axios'
import { AllocateType } from '@/views/sys/role.vue'

const SysUser = defineComponent({
  name: 'sys-user',
  components: {
    CommonButton,
    CommonDrawer,
    PlusOutlined,
    LoadingOutlined,
  },
  setup() {
    const selectkeys = ref<string[]>([])
    const roleList = ref<RoleType[]>([])
    const allocationTree = reactive<AllocateType>({
      visible: false,
      loading: false,
      allocateId: '',
    })
    const modelRef = reactive<UserType>({
      account: '',
      pass_word: '',
      nick_name: '',
      email: '',
      status: 1,
      avatar: '',
      dept_id: '',
      phone_num: '',
    })
    const treeOptions = reactive<{ options: DepartType[] }>({ options: [] })
    const editId = reactive<{ id: string | undefined }>({ id: '' })
    const commonDrawerData = reactive<DrawerProps>({
      title: '添加',
      loading: false,
      visible: false,
    })
    const uploadImageData = reactive({
      fileList: [],
      loading: false,
      imageUrl: '',
      action: 'test',
    })
    const rulesRef = reactive({
      account: [
        {
          required: true,
          message: '请输入账号',
        },
      ],
      nick_name: [
        {
          required: true,
          message: '请输入昵称',
        },
      ],
      dept_id: [
        {
          required: true,
          messgae: '请选择部门',
        },
      ],
    })
    const tableData = reactive<TableDataType<UserType>>({
      data: [],
      page: 1,
      page_size: 10,
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
          key: 'nick_name',
          dataIndex: 'nick_name',
        },
        {
          title: '头像',
          key: 'avatar',
          dataIndex: 'avatar',
          slots: {
            customRender: 'avatar',
          },
        },
        {
          title: '部门',
          key: 'dept_id',
          dataIndex: 'dept_id',
          slots: {
            customRender: 'depart',
          },
        },
        {
          title: '状态',
          key: 'status',
          dataIndex: 'status',
          slots: {
            customRender: 'status',
          },
        },
        {
          title: '邮箱',
          key: 'email',
          dataIndex: 'email',
        },
        {
          title: '操作',
          key: 'action',
          slots: {
            customRender: 'action',
          },
        },
      ],
    })

    function getList() {
      tableData.loading = true
      http<UserType>({
        url: 'user',
        method: 'GET',
        params: {
          page: tableData.page,
          page_size: tableData.page_size,
        },
      }).then((res) => {
        tableData.loading = false
        tableData.page = res.page
        tableData.page_size = res.page_size
        tableData.total = res.total
        tableData.data = res.list
      })
    }
    function getDepartList() {
      http<DepartType>({
        url: '/depart',
        method: 'GET',
        params: {
          page: 1,
          page_size: 1000,
        },
      }).then((res) => {
        const list = res.list.sort(ListObjCompare('order_num'))
        list.map((item) => {
          item.title = item.name
          item.value = item.id
          item.key = item.id
        })
        treeOptions.options = ListToTree(list)
      })
    }
    function getRoleList() {
      http<RoleType>({
        url: '/role',
        method: 'get',
        params: {
          page: 1,
          page_size: 1000,
        },
      }).then((res) => {
        roleList.value = res.list
      })
    }
    onMounted(() => {
      getList()
      getDepartList()
      getRoleList()
    })

    function formatStatus(text: number): string {
      let result = ''
      switch (text) {
        case 0:
          result = '失效'
          break
        case 1:
          result = '有效'
          break
        default:
          result = '未知'
      }
      return result
    }

    const { resetFields, validate, validateInfos } = useForm(modelRef, rulesRef)
    function Submit() {
      validate().then(() => {
        let url = 'user'
        let method: Method = 'POST'
        const data = toRaw(modelRef)
        commonDrawerData.loading = true
        if (editId.id) {
          url = `user/${editId.id}`
          method = 'PUT'
        }
        http({
          url,
          method,
          body: data,
        }).then(() => {
          message.success(`${commonDrawerData.title}成功`)
          commonDrawerData.loading = false
          commonDrawerData.visible = false
          getList()
        })
      })
    }
    function ChangAdd() {
      resetFields()
      commonDrawerData.visible = true
      editId.id = ''
    }

    function getBase64(img: any, callback: Function) {
      const reader = new FileReader()
      reader.addEventListener('load', () => callback(reader.result))
      reader.readAsDataURL(img)
    }

    function handleChange(info: any) {
      if (info.file.status === 'uploading') {
        uploadImageData.loading = true
        return
      }
      if (info.file.status === 'done') {
        // Get this url from response in real world.
        getBase64(info.file.originFileObj, (imageUrl: string) => {
          uploadImageData.imageUrl = imageUrl
          uploadImageData.loading = false
        })
      }
      if (info.file.status === 'error') {
        uploadImageData.loading = false
      }
    }
    function beforeUpload(file: any) {
      const isJpgOrPng = file.type === 'image/jpeg' || file.type === 'image/png'
      if (!isJpgOrPng) {
        message.error('请上传jpg格式图片')
      }
      const isLt2M = file.size / 1024 / 1024 < 2
      if (!isLt2M) {
        message.error('图片上传最大为2M')
      }
      return isJpgOrPng && isLt2M
    }

    function Editor(record: UserType) {
      if (record.id) {
        editId.id = record.id
        commonDrawerData.title = '修改'
        commonDrawerData.visible = true
        modelRef.account = record.account
        modelRef.nick_name = record.nick_name
        modelRef.email = record.email
        modelRef.status = record.status
        modelRef.avatar = record.avatar
        modelRef.dept_id = record.dept_id
        modelRef.phone_num = record.phone_num
      }
    }
    function Del(record: UserType) {
      http({ url: 'user/' + record.id, method: 'delete' }).then((res) => {
        message.success('删除成功')
        getList()
      })
    }
    function Change(pagin: TablePaginType) {
      tableData.page = pagin.current
      getList()
    }
    function Allocate(record: UserType) {
      allocationTree.loading = true
      allocationTree.visible = true
      if (record.id != null) {
        allocationTree.allocateId = record.id
      }
      http<RoleType>({
        url: '/user/permissions/' + record.id,
        method: 'get',
      }).then((res) => {
        const list: string[] = []
        res.list.forEach((item) => {
          if (item.id != null) {
            list.push(item.id)
          }
        })
        selectkeys.value = list
        allocationTree.loading = false
      })
    }
    function Close() {
      allocationTree.visible = false
    }
    function SubmitOk() {
      const data = {
        user_id: allocationTree.allocateId,
        role_id: selectkeys.value.join(','),
      }
      allocationTree.loading = true
      http<UserType>({
        url: '/user/permissions',
        method: 'post',
        body: data,
      })
        .then(() => {
          message.success('更新成功')
          allocationTree.loading = false
          allocationTree.visible = false
        })
        .catch(() => {
          allocationTree.loading = false
        })
    }
    return {
      //data
      tableData,
      commonDrawerData,
      modelRef,
      treeOptions,
      uploadImageData,
      allocationTree,
      selectkeys,
      roleList,

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
      // form
      validateInfos,
    }
  },
})

export default SysUser
</script>
<style scoped lang="less">
.checkoutContainer {
  display: flex;
  flex: 1;
  flex-direction: row;
  flex-wrap: wrap;
}
</style>
