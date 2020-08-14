<template>
  <div class="box">
    <table class="table is-striped is-fullwidth has-text-centered">
      <thead>
        <tr>
          <th>
            <abbr title="服务类型">服务类型</abbr>
          </th>
          <th>
            <abbr title="服务名">服务名</abbr>
          </th>
          <th>
            <abbr title="地址">地址</abbr>
          </th>
          <th>
            <abbr title="路径">路径</abbr>
          </th>
          <th>
            <a class="button is-small is-rounded is-success is-vcentered" @click="newDepIni">新建</a>
          </th>
        </tr>
      </thead>
      <tbody class>
        <tr v-for="(k,i) in depIniList" :key="i">
          <td>
            <div class="dropdown is-hoverable">
              <div class="dropdown-trigger">
                <button class="button" aria-haspopup="true" aria-controls="dropdown-menu">
                  {{k[1]}}
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
                      :class="{'is-active':t==k[1]}"
                      @click="changeServerType(i,t)"
                    >{{t}}</a>
                  </tr>
                </div>
              </div>
            </div>
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
  name: "DepIni",
  data: function () {
    return {
      title: "DepIni",
      currDropIdx: undefined,
      dbDepTableName: "sys:ops:depini",
      depIniList: [
        ["1", "online", "example", "10.1.1.1", "/opt/splan/example"],
      ],
      depTypeList: [],
    };
  },
  methods: {
    getAllDepIni: async function () {
      try {
        const ret = await this.$mojoapi.get(`/web/db/${this.dbDepTableName}`);
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
            await this.$mojoapi.put(
              `/web/db/${this.dbDepTableName}/${this.depIniList[idx][0]}`,
              {
                value: JSON.stringify(this.depIniList[idx].slice(1)),
              }
            );
            this.$store.commit("info", `修改部署配置成功`);
          } else {
            await this.$mojoapi.post(`/web/db/${this.dbDepTableName}`, {
              value: JSON.stringify(this.depIniList[idx].slice(1)),
            });
            this.$store.commit("info", `新建部署配置成功`);
          }
          await this.getAllDepIni();
        } catch (e) {
          this.$store.commit(
            "error",
            `修改部署配置错误 : ${e.data || e.message}`
          );
        }
      } else {
        this.$store.commit("warn", `修改部署配置错误 : 请检查你的数据格式`);
      }
    },
    deleteIni: async function (idx) {
      if (this.depIniList[idx][0]) {
        try {
          await this.$mojoapi.del(
            `/web/db/${this.dbDepTableName}/${this.depIniList[idx][0]}`
          );
          this.depIniList.splice(idx, 1);
          this.$store.commit("info", `删除部署配置成功`);
        } catch (e) {
          this.$store.commit(
            "error",
            `删除部署配置错误 : ${e.data || e.message}`
          );
        }
      } else {
        this.$store.commit("warn", `删除部署配置错误 : 无效项`);
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
  },
  beforeMount: function () {
    this.getAllDepIni();
    this.getAllDepType();
  },
};
</script>

<style scoped>
#opsbtn {
  margin-top: 4px;
}
</style>
