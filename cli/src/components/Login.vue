<template>
  <div class="column is-one-fifth-desktop is-offset-5" v-if="this.$store.state.visible.Login">
    <br />
    <br />
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
        <input class="input" type="text" placeholder="账号" v-model="usr" />
        <br />
        <br />
        <input class="input" type="text" placeholder="密码" v-model="passwd" />
        <br />
        <br />
        <a class="button is-primary is-fullwidth" @click="login">登陆</a>
        <br />
        <a class="button is-light is-fullwidth">注册</a>
        <br />
        <p class="help is-danger">{{this.err}}</p>
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
      loginBoxVisible: false,
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
        this.err = e.data || "network error";
        this.loginBoxVisible = true;
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

