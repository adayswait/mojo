<template>
  <div>
    <Login />
    <div class="container" v-if="this.$store.state.visible.Home">
      <div class="columns">
        <div class="column is-2">
          <aside class="menu">
            <p class="menu-label">General</p>
            <ul class="menu-list">
              <li>
                <router-link to="/devops/home">主页</router-link>
              </li>
            </ul>
            <p class="menu-label">Transactions</p>
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
            <p class="menu-label">Administration</p>
            <ul class="menu-list">
              <li>
                <router-link to="/devops/newini">项目配置</router-link>
              </li>
              <li>
                <a>成员管理</a>
                <ul>
                  <li>
                    <router-link to="/devops/queryusr">所有成员</router-link>
                  </li>
                  <li>
                    <router-link to="/devops/newusr">新建角色</router-link>
                  </li>
                </ul>
              </li>
            </ul>
            <p class="menu-label">Others</p>
            <ul class="menu-list">
              <li>
                <router-link to="/devops/otherops">其他操作</router-link>
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
  data: function() {
    return {};
  },
  components: {
    Login
  }
};
</script>
