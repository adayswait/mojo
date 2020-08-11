import Vue from 'vue';
import Router from 'vue-router';
import Home from "@/pages/Home.vue"
import DevOps from "@/pages/DevOps.vue";
import OpsHome from "@/pages/DevOps/Home.vue"
import AllDep from "@/pages/DevOps/AllDep.vue";
import MyDep from "@/pages/DevOps/MyDep.vue";
import ManageUsr from "@/pages/DevOps/ManageUsr.vue";
import NewDep from "@/pages/DevOps/NewDep.vue";
import MacIni from "@/pages/DevOps/MacIni.vue";
import DepIni from "@/pages/DevOps/DepIni.vue";
import DevIni from "@/pages/DevOps/DevIni.vue";
import NewUsr from "@/pages/DevOps/NewUsr.vue";
import DBview from "@/pages/DevOps/DBview.vue";
import DevTools from '@/pages/DevOps/DevTools.vue';
import NoLogin from "@/pages/NoLogin.vue";
import BreakDep from "@/pages/NoLogin/BreakDep.vue";
import Data from "@/pages/Data.vue";
import e404 from '@/pages/E404.vue';
import UsrInfo from '@/pages/UsrInfo.vue';

Vue.use(Router)

const router = new Router({
    routes: [
        {
            path: "/",
            redirect: "/devops"
        },
        {
            path: "/home",
            component: Home
        },
        {
            path: "/devops",
            component: DevOps,
            children: [
                {
                    path: "",
                    redirect: "home"
                },
                {
                    path: "home",
                    component: OpsHome
                },
                {
                    path: "mydep",
                    component: MyDep
                },
                {
                    path: "alldep",
                    component: AllDep
                },
                {
                    path: "manageuser",
                    component: ManageUsr
                },
                {
                    path: "newdep",
                    component: NewDep
                },
                {
                    path: "macini",
                    component: MacIni
                },
                {
                    path: "depini",
                    component: DepIni
                },
                {
                    path: "devini",
                    component: DevIni
                },
                {
                    path: "newusr",
                    component: NewUsr
                },
                {
                    path: "dbview",
                    component: DBview
                },
                {
                    path: "devtools",
                    component: DevTools
                },
            ]
        },
        {
            path: "/nologin",
            component: NoLogin,
            children: [
                {
                    path: "breakdep",
                    component: BreakDep
                },
            ]
        },
        {
            path: "/data",
            component: Data
        },
        {
            path: "/usrinfo",
            component: UsrInfo
        },
        {
            path: '*',
            component: e404
        }
    ]
})
export default router