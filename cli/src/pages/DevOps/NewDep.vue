<template>
  <div class="box">
    <div class="columns">
      <div class="column is-2">
        <input
          class="input has-text-centered"
          type="text"
          value="上线单标题"
          disabled
        />
      </div>
      <div class="column is-4">
        <input
          class="input is-primary"
          type="text"
          v-model="depTitle"
          placeholder="请输入此上线单标题"
        />
      </div>
    </div>
    <div class="columns">
      <div class="column is-2">
        <input
          class="input has-text-centered"
          type="text"
          value="服务类型"
          disabled
        />
      </div>
      <div class="column is-4">
        <div class="dropdown is-hoverable">
          <div class="dropdown-trigger">
            <button
              class="button"
              aria-haspopup="true"
              aria-controls="dropdown-menu"
            >
              {{ currServerType || "点击选择服务类型" }}
              <span class="icon is-small">
                <i class="fas fa-angle-down" aria-hidden="true"></i>
              </span>
            </button>
          </div>
          <div class="dropdown-menu" id="dropdown-menu" role="menu">
            <div class="dropdown-content">
              <tr v-for="t in depTypeList" :key="t">
                <a
                  class="dropdown-item"
                  :class="{ 'is-active': t == currServerType }"
                  @click="changeServerType(t)"
                  >{{ t }}</a
                >
              </tr>
            </div>
          </div>
        </div>
      </div>
    </div>
    <div class="columns">
      <div class="column is-2">
        <input
          class="input has-text-centered"
          type="text"
          value="版本提交时间"
          disabled
        />
      </div>
      <div class="column is-3">
        <date-picker
          v-model="versionTimeRange"
          range
          :lang="lang"
        ></date-picker>
        <span class="tag is-dark is-small">默认最新15条</span>
      </div>
    </div>
    <div class="columns">
      <div class="column is-2">
        <a
          class="button is-primary is-fullwidth"
          @click="updateSvnLog(true)"
          :class="{ 'is-loading': isLoading }"
          >刷新版本</a
        >
      </div>
      <div class="column">
        <div
          class="dropdown"
          :class="{ 'is-hoverable': versionMap[currServerType] }"
        >
          <div class="dropdown-trigger">
            <button
              class="button"
              aria-haspopup="true"
              aria-controls="dropdown-menu"
              @click="showVersion"
            >
              {{ currVersion || "点击选择待发布版本" }}
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
                  :class="{ 'is-active': currVersion === version }"
                  @click="selectVersion(version)"
                  >{{ version }}</a
                >
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
            <tr v-for="(msg, i) in currMessage" :key="i">
              <p>{{ msg }}</p>
            </tr>
          </div>
        </div>
      </div>
    </div>
    <div class="columns">
      <div class="column">
        <div class="control">
          <label class="radio">
            <input
              type="radio"
              name="answer"
              :checked="depRadio == 1"
              @click="changeRadio(1)"
            />
            全服上线
          </label>
          <label class="radio">
            <input
              type="radio"
              name="answer"
              :checked="depRadio == 2"
              @click="changeRadio(2)"
            />
            自定义上线
          </label>
        </div>
      </div>
    </div>
    <div class="columns" v-if="depRadio == 2">
      <div class="column is-12">
        <label class="checkbox" v-for="(k, i) in depRadioList" :key="i">
          <input type="checkbox" v-model="checkedDepRadioList" :value="k" />
          {{ k }}
        </label>
      </div>
    </div>
    <div class="columns">
      <div class="column is-2">
        <a class="button is-primary is-fullwidth" @click="submit">提交上线单</a>
      </div>
      <div class="column is-3">
        <a
          class="button is-primary is-fullwidth is-warning"
          @click="submitNewest"
          >一键提交最新版</a
        >
      </div>
    </div>
  </div>
</template>

<script>
import DatePicker from "vue2-datepicker";
import "vue2-datepicker/index.css";
import "vue2-datepicker/locale/zh-cn";
export default {
  name: "NewDep",
  data: function () {
    return {
      title: "NewDep",
      depTitle: "",
      depTypeList: [],
      currServerType: "",
      currVersion: "",
      currUrl: "",
      currMessage: [],
      dbDevTableName: "sys:ops:devini",
      versionList: [],
      versionMap: {},
      messageMap: {},
      depRadioList: [],
      checkedDepRadioList: [],
      urlMap: {},
      isLoading: false,
      revision: null,
      versionTimeRange: null,
      defaultSvnLogLimit: 15,
      depRadio: 1,

      lang: {
        formatLocale: {
          firstDayOfWeek: 1,
        },
        monthBeforeYear: false,
      },
    };
  },
  components: {
    DatePicker,
  },
  methods: {
    changeServerType: async function (type) {
      this.currServerType = type;
      this.currVersion = "";
      this.currMessage = [];
      this.depRadioList = [];
      this.versionList = [];
      if (this.versionMap[this.currServerType]) {
        this.revision = null;
        this.currUrl = this.urlMap[this.currServerType];
        this.versionList = JSON.parse(this.versionMap[this.currServerType]);
      }
    },
    updateSvnLog: async function (buttonClicked) {
      if (!this.currServerType) {
        return this.$store.commit("warn", `请先选择服务类型`);
      }
      let svnStartTime, svnEndTime;
      const reg = /^([0-9]{4})-((?:0[1-9]|[1-9]|1[1-2]))-((?:(?:0[1-9]|[1-9])|1[0-9]|2[0-9]|3[0-1]))$|^([0-9]{4})-((?:0[1-9]|[1-9]|1[1-2]))-((?:(?:0[1-9]|[1-9])|1[0-9]|2[0-9]|3[0-1]))\s((?:[0-1]?[0-9]{1}|2[0-3])):([0-5]?[0-9]{1}):([0-5]?[0-9]{1})$|^([0-9]{4})-((?:0[1-9]|[1-9]|1[1-2]))-((?:(?:0[1-9]|[1-9])|1[0-9]|2[0-9]|3[0-1]))\s((?:[0-1]?[0-9]{1}|2[0-3])):([0-5]?[0-9]{1}):([0-5]?[0-9]{1})\.?(\d{3})+$/;
      if (this.versionTimeRange && this.versionTimeRange.length == 2) {
        svnStartTime = this.format("yyyy-MM-dd", this.versionTimeRange[0]);
        svnEndTime = this.format("yyyy-MM-dd", this.versionTimeRange[1]);
        if (
          reg.test(svnStartTime) === false ||
          reg.test(svnEndTime) === false
        ) {
          return this.$store.commit("warn", `请检查你的日期格式yyyy:mm:dd`);
        }
      }
      this.isLoading = buttonClicked || false;
      this.updateMask(1);
      try {
        this.currUrl = await this.getServerSvnUrl(this.currServerType);
        this.urlMap[this.currServerType] = this.currUrl;
        let params = {
          repourl: this.currUrl,
          limit: this.defaultSvnLogLimit,
        };

        if (svnStartTime && svnEndTime) {
          params.period = `{${svnStartTime}}:{${svnEndTime}}`;
          delete params.limit;
        }
        let ret = await this.$mojoapi.get("/web/dep/commithistory", params);
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

      this.revision = this.currVersion.split("|")[0].substr(1);
      this.currMessage = [];
      if (
        this.messageMap[this.currServerType] &&
        this.messageMap[this.currServerType][this.revision]
      ) {
        this.currMessage = JSON.parse(
          this.messageMap[this.currServerType][this.revision]
        );
        return;
      }
      this.updateMask(1);
      try {
        let ret = await this.$mojoapi.get("/web/dep/commithistory", {
          repourl: this.currUrl,
          revision: this.revision,
        });
        this.currMessage = ret.data.split("\n");
        if (!this.messageMap[this.currServerType]) {
          this.messageMap[this.currServerType] = {};
        }
        this.messageMap[this.currServerType][this.revision] = JSON.stringify(
          this.currMessage
        );
        this.updateMask(-1);
      } catch (e) {
        this.$store.commit(
          "error",
          `加载服务类型<${this.currVersion}>版本<${
            this.revision
          }>说明信息失败 : ${e.data || e.message}`
        );
        this.updateMask(-1);
      }
    },
    showVersion: async function () {
      if (this.versionMap[this.currServerType]) {
        this.revision = null;
        this.currUrl = this.urlMap[this.currServerType];
        this.versionList = JSON.parse(this.versionMap[this.currServerType]);
      } else {
        await this.updateSvnLog();
      }
    },
    updateMask: function (n) {
      this.$store.commit("updateDevOpsMask", n);
    },
    getServerSvnUrl: async function (type) {
      try {
        let ret = await this.$mojoapi.get(
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
        if (this.depTitle.length < 6) {
          return this.$store.commit("warn", `提交上线单失败 : 标题最少6个字符`);
        }
        if (this.depRadio != 1) {
          if (this.checkedDepRadioList.length === 0) {
            return this.$store.commit("warn", `自定义上线至少选择一个服务`);
          }
        }
        await this.$mojoapi.put(`/web/dep/create`, {
          value: JSON.stringify({
            title: this.depTitle,
            type: this.currServerType,
            revision: this.revision,
            repourl: this.currUrl,
            desc: this.currMessage,
            list: this.checkedDepRadioList,
          }),
        });
        this.$store.commit("info", `提交上线单成功 : ${this.depTitle}`);
      } catch (e) {
        this.$store.commit("error", `提交上线单失败 : ${e.data || e.message}`);
      }
    },
    submitNewest: async function () {
      await this.updateSvnLog();
      await this.selectVersion(this.versionList[0]);
      this.depTitle = `${this.currServerType}-${
        this.revision
      }-${new Date().toLocaleString()}`;
      await this.submit();
    },
    changeRadio: async function (i) {
      if (i == 1) {
        this.depRadio = 1;
        return;
      }
      if (!this.currServerType) {
        return this.$store.commit("warn", `请先选择服务类型`);
      }
      this.depRadio = 2;
      try {
        const ret = await this.$mojoapi.get(`/web/db/sys:ops:depini`);
        let tempList = [];
        for (let i = 0; i < ret.data.length; i += 2) {
          let info = JSON.parse(ret.data[i + 1]);
          if (info[0] == this.currServerType) {
            tempList.push(info[1]);
          }
        }
        this.depRadioList = tempList;
      } catch (e) {
        this.$store.commit(
          "error",
          `获取数据库表错误 : ${e.data || e.message}`
        );
      }
    },
    getAllDepType: async function () {
      try {
        const ret = await this.$mojoapi.get(`/web/db/sys:ops:devini`);
        let tempList = [];
        for (let i = 0; i < ret.data.length; i += 2) {
          tempList.push(ret.data[i]);
        }
        this.depTypeList = tempList;
      } catch (e) {
        this.$store.commit(
          "error",
          `获取数据库表错误 : ${e.data || e.message}`
        );
      }
    },
    format: function (format, date) {
      if (!date) {
        date = new Date();
      }
      let o = {
        "M+": date.getMonth() + 1, // month
        "d+": date.getDate(), // day
        "h+": date.getHours(), // hour
        "m+": date.getMinutes(), // minute
        "s+": date.getSeconds(), // second
        "q+": Math.floor((date.getMonth() + 3) / 3), // quarter
        "S+": date.getMilliseconds(),
        // millisecond
      };

      if (/(y+)/.test(format)) {
        format = format.replace(
          RegExp.$1,
          (date.getFullYear() + "").substr(4 - RegExp.$1.length)
        );
      }

      for (let k in o) {
        if (new RegExp("(" + k + ")").test(format)) {
          let formatStr = "";
          for (let i = 1; i <= RegExp.$1.length; i++) {
            formatStr += "0";
          }

          let replaceStr = "";
          if (RegExp.$1.length == 1) {
            replaceStr = o[k];
          } else {
            formatStr = formatStr + o[k];
            let index = ("" + o[k]).length;
            formatStr = formatStr.substr(index);
            replaceStr = formatStr;
          }
          format = format.replace(RegExp.$1, replaceStr);
        }
      }
      return format;
    },
  },
  beforeMount: function () {
    this.getAllDepType();
  },
};
</script>

<style scoped>
</style>
