<template>
  <el-card class="strategy-card" shadow="hover">
    <div class="card-header">
      <span>{{ strategy.name }}</span>
      <el-button
        type="primary"
        round
        v-if="!strategy.isSub"
        :loading="loading"
        @click="subscribe(strategy)"
        >订阅策略</el-button
      >
      <el-button :loading="loading" round v-else @click="unsubscribe(strategy)"
        >取消订阅</el-button
      >
    </div>
    <div class="card-item">
      <span style="color: #409eff">策略ID: {{ strategy.id }}</span>
      <span style="color: #e6a23c">夏普率:{{ strategy.sharpeRatio }} </span>
    </div>
    <el-divider></el-divider>
    <div class="card-item">
      <span style="color: #f56c6c"
        >年化收益: {{ strategy.annualReturn + "%" }}</span
      >
      <span style="color: #f56c6c"
        >最大回撤: {{ strategy.maxDrawdown + "%" }}</span
      >
    </div>
    <div>
      <el-button type="text" @click="toMarket(strategy)">信号及持仓</el-button>
      <el-divider direction="vertical" />
      <el-button type="text" @click="toTrade(strategy)">交易记录</el-button>
    </div>
  </el-card>
</template>

<script>
import { defineComponent, ref } from "vue";
import { useRoute, useRouter } from "vue-router";
import { ElMessage } from "element-plus";
import { Subscribe, Unsubscribe } from "@/http/apis";

export default defineComponent({
  props: {
    strategy: Object,
  },

  setup(props, context) {
    const router = useRouter();
    const route = useRoute();

    const loading = ref(false);

    console.log("加载 strategy-card 组件");
    console.log(props.strategy);

    const toMarket = (strategy) => {
      console.log(strategy);
      ElMessage("操作成功!");
      if (!strategy.isSub && strategy.state === 1) {
        router.push({
          path: "/strategy/" + strategy.id + "/market",
          query: {
            display: false,
          },
        });
      } else {
        router.push({
          path: "/strategy/" + strategy.id + "/market",
          query: {
            display: true,
          },
        });
      }
    };
    const toTrade = (strategy) => {
      console.log(strategy);
      ElMessage("操作成功!");
      router.push({
        path: "/strategy/" + strategy.id + "/trade",
      });
    };

    const subscribe = (strategy) => {
      loading.value = true;
      Subscribe("strategy/" + strategy.id).then((Response) => {
        ElMessage.success("订阅成功!");
        loading.value = false;
        props.strategy.isSub = true;
      });
    };

    const unsubscribe = (strategy) => {
      loading.value = true;
      Unsubscribe("strategy/" + strategy.id).then((Response) => {
        ElMessage.success("取消成功!");
        loading.value = false;
        props.strategy.isSub = false;
      });
    };

    return {
      toMarket,
      toTrade,
      subscribe,
      unsubscribe,
      loading,
    };
  },
});
</script>

<style>
.strategy-card {
  width: 500px;
  height: 250px;
  margin: auto;
}
.card-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  font-family: "微软雅黑";
  font-size: larger;
}
.card-item {
  display: flex;
  justify-content: space-between;
  margin-bottom: 18px;
}
</style>