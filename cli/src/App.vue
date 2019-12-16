<template>
  <div id="app">
    <Header />
    <br />
    <router-view />
    <br />
    <Footer />
  </div>
</template>
<script>
import Header from "@/components/Header.vue";
import Footer from "@/components/Footer.vue";
export default {
  name: "app",
  components: {
    Header,
    Footer
  }
};
</script>
<style></style>;
