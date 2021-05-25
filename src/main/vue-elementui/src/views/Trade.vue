<template>
  <div>
    <el-container>
      <el-header> 交易记录 </el-header>
      <el-main>
        <!-- v-el-table-infinite-scroll="load" -->
        <el-table
          :data="Trades"
          v-elTableInfiniteScroll="load"
          v-loading="loading"
          border
          height="450"
          highlight-current-row
        >
          <el-table-column property="id" label="ID"></el-table-column>
          <el-table-column property="dateTime" label="日期"></el-table-column>
          <el-table-column property="stockID" label="股票"></el-table-column>
          <el-table-column
            property="commission"
            label="手续费"
          ></el-table-column>
          <el-table-column property="price" label="价格"></el-table-column>
          <el-table-column property="amount" label="数量"></el-table-column>
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
import { onMounted, defineComponent, ref, reactive } from "vue";
import { useRoute, useRouter } from "vue-router";
import { GetTradesPageByStrategyID } from "@/http/apis";

export default defineComponent({
  setup() {
    const router = useRouter();
    const route = useRoute();

    let Trades = ref([]);
    let loading = ref(false);

    let bookmark;
    let totalPage;
    let pageSize = 40;
    // let pageSize = reactive(0);
    // let totalPage = reactive([10, 20, 40]);
    onMounted(async () => {
      console.log(route.path);
      loading.value = true;
      GetTradesPageByStrategyID(route.path, { bookmark: bookmark }).then(
        (Response) => {
          loading.value = false;
          console.log(Response);
          bookmark = Response.data.data.bookmark;
          pageSize = Response.data.data.pageSize;
          totalPage = Response.data.data.totalPage;
          Response.data.data.data.forEach((element) => {
            Trades.value.push(element);
          });
        }
      );
    });

    const toHome = () => {
      router.push("/home");
    };

    const load = () => {
      loading.value = true;
      console.log("触发滚动加载");
      let params = {
        bookmark: bookmark,
        pageSize: 40,
      };
      console.log(params);
      GetTradesPageByStrategyID(route.path, params).then((Response) => {
        loading.value = false;
        console.log(Response);
        bookmark = Response.data.data.bookmark;
        pageSize = Response.data.data.pageSize;
        totalPage = Response.data.data.totalPage;
        Response.data.data.data.forEach((element) => {
          Trades.value.push(element);
        });
      });
    };

    return {
      toHome,
      Trades,
      totalPage,
      pageSize,
      loading,
      load,
    };
  },
});
</script>
