<template>
  <div id="app">
    <el-container ref="homePage">
      <el-header>
        <el-space>
          <div class="button">
            <i class="el-icon-s-home"></i>
          </div>
        </el-space>
        <el-button type="primary" @click="check">登录/注册</el-button>
      </el-header>
      <el-main>
        <login-diaglog
          :dialogFormVisible="showLogin"
          @closeDialog="this.showLogin = false"
        ></login-diaglog>
        <router-view></router-view>
      </el-main>
    </el-container>
  </div>
</template>

<script>
import LoginDialog from "./components/Login.vue";
export default {
  name: "App",
  components: {
    "login-diaglog": LoginDialog,
  },
  data() {
    return {
      clientHeight: "",
      loginState: false,
      showLogin: false,
      form: {
        name: "",
        password: "",
      },
    };
  },
  mounted() {
    // 获取浏览器可视区域高度
    this.clientHeight = `${document.documentElement.clientHeight}`;
    //document.body.clientWidth;
    //console.log(self.clientHeight);
    window.onresize = function temp() {
      this.clientHeight = `${document.documentElement.clientHeight}`;
    };
    document.title = "区块链量化投顾平台";
  },
  watch: {
    // 如果 `clientHeight` 发生改变，这个函数就会运行
    clientHeight: function () {
      this.changeFixed(this.clientHeight);
    },
  },
  methods: {
    changeFixed(clientHeight) {
      //动态修改样式
      // console.log(clientHeight);
      // console.log(this.$refs.homePage.$el.style.height);
      this.$refs.homePage.$el.style.height = clientHeight - 20 + "px";
    },
    check() {
        this.showLogin = true;
    },
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
.button {
  font-size: 50px;
}
</style>
