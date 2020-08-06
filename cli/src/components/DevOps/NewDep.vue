<template>
  <div class="box">
    <div class="columns">
      <div class="column is-2">
        <input class="input" type="text" value="上线单标题" disabled />
      </div>
      <div class="column is-4">
        <input class="input is-primary" type="text" />
      </div>
    </div>
    <div class="columns">
      <div class="column is-2">
        <input class="input" type="text" value="服务类型" disabled />
      </div>
      <div class="column is-4">
        <div class="dropdown is-hoverable">
          <div class="dropdown-trigger">
            <button class="button" aria-haspopup="true" aria-controls="dropdown-menu">
              {{currServerType||"点击选择服务类型"}}
              <span class="icon is-small">
                <i class="fas fa-angle-down" aria-hidden="true"></i>
              </span>
            </button>
          </div>
          <div class="dropdown-menu" id="dropdown-menu" role="menu">
            <div class="dropdown-content">
              <tr v-for="t in $store.state.SERVER_TYPE" :key="t">
                <a
                  class="dropdown-item"
                  :class="{'is-active':t==currServerType}"
                  @click="changeServerType(t)"
                >{{t}}</a>
              </tr>
            </div>
          </div>
        </div>
      </div>
    </div>
    <div class="columns">
      <div class="column is-2">
        <a class="button is-primary is-fullwidth" @click="updateSvnLog">刷新版本</a>
      </div>
      <div class="column">
        <div class="dropdown is-hoverable">
          <div class="dropdown-trigger">
            <button class="button" aria-haspopup="true" aria-controls="dropdown-menu">
              {{currVersion||"点击选择待发布版本"}}
              <span class="icon is-small">
                <i class="fas fa-angle-down" aria-hidden="true"></i>
              </span>
            </button>
          </div>
          <div class="dropdown-menu" id="dropdown-menu" role="menu">
            <div class="dropdown-content">
              <tr v-for="version in versionList" :key="version">
                <a
                  class="dropdown-item"
                  :class="{'is-active':currVersion===version}"
                  @click="selectVersion(version)"
                >{{version}}</a>
              </tr>
            </div>
          </div>
        </div>
      </div>
    </div>
    <div class="columns">
      <div class="column">
        <div class="box">
          <div class="content">
            <h4>版本信息</h4>
            <tr v-for="(msg,i) in currMessage" :key="i">
              <p>{{msg}}</p>
            </tr>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  name: "NewDep",
  data: function () {
    return {
      title: "NewDep",
      currServerType: "",
      currVersion: "",
      currUrl: "",
      currMessage: [],
      dbDevTableName: "sys:ops:devini",
      versionList: [],
      versionMap: {},
      messageMap: {},
      urlMap: {},
    };
  },
  methods: {
    changeServerType: async function (type) {
      this.currServerType = type;
      this.currVersion = "";
      this.currMessage = [];
      window.console.log(type, this.versionMap[type]);
      if (this.versionMap[type]) {
        this.currUrl = this.urlMap[type];
        this.versionList = JSON.parse(this.versionMap[type]);
      } else {
        await this.updateSvnLog();
      }
    },
    updateSvnLog: async function () {
      this.updateMask(1);
      try {
        this.currUrl = await this.getServerSvnUrl(this.currServerType);
        this.urlMap[this.currServerType] = this.currUrl;
        let ret = await this.$httpc.get("/web/svnlog", {
          repourl: this.currUrl,
          limit: 15,
        });
        const formatData = ret.data.split("\n");
        let tempList = [];
        //奇数下标的是有效信息
        for (let i = 1; i < formatData.length; i += 2) {
          if (formatData[i]) {
            tempList.push(formatData[i]);
          }
        }
        this.versionList = tempList;
        this.versionMap[this.currServerType] = JSON.stringify(tempList);
        this.updateMask(-1);
      } catch (e) {
        this.$store.commit("setVisible", {
          name: "DevOpsMask",
          visible: false,
        });
        this.$store.commit("error", `刷新svn版本失败:${e.data || e.message}`);
        this.updateMask(-1);
      }
    },
    selectVersion: async function (version) {
      this.currVersion = version;

      let rversion = this.currVersion.split("|")[0];
      this.currMessage = [];
      if (
        this.messageMap[this.currServerType] &&
        this.messageMap[this.currServerType][rversion]
      ) {
        this.currMessage = JSON.parse(
          this.messageMap[this.currServerType][rversion]
        );
        return;
      }
      this.updateMask(1);
      try {
        let ret = await this.$httpc.get("/web/svnlog", {
          repourl: this.currUrl,
          version: rversion,
        });
        this.currMessage = ret.data.split("\n");
        if (!this.messageMap[this.currServerType]) {
          this.messageMap[this.currServerType] = {};
        }
        this.messageMap[this.currServerType][rversion] = JSON.stringify(
          this.currMessage
        );
        this.updateMask(-1);
      } catch (e) {
        this.$store.commit(
          "error",
          `加载服务类型<${this.currVersion}>版本<${rversion}>说明信息失败 : ${
            e.data || e.message
          }`
        );
        this.updateMask(-1);
      }
    },
    updateMask: function (n) {
      this.$store.commit("updateDevOpsMask", n);
    },
    getServerSvnUrl: async function (type) {
      try {
        let ret = await this.$httpc.get(
          `/web/db/${this.dbDevTableName}/${type}`
        );
        return JSON.parse(ret.data).url;
      } catch (e) {
        this.$store.commit(
          "error",
          `加载服务类型<${type}>仓库URL失败 : ${e.data || e.message}`
        );
      }
    },
  },
  beforeMount: function () {
    // this.updateSvnLog();
  },
};
</script>

<style scoped>
</style>
