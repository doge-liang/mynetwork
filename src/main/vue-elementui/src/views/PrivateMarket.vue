<template>
  <div>
    <el-container>
      <el-header>计划交易</el-header>
      <el-main>
        <el-table
          v-loading="loadingPlanningTrades"
          :data="planningTrades"
          highlight-current-row
        >
          <el-table-column
            property="id"
            label="id"
            width="500"
          ></el-table-column>
          <el-table-column
            property="hash"
            label="hashcode"
            width="500"
          ></el-table-column>
        </el-table>
      </el-main>
      <el-header>持仓信息</el-header>
      <el-main>
        <el-table
          v-loading="loadingPositions"
          :data="positions"
          highlight-current-row
        >
          <el-table-column
            property="id"
            label="id"
            width="500"
          ></el-table-column>
          <el-table-column
            property="hash"
            label="hashcode"
            width="500"
          ></el-table-column>
        </el-table>
      </el-main>
      <el-footer>
        <el-button
          type="primary"
          style="margin-top: 20px; margin-bottom: 45px"
          @click="toHome"
          >返回</el-button
        >
      </el-footer>
    </el-container>
  </div>
</template>

<script>
import { onMounted, ref } from "vue";
import { useRoute, useRouter } from "vue-router";
import {
  getPlanningTradesByStrategyID,
  getPositionsByStrategyID,
} from "@/http/apis";

// import { defineComponent } from "vue";

export default {
  name: "PrivateMarket",
  setup() {
    const router = useRouter();
    const route = useRoute();

    let loadingPlanningTrades = ref(false);
    let loadingPositions = ref(false);
    let planningTrades = ref([]);
    let positions = ref([]);

    const toHome = () => {
      router.push("/home");
    };

    onMounted(async () => {
      loadingPlanningTrades = true;
      getPlanningTradesByStrategyID(route.path).then((Response) => {
        loadingPlanningTrades.value = false;
        console.log(Response);
        planningTrades.value = Response.data.data.planningTrades;
      });

      loadingPositions = true;
      getPositionsByStrategyID(route.path).then((Response) => {
        loadingPositions.value = false;
        console.log(Response);
        positions.value = Response.data.data.positions;
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
};
</script>
