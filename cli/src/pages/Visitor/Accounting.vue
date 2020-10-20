<template>
  <div>
    <Login />
    <div class="columns" v-if="!$store.state.visible.Login">
      <div class="column is-4 is-offset-1">
        <div class="tabs is-centered">
          <ul>
            <li
              :class="{ 'is-active': currRole == 'ERBAO' }"
              @click="changeCurrRole('ERBAO')"
            >
              <a>二宝</a>
            </li>
            <li
              :class="{ 'is-active': currRole == 'CHENGCHENG' }"
              @click="changeCurrRole('CHENGCHENG')"
            >
              <a>成成</a>
            </li>
            <li
              :class="{ 'is-active': currRole == 'PARENT' }"
              @click="changeCurrRole('PARENT')"
            >
              <a>爸妈</a>
            </li>
          </ul>
        </div>
        <tr v-for="(k, i) in CURRITEMS" :key="k" style="width: 100%">
          <div class="column" style="width: 50%; float: left">
            <div class="control has-icons-left">
              <input
                class="input"
                :placeholder="`${ITEMS[k].DES}(元)`"
                v-model="CURRDATA[i][0]"
              />
              <span class="icon is-medium is-left">
                <img :src="`/accounting/${k}.png`" />
              </span>
            </div>
          </div>
          <div class="column" style="float: right; width: 50%">
            <input
              class="input is-info"
              type="text"
              placeholder="备注"
              v-model="CURRDATA[i][1]"
            />
          </div>
        </tr>
        <tr>
          <br />
        </tr>

        <tr>
          <date-picker
            v-model="billTime"
            style="float: right; margin-right: 2%"
          ></date-picker>
        </tr>
        <tr>
          <br />
        </tr>

        <tr>
          <button
            class="button is-primary"
            style="float: right; margin-right: 2%"
            @click="bookkeep"
          >
            记账>>
          </button>
        </tr>
      </div>
      <div class="column is-6">
        <div class="tabs is-centered">
          <ul>
            <li
              :class="{ 'is-active': currChart == 'detail' }"
              @click="changeCurrChart('detail')"
            >
              <a>明细</a>
            </li>
            <li
              :class="{ 'is-active': currChart == 'item' }"
              @click="changeCurrChart('item')"
            >
              <a>按支出类型</a>
            </li>
            <li
              :class="{ 'is-active': currChart == 'role' }"
              @click="changeCurrChart('role')"
            >
              <a>按成员类型</a>
            </li>
          </ul>
        </div>
        <div class="column is-offset-2">
          选择时间 :
          <date-picker v-model="timeRange" range :lang="lang"></date-picker>
        </div>
        <br />
        <div
          id="vis"
          style="width: 80%; height: 80%; margin-left: 18%"
          v-if="currChart != 'detail'"
        ></div>
        <div v-if="currChart == 'detail'">
          <table class="table is-striped is-fullwidth">
            <tbody>
              <tr v-for="kv in allBills" :key="kv[0]">
                <td style="width: 15%">
                  {{
                    `${new Date(kv[1].time).getFullYear()}-${
                      new Date(kv[1].time).getMonth() + 1
                    }-${new Date(kv[1].time).getDate()}`
                  }}
                </td>
                <td style="width: 10%">{{ ROLEID[kv[1].role - 1] }}</td>
                <td style="width: 10%">
                  {{ ITEMS[ITEMID[kv[1].item - 1]].DES }}
                </td>
                <td style="width: 10%">{{ kv[1].price }}</td>
                <td style="width: 15%">{{ kv[1].remark }}</td>
                <td style="width: 20%">
                  <!-- <button class="button is-warning is-small is-vcentered">
                    修改
                  </button> -->
                  <button class="button is-danger is-small is-vcentered">
                    删除
                  </button>
                </td>
              </tr>
            </tbody>
          </table>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
const Login = () => import("@/pages/Login.vue");
import DatePicker from "vue2-datepicker";
import "vue2-datepicker/index.css";
export default {
  data: function () {
    return {
      priceReg: /^(([1-9]\d*)(\.\d{1,2})?)$|(0\.0?([1-9]\d?))$/,
      accountingBucketName: "sys:usr:acount",
      ITEMS: {
        HEALTH: { CODE: 1, DES: "健康" },
        FOOD: { CODE: 2, DES: "餐饮" },
        CLOTHE: { CODE: 3, DES: "衣服" },
        SHOPPING: { CODE: 4, DES: "购物" },
        LOVE: { CODE: 5, DES: "孝心" },
        LIVING: { CODE: 6, DES: "生活" },
        TRAVEL: { CODE: 7, DES: "出行" },
      },
      ITEMID: [
        "HEALTH",
        "FOOD",
        "CLOTHE",
        "SHOPPING",
        "LOVE",
        "LIVING",
        "TRAVEL",
      ],
      ROLEID: ["二宝", "成成", "爸妈"],
      ROLE: {
        ERBAO: 1,
        CHENGCHENG: 2,
        PARENT: 3,
      },
      ERBAO: [
        "HEALTH",
        "FOOD",
        "CLOTHE",
        "SHOPPING",
        "LOVE",
        "LIVING",
        "TRAVEL",
      ],
      CHENGCHENG: ["HEALTH", "FOOD", "CLOTHE", "LIVING"],
      PARENT: ["HEALTH", "FOOD", "CLOTHE", "SHOPPING", "LIVING"],
      ERBAO_DATA: [],
      CHENGCHENG_DATA: [],
      PARENT_DATA: [],
      CURRITEMS: [],
      CURRDATA: [],
      currRole: "",
      currChart: "",
      lang: {
        formatLocale: {
          firstDayOfWeek: 1,
        },
        monthBeforeYear: false,
      },
      billTime: new Date(),
      timeRange: [],
      allBills: [],
    };
  },
  components: {
    DatePicker,
    Login,
  },
  methods: {
    drawByStyle() {
      let option = {
        title: {
          text: "类型总览",
          subtext: "按类型近期支出状况一览",
          left: "center",
        },
        tooltip: {
          trigger: "item",
          formatter: "{a} <br/>{b} : {c}元 ({d}%)",
        },
        legend: {
          orient: "vertical",
          left: "left",
          data: [],
        },
        series: [
          {
            name: "金额",
            type: "pie",
            radius: "55%",
            center: ["50%", "60%"],
            data: undefined,
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
      for (let key in this.ITEMS) {
        option.legend.data.push(this.ITEMS[key].DES);
      }
      let data = [];
      let dataMap = {};
      for (let elem of this.allBills) {
        if (!dataMap[this.ITEMS[this.ITEMID[elem[1].item - 1]].DES]) {
          dataMap[this.ITEMS[this.ITEMID[elem[1].item - 1]].DES] = Number(
            elem[1].price
          );
        } else {
          dataMap[this.ITEMS[this.ITEMID[elem[1].item - 1]].DES] += Number(
            elem[1].price
          );
        }
      }
      for (let key in dataMap) {
        data.push({ value: dataMap[key], name: key });
      }
      option.series[0].data = data;
      let depChart = this.$echarts.init(document.getElementById("vis"));
      depChart.setOption(option);
    },
    drawByMember() {
      let option = {
        title: {
          text: "成员总览",
          subtext: "按成员近期支出状况一览",
          left: "center",
        },
        tooltip: {
          trigger: "item",
          formatter: "{a} <br/>{b} : {c}元 ({d}%)",
        },
        legend: {
          orient: "vertical",
          left: "left",
          data: [],
        },
        series: [
          {
            name: "金额",
            type: "pie",
            radius: "55%",
            center: ["50%", "60%"],
            data: undefined,
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
      for (let role of this.ROLEID) {
        option.legend.data.push(role);
      }
      let data = [];
      let dataMap = {};
      for (let elem of this.allBills) {
        if (!dataMap[this.ROLEID[elem[1].role - 1]]) {
          dataMap[this.ROLEID[elem[1].role - 1]] = Number(elem[1].price);
        } else {
          dataMap[this.ROLEID[elem[1].role - 1]] += Number(elem[1].price);
        }
      }
      for (let key in dataMap) {
        data.push({ value: dataMap[key], name: key });
      }
      window.console.log(data);
      option.series[0].data = data;
      let depChart = this.$echarts.init(document.getElementById("vis"));
      depChart.setOption(option);
    },
    changeCurrRole(key) {
      this.CURRITEMS = this[key];
      this.CURRDATA = this[`${key}_DATA`];
      this.currRole = key;
    },
    changeCurrChart(key) {
      this.currChart = key;
      if (key == "item") {
        setTimeout(() => {
          this.drawByStyle();
        }, 10);
      } else if (key == "role") {
        setTimeout(() => {
          this.drawByMember();
        }, 10);
      }
    },
    async bookkeep() {
      for (let i = 0; i < this.CURRDATA.length; i++) {
        if (!this.CURRDATA[i][0]) {
          continue;
        }
        if (this.priceReg.test(this.CURRDATA[i][0]) === false) {
          continue;
        }
        let data = {
          time: this.billTime.getTime(),
          role: this.ROLE[this.currRole],
          item: this.ITEMS[this[this.currRole][i]].CODE,
          price: this.CURRDATA[i][0],
          remark: this.CURRDATA[i][1],
        };
        await this.$mojoapi.post(`/web/db/${this.accountingBucketName}`, {
          value: JSON.stringify(data),
        });
      }
    },
    async getAllBills() {
      const ret = await this.$mojoapi.get(
        `/web/db/${this.accountingBucketName}`
      );
      let data = [];
      for (let i = 0; i < ret.data.length; i += 2) {
        data.push([ret.data[i], JSON.parse(ret.data[i + 1])]);
      }
      this.allBills = data;
    },
  },
  mounted() {
    let now = new Date();
    let aWeekAgo = new Date(now.getTime() - 7 * 24 * 3600000);
    aWeekAgo.setHours(0);
    aWeekAgo.setMinutes(0);
    aWeekAgo.setSeconds(0);
    this.timeRange = [aWeekAgo, now];
    for (let k in this.ITEMS) {
      void k;
      this.ERBAO_DATA.push([]);
      this.CHENGCHENG_DATA.push([]);
      this.PARENT_DATA.push([]);
    }
    this.changeCurrRole("ERBAO");
    this.changeCurrChart("detail");
    this.getAllBills();
  },
};
</script>

<style scoped>
img {
  height: 70%;
}
</style>
