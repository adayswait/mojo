<template>
  <div class="box">
    <table class="table is-striped is-fullwidth">
      <thead>
        <tr>
          <th>
            <abbr title="标题"></abbr>
          </th>
          <th>
            <abbr title="服务类型">服务类型</abbr>
          </th>
          <th>
            <abbr title="版本号">版本号</abbr>
          </th>
          <th>
            <abbr title="操作">操作</abbr>
          </th>
        </tr>
      </thead>
      <tbody>
        <tr v-for="(info,i) in allDeps" :key="info[0]">
          <td>
            <span class="tag is-black">未上线</span>
            {{info[1].title}}
          </td>
          <td>{{info[1].type}}</td>
          <td>{{info[1].rversion}}</td>
          <td>
            <button class="button is-small is-warning" @click="viewDetail(i)">查看详情</button>
            <button class="button is-small is-dark" @click="submit(i)">部署上线</button>
          </td>
        </tr>
      </tbody>
    </table>

    <div id="details" class="modal is-large" :class="{'is-active':modalActive}">
      <div class="modal-background"></div>
      <div class="modal-card">
        <header class="modal-card-head">
          <p class="modal-card-title">{{currFoucsIdx?allDeps[currFoucsIdx][1].title:""}}</p>
          <button class="delete" aria-label="close" @click="closeDetails"></button>
        </header>
        <section class="modal-card-body">
          <div class="content">
            <h4>服务类型</h4>
            <p>{{currFoucsIdx?allDeps[currFoucsIdx][1].type:"无"}}</p>
            <h4>checkout地址</h4>
            <p>{{currFoucsIdx?allDeps[currFoucsIdx][1].repourl:"无"}}</p>
            <h4>版本号</h4>
            <p>{{currFoucsIdx?allDeps[currFoucsIdx][1].rversion:"无"}}</p>

            <h4>版本描述</h4>
            <tr v-for="(msg,i) in currFoucsIdx?allDeps[currFoucsIdx][1].desc:[]" :key="i">
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
    };
  },
  methods: {
    // focus: function (idx) {},
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
  },
};
</script>

<style scoped>
#details {
  width: 100%;
}
</style>

