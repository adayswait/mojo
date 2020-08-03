import Vue from 'vue'
import Vuex from 'vuex'

Vue.use(Vuex)
//状态对象
const state = {
    userInfo: {
        user: undefined,
        group: undefined
    },
    visible: {
        Login: true,
        Home: false,
        Footer: true,
        UsrInfo: false
    },
    GROUP: {
        0: "管理员",
        1: "普通用户",
        2: "游客",
        3: "未激活"
    },
    footer: "Copyrights Reserved. All materials not authorized may not be redirected or for other usages."
};
const mutations = {//包含多个更新status函数的对象
    setVisible(state, { name, visible }) {
        state.visible[name] = visible;
    },
    setUserInfo(state, { user, group }) {
        state.userInfo = {
            user: user,
            group: group
        }
    }
};
const actions = {//包含多个对应事件回调函数的对象
    setVisibleAsync({ commit }) {
        commit('setVisible')
    }
};
const getters = {
    GROUP() {
        return state.GROUP
    }
};

export default new Vuex.Store({
    state,//状态对象,
    mutations,//包含多个更新status函数的对象
    actions,//包含多个对应事件回调函数的对象
    getters,//包含多个getter计算属性函数的对象
});