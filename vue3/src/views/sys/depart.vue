<template>
  <a-table
    :columns="columns"
    :data-source="data"
    :row-selection="rowSelection"
  />
</template>
<script lang="ts">
import { defineComponent, reactive } from 'vue'
const App = defineComponent({
  name: 'sys-depart',
  setup() {
    const tableCont = reactive({
      columns: [
        {
          title: '',
        },
      ],
    })
    return {}
  },
})
export default App
</script>
