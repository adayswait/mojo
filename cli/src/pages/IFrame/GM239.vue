<template>
  <div>
    <iframe
      title="GM239"
      width="100%"
      :height="iframeHeight"
      :src="`http://s.61.com/?ip=123.206.181.31&port=9999&log=true&showPP=false`"
    >
    </iframe>
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
</style>
