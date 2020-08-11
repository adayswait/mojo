<template>
  <div class="column">
    <div class="tabs">
      <ul>
        <li class="is-active">
          <a>Unix时间戳</a>
        </li>
        <li>
          <a>json</a>
        </li>
        <li>
          <a>Videos</a>
        </li>
        <li>
          <a>Documents</a>
        </li>
      </ul>
    </div>
    <div class="columns">
      <div class="column is-3">
        <input class="input" type="text" value="当前UNIX时间(秒)" disabled />
      </div>
      <div class="column is-3">
        <input class="input is-primary" type="text" v-model="currUnixTimeS" readonly />
      </div>
      <div class="column is-1">
        <a
          class="button is-primary "
          v-clipboard:copy="currUnixTimeS"
          v-clipboard:success="onCopySuccess"
          v-clipboard:error="onCopyError"
        >复制</a>
      </div>
    </div>
    <div class="columns">
      <div class="column is-3">
        <input class="input" type="text" value="当前UNIX时间(毫秒)" disabled />
      </div>
      <div class="column is-3">
        <input class="input is-primary" type="text" v-model="currUnixTimeMS" readonly />
      </div>
      <div class="column is-1">
        <a
          class="button is-primary"
          v-clipboard:copy="currUnixTimeMS"
          v-clipboard:success="onCopySuccess"
          v-clipboard:error="onCopyError"
        >复制</a>
      </div>
    </div>
    <div class="columns">
      <div class="column is-3">
        <input class="input" type="text" value="当前标准时间(秒)" disabled />
      </div>
      <div class="column is-3">
        <input class="input is-primary" type="text" v-model="currStdDateS" readonly />
      </div>
      <div class="column is-1">
        <a
          class="button is-primary"
          v-clipboard:copy="currStdDateS"
          v-clipboard:success="onCopySuccess"
          v-clipboard:error="onCopyError"
        >复制</a>
      </div>
    </div>
    <div class="columns">
      <div class="column is-3">
        <input class="input" type="text" value="当前标准时间(毫秒)" disabled />
      </div>
      <div class="column is-3">
        <input class="input is-primary" type="text" v-model="currStdDateMS" readonly />
      </div>
      <div class="column is-1">
        <a
          class="button is-primary"
          v-clipboard:copy="currStdDateMS"
          v-clipboard:success="onCopySuccess"
          v-clipboard:error="onCopyError"
        >复制</a>
      </div>
    </div>
    <br />
    <div class="columns">
      <div class="column is-3">
        <input
          class="input"
          type="text"
          placeholder="yyyy-mm-dd hh:mm:ss.SSS"
          v-model="inputStdDate"
          v-on:input="input"
        />
      </div>
      <div class="column is-3">
        <input
          class="input is-primary"
          type="text"
          placeholder="unix时间"
          v-model="transUnixMS"
          readonly
        />
      </div>
      <div class="column is-1">
        <a class="button is-primary" @click="copy(transUnixMS,true)">转换</a>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  name: "DevTools",
  data: function () {
    return {
      timerId: null,
      title: "DevTools",
      currStdDateMS: this.format("yyyy-MM-dd hh:mm:ss.SSS"),
      currStdDateS: this.format("yyyy-MM-dd hh:mm:ss"),
      currUnixTimeMS: Date.now(),
      currUnixTimeS: parseInt(Date.now() / 1000),
      inputStdDate: "",
      transUnixMS: "",
    };
  },
  methods: {
    copy: function (data, isTrans) {
      if (isTrans) {
        const reg = /^([0-9]{4})-((?:0[1-9]|[1-9]|1[1-2]))-((?:(?:0[1-9]|[1-9])|1[0-9]|2[0-9]|3[0-1]))$|^([0-9]{4})-((?:0[1-9]|[1-9]|1[1-2]))-((?:(?:0[1-9]|[1-9])|1[0-9]|2[0-9]|3[0-1]))\s((?:[0-1]?[0-9]{1}|2[0-3])):([0-5]?[0-9]{1}):([0-5]?[0-9]{1})$|^([0-9]{4})-((?:0[1-9]|[1-9]|1[1-2]))-((?:(?:0[1-9]|[1-9])|1[0-9]|2[0-9]|3[0-1]))\s((?:[0-1]?[0-9]{1}|2[0-3])):([0-5]?[0-9]{1}):([0-5]?[0-9]{1})\.?(\d{3})+$/;
        if (reg.test(this.inputStdDate) === false) {
          this.transUnixMS = null;
          this.$store.commit("error", "请检查你的时间格式");
          return;
        } else {
          this.transUnixMS = new Date(this.inputStdDate).getTime();
        }
      }
      this.$copyText(data).then(this.onCopySuccess, this.onCopyError);
    },
    input: function () {},
    onCopySuccess: function (e) {
      this.$store.commit("info", `成功复制:${e.text}`);
    },
    onCopyError: function (e) {
      this.$store.commit("error", `复制失败:${e.message}`);
    },
    format: function (format) {
      /*
       * eg:format="yyyy-MM-dd hh:mm:ss";
       */
      let date = new Date();
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
  mounted: function () {
    this.timerId = setInterval(() => {
      window.console.log("timer");
      this.currUnixTimeMS = Date.now();
      this.currUnixTimeS = parseInt(this.currUnixTimeMS / 1000);
      this.currStdDateS = this.format("yyyy-MM-dd hh:mm:ss");
      this.currStdDateMS = this.format("yyyy-MM-dd hh:mm:ss.SSS");
    }, 1000);
  },
  destroyed: function () {
    if (this.timerId) {
      clearInterval(this.timerId);
      this.timerId = null;
    }
  },
};
</script>