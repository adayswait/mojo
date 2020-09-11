<template>
  <div>
    <iframe title="GM239" width="100%" :height="iframeHeight" :src="srcURL"></iframe>
  </div>
</template>

<script>
export default {
  data() {
    return {
      iframeHeight: 0,
      srcURL: "",
    };
  },
  mounted() {
    //bulma navbar的默认高度为3.25rem 换算下来是52px
    this.iframeHeight = document.documentElement.clientHeight - 60;
  },
  beforeMount() {
    this.srcURL = this.$route.query.srcurl;
    window.console.log(this.$route.query.srcurl);
  },
};
</script>

<style scoped>
</style>
