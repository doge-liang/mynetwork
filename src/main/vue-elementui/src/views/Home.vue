<template>
  <div>
    <el-backtop :visibility-height="6"></el-backtop>
    <el-container id="strategy-card-list">
      <el-main v-loading="loading">
        <div v-for="item in strategies" :key="item.id">
          <strategy-card :strategy="item"></strategy-card>
        </div>
      </el-main>
    </el-container>
  </div>
</template>

<script>
import { getAllStrategies } from "@/http/apis";
import StrategyCard from "@/components/StrategyCard.vue";
import { onMounted, ref } from "vue";

export default {
  name: "Home",
  components: {
    "strategy-card": StrategyCard,
  },
  setup() {
    const strategies = ref([]);
    let loading = ref(true);
    // strategies.value = [
    //   {
    //     id: "6799979985630134272",
    //     name: "MA",
    //     provider: "",
    //     maxDrawdown: 16.82,
    //     annualReturn: 12.96,
    //     sharpeRatio: 0.63,
    //     state: 1,
    //     isSub: false,
    //   },
    //   {
    //     id: "6799980250357825536",
    //     name: "RSI",
    //     provider: "",
    //     maxDrawdown: 15.64,
    //     annualReturn: 33.48,
    //     sharpeRatio: 0.79,
    //     state: 0,
    //     isSub: false,
    //   },
    // ];
    onMounted(async () => {
      loading = true;
      getAllStrategies().then((resp) => {
        console.log(resp);
        if (resp.data.code === 200) {
          strategies.value = resp.data.data;
          loading = false;
        }
      });
    });

    return {
      strategies,
      loading,
    };
  },
};
</script>

<style>
#strategy-card-list {
  width: 550px;
  margin: auto;
}
</style>