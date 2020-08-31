<template>
  <div class="container">
    <!-- <div class="tile is-ancestor">
      <div class="tile is-parent">
        <article class="tile is-child box">
          <p class="title">One</p>
          <p class="subtitle">Subtitle</p>
        </article>
      </div>
      <div class="tile is-parent">
        <article class="tile is-child box">
          <p class="title">Two</p>
          <p class="subtitle">Subtitle</p>
        </article>
      </div>
      <div class="tile is-parent">
        <article class="tile is-child box">
          <p class="title">Three</p>
          <p class="subtitle">Subtitle</p>
        </article>
      </div>
      <div class="tile is-parent">
        <article class="tile is-child box">
          <p class="title">Four</p>
          <p class="subtitle">Subtitle</p>
        </article>
      </div>
    </div>-->

    <div class="tile is-ancestor" id="navigator">
      <div class="tile is-vertical is-9">
        <div class="tile">
          <div class="tile is-parent">
            <article class="tile is-child box">
              <p class="title">运维</p>
              <!-- <p class="subtitle">操作和维护线上服务</p> -->
              <div class="content">
                <p>
                  <strong>操作和维护线上服务</strong>
                </p>
              </div>
            </article>
          </div>
          <div class="tile is-8 is-vertical">
            <div class="tile">
              <div class="tile is-parent">
                <article class="tile is-child box">
                  <p class="title">center server</p>
                  <p class="subtitle">
                    <strong>操作raft集群</strong>
                  </p>
                </article>
              </div>
              <div class="tile is-parent">
                <article class="tile is-child box">
                  <p class="title">Seven</p>
                  <p class="subtitle">Subtitle</p>
                </article>
              </div>
            </div>
            <div class="tile is-parent">
              <article class="tile is-child box">
                <p class="title">Eight</p>
                <p class="subtitle">Subtitle</p>
              </article>
            </div>
          </div>
        </div>
        <div class="tile">
          <div class="tile is-8 is-parent">
            <article class="tile is-child box">
              <p class="title">Nine</p>
              <p class="subtitle">Subtitle</p>
              <div class="content">
                <p>facilisis.</p>
              </div>
            </article>
          </div>
          <div class="tile is-parent">
            <article class="tile is-child box">
              <p class="title">Ten</p>
              <p class="subtitle">Subtitle</p>
              <div class="content">
                <p>Lorem ipsum</p>
              </div>
            </article>
          </div>
        </div>
      </div>
      <div class="tile is-parent">
        <article class="tile is-child box">
          <div class="content">
            <p class="title">Eleven</p>
            <p class="subtitle">Subtitle</p>
            <div class="content">
              <p>Lorem ipsum dolor sit amet</p>
              <p>Integer sollicitudin, tortor a mattis commodo, velit urna</p>
            </div>
          </div>
        </article>
      </div>
    </div>

    <!-- <div class="tile is-ancestor">
      <div class="tile is-parent">
        <article class="tile is-child box">
          <p class="title">Twelve</p>
          <p class="subtitle">Subtitle</p>
          <div class="content">
            <p>Lorem ipsum dolor sit amet, consectetur adipiscing elit. Proin ornare magna eros, eu pellentesque tortor vestibulum ut.</p>
          </div>
        </article>
      </div>
      <div class="tile is-parent is-6">
        <article class="tile is-child box">
          <p class="title">Thirteen</p>
          <p class="subtitle">Subtitle</p>
          <div class="content">
            <p>Lorem ipsum dolor sit amet, consectetur adipiscing elit. Proin ornare magna eros, eu pellentesque tortor vestibulum ut. Maecenas non massa sem. Etiam finibus odio quis feugiat facilisis.</p>
          </div>
        </article>
      </div>
      <div class="tile is-parent">
        <article class="tile is-child box">
          <p class="title">Fourteen</p>
          <p class="subtitle">Subtitle</p>
          <div class="content">
            <p>Lorem ipsum dolor sit amet, consectetur adipiscing elit. Proin ornare magna eros, eu pellentesque tortor vestibulum ut.</p>
          </div>
        </article>
      </div>
    </div>-->
  </div>
</template>

<script>
export default {
  name: "Home",
  data: function () {
    return {
      title: "Home",
    };
  },
};
</script>

<style scoped>
@font-face {
  font-family: fzqiudaoyu;
  src: url("/fzqiudaoyu.ttf");
}
#navigator {
  margin-top: 2%;
}
article {
  cursor: pointer;
}
.subtitle {
  font-family: fzqiudaoyu;
}
.content {
  font-family: fzqiudaoyu;
}
</style>
