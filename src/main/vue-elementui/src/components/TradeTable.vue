<template>
  <el-table :data="Trades" highlight-current-row>
    <el-table-column property="id" label="id" width="100"></el-table-column>
    <el-table-column
      property="dateTime"
      label="日期"
      width="230"
    ></el-table-column>
    <el-table-column
      property="stockID"
      label="股票"
      width="100"
    ></el-table-column>
    <el-table-column
      property="commission"
      label="手续费"
      width="100"
    ></el-table-column>
    <el-table-column
      property="price"
      label="价格"
      width="100"
    ></el-table-column>
    <el-table-column property="amount" label="数量"></el-table-column>
  </el-table>
</template>

<script>
import { defineComponent } from '@vue/composition-api'

export default defineComponent({
    
    setup() {
        
    },
})
</script>
