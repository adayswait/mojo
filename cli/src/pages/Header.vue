<template>
  <header>
    <nav
      id="header-nav"
      class="navbar is-fixed-top has-shadow"
      role="navigation"
      aria-label="main navigation"
    >
      <div class="container">
        <div class="navbar-brand">
          <a class="navbar-item" href="/">
            <img src="./../assets/logo.png" width="100" height="30" />
          </a>
          <a
            role="button"
            class="navbar-burger burger"
            aria-label="menu"
            aria-expanded="false"
            data-target="navbarDataTarget"
            @click="showSmallMenu"
          >
            <span aria-hidden="true"></span>
            <span aria-hidden="true"></span>
            <span aria-hidden="true"></span>
          </a>
        </div>

        <div id="navbarDataTarget" class="navbar-menu" :class="{'is-active':activeMenu}">
          <div class="navbar-start">
            <a class="navbar-item" href="/">Home</a>
            <a class="navbar-item" href="/fs">Documentation</a>
            <div class="navbar-item has-dropdown is-hoverable">
              <a class="navbar-link">More</a>
              <div class="navbar-dropdown">
                <!-- <a class="navbar-item" target="_blank" href="http://10.1.1.239">239GM</a>
                <a class="navbar-item" target="_blank" href="http://plan_s">plan_s</a> -->
                <a
                  class="navbar-item"
                  target="_blank"
                  href="http://s.61.com/?ip=123.206.181.31&port=9999&log=true&showPP=false"
                >外网白名单入口</a>
                <hr class="navbar-divider" />
                <a class="navbar-item">Report an issue</a>
              </div>
            </div>
          </div>

          <div class="navbar-end" v-if="!this.$store.state.visible.Login">
            <div class="navbar-item has-dropdown is-hoverable">
              <a class="navbar-link">
                <p>
                  <strong>{{this.$store.state.userInfo.user}}</strong>
                  <small>@{{this.$store.state.GROUP[this.$store.state.userInfo.group]}}</small>
                </p>
              </a>
              <div class="navbar-dropdown is-right">
                <!-- <a class="navbar-item">About</a>
                <a class="navbar-item">Edit</a>
                <a class="navbar-item">Message</a>-->
                <hr class="navbar-divider" />
                <a href="./" class="navbar-item" @click="logout">退出</a>
              </div>
            </div>
          </div>
        </div>
      </div>
    </nav>
  </header>
</template>

<script>
export default {
  name: "Header",
  data: function () {
    return {
      title: "Header",
      activeMenu: false,
    };
  },
  methods: {
    showSmallMenu: function () {
      this.activeMenu = !this.activeMenu;
    },
    logout: async function () {
      try {
        await this.$mojoapi.get("/web/auth/logout");
        this.$store.commit("setUserInfo", {
          user: undefined,
          group: undefined,
        });
        this.$store.commit("setVisible", {
          name: "Login",
          visible: true,
        });
        this.$store.commit("setVisible", {
          name: "Home",
          visible: false,
        });
        this.$store.commit("setVisible", {
          name: "UsrInfo",
          visible: false,
        });
      } catch (e) {
        this.err = e.data;
      }
    },
  },
};
</script>

<style scoped>
#header-nav {
  background-color: #f1f1f1;
}
</style>