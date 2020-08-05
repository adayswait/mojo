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
import Header from "@/components/Header.vue";
import Notification from "@/components/Notification.vue";
// import Footer from "@/components/Footer.vue";

export default {
  name: "app",
  components: {
    Header,
    // Footer
    Notification,
  },
};
</script>
<style>
</style>;
