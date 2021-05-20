<template>
  <div>
    <el-container>
      <el-header>计划交易</el-header>
      <el-main>
        <el-table :data="planningTrades" highlight-current-row>
          <!-- <el-table-column property="id" label="id" width="150"></el-table-column> -->
          <!-- <el-table-column
        property="date"
        label="日期"
        width="150"
      ></el-table-column> -->
      <el-table-column
            property="ID"
            label="ID"
            width="200"
          ></el-table-column>
          <el-table-column
            property="stockID"
            label="股票"
            width="200"
          ></el-table-column>
          <el-table-column property="amount" label="交易份额"></el-table-column>
        </el-table>
      </el-main>
      <el-header>持仓信息</el-header>
      <el-main>
        <el-table :data="positions" highlight-current-row>
          <!-- <el-table-column property="id" label="id" width="150"></el-table-column> -->
          <!-- <el-table-column
        property="date"
        label="日期"
        width="150"
      ></el-table-column> -->
            <el-table-column
            property="ID"
            label="ID"
            width="200"
          ></el-table-column>
          <el-table-column
            property="stockID"
            label="股票"
            width="200"
          ></el-table-column>
          <el-table-column property="amount" label="价值"></el-table-column>
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
// import { defineComponent } from "@vue/composition-api";

export default {
  name: "Market",
  mounted() {},
  data() {
    return {
      //   strategyId: this.$router.params.strategyId,
      planningTrades: [],
      positions: [],
    };
  },
  created() {

  },
  methods: {
    toHome() {
      this.$router.push("/home");
    },
  },
};
</script>
