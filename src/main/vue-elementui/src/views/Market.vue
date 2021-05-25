<template>
  <div>
    <el-container>
      <el-header>计划交易</el-header>
      <el-main>
        <component
          :is="displayPlanningTradeTable"
          :loading="loadingPlanningTrades"
          :data="PlanningTrades"
        ></component>
      </el-main>
      <el-header>持仓信息</el-header>
      <el-main>
        <component
          :is="displayPositionTable"
          :loading="loadingPositions"
          :data="Positions"
        ></component>
      </el-main>
    </el-container>
    <el-footer>
      <el-button
        type="primary"
        style="margin-top: 20px; margin-bottom: 45px"
        @click="toHome"
        >返回</el-button
      >
    </el-footer>
  </div>
</template>

<script>
import { defineComponent, reactive } from "vue";
import { onMounted, ref } from "vue";
import { useRoute, useRouter } from "vue-router";
import {
  getPlanningTradesByStrategyID,
  getPositionsByStrategyID,
} from "@/http/apis";
import privateTable from "@/components/PrivateTable.vue";
import planningTradeTable from "@/components/PlanningTradeTable.vue";
import positionTable from "@/components/PositionTable.vue";

export default defineComponent({
  components: {
    privateTable,
    planningTradeTable,
    positionTable,
  },
  setup() {
    const router = useRouter();
    const route = useRoute();

    let loadingPlanningTrades = ref(false);
    let loadingPositions = ref(false);

    let PlanningTrades = ref([]);
    let Positions = ref([]);

    let displayPlanningTradeTable;
    let displayPositionTable;

    if (route.query.display === "true") {
      displayPlanningTradeTable = planningTradeTable;
    } else {
      displayPlanningTradeTable = privateTable;
    }

    if (route.query.display === "true") {
      displayPositionTable = positionTable;
    } else {
      displayPositionTable = privateTable;
    }

    const toHome = () => {
      router.push("/home");
    };

    onMounted(async () => {
      let url = route.path.split("/").slice(0, 3).join("/");

      loadingPlanningTrades.value = true;
      console.log(url + "/planningTrade");
      getPlanningTradesByStrategyID(url + "/planningTrade").then((Response) => {
        console.log(Response);
        loadingPlanningTrades.value = false;
        if (route.query.display === "true") {
          Response.data.data.planningTrades.forEach((element) => {
            PlanningTrades.value.push(element);
          });
        } else {
          console.log(Response);
          Response.data.data.planningTradesHash.forEach((element) => {
            PlanningTrades.value.push(element);
          });
        }
      });

      loadingPositions.value = true;
      console.log(url + "/position");
      getPositionsByStrategyID(url + "/position").then((Response) => {
        console.log(Response);
        loadingPositions.value = false;
        if (route.query.display === "true") {
          Response.data.data.positions.forEach((element) => {
            Positions.value.push(element);
          });
        } else {
          Response.data.data.positionsHash.forEach((element) => {
            Positions.value.push(element);
          });
        }
      });
    });

    return {
      loadingPlanningTrades,
      loadingPositions,
      PlanningTrades,
      Positions,
      toHome,
      displayPlanningTradeTable,
      displayPositionTable,
      privateTable,
    };
  },
});
</script>
