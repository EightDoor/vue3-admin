<template>
  <common-drawer
    title="字典配置"
    cancelText="取消"
    :visible="visible"
    @on-close="Close()"
  >
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
  </common-drawer>
</template>
<script lang="ts">
import CommonDrawer from '@/components/Drawer/Drawer.vue'
import { SysDict, SysDictItem } from '@/types/sys/dict'
import { TableDataType } from '@/types/type'
import { http } from '@/utils/request'
import { defineComponent, ref, onMounted, reactive } from 'vue'
const DictDrawer = defineComponent({
  components: {
    CommonDrawer,
  },
  setup() {
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
          name: 'label',
        },
        {
          title: '数据值',
          dataIndex: 'value',
          name: 'value',
        },
        {
          title: '描述',
          dataIndex: 'describe',
          name: 'describe',
        },
        {
          title: '操作',
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
    return {
      // data
      visible,
      tableData,
      // methods
      Close,
      IsShow,
      Edit,
    }
  },
})
export default DictDrawer
</script>
