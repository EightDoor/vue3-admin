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
      :title="modalForm.title"
      v-model:visible="modalForm.visible"
      :confirm-loading="modalForm.loading"
      @ok="handleOk"
    >
      <a-form :label-col="labelCol" :wrapper-col="wrapperCol">
        <a-form-item label="字典名称" v-bind="validateInfos.label">
          <a-input v-model:value="modalForm.label" />
        </a-form-item>
        <a-form-item label="字典值" v-bind="validateInfos.value">
          <a-input v-model:value="modalForm.value" />
        </a-form-item>
        <a-form-item label="描述">
          <a-textarea allow-clear v-model:value="modalForm.describe" />
        </a-form-item>
      </a-form>
    </a-modal>
  </common-drawer>
</template>
<script lang="ts">
import CommonDrawer from '@/components/Drawer/Drawer.vue'
import { SysDict } from '@/types/sys/dict'
import { TableDataType } from '@/types/type'
import { http } from '@/utils/request'
import { defineComponent, ref, onMounted, reactive } from 'vue'
import { useForm } from '@ant-design-vue/use'

const DictDrawer = defineComponent({
  components: {
    CommonDrawer,
  },
  setup() {
    const modalForm = reactive({
      title: '添加',
      visible: false,
      loading: false,
    })
    const ruleRef = reactive({
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
    })
    const { resetFields, validate, validateInfos } = useForm(modalForm, ruleRef)
    const tableData = reactive<TableDataType<SysDict>>({
      data: [],
      page: 1,
      page_size: 10,
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
          slots: { customRender: 'action' },
        },
      ],
    })
    const visible = ref(false)
    onMounted(() => {
      http({ url: '/dict', method: 'GET' }).then((res) => {
        console.log('🚀 ~ file: dict-drawer.vue ~ line 42 ~ http ~ res', res)
      })
      console.log(123)
    })
    const Close = () => {
      visible.value = false
    }
    const IsShow = (record: SysDict) => {
      console.log(
        '🚀 ~ file: dict-drawer.vue ~ line 26 ~ IsShow ~ record',
        record
      )
      visible.value = true
    }
    const Edit = (record: SysDict) => {
      console.log(record, 'record')
    }
    const Add = () => {
      modalForm.visible = true
    }
    const handleOk = () => {
      validate()
        .then((res) => {
          console.log(res, 'res')
        })
        .catch((err) => {
          console.error(err)
        })
    }
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
      resetFields,
      validateInfos,
      handleOk,
    }
  },
})
export default DictDrawer
</script>
