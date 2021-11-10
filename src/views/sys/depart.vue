<template>
  <div class="space-margin-bottom">
    <common-button
      v-bt-auth:add="{ title: true }"
      title="添加"
      @change="ChangeClick()"
      icon-name="add"
    />
  </div>
  <a-table
    :columns="tableCont.columns"
    row-key="id"
    :data-source="tableCont.data"
    :pagination="{
      total: tableCont.total,
    }"
    @change="Change"
    :loading="tableCont.loading"
  >
    <template #action="{ record }">
      <a-button
        type="primary"
        style="margin-right: 15px"
        @click="Editor(record)"
        v-bt-auth:edit
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
    :title="drawerData.title"
    :visible="drawerData.visible"
    @onClose="drawerData.visible = false"
    cancel-text="取消"
    ok-text="确定"
    @onOk="onSubmit"
    :loading="drawerData.loading"
  >
    <a-form :label-col="{ span: 4 }" :wrapper-col="{ span: 20 }">
      <a-form-item label="父级id" v-bind="validateInfos.parent_id">
        <a-tree-select
          v-model:value="modelRef.parent_id"
          style="width: 100%"
          :dropdown-style="{ maxHeight: '400px', overflow: 'auto' }"
          :tree-data="treeOptions.options"
          placeholder="请选择"
          tree-default-expand-all
        >
        </a-tree-select>
      </a-form-item>
      <a-form-item label="名称" v-bind="validateInfos.name">
        <a-input v-model:value="modelRef.name"></a-input>
      </a-form-item>
      <a-form-item label="排序" v-bind="validateInfos.order_num">
        <a-input-number v-model:value="modelRef.order_num"></a-input-number>
      </a-form-item>
    </a-form>
  </common-drawer>
</template>
<script lang="ts">
import { defineComponent, reactive, onMounted, toRaw } from 'vue'
import CommonButton from '@/components/Button/Button.vue'
import CommonDrawer, { DrawerProps } from '@/components/Drawer/Drawer.vue'
import { http } from '@/utils/request'
import { DepartType } from '@/types/sys'
import { useForm } from '@ant-design-vue/use'
import { message } from 'ant-design-vue'
import { TableDataType, TablePaginType } from '@/types/type'
import { Method } from 'axios'
import { cloneDeep } from 'lodash'
import { ListObjCompare, ListToTree } from '@/utils'

const SysDepart = defineComponent({
  components: { CommonButton, CommonDrawer },
  name: 'SysDepart',
  setup() {
    const tableCont = reactive<TableDataType<DepartType>>({
      page: 1,
      page_size: 10,
      total: 0,
      loading: false,
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
        {
          title: '操作',
          key: 'action',
          slots: {
            customRender: 'action',
          },
        },
      ],
      data: [],
    })
    const treeOptions = reactive<{ options: DepartType[] }>({ options: [] })
    const drawerData = reactive<DrawerProps>({
      title: '添加',
      visible: false,
      loading: false,
    })
    const modelRef = reactive<DepartType>({
      parent_id: '',
      name: '',
      order_num: 0,
      id: '',
    })
    const ruleRef = reactive({
      parent_id: [
        {
          required: true,
          message: '请选择父级',
        },
      ],
      name: [
        {
          required: true,
          message: '请输入名称',
        },
      ],
      order_num: [
        {
          required: true,
          message: '请输入排序',
        },
      ],
    })
    const { resetFields, validate, validateInfos } = useForm(modelRef, ruleRef)
    function getList() {
      tableCont.loading = true
      http<DepartType>({
        url: '/depart',
        method: 'GET',
        params: {
          page: tableCont.page,
          page_size: tableCont.page_size,
        },
      }).then((res) => {
        const list = res.list.sort(ListObjCompare('order_num'))
        tableCont.loading = false
        tableCont.data = ListToTree(list)
        tableCont.total = res.total
        list.map((item) => {
          item.title = item.name
          item.value = item.id
        })
        treeOptions.options = ListToTree(list)
      })
    }
    onMounted(() => {
      getList()
    })
    const onSubmit = () => {
      validate<DepartType>()
        .then(() => {
          drawerData.loading = true
          const data = cloneDeep(toRaw(modelRef))
          // @ts-ignore
          delete data.id
          let method: Method = 'POST'
          if (modelRef.id) {
            method = 'PUT'
          }
          http<DepartType>({
            url: modelRef.id ? `depart/${modelRef.id}` : 'depart',
            method,
            body: data,
          }).then(() => {
            message.success(`${drawerData.title}成功`)
            drawerData.loading = false
            drawerData.visible = false
            getList()
          })
        })
        .catch((err) => {
          console.log('error', err)
        })
    }
    function ChangeClick() {
      drawerData.title = '添加'
      drawerData.visible = true
      resetFields()
      modelRef.parent_id = '0'
    }
    function Editor(record: DepartType) {
      const data = toRaw(record)
      drawerData.title = '修改'
      drawerData.visible = true
      modelRef.name = data.name
      modelRef.order_num = data.order_num
      modelRef.parent_id = data.parent_id
      modelRef.id = data.id
    }
    function Del(record: DepartType) {
      const data = toRaw(record)
      http({ url: `depart/${data.id}`, method: 'delete' }).then(() => {
        message.success('删除成功')
        getList()
      })
    }
    function Change(pagination: TablePaginType) {
      tableCont.page = pagination.current
      getList()
    }

    return {
      tableCont,
      ChangeClick,
      drawerData,
      treeOptions,
      Change,

      // table
      Editor,
      Del,

      //form
      validateInfos,
      resetFields,
      modelRef,
      onSubmit,
    }
  },
})
export default SysDepart
</script>
<style scoped lang="less">
@import './depart.less';
</style>
