<template>
  <div>
    <Login />
    <div
      style="position: fixed; border:15px solid red; width: 100%; height: 100%;"
      v-if="!$store.state.visible.Login"
    >
      <p class="alertleft">外网正式服 小心操作</p>
      <p class="alertright">外网正式服 小心操作</p>
      <p class="alertbottom">😶</p>
      <div>
        <iframe title="GM" width="100%" height="850px" src="http://10.1.1.248:8001"></iframe>
      </div>
    </div>
  </div>
</template>

<script>
import Login from "@/pages/Login.vue";
export default {
  components: {
    Login,
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
  width:100%;
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
