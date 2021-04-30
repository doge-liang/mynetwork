<template>
  <div>
    <el-table :data="OpsSignalData" highlight-current-row>
      <el-table-column property="id" label="ID" width="100"></el-table-column>
      <el-table-column
        property="date"
        label="日期"
        width="100"
      ></el-table-column>
      <el-table-column
        property="stock"
        label="股票"
        width="100"
      ></el-table-column>
      <el-table-column
        property="ops"
        label="操作"
        width="230"
      ></el-table-column>
      <el-table-column property="amount" label="数量"></el-table-column>
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

export default {
  mounted() {},
  data() {
    return {
      OpsSignalData: [
        {
          id: "00010001",
          date: "2020-05-04",
          stock: "Stock_A",
          ops: "buy at 801.02, sell at 812.0",
          amount: "1",
        },
        {
          id: "00010002",
          date: "2020-05-06",
          stock: "Stock_A",
          ops: "buy at 810.22, sell at 819.98",
          amount: "1",
        },
        {
          id: "00010003",
          date: "2020-05-12",
          stock: "Stock_A",
          ops: "buy at 828.41, sell at 965.0",
          amount: "1",
        },
        {
          id: "00010004",
          date: "2020-05-18",
          stock: "Stock_A",
          ops: "buy at 1010.2, sell at 998.88",
          amount: "1",
        },
        {
          id: "00010005",
          date: "2020-05-19",
          stock: "Stock_A",
          ops: "buy at 994.1, sell at 1118.5",
          amount: "1",
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
