<template>
  <div id="app">
    <Header />
    <br/>
    <Login />
    <Home />
    <Footer/>
  </div>
</template>

<script>
import Header from "./components/Header.vue";
import Login from "./components/Login.vue";
import Home from "./components/Home.vue";
import Footer from "./components/Footer.vue"
export default {
  name: "app",
  components: {
    Header,
    Login,
    Home,
    Footer
  }
};
</script>

<style>
#app {
  background-color: #ffffff;
}
</style>
