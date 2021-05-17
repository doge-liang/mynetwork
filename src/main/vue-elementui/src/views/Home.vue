<template>
  <div>
    <el-backtop :visibility-height="6"></el-backtop>
    <div v-for="item in strategys" :key="item.id">
      <el-row>
        <el-col :span="6" :offset="7">
          <el-card class="productOwn-card" shadow="hover">
            <template #header>
              <div class="card-header">
                <span>{{ item.name }}</span>
                <el-button type="primary" round @click="subscribe(item.id)"
                  >订阅策略</el-button
                >
              </div>
            </template>
            <div class="item">
              <span style="color: #409eff">策略ID: {{ item.id }}</span>
              <span style="color: #e6a23c">夏普率:{{ item.sharpeRatio }} </span>
            </div>
            <el-divider></el-divider>
            <div class="item">
              <span style="color: #f56c6c"
                >年化收益: {{ item.annualReturn + "%" }}</span
              >
              <span style="color: #f56c6c"
                >最大回撤: {{ item.maxDrawDown + "%" }}</span
              >
            </div>
            <div>
              <el-button type="text" @click="toMarketDetails(item.id)"
                >计划交易及持仓</el-button
              >
              <el-divider direction="vertical" />
              <el-button type="text" @click="toTradeDetails(item.id)"
                >交易记录</el-button
              >
            </div>
          </el-card>
        </el-col>
      </el-row>
    </div>
  </div>
</template>

<script>

import { getAllStrategies } from "@/http/apis";

export default {
  name: "Home",
  components: {},
  mounted() {
  },
  data() {
    return {
      strategys: [
        {
          id: "1",
          name: "RSI",
          annualReturn: 24.08,
          sharpeRatio: 0.65,
          maxDrawDown: 23.44,
        },
        {
          id: "2",
          name: "MA",
          annualReturn: 12.01,
          sharpeRatio: 3.67,
          maxDrawDown: 17.4,
        },
      ],
    };
  },
  create() {
    this.$data.strategys = getAllStrategies()
  },
  methods: {
    toMarketDetails(strategy) {
      console.log(strategy);
      this.$message("操作成功!");
      this.$router.push({
        path: "/strategy/" + strategy + "/private-market",
      });
    },
    toTradeDetails(strategy) {
      console.log(strategy);
      this.$message("操作成功!");
      this.$router.push({
        path: "/strategy/" + strategy + "/private-trade",
      });
    },
    subscribe() {
      this.$message.success("订阅成功!");
    },
  },
};
</script>

<style>
.el-row {
  margin-bottom: 70px;
}
.el-row:last-child {
  margin-bottom: 0;
}
.el-col {
  border-radius: 4px;
  margin-bottom: 70px;
}
.productOwn-card {
  width: 500px;
  height: 250px;
  margin: 18px;
}
.card-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  font-family: "微软雅黑";
  font-size: larger;
}
.item {
  display: flex;
  justify-content: space-between;
  margin-bottom: 18px;
}
</style>