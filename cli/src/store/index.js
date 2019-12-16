import Vue from 'vue'
import Vuex from 'vuex'

Vue.use(Vuex)
//状态对象
const state = {
    visible: {
        Login: true,
        Home: false,
        Footer: true
    },
    mock: {
        query:[
            [1, 'crespo', 'S计划', 'online', '2019-10-01 12:02:23', '/opt/splan/game'],
            [2, 'messi', 'S计划', 'online', '2019-10-01 12:02:23', '/opt/splan/game'],
            [3, 'rui costa', 'S计划', 'online', '2019-10-01 12:02:23', '/opt/splan/game'],
            [4, 'cocacola', 'S计划', 'online', '2019-10-01 12:02:23', '/opt/splan/game'],
            [5, 'inzaghi', 'S计划', 'online', '2019-10-01 12:02:23', '/opt/splan/game'],
            [6, 'crespo', 'S计划', 'online', '2019-10-01 12:02:23', '/opt/splan/game'],
            [7, 'messi', 'S计划', 'online', '2019-10-01 12:02:23', '/opt/splan/game'],
            [8, 'rui costa', 'S计划', 'online', '2019-10-01 12:02:23', '/opt/splan/game'],
            [9, 'cocacola', 'S计划', 'online', '2019-10-01 12:02:23', '/opt/splan/game'],
            [10, 'inzaghi', 'S计划', 'online', '2019-10-01 12:02:23', '/opt/splan/game'],
        ] 
    },
    footer: "Copyrights Reserved. All materials not authorized may not be redirected or for other usages."
};
const mutations = {//包含多个更新status函数的对象
    setVisible(state, { name, visible }) {
        state.visible[name] = visible;
    }
};
const actions = {//包含多个对应事件回调函数的对象
    setVisibleAsync({ commit }) {
        commit('setVisible')
    }
};
const getters = {//包含多个getter计算属性函数的对象
    evenOrOdd(state) {//不需要调用
        return state.count & 0b1 ? '奇数' : '偶数'
    }
};

export default new Vuex.Store({
    state,//状态对象,
    mutations,//包含多个更新status函数的对象
    actions,//包含多个对应事件回调函数的对象
    getters,//包含多个getter计算属性函数的对象
});