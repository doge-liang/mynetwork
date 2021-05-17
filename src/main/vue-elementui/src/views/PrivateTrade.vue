<template>
  <div>
    <el-container>
      <el-header> 交易记录 </el-header>
      <el-main>
        <el-table :data="Trades" highlight-current-row>
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
export default {
  name: "PrivateTrade",
  mounted() {},
  data() {
    return {
      Trades: [{
        hash: "02c22341c22e6d4379fe716f51999638db9d63c504259ed64fdff6afcd65b6f3",
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
