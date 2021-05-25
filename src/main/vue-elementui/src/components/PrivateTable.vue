<template>
  <el-table
    v-loading="loading"
    :data="data"
    border
    height="180"
    highlight-current-row
  >
    <el-table-column property="id" label="id"></el-table-column>
    <el-table-column property="hashcode" label="hashcode"></el-table-column>
  </el-table>
</template>

<script>
import { defineComponent } from "vue";
import { reactive, ref } from "vue";

export default defineComponent({
  props: {
    data: Object,
    loading: Boolean,
  },
  setup(props, context) {},
});
</script>
