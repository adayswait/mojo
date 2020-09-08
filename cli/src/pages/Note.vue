<template>
  <div class="container">
    <Login />
    <table class="table is-striped is-fullwidth" v-if="!$store.state.visible.Login">
      <tr v-for="(note,i) in notes" :key="i">
        <div class="box">
          <article class="media is-fullwidth">
            <!-- <figure class="media-left">
              <p class="image is-64x64">
                <img src="./../assets/logo.png" />
              </p>
            </figure>-->
            <div class="media-content">
              <div class="content">
                <p>
                  <strong>{{note.user}}</strong>
                  <small>
                    发布于
                    {{new Date(note.time).toLocaleString()}}
                  </small>
                </p>
                <p v-html="note.content"></p>
              </div>
            </div>
          </article>
        </div>
      </tr>
    </table>
    <br />
    <br />
    <article class="media" v-if="!$store.state.visible.Login">
      <!-- <figure class="media-left">
        <p class="image is-64x64">
          <img src="./../assets/logo.png" />
        </p>
      </figure>-->
      <div class="media-content">
        <div class="field" style="height:350px;">
          <p class="control">
            <quill-editor
              ref="myTextEditor"
              v-model="noteContent"
              :options="editorOption"
              style="height:300px;"
            ></quill-editor>
          </p>
        </div>
        <div class="field">
          <div class="level-item has-text-centered">
            <p class="control">
              <button
                class="button is-primary"
                @click="submit"
                :class="{'is-loading':submitting}"
              >记录笔记</button>
            </p>
          </div>
        </div>
      </div>
    </article>
  </div>
</template>

<script>
import { quillEditor } from "vue-quill-editor";
import "quill/dist/quill.core.css";
import "quill/dist/quill.snow.css";
import "quill/dist/quill.bubble.css";
import Login from "@/pages/Login.vue";
export default {
  data: function () {
    return {
      tableName: "sys:usr:notepd",
      noteContent: "",
      notes: [],
      editorOption: {
        placeholder: "编辑内容",
      },
      submitting: false,
    };
  },
  components: {
    quillEditor,
    Login,
  },
  methods: {
    loadNotes: async function () {
      const ret = await this.$mojoapi.get(`/web/db/${this.tableName}`);
      let tempList = [];
      for (let i = 0; i < ret.data.length; i += 2) {
        tempList.splice(0, 0, JSON.parse(ret.data[i + 1]));
      }
      this.notes = tempList;
    },
    submit: async function () {
      if (this.noteContent.length === 0) {
        return this.$store.commit("warn", `内容不能为空`);
      }
      try {
        this.submitting = true;
        await this.$mojoapi.put(`/web/db/${this.tableName}`, {
          value: JSON.stringify({
            user: this.$store.getters.userInfo.user,
            time: Date.now(),
            content: this.noteContent,
          }),
        });
        this.submitting = false;
        this.noteContent = "";
        this.$store.commit("info", `笔记已成功记录`);
      } catch (e) {
        this.submitting = false;
        this.$store.commit(
          "error",
          `记录笔记数据错误 : ${e.data || e.message}`
        );
      }
      await this.loadNotes();
    },
  },
  beforeMount: function () {
    this.loadNotes();
  },
};
</script>

<style scoped>
.container {
  margin-top: 2%;
}
.box {
  margin-bottom: 10px;
}
tr {
  margin: 0;
}
</style>