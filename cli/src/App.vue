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
const Header = () => import("@/pages/Header.vue");
const Notification = () => import("@/pages/Notification.vue");
export default {
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
