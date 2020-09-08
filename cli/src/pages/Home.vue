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
          <div class="tile is-parent" id="card">
            <router-link to="/devops" class="tile is-child box" id="card_devops">
              <p class="title">运维</p>
              <div class="content">
                <p>
                  <strong>操作和维护线上服务</strong>
                </p>
              </div>
            </router-link>
          </div>
          <div class="tile is-8 is-vertical">
            <div class="tile">
              <div class="tile is-parent" id="card">
                <router-link to="/e404" class="tile is-child box" id="card_raft">
                  <p class="title">Raft集群</p>
                  <div class="content">
                    <strong></strong>
                  </div>
                </router-link>
              </div>
              <div class="tile is-parent" id="card">
                <router-link to="/iframe/taomeedoc" class="tile is-child box" id="card_taomeedoc">
                  <p class="title">淘米账号平台文档</p>
                  <div class="content">
                    <strong>登录验证等账号api文档</strong>
                  </div>
                </router-link>
              </div>
            </div>
            <div class="tile is-parent" id="card">
              <router-link to="/elkstack/kibana" class="tile is-child box" id="card_elk">
                <p class="title">ELK Stack</p>
                <div class="content">
                  <strong>查看线上日志</strong>
                </div>
              </router-link>
            </div>
          </div>
        </div>
        <div class="tile">
          <div class="tile is-8 is-parent" id="card">
            <router-link to="/iframe/gm239" class="tile is-child box" id="card_gm239">
              <p class="title">内网GM</p>
              <p class="subtitle">10.1.1.239</p>
              <div class="content">
                <p>内网GM工具</p>
              </div>
            </router-link>
          </div>
          <div class="tile is-parent" id="card">
            <router-link to="/iframe/plans" class="tile is-child box is-info" id="card_plans">
              <p class="title">QA多开</p>
              <div class="content">
                <p>QA前端多开页面</p>
              </div>
            </router-link>
          </div>
        </div>
      </div>
      <div class="tile is-parent" id="card">
        <router-link to="/iframe/devdoc" class="tile is-child box" id="card_devdoc">
          <div class="content">
            <p class="title">开发文档</p>
            <div class="content">
              <p>各种各样的文档</p>
            </div>
          </div>
        </router-link>
      </div>
    </div>

    <div class="tile is-ancestor">
      <div class="tile is-parent" id="card">
        <article class="tile is-child box">
          <p class="title">Twelve</p>
          <p class="subtitle">Subtitle</p>
          <div class="content">
            <p>Lorem ipsum dolor sit amet, consectetur adipiscing elit. Proin ornare magna eros, eu pellentesque tortor vestibulum ut.</p>
          </div>
        </article>
      </div>
      <div class="tile is-parent is-6" id="card">
        <article class="tile is-child box">
          <p class="title">Thirteen</p>
          <p class="subtitle">Subtitle</p>
          <div class="content">
            <p>Lorem ipsum dolor sit amet, consectetur adipiscing elit. Proin ornare magna eros, eu pellentesque tortor vestibulum ut. Maecenas non massa sem. Etiam finibus odio quis feugiat facilisis.</p>
          </div>
        </article>
      </div>
      <div class="tile is-parent" id="card">
        <article class="tile is-child box">
          <p class="title">Fourteen</p>
          <p class="subtitle">Subtitle</p>
          <div class="content">
            <p>Lorem ipsum dolor sit amet, consectetur adipiscing elit. Proin ornare magna eros, eu pellentesque tortor vestibulum ut.</p>
          </div>
        </article>
      </div>
    </div>
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
  font-family: FZMingSTJW;
  src: url("/FZMingSTJW.ttf");
}
#navigator {
  margin-top: 2%;
}
article {
  cursor: pointer;
}
router-link {
  text-decoration: none;
}

#card:hover {
  filter: brightness(0.9);
}
/* #card_devops {
  background-color: #eed6bb;
}
#card_raft {
  background-color: rgb(160, 191, 124);
}
#card_elk {
  background-color: #c298b8;
}
#card_gm239 {
  background-color: #6fc7e1;
}
#card_devdoc {
  background-color: rgb(204, 202, 169);
}
#card_taomeedoc {
  background-color: rgb(186, 227, 217);
}
#card_plans {
  background-color: rgb(219, 207, 202);
} */
.title {
  color: #b8860b;
  font-family: FZMingSTJW;
}
.subtitle {
  font-family: FZMingSTJW;
}
.content {
  font-family: FZMingSTJW;
}
</style>
