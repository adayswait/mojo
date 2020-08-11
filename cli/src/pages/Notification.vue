<template>
  <div id="blackboard" class="columns is-multiline is-vcentered">
    <div class="column">
      <tr v-for="message in $store.state.messageList" :key="message[0]">
        <div
          id="notice"
          class="notification"
          :class="{
            'is-danger':message[1]==$store.state.MESSAGE_TYPE.ERROR,
            'is-warning':message[1]==$store.state.MESSAGE_TYPE.WARN,
            'is-info':message[1]==$store.state.MESSAGE_TYPE.INFO,
          }"
        >
          <!-- <button class="delete"></button> -->
          {{message[2]}}
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
  methods: {},
};
</script>

<style scoped>
#blackboard {
  position: absolute;
  right: 0px;
  z-index: 100;
  filter: alpha(Opacity=50);
  -moz-opacity: 0.5;
  opacity: 0.5;
  pointer-events: none;
  margin-top: 20px;
}
#notice {
  margin-bottom: 2px;
}
</style>
