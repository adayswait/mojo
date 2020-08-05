<template>
  <div>
    <Login />
    <div class="container" v-if="this.$store.state.visible.Home">
      <div class="columns">
        <div class="column is-2">
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
                  <li>
                    <router-link to="/devops/queryops">所有部署</router-link>
                  </li>
                  <li>
                    <router-link to="/devops/queryops">我的部署</router-link>
                  </li>
                  <li>
                    <router-link to="/devops/newdep">新建部署</router-link>
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
                </ul>
              </li>
            </ul>
          </aside>
        </div>
        <router-view />
      </div>
    </div>
  </div>
</template>

<script>
import Login from "@/components/Login.vue";
export default {
  name: "DevOps",
  data: function () {
    return {
      currFocus: "主页",
    };
  },
  methods: {
    focus: function (target) {
      this.currFocus = target;
    },
  },
  components: {
    Login,
  },
};
</script>
