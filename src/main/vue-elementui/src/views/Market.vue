<template>
  <div>
    <el-table
      :data="marketSignalData"
      highlight-current-row
    >
      <el-table-column property="id" label="ID" width="150"></el-table-column>
      <el-table-column
        property="date"
        label="日期"
        width="150"
      ></el-table-column>
      <el-table-column
        property="stock"
        label="股票"
        width="200"
      ></el-table-column>
      <el-table-column property="price" label="价格"></el-table-column>
    </el-table>
    <el-button
      type="primary"
      style="margin-top: 20px; margin-bottom: 45px"
      @click="toHome"
      >返回</el-button
    >
  </div>
</template>

<script>
// import { defineComponent } from "@vue/composition-api";

export default {
  name: "Market",
  mounted() {},
  data() {
    return {
    //   strategyId: this.$router.params.strategyId,
      marketSignalData: [
        {
          id: "00000001",
          date: "2020-05-04",
          stock: "Stock_A",
          price: 801.02,
        },
        {
          id: "00000002",
          date: "2020-05-06",
          stock: "Stock_A",
          price: 810.22,
        },
        {
          id: "00000003",
          date: "2020-05-12",
          stock: "Stock_A",
          price: 828.41,
        },
        {
          id: "00000004",
          date: "2020-05-18",
          stock: "Stock_A",
          price: 1010.2,
        },
        {
          id: "00000005",
          date: "2020-05-19",
          stock: "Stock_A",
          price: 994.1,
        },
      ],
    };
  },
  methods: {
    toHome() {
      this.$router.push("/home");
    },
  },
};
</script>
