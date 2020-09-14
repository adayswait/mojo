<template>
  <div>
    <div>
      <iframe
        title="raft"
        width="100%"
        :height="iframeHeight"
        src="http://10.1.1.248:9500/raft_stat"
      ></iframe>
    </div>
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
