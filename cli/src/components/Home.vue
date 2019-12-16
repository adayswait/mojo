<template>
  <div class="container" v-if="this.$store.state.visible.Home">
    <div class="columns">
      <div class="column is-2">
        <aside class="menu">
          <p class="menu-label">General</p>
          <ul class="menu-list">
            <li>
              <a>主页</a>
            </li>
          </ul>
          <p class="menu-label">Transactions</p>
          <ul class="menu-list">
            <li>
              <a class="is-active">部署</a>
              <ul>
                <li>
                  <a>所有部署</a>
                </li>
                <li>
                  <a>我的部署</a>
                </li>
                <li>
                  <a>新建部署</a>
                </li>
              </ul>
            </li>
          </ul>
          <p class="menu-label">Administration</p>
          <ul class="menu-list">
            <li>
              <a>项目配置</a>
            </li>
            <li>
              <a>成员管理</a>
              <ul>
                <li>
                  <a>所有成员</a>
                </li>
                <li>
                  <a>新建成员</a>
                </li>
              </ul>
            </li>
            <li>
              <a>权限管理</a>
            </li>
          </ul>

          <p class="menu-label">Others</p>
          <ul class="menu-list">
            <li>
              <a>其他操作</a>
            </li>
          </ul>
        </aside>
      </div>
      <Query />
    </div>
  </div>
</template>

<script>
import Query from "./Home/Query.vue";
export default {
  name: "Home",
  data: function() {
    return {};
  },
  components: {
    Query
  }
};
</script>
