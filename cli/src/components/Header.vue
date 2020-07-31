<template>
  <header>
    <nav class="navbar" role="navigation" aria-label="main navigation">
      <div class="navbar-brand">
        <a class="navbar-item" href="https://github.com/adayswait">
          <img src="./../assets/logo.png" width="33" height="33" />
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

      <div id="navbarDataTarget" class="navbar-menu">
        <div class="navbar-start">
          <a class="navbar-item">Home</a>

          <a class="navbar-item" href="http://10.1.1.248:3000/fs">Documentation</a>

          <div class="navbar-item has-dropdown is-hoverable">
            <a class="navbar-link">More</a>

            <div class="navbar-dropdown">
              <a class="navbar-item">About</a>
              <a class="navbar-item">Jobs</a>
              <a class="navbar-item">Contact</a>
              <hr class="navbar-divider" />
              <a class="navbar-item">Report an issue</a>
            </div>
          </div>
        </div>

        <div class="navbar-end">
          <div class="navbar-item has-dropdown is-hoverable">
            <a class="navbar-link">
              <p>
                <strong>{{this.$store.state.userInfo.user}}</strong>
                <small>@{{this.GROUP[this.$store.state.userInfo.group]}}</small>
              </p>
            </a>
            <div class="navbar-dropdown">
              <a class="navbar-item">About</a>
              <a class="navbar-item">Edit</a>
              <a class="navbar-item">Message</a>
              <hr class="navbar-divider" />
              <a href="./" class="navbar-item" @click="logout">退出</a>
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
      GROUP: {
        0: "未激活",
        1: "管理员",
        2: "未定义",
      },
    };
  },
  methods: {
    showSmallMenu: async function () {
      window.console.log("显示小菜单");
    },
    logout: async function () {
      try {
        await this.$httpc.get("/web/auth/logout");
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
