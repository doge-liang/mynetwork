<template>
  <div id="login-wrapper">
    <el-container>
      <el-header> 区块链智能投顾平台 </el-header>
      <el-main>
        <el-form
          ref="form"
          status-icon
          :rules="rulesLogin"
          :label-position="labelPosition"
          :model="form"
          label-width="80px"
        >
          <el-form-item prop="username" label="账 号">
            <el-input
              size="medium"
              v-model="form.username"
              clearable
              placeholder="请输入账号"
            ></el-input>
          </el-form-item>

          <el-form-item prop="password" label="密 码">
            <el-input
              clearable
              show-password
              size="medium"
              v-model="form.password"
              placeholder="请输入密码"
            ></el-input>
          </el-form-item>

          <el-form-item prop="orgName" label="登录身份">
            <el-select v-model="form.orgName" placeholder="请选择登录身份">
              <el-option label="策略发布者" value="Provider"></el-option>
              <el-option label="策略订阅者" value="Subscriber"></el-option>
            </el-select>
          </el-form-item>
        </el-form>
      </el-main>
      <el-footer height="40px" id="login-footer">
        <el-button @click="requestLogin" type="primary" :loading="loading"
          >登 录</el-button
        >
        <el-button @click="registe" :loading="loading">注 册</el-button>
      </el-footer>
    </el-container>
  </div>
</template>

<script>
import { login } from "@/http/apis";
import { register } from "@/http/apis";

export default {
  data() {
    return {
      labelPosition: "left",
      loading: false,
      form: {
        username: "",
        password: "",
        orgName: "Subscriber",
      },
      // 校验表单规则
      rulesLogin: {
        username: [
          // FormItem标签中的 prop 属性预期值
          { required: true, message: "用户名不能为空", trigger: "blur" },
        ],
        password: [
          // FormItem标签中的 prop 属性预期值
          { required: true, message: "密码不能为空", trigger: "blur" },
        ],
      },
      //   formLabelWidth: "120px",
    };
  },
  methods: {
    requestLogin(event) {
      // this.isLogin = true;
      // this.$message.success("登录成功");
      // sessionStorage.setItem("isLogin", this.isLogin);
      // this.$router.push({
      //   path: "/home",
      // });

      this.$refs.form.validate((valid) => {
        if (valid) {
          this.loading = true;
          let params = {
            userName: this.form.username,
            userSecret: this.form.password,
            orgName: this.form.orgName,
          };
          console.log(params);
          login(params)
            .then((resp) => {
              console.log(resp);
              if (resp.data.code === 200) {
                this.loading = false;
                this.isLogin = true;
                sessionStorage.setItem("isLogin", this.isLogin);
                this.$message.success("登录成功");
                this.$router.push({
                  path: "/home",
                });
              }
            })
            .catch((err) => {
              console.log(err);
              this.loading = false;
              this.$alert("username or password wrong!", "info", {
                confirmButtonText: "ok",
              });
            });
        } else {
          console.log("error submit!");
          return false;
        }
      });
    },
    
    registe(event) {
      let username = this.form.username;
      let password = this.form.password;
      this.$refs.form.validate((valid) => {
        if (valid) {
          this.loading = true;
          let params = {
            userName: username,
            userSecret: password,
          };
          console.log(params);
          register(params).then((resp) => {
            console.log(resp);
            if (resp.data.code === 200) {
              this.loading = false;
              this.isLogin = true;
              sessionStorage.setItem("isLogin", this.isLogin);
              this.$message.success("登录成功");
              this.closeDialog();
            } else {
              this.loading = false;
              this.$alert("registration failed!", "info", {
                confirmButtonText: "ok",
              });
            }
          });
        } else {
          console.log("error submit!");
          return false;
        }
      });
    },
  },
};
</script>

<style>
#login-wrapper {
  margin: 7% 35%;
  padding: 1%;
  border-radius: 4px;
  box-shadow: 0 2px 4px rgba(0, 0, 0, 0.12), 0 0 6px rgba(0, 0, 0, 0.04);
}
#login-footer {
  padding: 0% 20%;
  margin: 2% 0%;
}
.el-header {
  display: block;
  text-align: center;
}
</style>
