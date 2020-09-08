<template>
  <div class="container">
    <iframe title="kibana" width="100%" height="850px" src="http://10.1.1.248:8002/app/logs"></iframe>
  </div>
</template>

<script>
export default {
  name: "Data",
  data: function () {
    return {
      title: "Data",
    };
  },
};
</script>

<style scoped>
iframe {
  margin-top: -20px;
}
</style>
