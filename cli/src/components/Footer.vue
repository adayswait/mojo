<template>
  <footer>copyrights reserved</footer>
</template>

<script>
export default {
  name: "Footer",
  data: function() {
    return {
      title: "Header"
    };
  }
};
</script>

<style scoped>
footer {
  height: 40px; /*设定footer高度*/
  position: absolute; /*设定footer绝对位置在底部*/
  bottom: 0;
  width: 100%; /*展开footer宽度*/
  background-color:#000000;
  line-height: 40px;
  display: flex; 
  justify-content: center;
}
</style>
