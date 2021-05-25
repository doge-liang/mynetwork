<template>
  <div>
    <el-container>
      <el-header>计划交易</el-header>
      <el-main>
        <private-table
          :is="!display"
          :loading="loadingPlanningTrades"
        ></private-table>
        <planningTrade-table
          :is="display"
          :loading="loadingPlanningTrades"
        ></planningTrade-table>
      </el-main>
      <el-header>持仓信息</el-header>
      <el-main>
        <private-table
          :is="!display"
          :loading="loadingPositions"
        ></private-table>
        <position-table
          :is="display"
          :loading="loadingPositions"
        ></position-table>
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
import { defineComponent } from "vue";
import { onMounted, ref } from "vue";
import { useRoute, useRouter } from "vue-router";
import {
  getPlanningTradesByStrategyID,
  getPositionsByStrategyID,
} from "@/http/apis";
import { privateTable } from "@/components/PrivateTable.vue";
import { planningTradeTable } from "@/components/PlanningTradeTable.vue";
import { positionTable } from "@/components/PositionTable.vue";

export default defineComponent({
  components: {
    "private-table": privateTable,
    "planningTrade-table": planningTradeTable,
    "position-table": positionTable,
  },
  setup() {
    const router = useRouter();
    const route = useRoute();

    let loadingPlanningTrades = ref(false);
    let loadingPositions = ref(false);
    let planningTrades = ref([]);
    let positions = ref([]);

    let display = ref(false);

    const toHome = () => {
      router.push("/home");
    };

    onMounted(async () => {
      display.value = route.query.display;
      loadingPlanningTrades = true;
      getPlanningTradesByStrategyID(route.path).then((Response) => {
        loadingPlanningTrades.value = false;
        console.log(Response);
        const PlanningTradeOutput = Response.data.data;
        if (display.value) {
          planningTrades.value = PlanningTradeOutput.planningTrades;
        } else {
          planningTrades.value = PlanningTradeOutput.planningTradesHash;
        }
      });

      loadingPositions = true;
      getPositionsByStrategyID(route.path).then((Response) => {
        loadingPositions.value = false;
        console.log(Response);
        const PositionOutput = Response.data.data;
        if (display.value) {
          positions.value = PositionOutput.positions;
        } else {
          positions.value = PositionOutput.positionsHash;
        }
      });
    });

    return {
      loadingPlanningTrades,
      loadingPositions,
      planningTrades,
      positions,
      toHome,
    };
  },
});
</script>
