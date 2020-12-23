<template>
  <div class="space-margin-bottom">
    <common-button title="添加" @change="ChangeClick()" icon-name="add" />
  </div>
  <a-table
    :columns="tableCont.columns"
    rowKey="id"
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
      <a-form-item label="菜单类型" v-bind="validateInfos.type">
        <a-select v-model:value="modelRef.type">
          <a-select-option
            v-for="(item, index) in optionsType"
            :key="index"
            :value="item.value"
          >
            {{ item.label }}
          </a-select-option>
        </a-select>
      </a-form-item>
      <a-form-item label="权限标识">
        <a-input v-model:value="modelRef.perms"></a-input>
      </a-form-item>
      <a-form-item label="菜单标识(name)" v-bind="validateInfos.code">
        <a-input v-model:value="modelRef.code"></a-input>
      </a-form-item>
      <a-form-item label="排序" v-bind="validateInfos.order_num">
        <a-input-number v-model:value="modelRef.order_num"></a-input-number>
      </a-form-item>
    </a-form>
  </common-drawer>
</template>
<script lang="ts">
import { defineComponent, reactive, onMounted, toRaw, ref } from 'vue'
import CommonButton from '@/components/Button/Button.vue'
import CommonDrawer, { DrawerProps } from '@/components/Drawer/Drawer.vue'
import { http } from '@/utils/request'
import { MenuType } from '@/types/sys'
import { useForm } from '@ant-design-vue/use'
import { message } from 'ant-design-vue'
import { TableDataType, TablePaginType } from '@/types/type'
import { Method } from 'axios'
import { cloneDeep } from 'lodash'
import { ListObjCompare, ListToTree } from '@/utils'

const SysMenu = defineComponent({
  components: { CommonButton, CommonDrawer },
  name: 'sys-menu',
  setup() {
    const optionsType = ref([
      {
        value: 1,
        label: '目录',
      },
      {
        value: 2,
        label: '菜单',
      },
      {
        value: 3,
        label: '按钮',
      },
    ])
    const tableCont = reactive<TableDataType<MenuType>>({
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
          title: '权限标识',
          dataIndex: 'perms',
        },
        {
          title: '菜单标识(name)',
          dataIndex: 'code',
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
    const treeOptions = reactive<{ options: MenuType[] }>({ options: [] })
    const drawerData = reactive<DrawerProps>({
      title: '添加',
      visible: false,
      loading: false,
    })
    const modelRef = reactive<MenuType>({
      parent_id: '',
      name: '',
      perms: '',
      code: '',
      type: 1,
      order_num: 1,
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
      type: [
        {
          required: true,
          message: '请选择类型',
        },
      ],
      code: [
        {
          required: true,
          message: '请输入菜单标识，前端路由name',
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
      http<MenuType>({
        url: '/menu',
        method: 'GET',
        params: {
          page: tableCont.page,
          page_size: tableCont.page_size,
        },
      }).then((res) => {
        const list = res.list.sort(ListObjCompare('order_num'))
        tableCont.loading = false
        tableCont.data = ListToTree<MenuType>(list)
        tableCont.total = res.total
        list.map((item) => {
          item.title = item.name
          item.value = item.id
        })
        treeOptions.options = ListToTree<MenuType>(list)
      })
    }
    onMounted(() => {
      getList()
    })
    const onSubmit = () => {
      validate<MenuType>()
        .then(() => {
          drawerData.loading = true
          const data = cloneDeep(toRaw(modelRef))
          delete data.id
          let method: Method = 'POST'
          if (modelRef.id) {
            method = 'PUT'
          }
          http<MenuType>({
            url: modelRef.id ? `menu/${modelRef.id}` : 'menu',
            method,
            data,
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
    function Editor(record: MenuType) {
      const data = toRaw(record)
      drawerData.title = '修改'
      drawerData.visible = true
      modelRef.name = data.name
      modelRef.order_num = data.order_num
      modelRef.parent_id = data.parent_id
      modelRef.id = data.id
      modelRef.type = data.type
      modelRef.code = data.code
      modelRef.perms = data.perms
    }
    function Del(record: MenuType) {
      const data = toRaw(record)
      http({ url: `menu/${data.id}`, method: 'delete' }).then(() => {
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
      optionsType,

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
export default SysMenu
</script>
<style scoped lang="less">
@import './depart.less';
</style>
