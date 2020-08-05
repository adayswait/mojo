<template>
  <div class="column is-10">
    <div class="box">
      <table class="table is-striped is-fullwidth has-text-centered">
        <thead>
          <tr>
            <th>
              <abbr title="机器ip">地址</abbr>
            </th>
            <th>
              <abbr title="ssh端口">ssh端口</abbr>
            </th>
            <th>
              <abbr title="用户名">用户名</abbr>
            </th>
            <th>
              <abbr title="密码">密码</abbr>
            </th>
            <th>
              <a class="button is-small is-rounded is-success is-vcentered" @click="newMacIni">新建</a>
            </th>
          </tr>
        </thead>
        <tbody class>
          <tr v-for="k in this.macList" :key="k[0]">
            <td>{{k[0]}}</td>
            <td>{{k[1]}}</td>
            <td>{{k[2]}}</td>
            <td>{{k[3]}}</td>
            <td>
              <button class="button is-primary is-small is-vcentered">修改</button>
              <button class="button is-danger is-small">删除</button>
            </td>
          </tr>
        </tbody>
      </table>
    </div>
  </div>
</template>

<script>
export default {
  name: "MacIni",
  data: function () {
    return {
      title: "MacIni",
      macList: [
        ["10.1.1.1", 22, "splan", "password"],
        ["10.1.1.2", 56000, "splan", "password"],
      ],
    };
  },
  methods: {
    newMacIni: function () {
      this.macList.splice(0, 0, []);
    },
  },
};
</script>

<style scoped>
</style>
