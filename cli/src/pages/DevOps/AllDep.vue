<template>
  <div class="box">
    <div class="box" v-if="progressList.length!==0">
      <table class="table is-striped is-fullwidth">
        <tbody>
          <tr v-for="(info,i) in progressList" :key="i">
            <td @click="showProgressDetails(info[1])">
              <div id="progressTag" class="columns">
                <div class="column">
                  <div class="field is-grouped is-grouped-multiline">
                    <div class="control">
                      <div class="tags has-addons">
                        <span class="tag is-dark">depid</span>
                        <span class="tag is-primary">{{info[1]}}</span>
                      </div>
                    </div>

                    <div class="control">
                      <div class="tags has-addons">
                        <span class="tag is-dark">type</span>
                        <span class="tag is-info">{{getDepInfoByDepid(info[1])[1].type}}</span>
                      </div>
                    </div>

                    <div class="control">
                      <div class="tags has-addons">
                        <span class="tag is-dark">time</span>
                        <span
                          class="tag is-light"
                        >{{new Date(parseInt(info[0])*1000).toLocaleString()}}</span>
                      </div>
                    </div>

                    <div class="control">
                      <div class="tags has-addons">
                        <span class="tag is-dark">status</span>
                        <span
                          class="tag"
                          :class="{'is-success':info[3]==progressDesc.length-1,
                        'is-warning':info[3]!=progressDesc.length-1&& info[3]>=0,
                        'is-danger':info[3]<0}"
                        >{{progressDesc[info[3]]}}</span>
                      </div>
                    </div>

                    <div class="control" v-if="parseInt(info[4])>0">
                      <div class="tags has-addons">
                        <span class="tag is-dark">awake</span>
                        <span class="tag is-warning">{{parseInt(info[4])}}</span>
                      </div>
                    </div>
                  </div>
                </div>
              </div>

              <div class="columns">
                <div class="column">
                  <progress
                    class="progress"
                    :class="{'is-danger':info[3]<0,
                    'is-success':info[3]==progressDesc.length-1}"
                    :value="info[3]>0?(info[3]/(progressDesc.length-1)):1"
                    max="1"
                  ></progress>
                </div>
              </div>
            </td>
          </tr>
        </tbody>
      </table>
    </div>
    <div class="box">
      <table class="table is-striped is-fullwidth">
        <tbody>
          <tr v-for="(info,i) in allDeps" :key="info[0]">
            <td>{{info[0]}}</td>
            <td>{{info[1].title}}</td>
            <td>
              <div class="field is-grouped is-grouped-multiline">
                <div class="control">
                  <div class="tags has-addons">
                    <span class="tag is-dark">type</span>
                    <span class="tag is-info">{{info[1].type}}</span>
                  </div>
                </div>

                <div class="control">
                  <div class="tags has-addons">
                    <span class="tag is-dark">version</span>
                    <span class="tag is-success">{{info[1].revision}}</span>
                  </div>
                </div>
              </div>
            </td>
            <td>
              <button class="button is-small is-warning" @click="viewDetail(i)">查看详情</button>
              <button class="button is-small is-dark" @click="submit(i)">部署上线</button>
              <button class="button is-small is-danger" @click="submit(i,true)">强制上线</button>
            </td>
          </tr>
        </tbody>
      </table>
    </div>
    <div id="details" class="modal is-large" :class="{'is-active':modalActive}">
      <div class="modal-background"></div>
      <div class="modal-card">
        <header class="modal-card-head">
          <p
            class="modal-card-title"
          >{{typeof currFoucsIdx==='number'?allDeps[currFoucsIdx][1].title:"非法上线单"}}</p>
          <button class="delete" aria-label="close" @click="closeDetails"></button>
        </header>
        <section class="modal-card-body">
          <div class="content">
            <h4>上线单号</h4>
            <p>{{typeof currFoucsIdx==='number'?allDeps[currFoucsIdx][0]:"无"}}</p>
            <h4>服务类型</h4>
            <p>{{typeof currFoucsIdx==='number'?allDeps[currFoucsIdx][1].type:"无"}}</p>
            <h4>checkout地址</h4>
            <p>{{typeof currFoucsIdx==='number'?allDeps[currFoucsIdx][1].repourl:"无"}}</p>
            <h4>发布至</h4>
            <p>{{typeof currFoucsIdx==='number'?(allDeps[currFoucsIdx][1].list.length===0?"全服":allDeps[currFoucsIdx][1].list):"无"}}</p>
            <h4>版本号</h4>
            <p>{{typeof currFoucsIdx==='number'?allDeps[currFoucsIdx][1].revision:"无"}}</p>

            <h4>版本描述</h4>
            <tr
              v-for="(msg,i) in typeof currFoucsIdx==='number'?allDeps[currFoucsIdx][1].desc.slice(1,-2):[]"
              :key="i"
            >
              <p>{{msg}}</p>
            </tr>
          </div>
        </section>
        <footer class="modal-card-foot">
          <button class="button is-success">部署</button>
          <button class="button is-danger">删除</button>
        </footer>
      </div>
    </div>
  </div>
</template>


<script>
export default {
  name: "AllDep",
  data: function () {
    return {
      currFoucsIdx: null,
      modalActive: false,
      allDeps: [],
      timerId: null,
      progressList: [],
      progressDesc: [
        "初始化",
        "检出代码",
        "睡眠中断",
        "部署重启",
        "停止旧服务",
        "启动新服务",
        "发布成功 ",
      ],
    };
  },
  methods: {
    viewDetail: function (idx) {
      this.currFoucsIdx = idx;
      this.modalActive = true;
    },
    closeDetails: function () {
      this.currFoucsIdx = null;
      this.modalActive = false;
    },
    getAllDeps: async function () {
      try {
        const ret = await this.$mojoapi.get(`/web/db/sys:ops:depbil`);
        let tempList = [];
        //奇数下标的是有效信息
        for (let i = 0; i < ret.data.length; i += 2) {
          tempList[i / 2] = [ret.data[i], JSON.parse(ret.data[i + 1])];
        }
        tempList.sort((a, b) => {
          return parseInt(b[0]) - parseInt(a[0]);
        });
        this.allDeps = tempList;
      } catch (e) {
        this.$store.commit("error", `获取上线单失败 : ${e.data || e.message}`);
      }
    },
    getProgressList: async function () {
      try {
        const ret = await this.$mojoapi.get("/web/dep/progress");
        let tempList = [];
        for (let i = 0; i < ret.data.length; i += 5) {
          tempList[i / 5] = [
            ret.data[i],
            ret.data[i + 1],
            ret.data[i + 2],
            ret.data[i + 3],
            ret.data[i + 4],
          ];
        }
        tempList.sort((a, b) => {
          return parseInt(b[0]) - parseInt(a[0]);
        });
        this.progressList = tempList;
      } catch (e) {
        this.$store.commit(
          "error",
          `获取任务进度列表失败 : ${e.data || e.message}`
        );
      }
    },
    submit: async function (idx, force) {
      const ret = await this.$mojoapi.get(`/web/dep/submit`, {
        depid: this.allDeps[idx][0],
        force: force || false,
      });
      this.$store.commit("info", `成功 : ${ret.data}`);
    },
    showProgressDetails: function (depid) {
      for (let i = 0; i < this.allDeps.length; i++) {
        if (depid == this.allDeps[i][0]) {
          this.viewDetail(i);
          break;
        }
      }
    },
    getDepInfoByDepid: function (depid) {
      for (let i = 0; i < this.allDeps.length; i++) {
        if (depid == this.allDeps[i][0]) {
          return this.allDeps[i];
        }
      }
    },
  },
  beforeMount: function () {
    this.getAllDeps();
    this.getProgressList();
    this.timerId = setInterval(this.getProgressList, 1000);
  },
  destroyed: function () {
    if (this.timerId) {
      clearInterval(this.timerId);
      this.timerId = null;
    }
  },
};
</script>

<style scoped>
#details {
  width: 100%;
}
#progressTag {
  margin-top: 5px;
}
.progress {
  margin-top: -15px;
  margin-bottom: 5px;
}
</style>

