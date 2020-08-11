<template>
    <div class="box">
      <table class="table is-striped is-fullwidth has-text-centered">
        <thead>
          <tr>
            <th>
              <abbr title="ip地址">ipv4地址</abbr>
            </th>
            <th>
              <abbr title="用于ssh登录的端口">ssh端口</abbr>
            </th>
            <th>
              <abbr title="用于ssh登录和操作用户名">用户名</abbr>
            </th>
            <th>
              <abbr title="ssh账号的密码">密码</abbr>
            </th>
            <th>
              <a class="button is-small is-rounded is-success is-vcentered" @click="newDepIni">新建</a>
            </th>
          </tr>
        </thead>
        <tbody class>
          <tr v-for="(k,i) in depIniList" :key="i">
            <td>
              <input class="input" type="text" v-model="k[1]" />
            </td>
            <td>
              <input class="input" type="text" v-model="k[2]" />
            </td>
            <td>
              <input class="input" type="text" v-model="k[3]" />
            </td>
            <td>
              <input class="input" type="text" v-model="k[4]" />
            </td>
            <td>
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
  name: "MacIni",
  data: function () {
    let dbMacTableName = "sys:ops:macini";
    return {
      title: "MacIni",
      currDropIdx: undefined,
      serverType: ["online", "battle", "match", "team", "center"],
      dbMacTableName: dbMacTableName,
      depIniList: [["1", "10.1.1.1", "22", "user", "password"]],
    };
  },
  methods: {
    getAllMacIni: async function () {
      try {
        const ret = await this.$httpc.get(`/web/db/${this.dbMacTableName}`);
        let tempList = [];
        for (let i = 0; i < ret.data.length; i += 2) {
          tempList[i / 2] = [ret.data[i]].concat(JSON.parse(ret.data[i + 1]));
        }
        this.depIniList = tempList;
      } catch (e) {
        this.$store.commit(
          "error",
          `获取数据库表错误 : ${e.data || e.message}`
        );
      }
    },
    newDepIni: function () {
      this.depIniList.splice(0, 0, [undefined]);
    },
    changeServerType: function (idx, type) {
      this.depIniList[idx][1] = type;
      this.$forceUpdate();
    },
    uploadChange: async function (idx) {
      if (
        this.depIniList[idx][1] &&
        this.depIniList[idx][2] &&
        this.depIniList[idx][3] &&
        this.depIniList[idx][4]
      ) {
        try {
          if (this.depIniList[idx][0]) {
            await this.$httpc.put(
              `/web/db/${this.dbMacTableName}/${this.depIniList[idx][0]}`,
              {
                value: JSON.stringify(this.depIniList[idx].slice(1)),
              }
            );
            this.$store.commit("info", `修改机器配置成功`);
          } else {
            await this.$httpc.post(`/web/db/${this.dbMacTableName}`, {
              value: JSON.stringify(this.depIniList[idx].slice(1)),
            });
            this.$store.commit("info", `新建机器配置成功`);
          }
          await this.getAllMacIni();
        } catch (e) {
          this.$store.commit(
            "error",
            `修改机器配置错误 : ${e.data || e.message}`
          );
        }
      } else {
        this.$store.commit("warn", `修改机器配置错误 : 请检查你的数据格式`);
      }
    },
    deleteIni: async function (idx) {
      if (this.depIniList[idx][0]) {
        try {
          await this.$httpc.del(
            `/web/db/${this.dbMacTableName}/${this.depIniList[idx][0]}`
          );
          this.depIniList.splice(idx, 1);
          this.$store.commit("info", `删除机器配置成功`);
        } catch (e) {
          this.$store.commit(
            "error",
            `删除机器配置错误 : ${e.data || e.message}`
          );
        }
      } else {
        this.$store.commit("warn", `删除机器配置错误 : 无效项`);
      }
    },
  },
  beforeMount: function () {
    this.getAllMacIni();
  },
};
</script>

<style scoped>
#opsbtn {
  margin-top: 4px;
}
</style>
