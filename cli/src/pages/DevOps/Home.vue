<template>
  <div class="box">
    <div id="dep" style="width:60%;height:500%;margin-left:18%"></div>
  </div>
</template>

<script>
export default {
  name: "DevOpsHome",
  methods: {
    drawDeployTotal(data) {
      let option = {
        title: {
          text: "部署总览",
          subtext: "近期全部部署的情况概览",
          left: "center",
        },
        tooltip: {
          trigger: "item",
          formatter: "{a} <br/>{b} : {c} ({d}%)",
        },
        legend: {
          orient: "vertical",
          left: "left",
          data: ["成功", "失败"],
        },
        series: [
          {
            name: "次数",
            type: "pie",
            radius: "55%",
            center: ["50%", "60%"],
            data: data,
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
    getAndDrawTaskList: async function () {
      try {
        const ret = await this.$httpc.get("/web/dep/progress");
        let tempList = [];
        let success = 0,
          failed = 0,
          processing = 0;
        for (let i = 0; i < ret.data.length; i += 5) {
          tempList[i / 5] = [
            ret.data[i], //start time
            ret.data[i + 1], //depid
            ret.data[i + 2], //depuuid
            ret.data[i + 3], //depstatus
            ret.data[i + 4], //awake countdown
          ];
          if (ret.data[i + 3] == 6) {
            success += 1;
          } else if (ret.data[i + 3] < 0) {
            failed += 1;
          } else {
            processing += 1;
          }
        }

        let data = [];
        if (success > 0) {
          data.push({ value: success, name: "成功" });
        }
        if (failed > 0) {
          data.push({ value: failed, name: "失败" });
        }
        if (processing > 0) {
          data.push({ value: processing, name: "进行中" });
        }
        this.drawDeployTotal(data);
      } catch (e) {
        this.$store.commit(
          "error",
          `获取任务进度列表失败 : ${e.data || e.message}`
        );
      }
    },
  },
  mounted: function () {
    this.getAndDrawTaskList();
  },
};
</script>