
<template>
  <div class="tile is-ancestor">
    <div class="tile is-vertical is-8">
      <div class="tile">
        <div class="tile is-parent is-vertical">
          <article class="tile is-child notification is-primary">
            <p class="title">对开发者提点建议</p>
            <p class="subtitle">( ´◔ ‸◔`)</p>
            <textarea class="textarea" placeholder="我觉得, 这个地方有问题" v-model="developerText"></textarea>
            <br />
            <nav class="level is-mobile">
              <div class="level-item has-text-centered">
                <button
                  class="button is-small is-rounded is-dark is-vcentered"
                  :class="{'is-loading':btnDevLoading}"
                  @click="chat2Developer"
                >(ﾉ◕ヮ◕)ﾉ*:･ﾟ✧</button>
              </div>
            </nav>
          </article>
          <article class="tile is-child notification is-warning">
            <p class="title">(๑◕ܫ￩๑)b</p>
            <p class="subtitle">under developing</p>
          </article>
        </div>
        <div class="tile is-parent">
          <article class="tile is-child notification is-info">
            <p class="title">匿名聊天</p>
            <p class="subtitle">◍'ㅅ'◍</p>
            <textarea
              class="textarea"
              rows="20"
              placeholder="随便说两句吧, 反正没人知道我是谁ᐕ)⁾⁾"
              v-model="groupText"
            ></textarea>
            <br />
            <nav class="level is-mobile">
              <div class="level-item has-text-centered">
                <button
                  class="button is-small is-rounded is-dark is-vcentered"
                  :class="{'is-loading':btnGroupLoading}"
                  @click="chat2Group"
                >匿名发布到 ❥S计划讨论群</button>
              </div>
            </nav>
          </article>
        </div>
      </div>
      <div class="tile is-parent">
        <article class="tile is-child notification is-danger">
          <p class="title">..(｡•ˇ‸ˇ•｡)…</p>
          <p class="subtitle">under developing</p>
          <div class="content">
            <!-- Content -->
          </div>
        </article>
      </div>
    </div>
    <div class="tile is-parent">
      <article class="tile is-child notification is-success">
        <div class="content">
          <p class="title">☝ᖗ乛◡乛ᖘ☝</p>
          <p class="subtitle">under developing</p>
          <div class="content">
            <!-- Content -->
          </div>
        </div>
      </article>
    </div>
  </div>
</template>


<script>
export default {
  name: "Issue",
  data: function () {
    return {
      title: "Issue",
      developerText: "",
      groupText: "",
      btnDevLoading: false,
      btnGroupLoading: false,
    };
  },
  methods: {
    chat2Developer: async function () {
      if (this.developerText.length === 0) {
        return this.$store.commit("warn", `什么都没说呀`);
      }
      try {
        this.btnDevLoading = true;
        await this.$mojoapi.post(`/web/chat/dev`, {
          message: this.developerText,
        });
        this.$store.commit("info", `我们已经收到你的建议啦(●･◡･●)ﾉ♥`);
        this.developerText = "";
      } catch (e) {
        this.$store.commit("error", `聊天错误 : ${e.data || e.message}`);
      }
      this.btnDevLoading = false;
    },
    chat2Group: async function () {
       if (this.groupText.length === 0) {
        return this.$store.commit("warn", `什么都没说呀`);
      }
      try {
        this.btnGroupLoading = true;
        await this.$mojoapi.post(`/web/chat/group`, {
          message: this.groupText,
        });
        this.$store.commit("info", `☝ᖗ乛◡乛ᖘ☝ 完美`);
        this.groupText = "";
      } catch (e) {
        this.$store.commit("error", `聊天错误 : ${e.data || e.message}`);
      }
      this.btnGroupLoading = false;
    },
  },
};
</script>