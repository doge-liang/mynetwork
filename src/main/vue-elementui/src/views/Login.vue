<template>
  <div id="login-wrapper">
    <el-container>
      <el-header> 区块链智能投顾平台 </el-header>
      <el-main>
        <el-form
          ref="ruleForm"
          status-icon
          :rules="rules"
          :label-position="labelPosition"
          :model="formModel"
          label-width="80px"
        >
          <el-form-item prop="username" label="账 号">
            <el-input
              size="medium"
              v-model="formModel.username"
              clearable
              placeholder="请输入账号"
            ></el-input>
          </el-form-item>

          <el-form-item prop="password" label="密 码">
            <el-input
              clearable
              show-password
              size="medium"
              v-model="formModel.password"
              placeholder="请输入密码"
            ></el-input>
          </el-form-item>

          <el-form-item prop="orgName" label="登录身份">
            <el-select v-model="formModel.orgName" placeholder="请选择登录身份">
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
import { reactive, ref, unref } from "vue";
import { ElMessage } from "element-plus";
import { useRouter } from "vue-router";

export default {
  // emits: ["isLogin"],
  setup(props, context) {
    let loading = false;
    const labelPosition = "left";
    const ruleForm = ref(null);
    const formModel = reactive({
      username: "",
      password: "",
      orgName: "Subscriber",
    });
    // 校验表单规则
    const rules = {
      username: [
        // FormItem标签中的 prop 属性预期值
        { required: true, message: "用户名不能为空", trigger: "blur" },
      ],
      password: [
        // FormItem标签中的 prop 属性预期值
        { required: true, message: "密码不能为空", trigger: "blur" },
      ],
    };
    const router = useRouter();

    const requestLogin = async () => {
      ruleForm.value.validate((valid) => {
        if (valid) {
          loading = true;
          let params = {
            userName: formModel.username,
            userSecret: formModel.password,
            orgName: formModel.orgName,
          };
          console.log(params);
          login(params)
            .then((resp) => {
              console.log(resp);
              if (resp.data.code === 200) {
                loading = false;
                sessionStorage.setItem("isLogin", true);
                ElMessage.success("登录成功");
                // context.emit("isLogin");
                router.push({
                  path: "/home",
                });
              }
            })
            .catch((err) => {
              console.log(err);
              loading = false;
              ElMessage.error("username or password wrong!");
            });
        } else {
          console.log("error submit!");
          return false;
        }
      });
    };

    const registe = async () => {
      ruleForm.validate((valid) => {
        if (valid) {
          loading = true;
          let params = {
            userName: formModel.username,
            userSecret: formModel.password,
            orgName: formModel.orgName,
          };
          console.log(params);
          register(params).then((resp) => {
            console.log(resp);
            if (resp.data.code === 200) {
              loading = false;
              sessionStorage.setItem("isLogin", true);
              ElMessage.success("登录成功");
            } else {
              loading = false;
              ElMessage.error("registration failed!");
            }
          });
        } else {
          console.log("error submit!");
          return false;
        }
      });
    };

    return {
      requestLogin,
      registe,
      labelPosition,
      loading,
      formModel,
      rules,
      ruleForm,
    };
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
  justify-content: space-between;
}
.el-header {
  display: block;
  text-align: center;
}
</style>
