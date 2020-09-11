<template>
  <div class="container">
    <!-- <audio src="/hua.mp3" loop autoplay="autoplay"></audio> -->
    <div class="tile is-ancestor">
      <div class="tile is-vertical is-9">
        <div class="tile">
          <div class="tile is-parent" id="card">
            <router-link to="/devops" class="tile is-child box" id="card_devops" hidefocus="true">
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
                <router-link to="/centersvr" class="tile is-child box" id="card_raft">
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
            <!-- :to="{path:'/iframe/basic',query:{srcurl:'http://10.1.1.239'}}" -->
            <router-link to="/iframe/gm239" class="tile is-child box" id="card_gm239">
              <p class="title">内网GM</p>
              <div class="content">
                <p>
                  <strong>内网QA服</strong>
                </p>
              </div>
            </router-link>
          </div>
          <div class="tile is-parent" id="card">
            <router-link to="/iframe/plans" class="tile is-child box is-info" id="card_plans">
              <p class="title">QA多开</p>
              <div class="content">
                <p>
                  <strong>前端多开页面</strong>
                </p>
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
              <p>
                <br />
                <br />
                <strong>语言, 标准库, 第三方库的文档</strong>
              </p>
              <p>
                <strong>本目录下的文档存储在内网服务器上</strong>
              </p>
              <p>
                <strong>访问速度比官网快很多</strong>
              </p>
            </div>
          </div>
        </router-link>
      </div>
    </div>

    <div class="tile is-ancestor">
      <div class="tile is-parent" id="card">
        <router-link to="/iframe/clienthub" class="tile is-child box">
          <p class="title">前端目录</p>
          <div class="content">
            <p>
              <strong>客户端汇总</strong>
            </p>
          </div>
        </router-link>
      </div>
      <div class="tile is-parent is-6" id="card">
        <router-link to="/note" class="tile is-child box">
          <p class="title">记事本or留言板?</p>
          <div class="content">
            <p>
              <strong>......</strong>
            </p>
          </div>
        </router-link>
      </div>

      <div class="tile is-parent" id="card">
        <router-link to="/iframe/gm" class="tile is-child box">
          <p class="title">外网GM</p>
          <div class="content">
            <p>
              <strong>外网正式服</strong>
            </p>
          </div>
        </router-link>
      </div>
    </div>
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
    </div> -->
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
/* 方正黑隶简体 */
@font-face {
  font-family: FZHeiLJW;
  src: url("/FZHeiLJW.TTF");
}
.container {
  margin-top: 2%;
}
article {
  cursor: pointer;
}
#card:hover {
  filter: brightness(1);
}

#card:hover::before {
  content: "";
  position: absolute;
  top: -1px;
  left: -1px;
  right: -1px;
  bottom: -1px;
  z-index: -1;
}
#card:hover::after {
  content: "";
  position: absolute;
  top: -1px;
  left: -1px;
  right: -1px;
  bottom: -1px;
  z-index: -2;
  filter: blur(5px);
}
/* #card:hover::before, */
#card:hover::after {
  background: linear-gradient(
    rgba(51, 230, 202, 0.1),
    rgba(51, 230, 202, 0.1),
    rgba(51, 230, 202, 0.1)
  );
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
  font-family: FZHeiLJW;
}
.subtitle {
  font-family: FZHeiLJW;
}
.content {
  font-family: FZHeiLJW;
}
</style>
