<template>
  <div class="box">
    <div id="dep" style="width:60%;height:500%;"></div>
  </div>
</template>

<script>
export default {
  name: "DevOpsHome",
  methods: {
    drawDeployTotal() {
      let option = {
        title: {
          text: "某站点用户访问来源",
          subtext: "纯属虚构",
          left: "center",
        },
        tooltip: {
          trigger: "item",
          formatter: "{a} <br/>{b} : {c} ({d}%)",
        },
        legend: {
          orient: "vertical",
          left: "left",
          data: ["直接访问", "邮件营销", "联盟广告", "视频广告", "搜索引擎"],
        },
        series: [
          {
            name: "访问来源",
            type: "pie",
            radius: "55%",
            center: ["50%", "60%"],
            data: [
              { value: 335, name: "直接访问" },
              { value: 310, name: "邮件营销" },
              { value: 234, name: "联盟广告" },
              { value: 135, name: "视频广告" },
              { value: 1548, name: "搜索引擎" },
            ],
            emphasis: {
              itemStyle: {
                shadowBlur: 10,
                shadowOffsetX: 0,
                shadowColor: "rgba(0, 0, 0, 0.5)",
              },
            },
          },
        ],
      };
      let depChart = this.$echarts.init(document.getElementById("dep"));
      depChart.setOption(option);
    },
    drawDeployDetails() {
      let optionDepChart = {
        title: {
          text: "柱状图",
        },
        yAxis: {
          type: "category",
          data: [
            "online1",
            "online2",
            "online3",
            "battle1",
            "battle2",
            "battle3",
            "center",
            "online1",
            "online2",
            "online3",
            "battle1",
            "battle2",
            "battle3",
            "center",
          ],
        },
        // backgroundColor: "#fff000",
        color: ["#000000", "red", "green"],
        xAxis: {
          type: "value",
        },
        series: [
          {
            data: [60, 7, 32, 2, 100, 23, 87, 98, 90, 45, 7, 34, 99, 0],
            type: "bar",
            // animationDelay: function (idx) {
            //   return idx * 10;
            // },
          },
        ],
      };
      let depChart = this.$echarts.init(document.getElementById("dep"));
      depChart.setOption(optionDepChart);
    },
    drawChart() {
      // 基于准备好的dom，初始化echarts实例
      let payload = this.$echarts.init(document.getElementById("payload"));
      // 指定图表的配置项和数据
      let optionPayload = {
        title: {
          text: "柱状图",
        },
        legend: {},
        tooltip: {},
        dataset: {
          source: [
            ["product", "内存", "在线人数", "数据包"],
            ["10.30.14.109", 23123, 32443, 33423],
            ["10.30.14.239", 34324, 23223, 7656],
            ["10.30.14.101", 8797, 6544, 9878],
            ["10.30.14.185", 63354, 5435, 33423],
          ],
        },
        xAxis: { type: "category" },
        yAxis: {},
        // Declare several bar series, each will be mapped
        // to a column of dataset.source by default.
        series: [{ type: "bar" }, { type: "bar" }, { type: "bar" }],
      };

      // 使用刚指定的配置项和数据显示图表。
      payload.setOption(optionPayload);

      let nOption = {
        title: {
          text: "压力图(qps)",
        },
        tooltip: {
          trigger: "axis",
        },
        legend: {
          data: ["online", "battle", "api", "mysql", "redis"],
        },
        grid: {
          left: "3%",
          right: "4%",
          bottom: "3%",
          containLabel: true,
        },
        toolbox: {
          feature: {
            saveAsImage: {},
          },
        },
        xAxis: {
          type: "category",
          boundaryGap: false,
          data: ["13:00", "14:00", "15:00", "16:00", "17:00", "18:00", "19:00"],
        },
        yAxis: {
          type: "value",
        },
        series: [
          {
            name: "online",
            type: "line",
            stack: "总量",
            data: [120, 132, 101, 134, 90, 230, 210],
          },
          {
            name: "battle",
            type: "line",
            stack: "总量",
            data: [220, 182, 191, 234, 290, 330, 310],
          },
          {
            name: "api",
            type: "line",
            stack: "总量",
            data: [150, 232, 201, 154, 190, 330, 410],
          },
          {
            name: "redis",
            type: "line",
            stack: "总量",
            data: [320, 332, 301, 334, 390, 330, 320],
          },
          {
            name: "mysql",
            type: "line",
            stack: "总量",
            data: [820, 932, 901, 934, 1290, 1330, 1320],
          },
        ],
      };
      let np = this.$echarts.init(document.getElementById("np"));
      np.setOption(nOption);
    },
  },
  mounted() {
    this.drawDeployTotal()
    // this.drawDeployDetails();
    // this.drawChart();
  },
};
</script>