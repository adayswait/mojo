<template>
  <div class="column is-2 is-offset-5" v-if="this.$store.state.visible.Login">
    <br />
    <br />
    <br />
    <div class="box">
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
      无脑点登录
    </div>
  </div>
</template>

<script>
export default {
  name: "Login",
  data: function() {
    return {
      title: "Login",
      usr: "",
      passwd: "",
      visible: true
    };
  },
  methods: {
    login: function() {
      this.$store.commit("setVisible", {
        name: "Login",
        visible: false
      });
      this.$store.commit("setVisible", {
        name: "Home",
        visible: true
      });
      this.$store.commit("setVisible", {
        name: "UsrInfo",
        visible: true
      });
    }
  }
};
</script>
