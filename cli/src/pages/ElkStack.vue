<template>
  <div class="column">
    <Login />
    <router-view id="router" />
  </div>
</template>

<script>
import Login from "@/pages/Login.vue";
export default {
  data: function () {
    return {};
  },
  components: {
    Login,
  },
};
</script>