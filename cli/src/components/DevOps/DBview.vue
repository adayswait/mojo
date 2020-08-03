<template>
  <div class="column is-10">
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
    <div class="box">
      <table class="table is-striped is-fullwidth has-text-centered" v-if="this.depth==0">
        <thead>
          <tr>
            <th>
              <abbr title="bucket名称">表名称</abbr>
            </th>
            <th>
              <abbr title="查看全表">查看</abbr>
            </th>
            <th>
              <abbr title="删除表">删除</abbr>
            </th>
          </tr>
        </thead>
        <tbody class>
          <tr v-for="item in this.tableList" :key="item[0]">
            <td>
              <input class="input" type="text" v-model="item[0]" readonly />
            </td>
            <td>
              <button class="button is-primary is-vcentered" @click="viewTable(item)">查看</button>
            </td>
            <td>
              <button class="button is-danger" @click="deleteTable(item)">删除</button>
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
              <abbr title="修改值">修改</abbr>
            </th>
            <th>
              <a class="button is-small is-rounded is-success is-vcentered" @click="newKv">新建</a>
            </th>
          </tr>
        </thead>
        <tbody class>
          <tr v-for="k in this.keyValueList" :key="k[0]">
            <td>
              <input class="input" type="text" v-model="k[0]" />
            </td>
            <td>
              <input class="input is-info" type="text" v-model="k[1]" />
            </td>
            <td>
              <button class="button is-primary is-vcentered" @click="changeKey(k)">修改</button>
            </td>
            <td>
              <button class="button is-danger" @click="deleteKey(k)">删除</button>
            </td>
          </tr>
        </tbody>
      </table>
    </div>
  </div>
</template>

<script>
export default {
  name: "DataView",
  data: function () {
    this.getAllDBs();
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
        const ret = await this.$httpc.get("/web/db");
        for (let key of ret.data) {
          this.tableList.push([key]);
        }
      } catch (e) {
        window.alert(e);
      }
    },
    viewTable: async function (tableName) {
      this.depth = 1;
      this.breadcrumbPath[1] = tableName;
      this.keyValueList = [];
      const ret = await this.$httpc.get(`/web/db/${tableName}`);
      let tempList = [];
      for (let i = 0; i < ret.data.length; i += 2) {
        tempList[i / 2] = [ret.data[i], ret.data[i + 1]];
      }
      this.keyValueList = tempList;
      window.console.log(this.keyValueList);
    },
    changeKey: async function (kv) {
      try {
        await this.$httpc.put(`/web/db/${this.breadcrumbPath[1]}/${kv[0]}`, {
          value: kv[1],
        });
      } catch (e) {
        window.console.log(e);
      }
    },
    deleteKey: async function (kv) {
      try {
        if (kv[0] === undefined || kv[0] === null || kv[0].toString() === "") {
          throw new Error("can not delete ''");
        }
        await this.$httpc.del(`/web/db/${this.breadcrumbPath[1]}/${kv[0]}`);
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
        window.console.log(e);
      }
    },
    deleteTable: async function (tableName) {
      window.console.log("删除", tableName);
    },
    setDepth: function (depth) {
      window.console.log("depth", depth);
      this.depth = depth;
    },
    newKv: function () {
      this.keyValueList = this.keyValueList.concat([[]]);
    },
    // copyValue: function (value) {
    //   window.console.log("depth", value);
    // },
  },
};
</script>

<style scoped>
</style>
