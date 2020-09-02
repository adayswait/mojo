<template>
  <div class="column">
    <div class="tabs">
      <ul>
        <li :class="{'is-active':currFocusTitle=='快捷操作'}" @click="focus('快捷操作')">
          <a>快捷操作</a>
        </li>
        <li :class="{'is-active':currFocusTitle=='全服邮件'}" @click="focus('全服邮件')">
          <a>全服邮件</a>
        </li>
        <li :class="{'is-active':currFocusTitle=='全服公告'}" @click="focus('全服公告')">
          <a>全服公告</a>
        </li>
        <li :class="{'is-active':currFocusTitle=='历史公告'}" @click="focus('历史公告')">
          <a>历史公告</a>
        </li>
        <li :class="{'is-active':currFocusTitle=='全服跑马灯'}" @click="focus('全服跑马灯')">
          <a>全服跑马灯</a>
        </li>
      </ul>
    </div>
    <div class="container" v-if="currFocusTitle=='快捷操作'">
      <div class="columns">
        <div class="column is-2">
          <input class="input has-text-centered" type="text" value="配表热更" disabled />
        </div>
        <div class="column is-3">
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
                <tr v-for="t in depTypeList" :key="t">
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
        <div class="column is-4">
          <a
            class="button is-primary"
            @click="hotUpdateConfig"
            :class="{'is-loading':hotUpdating}"
          >一键热更</a>
          <label class="checkbox">
            <input type="checkbox" v-model="withNotify" :value="true" />
            <strong>
              <small>通知</small>
            </strong>
          </label>
        </div>
      </div>
      <div class="columns">
        <div class="column is-2">
          <input class="input has-text-centered" type="text" value="修改服务器时间" disabled />
        </div>
        <div class="column is-3">
          <date-picker class="datepicker" type="datetime" v-model="serverTime"></date-picker>
        </div>
        <div class="column is-2">
          <a class="button is-primary is-warning" @click="changeServerTime">立即修改</a>
        </div>
      </div>
    </div>

    <div class="container" v-if="currFocusTitle=='全服邮件'">
      <div class="columns">
        <div class="column is-2">
          <input class="input has-text-centered is-small" type="text" value="邮件标题" disabled />
        </div>
        <div class="column is-3">
          <input class="input is-small" type="text" placeholder="邮件标题" v-model="mailTitle" />
        </div>
        <div class="column is-2 is-offset-1">
          <input class="input has-text-centered is-small" type="text" value="发件人" disabled />
        </div>
        <div class="column is-3">
          <input class="input is-small" type="text" placeholder="发件人" v-model="mailSender" />
        </div>
      </div>
      <div class="columns">
        <div class="column is-2">
          <input class="input has-text-centered is-small" type="text" value="生效日期" disabled />
        </div>
        <div class="column is-4">
          <date-picker v-model="timeEffect"></date-picker>
          <span class="tag is-small is-dark">00:00:00</span>
        </div>
        <div class="column is-2">
          <input class="input has-text-centered is-small" type="text" value="失效日期" disabled />
        </div>
        <div class="column is-4">
          <date-picker v-model="timeExpire"></date-picker>
          <span class="tag is-small is-dark">23:59:59</span>
        </div>
      </div>
      <div class="columns">
        <div class="column is-12" style="height:400px;">
          <div>
            <quill-editor
              ref="myTextEditor"
              v-model="mailContent"
              :options="editorOption"
              style="height:300px;"
            ></quill-editor>
          </div>
        </div>
      </div>
      <nav class="level">
        <div class="level-item has-text-centered">
          <a class="button is-primary" @click="sendMail" disabled>发送</a>
        </div>
      </nav>
    </div>

    <div class="container" v-if="currFocusTitle=='全服公告'">
      <div class="columns">
        <div class="column is-2">
          <input class="input has-text-centered" type="text" value="公告标题" disabled />
        </div>
        <div class="column is-4">
          <input class="input has-text-centered" type="text" placeholder="公告标题" v-model="mailTitle" />
        </div>
      </div>
      <div class="columns">
        <div class="column is-12" style="height:400px;">
          <div>
            <quill-editor
              ref="myTextEditor"
              v-model="announcementContent"
              :options="editorOption"
              style="height:300px;"
            ></quill-editor>
          </div>
        </div>
      </div>
      <nav class="level">
        <div class="level-item has-text-centered">
          <a class="button is-primary" @click="sendMail" disabled>发送</a>
        </div>
      </nav>
    </div>
    <div class="container" v-if="currFocusTitle=='历史公告'">
      <p>历史公告 under developing</p>
    </div>
    <div class="container" v-if="currFocusTitle=='全服跑马灯'">
      <p>全服跑马灯 under developing</p>
    </div>
  </div>
</template>

<script>
import DatePicker from "vue2-datepicker";
import { quillEditor } from "vue-quill-editor";
import "quill/dist/quill.core.css";
import "quill/dist/quill.snow.css";
import "quill/dist/quill.bubble.css";
import "vue2-datepicker/index.css";
export default {
  data: function () {
    return {
      mailContent: "",
      announcementContent: "",
      editorOption: {
        placeholder: "编辑内容",
      },
      depTypeList: [],
      currServerType: "online",
      currFocusTitle: "快捷操作",
      mailTitle: "",
      mailSender: "",
      timeEffect: new Date(),
      timeExpire: new Date(Date.now() + 7 * 24 * 3600 * 1000),
      serverTime: new Date(),
      hotUpdating: false,
      withNotify: false,
    };
  },
  components: {
    quillEditor,
    DatePicker,
  },
  methods: {
    getAllDepType: async function () {
      try {
        const ret = await this.$mojoapi.get(`/web/db/sys:ops:devini`);
        let tempList = [];
        for (let i = 0; i < ret.data.length; i += 2) {
          if (
            ret.data[i].indexOf("online") !== -1 ||
            ret.data[i].indexOf("battle") !== -1
          ) {
            tempList.push(ret.data[i]);
          }
        }
        this.depTypeList = tempList;
      } catch (e) {
        this.$store.commit(
          "error",
          `获取数据库表错误 : ${e.data || e.message}`
        );
      }
    },
    changeServerType: async function (type) {
      this.currServerType = type;
    },
    hotUpdateConfig: async function () {
      if (this.currServerType === "") {
        return this.$store.commit("warn", `请先选择服务类型`);
      }
      let updateModule = this.currServerType;
      try {
        this.hotUpdating = true;
        const ret = await this.$mojoapi.put(
          `/web/splan/config/${updateModule}`,
          { notify: this.withNotify }
        );
        const retData = JSON.parse(ret.data);
        if (retData.code == 0) {
          this.$store.commit("info", `热更${updateModule}配表成功`);
        } else {
          this.$store.commit(
            "error",
            `热更${updateModule}配表失败:${retData.desc}`
          );
        }
      } catch (e) {
        this.$store.commit(
          "error",
          `热更${updateModule}配表错误 : ${e.data || e.message}`
        );
      }
      this.hotUpdating = false;
    },
    changeServerTime: async function () {
      await this.$mojoapi.put(`/web/splan/changetime`, {
        ip: "10.1.1.239",
        time: this.format("yyyy-MM-dd hh:mm:ss", this.serverTime),
      });
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
    focus: function (title) {
      this.currFocusTitle = title;
    },
    sendMail: async function () {
      return this.$store.commit("warn", "under developing");
      // if (this.mailTitle.length === 0) {
      //   return this.$store.commit("warn", "邮件标题不能为空");
      // }
      // if (this.mailSender.length === 0) {
      //   return this.$store.commit("warn", "发件人不能为空");
      // }
      // if (this.mailContent.length === 0) {
      //   return this.$store.commit("warn", "邮件内容不能为空");
      // }
      // window.console.log("send mail", this.mailContent);
      // await this.$mojoapi.post("/web/splan/mail", {
      //   activetime: parseInt(this.timeEffect.getTime() / 1000).toString(),
      //   sender: this.mailSender,
      //   title: this.mailTitle,
      //   regendtime: "0",
      //   user: "",
      //   gmail_file: "",
      //   regstarttime: "0",
      //   refresh: "true",
      //   mailtype: "1",
      //   content: this.mailContent,
      //   addition: "",
      //   switch: "10.1.1.43:21010",
      //   deadtime: parseInt(this.timeExpire.getTime() / 1000).toString(),
      //   switch_key: "123456",
      //   attachment: "",
      // });
    },
    sendAnouncement: function () {
      window.console.log("send anouncement", this.announcementContent);
    },
  },
  beforeMount: function () {
    this.getAllDepType();
  },
};
</script>

<style scoped>
.tag {
  margin-left: 5px;
}
.checkbox {
  padding-left: 5px;
  padding-top: 24px;
}
.datepicker {
  margin-top: 3px;
}
</style>
