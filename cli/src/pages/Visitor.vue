<template>
  <div class="column">
    <router-view id="router" />
  </div>
</template>

<script>
export default {
  beforeMount() {
    this.$store.commit("setVisible", {
      name: "Header",
      visible: false,
    });
  },
  destroyed() {
    this.$store.commit("setVisible", {
      name: "Header",
      visible: true,
    });
  },
};
</script>
