<template>
  <div class="space-margin-bottom">
    <common-button title="添加" @change="ChangeClick()" icon-name="add" />
  </div>
  <a-table
    :columns="tableCont.columns"
    rowKey="id"
    :data-source="tableCont.data"
  />
</template>
<script lang="ts">
import { defineComponent, reactive, onMounted } from 'vue'
import CommonButton from '@/components/button/button.vue'
import { http } from '@/utils/request'
import { DepartType } from '@/types/sys/depart'

const SysDepart = defineComponent({
  components: { CommonButton },
  name: 'sys-depart',
  setup() {
    const tableCont = reactive<{
      data: DepartType[]
      columns: any
    }>({
      columns: [
        {
          title: '名称',
          dataIndex: 'name',
        },
        {
          title: '父级节点',
          dataIndex: 'parent_id',
        },
        {
          title: '排序',
          dataIndex: 'order_num',
        },
      ],
      data: [],
    })

    function ChangeClick() {
      console.log('add')
    }
    function getList() {
      http<DepartType>({ url: '/depart', method: 'GET' }).then((res) => {
        tableCont.data = res.list
      })
    }
    onMounted(() => {
      getList()
    })
    return {
      tableCont,
      ChangeClick,
    }
  },
})
export default SysDepart
</script>
<style scoped lang="less">
@import './depart.less';
</style>
