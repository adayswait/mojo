<template>
  <div>
    <Login />
    <div
      style="position: fixed; border-left:10px solid red;border-right:10px solid red;border-top:18px solid red; width: 100%; height: auto;"
      v-if="!$store.state.visible.Login"
    >
      <p class="alertleft">外网正式服 小心操作</p>
      <p class="alertright">外网正式服 小心操作</p>
      <p class="alertbottom">😶</p>
      <div>
        <iframe title="GM" width="100%" :height="iframeHeight" src="http://10.1.1.248:8001"></iframe>
      </div>
    </div>
  </div>
</template>

<script>
const Login = () => import("@/pages/Login.vue");
export default {
  data() {
    return {
      iframeHeight: 0,
    };
  },
  components: {
    Login,
  },
  mounted() {
    //bulma navbar的默认高度为3.25rem 换算下来是52px
    this.iframeHeight = document.documentElement.clientHeight - 60;
  },
};
</script>

<style scoped>
.alertleft {
  top: 15%;
  left: 5%;
  writing-mode: vertical-rl;
}
.alertright {
  top: 15%;
  right: 5%;
  writing-mode: vertical-rl;
}
.alertbottom {
  width: 100%;
  left: 50%;
  bottom: 10%;
  margin-left: -50%;
}
.alertleft,
.alertright,
.alertbottom {
  font-size: 30px;
  position: fixed;
  text-align: center;
  font-weight: bold;
  /* background: linear-gradient(to right, red, blue);
  background-clip: text; */
  color: red;
  animation: blink 1s linear infinite;
  pointer-events: none;
}
@keyframes blink {
  0% {
    opacity: 0;
  }
  50% {
    opacity: 1;
  }
  100% {
    opacity: 0;
  }
}
</style>
