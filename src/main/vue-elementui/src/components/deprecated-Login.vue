<template>
  <div>
    <el-dialog
      title="登  录"
      v-model="dialogFormVisible"
      :before-close="closeDialog"
      width="30%"
    >
      <el-form ref="form" :model="form" status-icon :rules="rulesLogin">
        <el-form-item
          label="账号"
          :label-width="formLabelWidth"
          prop="username"
        >
          <el-input
            type="text"
            v-model="form.username"
            clearable
            placeholder="请输入账号"
          ></el-input>
        </el-form-item>
        <el-form-item
          label="密码"
          :label-width="formLabelWidth"
          prop="password"
        >
          <el-input
            type="password"
            v-model="form.password"
            clearable
            placeholder="请输入密码"
            show-password
          ></el-input>
        </el-form-item>
      </el-form>
      <template #footer>
        <span class="dialog-footer">
          <el-button @click="registe" :loading="loading">注 册</el-button>
          <el-button @click="requestLogin" type="primary" :loading="loading"
            >确 定</el-button
          >
        </span>
      </template>
    </el-dialog>
  </div>
</template>

<script>
import { login } from "@/http/apis";
import { register } from "@/http/apis";

export default {
  name: "login-dialog",
  props: {
    dialogFormVisible: Boolean,
  },

  data() {
    return {
      loading: false,
      form: {
        username: "",
        password: "",
      },
      rulesLogin: {
        // 校验表单规则
        username: [
          // FormItem标签中的 prop 属性预期值
          { required: true, message: "用户名不能为空", trigger: "blur" },
        ],
        password: [
          // FormItem标签中的 prop 属性预期值
          { required: true, message: "密码不能为空", trigger: "blur" },
        ],
      },
      formLabelWidth: "120px",
    };
  },
  methods: {
    closeDialog() {
      this.$emit("closeDialog");
      this.form.username = "";
      this.form.password = "";
    },
    requestLogin(event) {
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
          login(params)
            .then((resp) => {
              console.log(resp);
              if (resp.data.code === 200) {
                this.loading = false;
                this.isLogin = true;
                sessionStorage.setItem("isLogin", this.isLogin);
                this.$message.success("登录成功");
                this.closeDialog();
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
  mounted() {},
};
</script>
<style scoped>
</style>