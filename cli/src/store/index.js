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
        UsrInfo: false,
    },
    DevOpsMask: 0,
    GROUP: {
        0: "whosyourdaddy",
        1: "管理员",
        2: "普通用户",
        3: "游客",
        4: "未激活"
    },
    SERVER_TYPE: ["online", "battle", "match", "team", "center"],
    REPO_TYPE: ["svn"],
    MESSAGE_TYPE: {
        INFO: 1,
        WARN: 2,
        ERROR: 3
    },
    messageIdx: 0,
    messageDelN: 0,
    messageList: [],
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
    },
    updateDevOpsMask(state, n) {
        state.DevOpsMask += n;
        if (state.DevOpsMask < 0) {
            state.DevOpsMask = 0;
        }
    },
    error(state, message) {
        state.messageList.push([state.messageIdx++, state.MESSAGE_TYPE.ERROR, message]);
        state.messageDelN += 1;
        setTimeout(() => {
            state.messageList.splice(state.messageList.length - state.messageDelN, 1);
            state.messageDelN -= 1;
        }, 5000);
    },
    warn(state, message) {
        state.messageList.push([state.messageIdx++, state.MESSAGE_TYPE.WARN, message]);
        state.messageDelN += 1;
        setTimeout(() => {
            state.messageList.splice(state.messageList.length - state.messageDelN, 1);
            state.messageDelN -= 1;
        }, 5000);
    },
    info(state, message) {
        state.messageList.push([state.messageIdx++, state.MESSAGE_TYPE.INFO, message]);
        state.messageDelN += 1;
        setTimeout(() => {
            state.messageList.splice(state.messageList.length - state.messageDelN, 1);
            state.messageDelN -= 1;
        }, 5000);
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
    },
    MESSAGE_TYPE() {
        return state.MESSAGE_TYPE
    },
    userInfo() {
        return state.userInfo
    }
};

export default new Vuex.Store({
    state,//状态对象,
    mutations,//包含多个更新status函数的对象
    actions,//包含多个对应事件回调函数的对象
    getters,//包含多个getter计算属性函数的对象
});