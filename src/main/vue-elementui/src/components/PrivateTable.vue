<template>
  <el-table v-loading="loading" :data="data" highlight-current-row>
    <el-table-column property="id" label="id" width="500"></el-table-column>
    <el-table-column
      property="hash"
      label="hashcode"
      width="500"
    ></el-table-column>
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
  setup(props, context) {
    console.log(props.data);
    console.log(props.loading);
  },
});
</script>
