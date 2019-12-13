<template>
  <div id="app">
    <Header />
    <Login />
  </div>
</template>

<script>
import Header from "./components/Header.vue";
import Login from "./components/Login.vue";


export default {
  name: "app",
  components: {
    Header,
    Login
  }
};
</script>

<style>
</style>
