<template>
  <div class="container">
    <a-tabs
      v-model:activeKey="activeKey"
      hide-add
      type="editable-card"
      @change="OnChange"
      @edit="OnEdit"
    >
      <a-tab-pane
        v-for="(pane, index) in panes"
        :key="index"
        :tab="pane.title"
        :closable="pane.closable"
      >
        <div class="content">
          <slot />
        </div>
      </a-tab-pane>
    </a-tabs>
  </div>
</template>
<script lang="ts">
import { PanesType } from '@/store/sys/sys-crumbs'
import { computed, defineComponent, ref, toRaw, watch } from 'vue'
import { useStore } from 'vuex'
import { useRouter } from 'vue-router'
import { localForage } from '@/utils/localforage'
import { STORELETMENUPATH } from '@/utils/constant'
import { SysTab, SysTabDel } from '@/types/sys/tab'
import { DELETETABS, DELETETABSACTION } from '@/store/mutation-types'
import _ from 'lodash'

export default defineComponent({
  setup() {
    const store = useStore()
    const router = useRouter()
    const activeKey = ref(0)

    const panes = computed(() => {
      const data = store.state.crumbs.panes
      localForage.getItem<SysTab>(STORELETMENUPATH).then((res) => {
        if (data && data.length > 0) {
          if (res) {
            const result = data.findIndex((item) => item.id === res.id)
            // 设置当前被选择项
            if (result !== -1) {
              activeKey.value = result
            }
          }
        }
      })
      return data
    })
    const OnChange = (val: number) => {
      const result = FormatData()[val]
      activeKey.value = val
      router.push(result.path)
      store.commit(DELETETABS, { selectData: toRaw(result) })
    }
    const FormatData = () => {
      const data = toRaw(store.state.crumbs.panes)
      const list: PanesType[] = []
      data.forEach((item: PanesType) => {
        list.push(toRaw(item))
      })
      return list
    }
    const OnEdit = (val: number) => {
      const value = _.cloneDeep(toRaw(activeKey.value))
      const result: SysTabDel = {
        delData: FormatData()[val],
        selectData: '',
      }
      if (FormatData().length > 1) {
        const index = value >= val ? value - 1 : value

        result.selectData = FormatData()[index]
      }
      store.dispatch(DELETETABSACTION, toRaw(result))
    }
    return {
      activeKey,
      OnChange,
      panes,
      OnEdit,
    }
  },
})
</script>
<style lang="less" scoped>
@import 'index.less';
</style>
