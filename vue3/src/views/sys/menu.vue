<template>
  <div class="space-margin-bottom">
    <common-button
      v-bt-auth:add
      title="添加"
      @change="ChangeClick()"
      icon-name="add"
    />
  </div>
  <a-table
    :scroll="{ x: 500 }"
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
        v-bt-auth:edit
        @click="Editor(record)"
        >编辑</a-button
      >
      <a-popconfirm
        title="确定删除吗?"
        ok-text="删除"
        cancel-text="取消"
        @confirm="Del(record)"
      >
        <a-button type="danger" v-bt-auth:del>删除</a-button>
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
      <a-form-item label="名称" v-bind="validateInfos.title">
        <a-input v-model:value="modelRef.title"></a-input>
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
      <a-form-item label="权限标识" v-if="modelRef.type === 3">
        <a-input v-model:value="modelRef.perms"></a-input>
      </a-form-item>
      <div v-if="modelRef.type !== 3">
        <a-form-item label="菜单name">
          <a-input v-model:value="modelRef.name"></a-input>
        </a-form-item>
        <a-form-item label="路径">
          <a-input v-model:value="modelRef.path"></a-input>
        </a-form-item>
        <a-form-item label="组件地址">
          <a-input v-model:value="modelRef.component"></a-input>
        </a-form-item>
        <a-form-item label="重定向地址">
          <a-input v-model:value="modelRef.redirect"></a-input>
        </a-form-item>
        <a-form-item label="图标">
          <a-input v-model:value="modelRef.icon"></a-input>
        </a-form-item>
        <a-form-item label="是否隐藏">
          <a-radio-group name="radioGroup" v-model:value="modelRef.hidden">
            <a-radio :value="0"> 否 </a-radio>
            <a-radio :value="1"> 是 </a-radio>
          </a-radio-group>
        </a-form-item>
      </div>
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
    const width = 150
    const tableCont = reactive<TableDataType<MenuType>>({
      page: 1,
      page_size: 1000,
      total: 0,
      loading: false,
      columns: [
        {
          title: '名称',
          dataIndex: 'title',
          width: width * 2,
          fixed: true,
        },
        {
          title: '类型',
          dataIndex: 'type',
          width: width / 2,
        },
        {
          title: '路径',
          dataIndex: 'path',
          width,
        },
        {
          title: '组件地址',
          dataIndex: 'component',
          width,
        },
        {
          title: '重定向地址',
          dataIndex: 'redirect',
          width,
        },
        {
          title: '图标',
          dataIndex: 'icon',
          width,
        },
        {
          title: '父级节点',
          dataIndex: 'parent_id',
          width,
        },
        {
          title: '权限标识',
          dataIndex: 'perms',
          width: width / 2,
        },
        {
          title: '菜单标识',
          dataIndex: 'name',
          width: width / 2,
        },
        {
          title: '排序',
          dataIndex: 'order_num',
          width: width / 2,
        },
        {
          title: '操作',
          key: 'action',
          fixed: 'right',
          width: 200,
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
    let modelRef = reactive<MenuType>({
      parent_id: '',
      path: '',
      component: '',
      redirect: '',
      icon: '',
      title: '',
      perms: '',
      name: '',
      type: 1,
      order_num: 1,
      id: '',
      hidden: 0,
    })
    const ruleRef = reactive({
      parent_id: [
        {
          required: true,
          message: '请选择父级',
        },
      ],
      title: [
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
        const list = cloneDeep(res.list.sort(ListObjCompare('order_num')))
        tableCont.loading = false
        tableCont.data = ListToTree(list)
        tableCont.total = res.total
        const treeMenus = cloneDeep(res.list)
        treeMenus.map((item) => {
          item.value = item.id
        })
        treeOptions.options = ListToTree(treeMenus)
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
    function Editor(record: MenuType) {
      drawerData.title = '修改'
      drawerData.visible = true
      modelRef = Object.assign(modelRef, record)
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
