<template>
  <div id="app">
    <Header />
    <Login />
    <Home />
  </div>
</template>

<script>
import Header from "./components/Header.vue";
import Login from "./components/Login.vue";
import Home from "./components/Home.vue";

export default {
  name: "app",
  components: {
    Header,
    Login,
    Home
  }
};
</script>

<style>
</style>
