<template>
  <table class="table">
    <thead>
      <tr>
        <th>
          <abbr title="编号">编号</abbr>
        </th>
        <th>
          <abbr title="操作人">用户名</abbr>
        </th>
        <th>
          <abbr title="Won">权限</abbr>
        </th>
        <th>
          <abbr title="Goals for">操作</abbr>
        </th>
      </tr>
    </thead>
    <tbody class>
      <tr v-for="item in this.$store.state.mock.query" :key="item[0]">
        <th>{{item[0]}}</th>
        <td>{{item[1]}}</td>
        <td>{{item[2]}}</td>
        <td>
          <button class="button is-small is-rounded is-primary is-vcentered">查看操作</button>
          <button class="button is-small is-rounded is-warn">删除</button>
          <button class="button is-small is-rounded is-danger">变更权限</button>
        </td>
      </tr>
    </tbody>
  </table>
</template>
