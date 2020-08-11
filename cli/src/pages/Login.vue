<template>
  <div class="column is-one-fifth-desktop is-offset-5" v-if="this.$store.state.visible.Login">
    <br />
    <br />
    <br />
    <div class="box">
      <div v-if="!this.loginBoxVisible">
        <br />
        <br />
        <br />
        <div class="logining">
          <div class="bounce bounce1"></div>
          <div class="bounce bounce2"></div>
        </div>
        <br />
        <br />
        <br />
      </div>
      <div v-if="this.loginBoxVisible">
        <br />
        <input class="input is-primary" type="email" placeholder="账号" v-model="usr" />
        <br />
        <br />
        <input class="input is-danger" type="password" placeholder="密码" v-model="passwd" />
        <br />
        <br />
        <a class="button is-primary is-fullwidth" @click="login" v-if="this.loginButtonVisible">登陆</a>
        <div v-if="!this.loginButtonVisible">
          <input class="input is-danger" type="password" placeholder="再次输入密码" v-model="repasswd" />
          <br />
        </div>
        <br />
        <a class="button is-warning is-fullwidth" @click="register">注册</a>
        <br />
        <strong>
          <p class="help is-danger">{{this.err}}</p>
        </strong>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  name: "Login",
  data: function () {
    this.login();
    return {
      title: "Login",
      usr: "",
      passwd: "",
      repasswd: "",
      loginBoxVisible: false,
      loginButtonVisible: true,
      err: "",
    };
  },
  methods: {
    login: async function () {
      try {
        this.loginBoxVisible = false;
        const ret = await this.$httpc.get("/web/auth/login", {
          user: this.usr || undefined,
          passwd: this.passwd || undefined,
        });
        ret.data = JSON.parse(ret.data);
        this.$store.commit("setUserInfo", {
          user: ret.data.user,
          group: ret.data.group,
        });
        this.$store.commit("setVisible", {
          name: "Login",
          visible: false,
        });
        this.$store.commit("setVisible", {
          name: "Home",
          visible: true,
        });
        this.$store.commit("setVisible", {
          name: "UsrInfo",
          visible: true,
        });
      } catch (e) {
        this.err = e.data || "network error, try later";
        this.loginBoxVisible = true;
      }
    },
    register: async function () {
      try {
        if (this.loginButtonVisible === true) {
          this.loginButtonVisible = false;
          return;
        }
        if (this.usr.length === 0) {
          this.err = "username can't be empty";
          return;
        }
        if (this.passwd.length < 6) {
          this.err = "password length < 6";
          return;
        }
        if (this.passwd !== this.repasswd) {
          this.err = "ensure your password is the same";
          return;
        }
        await this.$httpc.get("/web/auth/register", {
          user: this.usr,
          passwd: this.passwd,
        });
        this.login();
      } catch (e) {
        this.err = e.data || "register error, try later";
      }
    },
  },
};
</script>

<style scoped>
.logining {
  width: 50px;
  height: 50px;
  margin: 50px auto;
  position: relative;
}
.bounce {
  position: absolute;
  top: 0;
  left: 0;
  width: 100%;
  height: 100%;
  border-radius: 50%;
  background-color: #67cf22;
  opacity: 0.6;
  -webkit-animation: bounce 2s infinite ease-in-out;
  animation: bounce 2s infinite ease-in-out;
}
.bounce2 {
  -webkit-animation-delay: -1s;
  animation-delay: -1s;
}
@keyframes bounce {
  0% {
    transform: scale(0);
  }
  50% {
    transform: scale(1);
  }
  100% {
    transform: scale(0);
  }
}
@-webkit-keyframes bounce {
  0% {
    transform: scale(0);
  }
  50% {
    transform: scale(1);
  }
  100% {
    transform: scale(0);
  }
}
</style>

