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
      planningTrades: [
        {
          amount: 100,
          stockID: "600208.SH",
        },
        {
          amount: -100,
          stockID: "603658.SH",
        },
        {
          amount: -100,
          stockID: "600196.SH",
        },
        {
          amount: 100,
          stockID: "000723.SZ",
        },
        {
          amount: 100,
          stockID: "600655.SH",
        },
        {
          amount: -100,
          stockID: "000725.SZ",
        },
      ],
      positions: [
        {
          amount: 3431,
          stockID: "600161.SH",
        },
        {
          amount: 8175,
          stockID: "603799.SH",
        },
        {
          amount: 29548,
          stockID: "603392.SH",
        },
        {
          amount: 2581,
          stockID: "300433.SZ",
        },
        {
          amount: 32448,
          stockID: "002821.SZ",
        },
        {
          amount: 2171,
          stockID: "002555.SZ",
        },
        {
          amount: 2713,
          stockID: "300136.SZ",
        },
        {
          amount: 20577,
          stockID: "300122.SZ",
        },
        {
          amount: 3675,
          stockID: "002241.SZ",
        },
        {
          amount: 6289.467329999996,
          stockID: "cash",
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
