<template>
  <div class="container">
    <div class="columns">
      <div class="column">
        <p>
          <a :href="luyi" target="_blank">{{luyi}}</a>
        </p>
        <button class="delete is-danger" @click="changeVisible(0)"></button>
        <button class="button is-dark" v-if="!visibleExtra[0]" @click="changeVisible(0)">开始加载</button>
        <iframe width="640px" height="360px" :src="luyi" v-if="visibleExtra[0]"></iframe>
      </div>
      <div class="column">
        <p>
          <a :href="tinywu" target="_blank">{{tinywu}}</a>
        </p>
        <button class="delete is-danger" @click="changeVisible(1)"></button>
        <button class="button is-dark" v-if="!visibleExtra[1]" @click="changeVisible(1)">开始加载</button>
        <iframe width="640px" height="360px" :src="tinywu" v-if="visibleExtra[1]"></iframe>
      </div>
    </div>
    <div class="columns">
      <div class="column">
        <p>
          <a :href="liukaichang" target="_blank">{{liukaichang}}</a>
        </p>
        <button class="delete is-danger" @click="changeVisible(2)"></button>
        <button class="button is-dark" v-if="!visibleExtra[2]" @click="changeVisible(2)">开始加载</button>
        <iframe width="640px" height="360px" :src="liukaichang" v-if="visibleExtra[2]"></iframe>
      </div>
      <div class="column">
        <p>
          <a :href="qa" target="_blank">{{qa}}</a>
        </p>
        <button class="delete is-danger" @click="changeVisible(3)"></button>
        <button class="button is-dark" v-if="!visibleExtra[3]" @click="changeVisible(3)">开始加载</button>
        <iframe width="640px" height="360px" :src="qa" v-if="visibleExtra[3]"></iframe>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  data: function () {
    return {
      visibleExtra: [false, false, false, false],
      luyi: "http://luyi.taomee-ex.com/",
      tinywu: "http://tinywu.taomee-ex.com/",
      liukaichang: "http://liukaichang.taomee-ex.com/",
      qa: "http://plan_s.taomee-ex.com/",
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
  padding-top: 0px;
  width: 670px;
  height: 400px;
}
.button is-dark {
  /* position: absolute; */
  margin: auto auto;
}
.delete {
  position: absolute;
}
</style>
