<template>
  <common-drawer
    title="权限分配"
    ok-text="确定"
    cancel-text="取消"
    :visible="visible"
    :loading="loading"
    @on-close="DClose()"
    @on-ok="DSubmit()"
  >
    <a-spin :spinning="treeData.spinningLoading">
      <a-tree
        v-model:checkedKeys="treeData.checkedKeys"
        checkable
        :expanded-keys="treeData.expandedKeys"
        :auto-expand-parent="treeData.autoExpandParent"
        :selected-keys="treeData.selectedKeys"
        :tree-data="treeData.data"
        @expand="onExpand"
        @select="onSelect"
      />
    </a-spin>
  </common-drawer>
</template>
<script lang="ts">
import { defineComponent, reactive, onMounted, watch } from 'vue'
import CommonDrawer from '@/components/Drawer/Drawer.vue'
import { http } from '@/utils/request'
import { MenuType } from '@/types/sys'
import { ListObjCompare, ListToTree } from '@/utils'

interface TreeDataType {
  spinningLoading: boolean
  checkedKeys: string[]
  expandedKeys: string[]
  autoExpandParent: boolean
  selectedKeys: string[]
  data: MenuType[]
}
interface Props {
  visible: boolean
  data: string[]
}
const CommonTree = defineComponent({
  name: 'common-tree',
  components: {
    CommonDrawer,
  },
  props: {
    visible: {
      type: Boolean,
      default: false,
    },
    loading: {
      type: Boolean,
      default: false,
    },
  },
  setup(props: Props, { emit }: { emit: Function }) {
    const treeData = reactive<TreeDataType>({
      spinningLoading: false,
      checkedKeys: [],
      expandedKeys: [],
      autoExpandParent: true,
      selectedKeys: [],
      data: [],
    })

    function getList() {
      treeData.spinningLoading = true
      http<MenuType>({
        url: '/menu',
        method: 'GET',
        params: {
          page: 1,
          page_size: 1000,
        },
      }).then((res) => {
        const list = res.list.sort(ListObjCompare('order_num'))
        treeData.spinningLoading = false
        const result = ListToTree<MenuType>(list)
        result.map((item) => {
          item.title = item.name
          item.key = item.id
        })
        treeData.data = result
      })
    }
    onMounted(() => {
      getList()
    })
    function onExpand(expandedKeys: string[]) {
      console.log(expandedKeys)
    }
    function onSelect(selectedKeys: string[], info: any) {
      console.log(selectedKeys)
      console.log(info)
    }
    function DClose() {
      emit('on-close')
    }
    function DSubmit() {
      emit('on-submit')
    }
    watch(
      () => props.data,
      (newValue) => {
        treeData.selectedKeys = newValue
      }
    )
    return {
      // data
      treeData,

      // methods
      onExpand,
      onSelect,
      DClose,
      DSubmit,
    }
  },
})
export default CommonTree
</script>
