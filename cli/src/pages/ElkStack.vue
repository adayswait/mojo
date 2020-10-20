<template>
  <div class="column">
    <Login />
    <router-view id="router" />
  </div>
</template>

<script>
const Login = () => import("@/pages/Login.vue");
export default {
  data: function () {
    return {};
  },
  components: {
    Login,
  },
};
</script>

<style scoped>
.column {
  padding: 0;
}
div {
  height: 90%;
}
#router {
  height: 100%;
}
</style>