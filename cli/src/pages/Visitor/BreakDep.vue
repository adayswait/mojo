<template>
  <div class="column is-4 is-offset-4">
    <div id="breakdepprogress" class="box">
      <div class="level-item has-text-centered">
        <p>{{progressDescCN[depStatus]||"已过期"}}</p>
      </div>
      <progress
        class="progress"
        :value="depStatus==progressDesc.length-1?1:undefined"
        max="1"
        :class="{'is-warning':depStatus!=progressDesc.length-1,
        'is-success':depStatus==progressDesc.length-1}"
      ></progress>
    </div>
    <div id="breakdepstatus" class="box">
      <nav class="level is-mobile">
        <div class="level-item has-text-centered">
          <div>
            <p class="heading">模块</p>
            <p class="title">
              <small>
                <small>{{depType}}</small>
              </small>
            </p>
          </div>
        </div>
        <div class="level-item has-text-centered">
          <div>
            <p class="heading">倒计时(s)</p>
            <p class="title">
              <small>
                <small>{{countDown}}</small>
              </small>
            </p>
          </div>
        </div>
        <div class="level-item has-text-centered">
          <div>
            <p class="heading">状态</p>
            <p class="title">
              <small>
                <small>{{progressDesc[depStatus]||"unknown"}}</small>
              </small>
            </p>
          </div>
        </div>
      </nav>
    </div>
    <div id="breakdepops" class="box">
      <nav class="level is-mobile">
        <div class="level-item has-text-centered">
          <button
            class="button is-dark is-small is-rounded"
            :class="{'is-loading':btnCancelLoading}"
            @click="opDep('cancel')"
          >立即重启</button>
        </div>
        <div class="level-item has-text-centered">
          <button
            class="button is-warning is-small is-rounded"
            :class="{'is-loading':btnRenewLoading}"
            @click="opDep('renew')"
          >续期1分钟</button>
        </div>
      </nav>
    </div>
    <strong>
      <p class="help is-danger">{{this.err}}</p>
    </strong>
    <!-- <p>{{$route.query}}</p> -->
  </div>
</template>

<script>
export default {
  data: function () {
    return {
      title: "BreakDep",
      countDown: "-",
      depType: "unknown",
      depStatus: "unknown",
      err: undefined,
      btnCancelLoading: false,
      btnRenewLoading: false,
      timer: null,
      progressDesc: [
        "init",
        "checkout",
        "sleep",
        "deploy",
        "stop",
        "start",
        "over",
      ],
      progressDescCN: [
        "初始化中",
        "检出代码中",
        "休眠等待中",
        "部署中",
        "停止服务中",
        "启动服务中",
        "已发布完成",
      ],
    };
  },
  methods: {
    opDep: async function (op, auto) {
      if (auto !== true) {
        this.err = "";
      }
      if (op == "cancel") {
        this.btnCancelLoading = true;
      } else if (op == "renew") {
        this.btnRenewLoading = true;
      }
      try {
        const ret = await this.$mojoapi.get("/web/visitor/breakdep", {
          depuuid: this.$route.query.depuuid,
          op: op,
        });
        this.depType = ret.data[0];
        this.countDown = ret.data[1];
        this.depStatus = ret.data[2];
      } catch (e) {
        this.err = e.err || e.message;
        this.depType = e.data ? e.data[0] : "unknown";
        this.countDown = e.data ? e.data[1] : 0;
        this.depStatus = e.data ? e.data[2] : "unknown";
      }
      if (this.epStatus == this.progressDesc.length - 1) {
        if (this.timer) {
          clearInterval(this.timer);
          this.timer = null;
        }
        this.countDown = "-";
      }
      if (op == "cancel") {
        this.btnCancelLoading = false;
      } else if (op == "renew") {
        this.btnRenewLoading = false;
      }
    },
  },
  beforeMount: function () {
    this.opDep(this.$route.query.op);
    this.timer = setInterval(() => {
      this.opDep("view", true);
    }, 1000);
  },
  destroyed: function () {
    if (this.timer) {
      clearInterval(this.timer);
      this.timer = null;
    }
  },
};
</script>

<style scoped>
.box {
  background-color: #f1f1f1;
}
#breakdepprogress {
  margin-top: 10%;
  margin-bottom: 0px;
}

#breakdepstatus {
  margin-top: 2px;
  margin-bottom: 0px;
}

#breakdepops {
  margin-top: 2px;
}
</style>

