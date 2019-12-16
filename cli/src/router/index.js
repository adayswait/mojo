import Vue from 'vue';
import Router from 'vue-router';
import Home from "@/components/Home.vue"
import DevOps from "@/components/DevOps.vue";
import QueryOps from "@/components/DevOps/QueryOps.vue";
import QueryUsr from "@/components/DevOps/QueryUsr.vue";
import NewDep from "@/components/DevOps/NewDep.vue";
import NewIni from "@/components/DevOps/NewIni.vue";
import NewUsr from "@/components/DevOps/NewUsr.vue";
import OtherOps from "@/components/DevOps/OtherOps.vue";
import Data from "@/components/Data.vue";
import e404 from '@/components/E404.vue';

Vue.use(Router)

const router = new Router({
    routes: [
        {
            path: "/",
            component: DevOps
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
                    path: "queryops",
                    component: QueryOps
                },
                {
                    path:"queryusr",
                    component:QueryUsr
                },
                {
                    path:"newdep",
                    component:NewDep
                },
                {
                    path:"newini",
                    component:NewIni
                },
                {
                    path:"newusr",
                    component:NewUsr
                },
                {
                    path:"otherops",
                    component:OtherOps
                }
                
            ]
        },
        {
            path: "/data",
            component: Data
        },
        {
            path: '*',
            component: e404
        }
    ]
})
export default router