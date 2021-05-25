<template>
  <el-table
    v-loading="loading"
    :data="data"
    border
    height="180"
    highlight-current-row
  >
    <el-table-column property="id" label="id"></el-table-column>
    <el-table-column property="stockID" label="股票"></el-table-column>
    <el-table-column property="amount" label="交易份额"></el-table-column>
  </el-table>
</template>

<script>
import { defineComponent } from "vue";
import { ref } from "vue";

export default defineComponent({
  props: {
    loading: Boolean,
    data: Object,
  },
  setup() {},
});
</script>
