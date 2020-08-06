import Vue from 'vue';
import Router from 'vue-router';
import Home from "@/components/Home.vue"
import DevOps from "@/components/DevOps.vue";
import OpsHome from "@/components/DevOps/Home.vue"
import AllDep from "@/components/DevOps/AllDep.vue";
import MyDep from "@/components/DevOps/MyDep.vue";
import ManageUsr from "@/components/DevOps/ManageUsr.vue";
import NewDep from "@/components/DevOps/NewDep.vue";
import MacIni from "@/components/DevOps/MacIni.vue";
import DepIni from "@/components/DevOps/DepIni.vue";
import DevIni from "@/components/DevOps/DevIni.vue";
import NewUsr from "@/components/DevOps/NewUsr.vue";
import DBview from "@/components/DevOps/DBview.vue";
import DevTools from '@/components/DevOps/DevTools.vue';
import Data from "@/components/Data.vue";
import e404 from '@/components/E404.vue';
import UsrInfo from '@/components/UsrInfo.vue';

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
                }
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