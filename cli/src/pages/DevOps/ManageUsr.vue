<template>
  <div class="box">
    <table class="table is-striped is-fullwidth has-text-centered">
      <thead>
        <tr>
          <th>
            <abbr title="用户名">用户名</abbr>
          </th>
          <th>
            <abbr title="用户组">用户组</abbr>
          </th>
          <th>
            <abbr title="操作">修改用户组</abbr>
          </th>
        </tr>
      </thead>
      <tbody>
        <tr v-for="kv in this.userList" :key="kv[0]">
          <td>{{kv[1].user}}</td>
          <td>{{$store.state.GROUP[kv[1].group]}}</td>
          <td>
            <button class="button is-small is-warning" @click="changeRight(true,kv)">提权</button>
            <button class="button is-small is-dark" @click="changeRight(false,kv)">降权</button>
          </td>
        </tr>
      </tbody>
    </table>
  </div>
</template>


<script>
export default {
  name: "QueryUsr",
  data: function () {
    this.dbTable = "sys:token:info";
    this.getAllUser();
    return {
      title: "QueryUsr",
      userList: [],
    };
  },
  methods: {
    getAllUser: async function () {
      try {
        const ret = await this.$httpc.get(`/web/db/${this.dbTable}`);
        let tempList = [];
        for (let i = 0; i < ret.data.length; i += 2) {
          tempList[i / 2] = [ret.data[i], JSON.parse(ret.data[i + 1])];
        }
        this.userList = tempList;
      } catch (e) {
        this.$store.commit(
          "error",
          `获取所有注册用户错误 : ${e.data || e.message}`
        );
      }
    },
    changeRight: async function (b, kv) {
      let newGroup;
      if (b === true) {
        newGroup = kv[1].group - 1;
      } else {
        newGroup = kv[1].group + 1;
      }
      if (this.$store.getters.GROUP[newGroup]) {
        if (this.$store.getters.userInfo.group >= newGroup) {
          this.$store.commit("error", `改变用户权限组错误 : 权限不足`);
          return;
        }
        try {
          await this.$httpc.put(`/web/db/${this.dbTable}/${kv[0]}`, {
            value: JSON.stringify({
              user: kv[1].user,
              group: newGroup,
            }),
          });
          kv[1].group = newGroup;
          if (this.$store.getters.userInfo.user === kv[1].user) {
            this.$store.commit("setUserInfo", {
              user: kv[1].user,
              group: kv[1].group,
            });
          }
        } catch (e) {
          this.$store.commit(
            "error",
            `改变用户权限组错误 : ${e.data || e.message}`
          );
        }
      } else {
        this.$store.commit(
          "error",
          `改变用户权限组错误 : invalid group:${kv[1].group}`
        );
      }
    },
  },
};
</script>
