<template>
  <div>
    <el-container>
      <el-header> 交易记录 </el-header>
      <el-main>
        <el-table :data="Trades" highlight-current-row>
          <el-table-column
            property="id"
            label="id"
            width="100"
          ></el-table-column>
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
      </el-main>
    </el-container>
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
  name: "Trade",
  mounted() {},
  data() {
    return {
      Trades: [],
    };
  },
  created(){

  },
  methods: {
    toHome() {
      this.$router.push("/home");
    },
  },
};
</script>
