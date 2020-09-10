<template>
  <div v-if="!$store.state.visible.Login">
    <iframe
      id="ikibana"
      title="kibana"
      width="100%"
      :height="iframeHeight"
      frameborder="0"
      src="http://10.1.1.248:8002/app/logs"
    ></iframe>
  </div>
</template>

<script>
export default {
  data() {
    return {
      iframeHeight: 0,
    };
  },
  mounted() {
    //bulma navbar的默认高度为3.25rem 换算下来是52px
    this.iframeHeight = document.documentElement.clientHeight - 60;
  },
};
</script>

<style scoped>
.iframe {
  bottom: 0px;
}
</style>
