<template>
  <header>
    <div id="headermenu">
      <ul>
        <li>
          <router-link to="/home">Home</router-link>
        </li>
        <li>
          <router-link to="/devops">DevOps</router-link>
        </li>
        <li>
           <router-link to="/data">Data</router-link>
        </li>
        <li></li>
      </ul>
    </div>
    <div v-if="this.$store.state.visible.UsrInfo">
      <li id="usrinfo">
        <router-link to="/usrinfo">modu</router-link>
      </li>
    </div>
    
  </header>
</template>

<script>
export default {
  name: "Header",
  data: function() {
    return {
      title: "Header"
    };
  }
};
</script>

<style scoped>
header {
  font-family: Helvetica;
  background-color: #333333;
  height: 40px;
  width: 100%;
}
#headermenu {
  margin: 0px 10%;
}
#headermenu li {
  float: left;
  background-color: #333333;
  border-bottom: 4px solid #333333;
  margin: 0 30px 0 0;
  line-height: 36px;
  list-style: none;
}
#usrinfo {
  float:right;
  border-bottom: 4px solid #333333;
  margin: 0 30px 0 0;
  line-height: 36px;
  list-style: none;
}
#usrinfo a {
  display: block;
  color: #ffffff;
  font-size: 1.1em;
  text-decoration: none;
}
#headermenu li:hover {
  border-bottom: 4px solid #6dc2e9;
}
#headermenu a {
  display: block;
  color: #ffffff;
  font-size: 1.1em;
  text-decoration: none;
}
#headermenu a:hover {
  color: #6dc2e9;
}
</style>
