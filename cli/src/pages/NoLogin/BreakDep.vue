<template>
  <div class="column is-4 is-offset-4">
    <p>{{$route.query}}</p>
  </div>
</template>

<script>
export default {
  data: function () {
    return {
      title: "BreakDep",
    };
  },
  beforeMount: function () {
    window.console.log("params", this.$route.query);
  },
};
</script>