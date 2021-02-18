<template>
  <div id="userLayout" class="user-layout-wrapper">
    <div class="container">
      <div class="top">
        <div class="desc">
          <img class="logo_login" src="@/assets/unicom_logo.png" alt="" />
        </div>
        <div class="header">
          <a href="/">
            <!-- <img src="~@/assets/logo.png" class="logo" alt="logo" /> -->
            <span class="title">区局建维任务填报系统</span>
          </a>
        </div>
      </div>
      <div class="main">
        <el-form
          :model="loginForm"
          :rules="rules"
          ref="loginForm"
          @keyup.enter.native="submitForm"
        >
          <el-form-item prop="username">
            <el-input placeholder="请输入用户名" v-model="loginForm.username">
              <i class="el-input__icon el-icon-user" slot="suffix"></i
            ></el-input>
          </el-form-item>
          <el-form-item prop="password">
            <el-input
              :type="lock === 'lock' ? 'password' : 'text'"
              placeholder="请输入密码"
              v-model="loginForm.password"
            >
              <i
                :class="'el-input__icon el-icon-' + lock"
                @click="changeLock"
                slot="suffix"
              ></i>
            </el-input>
          </el-form-item>
          <el-form-item style="position: relative">
            <el-input
              v-model="loginForm.captcha"
              name="logVerify"
              placeholder="请输入验证码"
              style="width: 60%"
            />
            <div class="vPic">
              <img
                v-if="picPath"
                :src="picPath"
                width="100%"
                height="100%"
                alt="请输入验证码"
                @click="loginVefify()"
              />
            </div>
          </el-form-item>
          <el-form-item>
            <el-button type="primary" @click="submitForm" style="width: 100%"
              >登 录</el-button
            >
          </el-form-item>
        </el-form>
      </div>

      <div class="footer">

        <div class="copyright">Copyright &copy; {{ curYear }} 云网中心数据网室🐂ChinaUnicom</div>
      </div>
    </div>
  </div>
</template>

<script>
import { mapActions } from "vuex";
import { captcha } from "@/api/user";
export default {
  name: "Login",
  data() {
    const checkUsername = (rule, value, callback) => {
      if (value.length < 5 || value.length > 12) {
        return callback(new Error("请输入正确的用户名"));
      } else {
        callback();
      }
    };
    const checkPassword = (rule, value, callback) => {
      if (value.length < 6 || value.length > 12) {
        return callback(new Error("请输入正确的密码"));
      } else {
        callback();
      }
    };
    return {
      curYear: 0,
      lock: "lock",
      loginForm: {
        username: "",
        password: "",
        captcha: "",
        captchaId: "",
      },
      rules: {
        username: [{ validator: checkUsername, trigger: "blur" }],
        password: [{ validator: checkPassword, trigger: "blur" }],
      },
      logVerify: "",
      picPath: "",
    };
  },
  created() {
    this.loginVefify();
    this.curYear = new Date().getFullYear();
  },
  methods: {
    
    ...mapActions("user", ["LoginIn"]),
    async loginLog(){
      console.assert("登陆成功！")
    },
    async login() {
      return await this.LoginIn(this.loginForm);
    },
    async submitForm() {
      this.$refs.loginForm.validate(async (v) => {
        if (v) {
          const flag = await this.login();
          if (!flag) {
            this.loginVefify();
          }
        } else {
          this.$message({
            type: "error",
            message: "请正确填写登录信息",
            showClose: true,
          });
          this.loginVefify();
          return false;
        }
      });
    },
    changeLock() {
      this.lock === "lock" ? (this.lock = "unlock") : (this.lock = "lock");
    },
    loginVefify() {
      captcha({}).then((ele) => {
        this.picPath = ele.data.picPath;
        this.loginForm.captchaId = ele.data.captchaId;
      });
    },
  },
};
</script>

<style scoped lang="scss">
  .container{
    background-color: #00b7ff;
    background-image: url("data:image/svg+xml,%3Csvg xmlns='http://www.w3.org/2000/svg' width='540' height='450' viewBox='0 0 1080 900'%3E%3Cg fill-opacity='.1'%3E%3Cpolygon fill='%23444' points='90 150 0 300 180 300'/%3E%3Cpolygon points='90 150 180 0 0 0'/%3E%3Cpolygon fill='%23AAA' points='270 150 360 0 180 0'/%3E%3Cpolygon fill='%23DDD' points='450 150 360 300 540 300'/%3E%3Cpolygon fill='%23999' points='450 150 540 0 360 0'/%3E%3Cpolygon points='630 150 540 300 720 300'/%3E%3Cpolygon fill='%23DDD' points='630 150 720 0 540 0'/%3E%3Cpolygon fill='%23444' points='810 150 720 300 900 300'/%3E%3Cpolygon fill='%23FFF' points='810 150 900 0 720 0'/%3E%3Cpolygon fill='%23DDD' points='990 150 900 300 1080 300'/%3E%3Cpolygon fill='%23444' points='990 150 1080 0 900 0'/%3E%3Cpolygon fill='%23DDD' points='90 450 0 600 180 600'/%3E%3Cpolygon points='90 450 180 300 0 300'/%3E%3Cpolygon fill='%23666' points='270 450 180 600 360 600'/%3E%3Cpolygon fill='%23AAA' points='270 450 360 300 180 300'/%3E%3Cpolygon fill='%23DDD' points='450 450 360 600 540 600'/%3E%3Cpolygon fill='%23999' points='450 450 540 300 360 300'/%3E%3Cpolygon fill='%23999' points='630 450 540 600 720 600'/%3E%3Cpolygon fill='%23FFF' points='630 450 720 300 540 300'/%3E%3Cpolygon points='810 450 720 600 900 600'/%3E%3Cpolygon fill='%23DDD' points='810 450 900 300 720 300'/%3E%3Cpolygon fill='%23AAA' points='990 450 900 600 1080 600'/%3E%3Cpolygon fill='%23444' points='990 450 1080 300 900 300'/%3E%3Cpolygon fill='%23222' points='90 750 0 900 180 900'/%3E%3Cpolygon points='270 750 180 900 360 900'/%3E%3Cpolygon fill='%23DDD' points='270 750 360 600 180 600'/%3E%3Cpolygon points='450 750 540 600 360 600'/%3E%3Cpolygon points='630 750 540 900 720 900'/%3E%3Cpolygon fill='%23444' points='630 750 720 600 540 600'/%3E%3Cpolygon fill='%23AAA' points='810 750 720 900 900 900'/%3E%3Cpolygon fill='%23666' points='810 750 900 600 720 600'/%3E%3Cpolygon fill='%23999' points='990 750 900 900 1080 900'/%3E%3Cpolygon fill='%23999' points='180 0 90 150 270 150'/%3E%3Cpolygon fill='%23444' points='360 0 270 150 450 150'/%3E%3Cpolygon fill='%23FFF' points='540 0 450 150 630 150'/%3E%3Cpolygon points='900 0 810 150 990 150'/%3E%3Cpolygon fill='%23222' points='0 300 -90 450 90 450'/%3E%3Cpolygon fill='%23FFF' points='0 300 90 150 -90 150'/%3E%3Cpolygon fill='%23FFF' points='180 300 90 450 270 450'/%3E%3Cpolygon fill='%23666' points='180 300 270 150 90 150'/%3E%3Cpolygon fill='%23222' points='360 300 270 450 450 450'/%3E%3Cpolygon fill='%23FFF' points='360 300 450 150 270 150'/%3E%3Cpolygon fill='%23444' points='540 300 450 450 630 450'/%3E%3Cpolygon fill='%23222' points='540 300 630 150 450 150'/%3E%3Cpolygon fill='%23AAA' points='720 300 630 450 810 450'/%3E%3Cpolygon fill='%23666' points='720 300 810 150 630 150'/%3E%3Cpolygon fill='%23FFF' points='900 300 810 450 990 450'/%3E%3Cpolygon fill='%23999' points='900 300 990 150 810 150'/%3E%3Cpolygon points='0 600 -90 750 90 750'/%3E%3Cpolygon fill='%23666' points='0 600 90 450 -90 450'/%3E%3Cpolygon fill='%23AAA' points='180 600 90 750 270 750'/%3E%3Cpolygon fill='%23444' points='180 600 270 450 90 450'/%3E%3Cpolygon fill='%23444' points='360 600 270 750 450 750'/%3E%3Cpolygon fill='%23999' points='360 600 450 450 270 450'/%3E%3Cpolygon fill='%23666' points='540 600 630 450 450 450'/%3E%3Cpolygon fill='%23222' points='720 600 630 750 810 750'/%3E%3Cpolygon fill='%23FFF' points='900 600 810 750 990 750'/%3E%3Cpolygon fill='%23222' points='900 600 990 450 810 450'/%3E%3Cpolygon fill='%23DDD' points='0 900 90 750 -90 750'/%3E%3Cpolygon fill='%23444' points='180 900 270 750 90 750'/%3E%3Cpolygon fill='%23FFF' points='360 900 450 750 270 750'/%3E%3Cpolygon fill='%23AAA' points='540 900 630 750 450 750'/%3E%3Cpolygon fill='%23FFF' points='720 900 810 750 630 750'/%3E%3Cpolygon fill='%23222' points='900 900 990 750 810 750'/%3E%3Cpolygon fill='%23222' points='1080 300 990 450 1170 450'/%3E%3Cpolygon fill='%23FFF' points='1080 300 1170 150 990 150'/%3E%3Cpolygon points='1080 600 990 750 1170 750'/%3E%3Cpolygon fill='%23666' points='1080 600 1170 450 990 450'/%3E%3Cpolygon fill='%23DDD' points='1080 900 1170 750 990 750'/%3E%3C/g%3E%3C/svg%3E");
  }
@import "@/style/login.scss";
</style>
