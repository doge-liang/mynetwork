<template>
  <div>
    <el-backtop :visibility-height="6"></el-backtop>
    <div v-for="item in strategys" :key="item.ID">
      <el-row>
        <el-col :span="6" :offset="7">
          <el-card class="productOwn-card" shadow="hover">
            <template #header>
              <div class="card-header">
                <span>{{ item.name }}</span>
                <el-button type="primary" round @click="check()"
                  >订阅策略</el-button
                >
              </div>
            </template>
            <div class="item">
              <span style="color: #409eff">策略ID: {{ item.ID }}</span>
              <span style="color: #e6a23c">夏普率:{{ item.sharpRatio }} </span>
            </div>
            <el-divider></el-divider>
            <div class="item">
              <span style="color: #f56c6c"
                >年化收益: {{ item.annualReturn }}</span
              >
              <span style="color: #f56c6c"
                >最大回撤: {{ item.maxDrawDown }}</span
              >
            </div>
            <div>
              <el-button type="text" @click="toMarketDetails(item.ID)"
                >市场信号</el-button
              >
              <el-divider direction="vertical" />
              <el-button type="text" @click="toTradeDetails(item.ID)"
                >操作信号</el-button
              >
            </div>
          </el-card>
        </el-col>
      </el-row>
    </div>
  </div>
</template>

<script>
export default {
  name: "Home",
  components: {},
  mounted() {
    this.loginState = sessionStorage.getItem("state");
    this.strategys = {}
  },
  data() {
    return {
      loginState: false,
      form: {
        name: "",
        password: "",
      },
      strategys: [
        {
          ID: "394501928",
          name: "ADXStrategy",
          annualReturn: 1.1897,
          sharpRatio: "11.61",
          maxDrawDown: "0.1025",
        },
        {
          ID: "1788534970",
          name: "CCICorrectionStrategy",
          annualReturn: 1.317,
          sharpRatio: "16.54",
          maxDrawDown: " 0.0795",
        },
        {
          ID: "1405007352",
          name: "GlobalExtremaStrategy",
          annualReturn: 1.1008,
          sharpRatio: "27.63",
          maxDrawDown: "0.0398",
        },
        {
          ID: "105200825",
          name: "MovingMomentumStrategy",
          annualReturn: 1.03617,
          sharpRatio: "27.336",
          maxDrawDown: "0.03790",
        },
        {
          ID: "909777763",
          name: "RSI2Strategy",
          annualReturn: 1.32955,
          sharpRatio: "23.71",
          maxDrawDown: "0.0560",
        },
      ],
    };
  },
  methods: {
    toMarketDetails(strategy) {
      if (sessionStorage.getItem("state")) {
        console.log(strategy);
        this.$message("操作成功!");
        this.$router.push({
          path: "/strategy/" + strategy + "/market",
        });
      } else {
        this.dialogFormVisible = true;
      }
    },
    toTradeDetails(strategy) {
      if (sessionStorage.getItem("state")) {
        console.log(strategy);
        this.$message("操作成功!");
        this.$router.push({
          path: "/strategy/" + strategy + "/trade",
        });
      } else {
        this.loginState = true;
      }
    },
    check() {
      if (sessionStorage.getItem("state")) {
        this.$message.success("订阅成功!");
      } else {
        this.loginState = true;
      }
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