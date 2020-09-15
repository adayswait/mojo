<template>
  <div>
    <div class="box" v-if="!($store.getters.userInfo.group<=1)">
      <NoRights height="90%" width="100%" />
    </div>
    <div v-if="!$store.state.visible.Login && $store.getters.userInfo.group<=1">
      <iframe
        id="ikibana"
        title="kibana"
        width="100%"
        :height="iframeHeight"
        frameborder="0"
        src="http://10.1.1.248:8002/app/discover"
      ></iframe>
    </div>
  </div>
</template>

<script>
import NoRights from "@/pages/ErrorPage/NoRights.vue";
export default {
  components: {
    NoRights,
  },
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
.box {
  height: 100%;
}
</style>
