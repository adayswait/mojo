<template>
  <div class="column">
    <div class="tabs">
      <ul>
        <!-- <li :class="{'is-active':currFocusTitle=='配表'}" @click="focus('配表')">
          <a>配表</a>
        </li>-->
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
    <div class="columns" v-if="currFocusTitle=='配表'">
      <div class="column is-2">
        <input class="input has-text-centered" type="text" value="服务类型" disabled />
      </div>
      <div class="column is-2">
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
      <div class="column is-2">
        <a class="button is-primary is-fullwidth" @click="hotUpdate">一键热更</a>
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
          <date-picker v-model="timeEffect" valuetype="format"></date-picker>
          <span class="tag is-small is-dark">00:00:00</span>
        </div>
        <div class="column is-2">
          <input class="input has-text-centered is-small" type="text" value="失效日期" disabled />
        </div>
        <div class="column is-4">
          <date-picker v-model="timeExpire" valuetype="format"></date-picker>
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
      <div class="columns">
        <div class="column is-12">
          <a class="button is-primary is-fullwidth" @click="sendMail">发送</a>
        </div>
      </div>
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
      <div class="columns">
        <div class="column is-12">
          <a class="button is-primary is-fullwidth" @click="sendMail">发送</a>
        </div>
      </div>
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
      currServerType: "",
      currFocusTitle: "全服邮件",
      mailTitle: "",
      mailSender: "",
      timeEffect: new Date(),
      timeExpire: new Date(Date.now() + 7 * 24 * 3600 * 1000),
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
    changeServerType: async function (type) {
      this.currServerType = type;
    },
    hotUpdate: async function () {
      void this.currServerType;
    },
    focus: function (title) {
      this.currFocusTitle = title;
    },
    sendMail: async function () {
      if (this.mailTitle.length === 0) {
        return this.$store.commit("warn", "邮件标题不能为空");
      }
      if (this.mailSender.length === 0) {
        return this.$store.commit("warn", "发件人不能为空");
      }
      if (this.mailContent.length === 0) {
        return this.$store.commit("warn", "邮件内容不能为空");
      }
      window.console.log("send mail", this.mailContent);
      await this.$mojoapi.post("/web/splan/mail", {
        activetime: parseInt(this.timeEffect.getTime() / 1000).toString(),
        sender: this.mailSender,
        title: this.mailTitle,
        regendtime: "0",
        user: "",
        gmail_file: "",
        regstarttime: "0",
        refresh: "true",
        mailtype: "1",
        content: this.mailContent,
        addition: "",
        switch: "10.1.1.43:21010",
        deadtime: parseInt(this.timeExpire.getTime() / 1000).toString(),
        switch_key: "123456",
        attachment: "",
      });
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
</style>
