<template>
  <div class="box">
    <nav class="breadcrumb" aria-label="breadcrumbs">
      <ul>
        <li v-if="this.depth>=0">
          <a @click="setDepth(0)">ROOT</a>
        </li>
        <li v-if="this.depth>=1">
          <a href="#">表 {{this.breadcrumbPath[1]}}</a>
        </li>
        <li v-if="this.depth>=2">
          <a href="#">键 {{this.breadcrumbPath[2]}}</a>
        </li>
      </ul>
    </nav>
    <table class="table is-striped is-fullwidth has-text-centered" v-if="this.depth==0">
      <thead>
        <tr>
          <th>
            <abbr title="bucket名称">表名称</abbr>
          </th>
          <th>
            <abbr title="操作">操作</abbr>
          </th>
        </tr>
      </thead>
      <tbody class v-if="$store.getters.userInfo.group<=1">
        <tr v-for="item in this.tableList" :key="item[0]">
          <td>
            <input class="input" type="text" v-model="item[0]" readonly />
          </td>
          <td>
            <button
              id="opsbtn"
              class="button is-primary is-vcentered is-small"
              @click="viewTable(item)"
            >查看</button>
            <button id="opsbtn" class="button is-danger is-small" @click="deleteTable(item)">删除</button>
          </td>
        </tr>
      </tbody>
    </table>

    <table class="table is-striped is-fullwidth has-text-centered" v-if="this.depth==1">
      <thead>
        <tr>
          <th>
            <abbr title="键名称">键</abbr>
          </th>
          <th>
            <abbr title="值">值</abbr>
          </th>
          <th>
            <a class="button is-small is-rounded is-success is-vcentered" @click="newKv">新建</a>
          </th>
        </tr>
      </thead>
      <tbody class v-if="$store.getters.userInfo.group<=1">
        <tr v-for="k in this.keyValueList" :key="k[0]">
          <td style="width:20%">
            <input class="input" type="text" v-model="k[0]" />
          </td>
          <td>
            <input class="input is-info" type="text" v-model="k[1]" />
          </td>
          <td style="width:20%">
            <button
              id="opsbtn"
              class="button is-primary is-vcentered is-small"
              @click="changeKey(k)"
            >修改</button>
            <button id="opsbtn" class="button is-danger is-small" @click="deleteKey(k)">删除</button>
          </td>
        </tr>
      </tbody>
    </table>
  </div>
</template>

<script>
export default {
  name: "DataView",
  data: function () {
    return {
      title: "DataView",
      breadcrumbPath: [],
      tableList: [],
      keyValueList: [
        // [
        //   "00dd597f-40dc-43ff-852c-a6e7a236e025",
        //   { user: "jesse", group: 2 },
        // ],
      ],
      depth: 0,
    };
  },
  methods: {
    getAllDBs: async function () {
      try {
        const ret = await this.$mojoapi.get("/web/db");
        for (let key of ret.data) {
          this.tableList.push([key]);
        }
      } catch (e) {
        this.$store.commit(
          "error",
          `获取数据库表错误 : ${e.data || e.message}`
        );
      }
    },
    viewTable: async function (tableName) {
      try {
        const ret = await this.$mojoapi.get(`/web/db/${tableName}`);
        this.depth = 1;
        this.breadcrumbPath[1] = tableName;
        this.keyValueList = [];
        let tempList = [];
        for (let i = 0; i < ret.data.length; i += 2) {
          tempList[i / 2] = [ret.data[i], ret.data[i + 1]];
        }
        this.keyValueList = tempList;
      } catch (e) {
        this.$store.commit(
          "error",
          `查看表${tableName}错误 : ${e.data || e.message}`
        );
      }
    },
    changeKey: async function (kv) {
      if (this.$store.state.userInfo.group > 1) {
        return this.$store.commit("error", `修改数据库值错误 : 权限不足`);
      }
      try {
        await this.$mojoapi.put(`/web/db/${this.breadcrumbPath[1]}/${kv[0]}`, {
          value: kv[1],
        });
      } catch (e) {
        this.$store.commit(
          "error",
          `修改数据库值错误 : ${e.data || e.message}`
        );
      }
    },
    deleteKey: async function (kv) {
      try {
        if (kv[0] === undefined || kv[0] === null || kv[0].toString() === "") {
          throw new Error("无法删除空值");
        }
        await this.$mojoapi.del(`/web/db/${this.breadcrumbPath[1]}/${kv[0]}`);
        let delIndex = null;
        for (let i = 0; i < this.keyValueList.length; i++) {
          if (kv[0] === this.keyValueList[i][0]) {
            delIndex = i;
            break;
          }
        }
        if (delIndex !== null) {
          this.keyValueList.splice(delIndex, 1);
        }
      } catch (e) {
        this.$store.commit("error", `删除键值对错误 : ${e.data || e.message}`);
      }
    },
    deleteTable: async function (tableName) {
      this.$store.commit("warn", `删除表${tableName}错误 : 尚未实现`);
    },
    setDepth: function (depth) {
      this.depth = depth;
    },
    newKv: function () {
      this.keyValueList.splice(0, 0, []);
    },
  },
  beforeMount: function () {
    this.getAllDBs();
  },
};
</script>

<style scoped>
#opsbtn {
  margin-top: 4px;
}
</style>
