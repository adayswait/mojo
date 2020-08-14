<template>
  <div class="box">
    <table class="table is-striped is-fullwidth has-text-centered">
      <thead>
        <tr>
          <th>
            <abbr title="服务类型">服务类型</abbr>
          </th>
          <th>
            <abbr title="仓库地址">仓库地址</abbr>
          </th>
          <th>
            <a class="button is-small is-rounded is-success is-vcentered" @click="newDepIni">新建</a>
          </th>
        </tr>
      </thead>
      <tbody class>
        <tr v-for="(k,i) in devIniList" :key="i">
          <td style="width: 20%;">
            <input class="input" type="text" v-model="k[0]" />
          </td>
          <td>
            <input class="input" type="text" v-model="k[1].url" />
          </td>
          <td style="width: 20%;">
            <button
              id="opsbtn"
              class="button is-primary is-small is-vcentered"
              @click="uploadChange(i)"
            >修改</button>
            <button id="opsbtn" class="button is-danger is-small" @click="deleteIni(i)">删除</button>
          </td>
        </tr>
      </tbody>
    </table>
  </div>
</template>

<script>
export default {
  name: "DepIni",
  data: function () {
    return {
      title: "DepIni",
      currDropIdx: undefined,
      dbDevTableName: "sys:ops:devini",
      devIniList: [["online", "http://svn.xxx.com/server/online"]],
    };
  },
  methods: {
    getAllDevIni: async function () {
      try {
        const ret = await this.$mojoapi.get(`/web/db/${this.dbDevTableName}`);
        let tempList = [];
        for (let i = 0; i < ret.data.length; i += 2) {
          tempList[i / 2] = [ret.data[i], JSON.parse(ret.data[i + 1])];
        }
        this.devIniList = tempList;
      } catch (e) {
        this.$store.commit(
          "error",
          `获取数据库表错误 : ${e.data || e.message}`
        );
      }
    },
    newDepIni: function () {
      this.devIniList.splice(0, 0, [undefined, {}]);
    },
    changeServerType: function (idx, type) {
      this.devIniList[idx][0] = type;
      this.$forceUpdate();
    },
    uploadChange: async function (idx) {
      if (this.devIniList[idx][1] && this.devIniList[idx][1]) {
        try {
          if (this.devIniList[idx][0]) {
            await this.$mojoapi.put(
              `/web/db/${this.dbDevTableName}/${this.devIniList[idx][0]}`,
              {
                value: JSON.stringify({
                  url: this.devIniList[idx][1].url,
                }),
              }
            );
            this.$store.commit("info", `修改开发配置成功`);
          } else {
            this.$store.commit("error", `修改开发配置错误 : 服务名不能为空`);
          }
          await this.getAllDevIni();
        } catch (e) {
          this.$store.commit(
            "error",
            `修改开发配置错误 : ${e.data || e.message}`
          );
        }
      } else {
        this.$store.commit("warn", `修改开发配置错误 : 请检查你的数据格式`);
      }
    },
    deleteIni: async function (idx) {
      if (this.devIniList[idx][0] && this.devIniList[idx][1]) {
        try {
          await this.$mojoapi.del(
            `/web/db/${this.dbDevTableName}/${this.devIniList[idx][0]}`
          );
          this.devIniList.splice(idx, 1);
          this.$store.commit("info", `删除开发配置成功`);
        } catch (e) {
          this.$store.commit(
            "error",
            `删除开发配置错误 : ${e.data || e.message}`
          );
        }
      } else {
        this.$store.commit("warn", `删除开发配置错误 : 无效项`);
      }
    },
  },
  beforeMount: function () {
    this.getAllDevIni();
  },
};
</script>

<style scoped>
#opsbtn {
  margin-top: 4px;
}
</style>
