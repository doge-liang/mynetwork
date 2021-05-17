<template>
  <div>
    <el-container>
      <el-header>计划交易</el-header>
      <el-main>
        <el-table :data="planningTrades" highlight-current-row>
          <el-table-column
            property="hash"
            label="hashcode"
            width="500"
          ></el-table-column>
        </el-table>
      </el-main>
      <el-header>持仓信息</el-header>
      <el-main>
        <el-table :data="positions" highlight-current-row>
          <el-table-column
            property="hash"
            label="hashcode"
            width="500"
          ></el-table-column>
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
  name: "PrivateMarket",
  mounted() {},
  data() {
    return {
      //   strategyId: this.$router.params.strategyId,
      planningTrades: [{
        hash:
          "13cb230ff2557005cc61fc1f4bb6f6cab6a9df58b19c202cb70bdf52e92dfc3a",
      }],
      positions: [{
        hash:
          "f036e87b74b9796acec88d8d8981ec2ab98238658c8bc55fcd5ff94c4d0b3582",
      }],
    };
  },
  methods: {
    toHome() {
      this.$router.push("/home");
    },
  },
};
</script>
