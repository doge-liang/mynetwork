<template>
  <el-table v-loading="loading" :data="data" highlight-current-row>
    <el-table-column property="ID" label="ID" width="200"></el-table-column>
    <el-table-column
      property="stockID"
      label="股票"
      width="200"
    ></el-table-column>
    <el-table-column property="amount" label="价值"></el-table-column>
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
