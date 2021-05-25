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
    <el-table-column property="value" label="价值"></el-table-column>
  </el-table>
</template>

<script>
import { defineComponent, ref } from "vue";

export default defineComponent({
  props: {
    data: Object,
    loading: Boolean,
  },
  setup() {},
});
</script>
