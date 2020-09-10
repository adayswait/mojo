<template>
  <div id="devopsPage">
    <Login />

    <div class="container" v-if="this.$store.state.visible.Home">
      <div class="columns">
        <div class="assistive-wrap" v-if="!menuVisible" @click="showMenuOnMobile">
          <div class="assistive-touch">
            <span></span>
          </div>
        </div>
        <div class="column is-2" v-if="menuVisible">
          <div id="devopsmenu" class="box">
            <aside class="menu">
              <p class="menu-label">HOME</p>
              <ul class="menu-list">
                <li @click="focus('主页')">
                  <router-link to="/devops/home" :class="{'is-active':currFocus=='主页'}">主页</router-link>
                </li>
              </ul>
              <p class="menu-label">CI/CD</p>
              <ul class="menu-list">
                <li>
                  <a>部署</a>
                  <!-- <a class="is-active">部署</a> -->
                  <ul>
                    <li @click="focus('快捷操作')">
                      <router-link
                        to="/devops/quickops"
                        :class="{'is-active':currFocus=='快捷操作'}"
                      >快捷操作</router-link>
                    </li>
                    <li @click="focus('新建部署')">
                      <router-link to="/devops/newdep" :class="{'is-active':currFocus=='新建部署'}">新建部署</router-link>
                    </li>

                    <li @click="focus('所有部署')">
                      <router-link to="/devops/alldep" :class="{'is-active':currFocus=='所有部署'}">所有部署</router-link>
                    </li>
                  </ul>
                </li>
              </ul>
              <p class="menu-label">Configuration</p>
              <ul class="menu-list">
                <li>
                  <a>项目配置</a>
                  <ul>
                    <li @click="focus('机器配置')">
                      <router-link to="/devops/macini" :class="{'is-active':currFocus=='机器配置'}">机器配置</router-link>
                    </li>
                    <li @click="focus('部署配置')">
                      <router-link to="/devops/depini" :class="{'is-active':currFocus=='部署配置'}">部署配置</router-link>
                    </li>
                    <li @click="focus('开发配置')">
                      <router-link to="/devops/devini" :class="{'is-active':currFocus=='开发配置'}">开发配置</router-link>
                    </li>
                  </ul>
                </li>
              </ul>
              <p class="menu-label">Management</p>
              <ul class="menu-list">
                <li>
                  <a>成员管理</a>
                  <ul>
                    <li @click="focus('所有成员')">
                      <router-link
                        to="/devops/manageuser"
                        :class="{'is-active':currFocus=='所有成员'}"
                      >所有成员</router-link>
                    </li>
                  </ul>
                </li>
              </ul>
              <p class="menu-label">Tools</p>
              <ul class="menu-list">
                <li>
                  <a>工具箱</a>
                  <ul>
                    <li @click="focus('DataView')">
                      <router-link
                        to="/devops/dbview"
                        :class="{'is-active':currFocus=='DataView'}"
                      >数据视图</router-link>
                    </li>
                    <li @click="focus('DevTools')">
                      <router-link
                        to="/devops/devtools"
                        :class="{'is-active':currFocus=='DevTools'}"
                      >开发工具</router-link>
                    </li>
                    <li @click="focus('Issue')">
                      <router-link
                        to="/devops/issue"
                        :class="{'is-active':currFocus=='Issue'}"
                      >Issue</router-link>
                    </li>
                  </ul>
                </li>
              </ul>
            </aside>
          </div>
        </div>
        <div class="column" id="devopsBody" ref="devopsBody">
          <div id="mask" v-if="this.$store.state.DevOpsMask!=0"></div>
          <div id="loading" v-if="this.$store.state.DevOpsMask!=0">
            <div class="loading3">
              <div class="circle circle1">
                <span></span>
                <span></span>
                <span></span>
                <span></span>
              </div>
              <div class="circle circle2">
                <span></span>
                <span></span>
                <span></span>
                <span></span>
              </div>
              <div class="circle circle3">
                <span></span>
                <span></span>
                <span></span>
                <span></span>
              </div>
            </div>
          </div>
          <router-view id="router" />
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import Login from "@/pages/Login.vue";
export default {
  name: "DevOps",
  data: function () {
    return {
      currFocus: "主页",
      timer: null,
      menuVisible: true,
    };
  },
  methods: {
    focus: function (target) {
      this.currFocus = target;
      this.hideMenuOnMobile();
      if (!this.$mojoapi.isMobile()) {
        this.$refs.devopsBody &&
          this.$refs.devopsBody.scrollIntoView({
            behavior: "smooth",
            block: "start", //垂直方向
            inline: "center", //水平方向
          });
      }
    },
    hideMenuOnMobile: function () {
      if (this.$mojoapi.isMobile()) {
        this.menuVisible = false;
      }
    },
    showMenuOnMobile: function () {
      if (this.$mojoapi.isMobile()) {
        this.menuVisible = true;
      }
    },
  },
  mounted: function () {
    if (this.$mojoapi.isMobile()) {
      this.hideMenuOnMobile();
    }
  },
  components: {
    Login,
  },
};
</script>


<style scoped>
@font-face {
  font-family: consolas;
  src: url("/consola.ttf");
}

#devopsmenu {
  background-color: #f1f1f1;
}
#devopsPage {
  margin-top: 20px;
  font-family: consolas;
}

#devopsBody {
  min-height: 500px;
}

#mask {
  position: absolute;
  width: 100%;
  height: 100%;
  z-index: 98;
  /* filter: alpha(Opacity=15);
  -moz-opacity: 0.15;
  opacity: 0.15;
  background-color: #cdc9c9; */
}
.loading3 {
  width: 30px;
  height: 30px;
  margin: 25% 39%;
  z-index: 99;
  position: absolute;
}
.circle {
  position: absolute;
  top: 0;
  left: 0;
  width: 100%;
  height: 100%;
}
.circle span {
  width: 8px;
  height: 8px;
  display: inline-block;
  background: #1c1c1c;
  border-radius: 100%;
  position: absolute;
  -webkit-animation: mycircle 1.2s infinite ease-in-out;
  animation: mycircle 1.2s infinite ease-in-out;
  -webkit-animation-fill-mode: both;
  animation-fill-mode: both;
}
.circle2 {
  -webkit-transform: rotateZ(45deg);
  transform: rotateZ(45deg);
}
.circle3 {
  -webkit-transform: rotateZ(90deg);
  transform: rotateZ(90deg);
}
.circle > span:nth-child(1) {
  top: 0;
  left: 0;
}
.circle > span:nth-child(2) {
  top: 0;
  right: 0;
}
.circle > span:nth-child(3) {
  right: 0;
  bottom: 0;
}
.circle > span:nth-child(4) {
  left: 0;
  bottom: 0;
}
.circle2 > span:nth-child(1) {
  -webkit-animation-delay: -1.1s;
  animation-delay: -1.1s;
}
.circle3 > span:nth-child(1) {
  -webkit-animation-delay: -1s;
  animation-delay: -1s;
}
.circle1 > span:nth-child(2) {
  -webkit-animation-delay: -0.9s;
  animation-delay: -0.9s;
}
.circle2 > span:nth-child(2) {
  -webkit-animation-delay: -0.8s;
  animation-delay: -0.8s;
}
.circle3 > span:nth-child(2) {
  -webkit-animation-delay: -0.7s;
  animation-delay: -0.7s;
}
.circle1 > span:nth-child(3) {
  -webkit-animation-delay: -0.6s;
  animation-delay: -0.6s;
}
.circle2 > span:nth-child(3) {
  -webkit-animation-delay: -0.7s;
  animation-delay: -0.7s;
}
.circle3 > span:nth-child(3) {
  -webkit-animation-delay: -0.4s;
  animation-delay: -0.4s;
}
.circle1 > span:nth-child(4) {
  -webkit-animation-delay: -0.3s;
  animation-delay: -0.3s;
}
.circle2 > span:nth-child(4) {
  -webkit-animation-delay: -0.2s;
  animation-delay: -0.2s;
}
.circle3 > span:nth-child(4) {
  -webkit-animation-delay: -0.1s;
  animation-delay: -0.1s;
}
@-webkit-keyframes mycircle {
  0% {
    transform: scale(0);
  }
  40% {
    transform: scale(1);
  }
  80% {
    transform: scale(0);
  }
  100% {
    transform: scale(0);
  }
}
@keyframes mycircle {
  0% {
    transform: scale(0);
  }
  40% {
    transform: scale(1);
  }
  80% {
    transform: scale(0);
  }
  100% {
    transform: scale(0);
  }
}

.assistive-wrap {
  width: 58px;
  height: 58px;
  position: fixed;
  top: 50%;
  margin-top: -29px;
  left: 1px;
  z-index: 9;
}

.assistive-touch {
  width: 100%;
  height: 100%;
  background: #343434;
  border-radius: 10px;
  opacity: 0.3;
  position: relative;
}
.assistive-touch:before,
.assistive-touch:after,
.assistive-touch span {
  content: "";
  position: absolute;
  border-radius: 100%;
  box-shadow: 0 0 2px rgba(30, 30, 30, 0.5);
  display: block;
  background: rgba(255, 255, 255, 0.6);
}

.assistive-touch:before {
  width: 42px;
  height: 42px;
  left: 8px;
  top: 8px;
  opacity: 0.5;
}

.assistive-touch span {
  width: 34px;
  height: 34px;
  left: 12px;
  top: 12px;
}

.assistive-touch:after {
  width: 26px;
  height: 26px;
  left: 16px;
  top: 16px;
  background: #fff;
}

@keyframes rotate {
  from {
    transform: rotate(0);
  }
  to {
    transform: rotate(1turn);
  }
}
</style>
