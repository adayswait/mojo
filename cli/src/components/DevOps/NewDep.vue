<template>
  <div class="box">
    <div class="columns">
      <div class="column is-2">
        <input class="input has-text-centered" type="text" value="上线单标题" disabled />
      </div>
      <div class="column is-4">
        <input class="input is-primary" type="text" v-model="depTitle" placeholder="请输入此上线单标题" />
      </div>
    </div>
    <div class="columns">
      <div class="column is-2">
        <input class="input has-text-centered" type="text" value="服务类型" disabled />
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
        <a
          class="button is-primary is-fullwidth"
          @click="updateSvnLog(true)"
          :class="{'is-loading':isLoading}"
        >刷新版本</a>
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
    <div class="columns">
      <div class="column is-2">
        <a class="button is-primary is-fullwidth" @click="submit">提交上线单</a>
      </div>
      <div class="column is-3">
        <a class="button is-primary is-fullwidth" @click="submitNewest">一键提交最新版</a>
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
      depTitle: "",
      currServerType: "",
      currVersion: "",
      currUrl: "",
      currMessage: [],
      dbDevTableName: "sys:ops:devini",
      versionList: [],
      versionMap: {},
      messageMap: {},
      urlMap: {},
      isLoading: false,
      rversion: null,
    };
  },
  methods: {
    changeServerType: async function (type) {
      this.currServerType = type;
      this.currVersion = "";
      this.currMessage = [];
      window.console.log(type, this.versionMap[type]);
      if (this.versionMap[type]) {
        this.rversion = null;
        this.currUrl = this.urlMap[type];
        this.versionList = JSON.parse(this.versionMap[type]);
      } else {
        await this.updateSvnLog();
      }
    },
    updateSvnLog: async function (buttonClicked) {
      this.isLoading = buttonClicked || false;
      this.updateMask(1);
      try {
        this.currUrl = await this.getServerSvnUrl(this.currServerType);
        this.urlMap[this.currServerType] = this.currUrl;
        let ret = await this.$httpc.get("/web/cmd/svnlog", {
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
      } catch (e) {
        this.$store.commit("setVisible", {
          name: "DevOpsMask",
          visible: false,
        });
        this.$store.commit("error", `刷新svn版本失败:${e.data || e.message}`);
      }
      this.isLoading = false;
      this.updateMask(-1);
    },
    selectVersion: async function (version) {
      this.currVersion = version;

      this.rversion = this.currVersion.split("|")[0];
      this.currMessage = [];
      if (
        this.messageMap[this.currServerType] &&
        this.messageMap[this.currServerType][this.rversion]
      ) {
        this.currMessage = JSON.parse(
          this.messageMap[this.currServerType][this.rversion]
        );
        return;
      }
      this.updateMask(1);
      try {
        let ret = await this.$httpc.get("/web/cmd/svnlog", {
          repourl: this.currUrl,
          version: this.rversion,
        });
        this.currMessage = ret.data.split("\n");
        if (!this.messageMap[this.currServerType]) {
          this.messageMap[this.currServerType] = {};
        }
        this.messageMap[this.currServerType][this.rversion] = JSON.stringify(
          this.currMessage
        );
        this.updateMask(-1);
      } catch (e) {
        this.$store.commit(
          "error",
          `加载服务类型<${this.currVersion}>版本<${
            this.rversion
          }>说明信息失败 : ${e.data || e.message}`
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
    submit: async function () {
      try {
        await this.$httpc.put(`/web/db/sys:ops:depbil`, {
          value: JSON.stringify({
            title: this.depTitle,
            type: this.currServerType,
            rversion: this.rversion,
            repourl: this.currUrl,
            desc: this.currMessage,
          }),
        });
        this.$store.commit("info", `提交上线单成功 : ${this.depTitle}`);
      } catch (e) {
        this.$store.commit("error", `提交上线单失败 : ${e.data || e.message}`);
      }
    },
    submitNewest: async function () {
      this.$store.commit("info", `一键提交最新上线单成功`);
    },
  },
  beforeMount: async function () {
    // this.updateSvnLog();
  },
};
</script>

<style scoped>
</style>
