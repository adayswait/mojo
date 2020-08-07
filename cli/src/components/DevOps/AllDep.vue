<template>
  <div class="box">
    <div class="box">
      <table class="table is-striped is-fullwidth">
        <thead>
          <tr>
            <th>
              <abbr title="bucket名称">表名称</abbr>
            </th>
          </tr>
        </thead>
        <tbody>
          <tr v-for="(info,i) in progressList" :key="i">
            <td>
              <div class="columns">
                <div class="column is-1">
                  <p>{{info.depid}}</p>
                </div>
                <div class="column">
                  <progress
                    class="progress"
                    :class="{'is-danger':info.status===-1,
                    'is-success':info.status===5}"
                    :value="info.status/5"
                    max="1"
                  >15%</progress>
                </div>
                <div class="column is-3">
                  <p>{{info.desc}}</p>
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
                    <span class="tag is-success">{{info[1].rversion}}</span>
                  </div>
                </div>
                <span class="tag is-dark">status</span>
              </div>
            </td>
            <td>
              <button class="button is-small is-warning" @click="viewDetail(i)">查看详情</button>
              <button class="button is-small is-dark" @click="submit(i)">部署上线</button>
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
            <h4>版本号</h4>
            <p>{{typeof currFoucsIdx==='number'?allDeps[currFoucsIdx][1].rversion:"无"}}</p>

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
      progressList: [1, 2, 3, 4],
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
        const ret = await this.$httpc.get(`/web/db/sys:ops:depbil`);
        let tempList = [];
        //奇数下标的是有效信息
        for (let i = 0; i < ret.data.length; i += 2) {
          window.console.log(ret.data[i + 1]);
          tempList[i / 2] = [ret.data[i], JSON.parse(ret.data[i + 1])];
        }
        this.allDeps = tempList;
      } catch (e) {
        this.$store.commit("error", `获取上线单失败 : ${e.data || e.message}`);
      }
    },
    getProgressList: async function () {
      try {
        const ret = await this.$httpc.get("/web/dep/progress");
        this.progressList = ret.data;
      } catch (e) {
        this.$store.commit(
          "error",
          `获取任务进度列表失败 : ${e.data || e.message}`
        );
      }
    },
    submit: async function (idx) {
      window.console.log(this.allDeps[idx][0]);
      const ret = await this.$httpc.get(`/web/dep/submit`, {
        depid: this.allDeps[idx][0],
      });
      this.$store.commit("info", `成功 : ${ret.data}`);
    },
  },
  beforeMount: function () {
    this.getAllDeps();
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
</style>

