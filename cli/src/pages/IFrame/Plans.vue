<template>
  <div class="container">
    <br />
    <br />
    <div class="columns">
      <div class="column">
        <button class="delete is-danger" @click="changeVisible(0)"></button>
        <button
          class="button is-dark"
          v-if="!visibleExtra[0]"
          @click="changeVisible(0)"
        >{{buttonText}}</button>
        <iframe
          title="plan_s"
          width="640px"
          height="360px"
          src="http://plan_s.taomee-ex.com/"
          v-if="visibleExtra[0]"
        ></iframe>
      </div>
      <div class="column">
        <button class="delete is-danger" @click="changeVisible(1)"></button>
        <button
          class="button is-dark"
          v-if="!visibleExtra[1]"
          @click="changeVisible(1)"
        >{{buttonText}}</button>
        <iframe
          title="plan_s"
          width="640px"
          height="360px"
          src="http://plan_s.taomee-ex.com/"
          v-if="visibleExtra[1]"
        ></iframe>
      </div>
    </div>
    <div class="columns">
      <div class="column">
        <button class="delete is-danger" @click="changeVisible(2)"></button>
        <button
          class="button is-dark"
          v-if="!visibleExtra[2]"
          @click="changeVisible(2)"
        >{{buttonText}}</button>
        <iframe
          title="plan_s"
          width="640px"
          height="360px"
          src="http://plan_s.taomee-ex.com/"
          v-if="visibleExtra[2]"
        ></iframe>
      </div>
      <div class="column">
        <button class="delete is-danger" @click="changeVisible(3)"></button>
        <button
          class="button is-dark"
          v-if="!visibleExtra[3]"
          @click="changeVisible(3)"
        >{{buttonText}}</button>
        <iframe
          title="plan_s"
          width="640px"
          height="360px"
          src="http://plan_s.taomee-ex.com/"
          v-if="visibleExtra[3]"
        ></iframe>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  data: function () {
    return {
      buttonText: "再开一个",
      visibleExtra: [true, false, false, false],
    };
  },
  methods: {
    changeVisible(idx) {
      this.visibleExtra[idx] = !this.visibleExtra[idx];
      window.console.log(this.visibleExtra);
      this.$forceUpdate();
    },
  },
};
</script>

<style scoped>
.column {
  text-align: center;
  border: 1px solid #dcdcdc;
  margin-right: 20px;
  margin-bottom: 20px;
  width: 640px;
  height: 390px;
}
.button is-dark {
  position: absolute;
  margin: auto auto;
}
.delete {
  position: absolute;
}
</style>
