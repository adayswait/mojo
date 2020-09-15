<template>
  <div id="app">
    <Header />
    <Notification />
    <br />
    <br />
    <router-view />
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
<style scoped>
#app {
  height: 100%;
}
</style>;
