<template>
  <div id="notification" class="columns is-multiline is-vcentered">
    <div class="column">
      <tr v-for="message in $store.state.messageList" :key="message[0]">
        <div class="notification is-primary">
          <button class="delete"></button>
          {{message[1]}}
        </div>
      </tr>
    </div>
  </div>
</template>


<script>
export default {
  name: "DevOps",
  data: function () {
    return {};
  },
  methods: {
    addMessage: function () {
      this.$store.commit("pushMessage", "message");
    },
  },
};
</script>

<style scoped>
#notification {
  position: absolute; /*设定footer绝对位置在底部*/
  right: 0px;
  z-index: 100;
}
</style>
