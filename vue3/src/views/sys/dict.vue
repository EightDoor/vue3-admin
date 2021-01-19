<template>
  <common-button v-bt-auth:add icon-name="add" @change="ChangAdd" />
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
      {{ formatStatus(text) }}
    </template>
    <template #depart="{ record }">
      {{ record.depart_info.name }}
    </template>
    <template #action="{ record }">
      <a-button
        type="primary"
        style="margin-right: 15px"
        @click="PowerAllocation(record)"
        v-bt-auth:power
      />

      <a-button
        type="primary"
        style="margin-right: 15px"
        v-bt-auth:edit
        @click="Editor(record)"
      />
      <a-button
        type="primary"
        style="margin-right: 15px"
        v-bt-auth:setting
        @click="Setting(record)"
      />
      <a-popconfirm
        title="确定删除吗?"
        ok-text="删除"
        cancel-text="取消"
        @confirm="Del(record)"
      >
        <a-button type="danger" v-bt-auth:del />
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
      <a-form-item label="字典名称" v-bind="validateInfos.name">
        <a-input v-model:value="modelRef.name"></a-input>
      </a-form-item>
      <a-form-item label="字典编号">
        <a-input v-model:value="modelRef.serial_number"></a-input>
      </a-form-item>
      <a-form-item label="描述">
        <a-input v-model:value="modelRef.describe"></a-input>
      </a-form-item>
    </a-form>
  </common-drawer>
  <dict-drawer ref="RefDictDrawer" />
</template>
<script lang="ts">
import { defineComponent, onMounted, reactive, toRaw, ref } from 'vue'
import { useForm } from '@ant-design-vue/use'
import { TableDataType, TablePaginType } from '@/types/type'
import { SysDict } from '@/types/sys/dict'
import { http } from '@/utils/request'
import CommonDrawer, { DrawerProps } from '@/components/Drawer/Drawer.vue'
import CommonButton from '@/components/Button/Button.vue'
import DictDrawer from './dict-drawer.vue'
import { message } from 'ant-design-vue'
import { Method } from 'axios'

export interface AllocateType {
  visible: boolean
  loading: boolean
  data?: string[]
  allocateId: string
}
const SysDictView = defineComponent({
  name: 'sys-dict',
  components: {
    CommonButton,
    CommonDrawer,
    DictDrawer,
  },
  setup() {
    const modelRef = reactive<SysDict>({
      name: '',
      serial_number: '',
      describe: '',
    })
    const RefDictDrawer = ref()
    const editId = reactive<{ id: string | undefined }>({ id: '' })
    const commonDrawerData = reactive<DrawerProps>({
      title: '添加',
      loading: false,
      visible: false,
    })
    const rulesRef = reactive({
      name: [
        {
          required: true,
          message: '请输入名称',
        },
      ],
    })
    const tableData = reactive<TableDataType<SysDict>>({
      data: [],
      page: 1,
      page_size: 10,
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
          slots: {
            customRender: 'action',
          },
        },
      ],
    })

    function getList() {
      tableData.loading = true
      http<SysDict>({
        url: 'dict',
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
    onMounted(() => {
      getList()
    })

    const { resetFields, validate, validateInfos } = useForm(modelRef, rulesRef)
    function Submit() {
      validate().then(() => {
        let url = 'dict'
        let method: Method = 'POST'
        const body = toRaw(modelRef)
        commonDrawerData.loading = true
        if (editId.id) {
          url = `dict/${editId.id}`
          method = 'PUT'
        }
        http({
          url,
          method,
          body,
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

    function Editor(record: SysDict) {
      if (record.id) {
        editId.id = record.id
        commonDrawerData.title = '修改'
        commonDrawerData.visible = true
        modelRef.name = record.name
        modelRef.serial_number = record.serial_number
        modelRef.describe = record.describe
      }
    }
    function Del(record: SysDict) {
      http({ url: 'dict/' + record.id, method: 'delete' }).then((res) => {
        message.success('删除成功')
        getList()
      })
    }
    function Change(pagin: TablePaginType) {
      tableData.page = pagin.current
      getList()
    }

    const Setting = (record: SysDict) => {
      RefDictDrawer.value.IsShow(record)
    }

    return {
      //data
      tableData,
      commonDrawerData,
      modelRef,
      RefDictDrawer,
      // methods
      ChangAdd,
      Submit,
      Editor,
      Del,
      Change,
      Setting,

      // form
      validateInfos,
    }
  },
})

export default SysDictView
</script>
