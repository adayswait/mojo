<template>
  <div id="app">
    <Header />
    <br />
    <Notification />
    <br />
    <br />
    <router-view />
    <br />
    <Footer />
  </div>
</template>
<script>
import Header from "@/pages/Header.vue";
import Notification from "@/pages/Notification.vue";
// import Footer from "@/pages/Footer.vue";

export default {
  name: "app",
  components: {
    Header,
    // Footer,
    Notification,
  },
};
</script>
<style>
/* @font-face {
  font-family: FZHeiLJW;
  src: url("/FZHeiLJW.TTF");
}
#app {
  font-family: FZHeiLJW;
} */
</style>;
