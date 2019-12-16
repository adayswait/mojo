<template>
  <header>
    <div id="headermenu">
      <ul>
        <li>
          <a href="../index.html">Home</a>
        </li>
        <li>
          <a href="../Lab/index.html">DevOps</a>
        </li>
        <li>
          <a href="../Buzz/index.html">Data</a>
        </li>
        <li>
          <a href="../About/index.html">About</a>
        </li>
      </ul>
    </div>
  </header>
</template>

<script>
export default {
  name: "Login",
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
#headermenu
{
	margin:0px 10%;
}
#headermenu li
{
	float:left;
	background-color: #333333;
	border-bottom:4px solid #333333;
	margin:0 30px 0 0;
	line-height:36px;
	list-style:none;
}
#headermenu li:hover
{
	border-bottom:4px solid #6DC2E9;
}
#headermenu a
{
	display: block;
	color:#ffffff;
	font-size: 1.1em;
	text-decoration: none;
}
#headermenu a:hover
{
	color:#6DC2E9;
}
</style>
