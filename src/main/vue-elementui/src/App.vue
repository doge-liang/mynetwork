<template>
  <div id="app">
    <el-container ref="homePage">
      <el-header id="banner">
        <el-space>
          <div id="home-btn" @click="toHome">
            <i class="el-icon-s-home"> </i>
          </div>
        </el-space>
        <el-button v-if="isLogin" type="primary" @click="logout">
          注销
        </el-button>
        <el-button v-else type="primary" @click="toLogin">登录/注册</el-button>
      </el-header>
      <el-main>
        <router-view></router-view>
      </el-main>
    </el-container>
  </div>
</template>

<script>
import { onMounted, ref, reactive } from "vue";
import { useRouter } from "vue-router";
export default {
  setup() {
    const router = useRouter();
    const isLogin = sessionStorage.getItem("isLogin");

    onMounted(() => {
      document.title = "区块链智能投顾平台";
    });

    // TODO 发起注销请求
    const logout = () => {
      sessionStorage.setItem("isLogin", false);
    };

    const toLogin = () => {
      router.push({
        path: "/login",
      });
    };

    const toHome = () => {
      router.push({
        path: "/home",
      });
    };

    return {
      isLogin,
      logout,
      toHome,
    };
  },
};
</script>

<style>
#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
}
.el-header,
.el-footer {
  /* background-color: #409eff; */
  line-height: 60px;
  display: flex;
  font-size: x-large;
  justify-content: space-between;
}
.el-footer {
  justify-content: flex-end;
}
#banner.el-header {
  background-color: #409eff;
  line-height: 60px;
  display: flex;
  font-size: x-large;
  justify-content: space-between;
}
.el-aside {
  background-color: #d3dce6;
  color: #333;
  text-align: center;
  line-height: 200px;
}
.el-main {
  color: #333;
  text-align: center;
}
#home-btn {
  font-size: 30px;
}
body {
  margin: 0;
  padding: 0;
}
i {
  transition: all 0.5s;
}
.el-icon-s-home:hover {
  transition: all 0.5s;
  color: rgb(231, 231, 231);
}
</style>
