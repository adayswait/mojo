import Vue from 'vue'
import Vuex from 'vuex'

Vue.use(Vuex)
//状态对象
const state = {
    visible: {
        Login: false,
        Home: true,
        Footer: false
    },
    footer: "Copyrights Reserved. All materials not authorized may not be redirected or for other usages."
};
const mutations = {//包含多个更新status函数的对象
    setVisible(state, payload) {
        window.console.log(payload);
        state.visible[payload.name] = payload.visible;
    }
};
const actions = {//包含多个对应事件回调函数的对象
    increment({ commit }, { n, m }) {//参数对象化传递
        commit('inCrement', { n, m })
    },
    decrement({ commit }) {
        commit('deCrement')
    },
    //带条件的action
    incrementIfOdd({ commit, state }) {
        if (state.count % 2 === 1) {
            commit('inCrement', { n: 1, m: 1 })
        }
    },

    //异步的action
    incrementAsync({ commit }) {
        setTimeout(() => {
            commit('inCrement', { m: 1, n: 1 })
        }, 2000)
    }
};
const getters = {//包含多个getter计算属性函数的对象
    evenOrOdd(state) {//不需要调用
        return state.count % 2 === 0 ? '偶数' : '奇数'
    }
};

export default new Vuex.Store({
    state,//状态对象,
    mutations,//包含多个更新status函数的对象
    actions,//包含多个对应事件回调函数的对象
    getters,//包含多个getter计算属性函数的对象
});