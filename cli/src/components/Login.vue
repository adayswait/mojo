<template>
  <div class="column is-2 is-offset-5" v-if="visible">
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
      this.visible=false;
    }
  }
};
</script>
