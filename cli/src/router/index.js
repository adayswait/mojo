import Vue from 'vue';
import Router from 'vue-router';
const Home = () => import("@/pages/Home.vue")
const DevOps = () => import("@/pages/DevOps.vue")
const OpsHome = () => import("@/pages/DevOps/Home.vue")
const AllDep = () => import("@/pages/DevOps/AllDep.vue")
const MyDep = () => import("@/pages/DevOps/MyDep.vue")
const Issue = () => import("@/pages/DevOps/Issue.vue")
const ManageUsr = () => import("@/pages/DevOps/ManageUsr.vue")
const NewDep = () => import("@/pages/DevOps/NewDep.vue")
const QuickOps = () => import("@/pages/DevOps/QuickOps.vue")
const MacIni = () => import("@/pages/DevOps/MacIni.vue")
const DepIni = () => import("@/pages/DevOps/DepIni.vue")
const DevIni = () => import("@/pages/DevOps/DevIni.vue")
const NewUsr = () => import("@/pages/DevOps/NewUsr.vue")
const DBview = () => import("@/pages/DevOps/DBview.vue")
const DevTools = () => import('@/pages/DevOps/DevTools.vue')
const Note = () => import("@/pages/Note.vue")
const ElkStack = () => import("@/pages/ElkStack.vue")
const Kibana = () => import('@/pages/ElkStack/Kibana.vue')
const IFrame = () => import("@/pages/IFrame.vue")
const Raft = () => import('@/pages/IFrame/Raft.vue')
const GM239 = () => import('@/pages/IFrame/GM239.vue')
const GM = () => import('@/pages/IFrame/GM.vue')
const Plans = () => import('@/pages/IFrame/Plans.vue')
const ClientHub = () => import('@/pages/IFrame/ClientHub.vue')
const DevDoc = () => import('@/pages/IFrame/DevDoc.vue')
const TaomeeDoc = () => import('@/pages/IFrame/TaomeeDoc.vue')
const Visitor = () => import("@/pages/Visitor.vue")
const BreakDep = () => import("@/pages/Visitor/BreakDep.vue")
const Accounting = () => import("@/pages/Visitor/Accounting.vue")
const Data = () => import("@/pages/Data.vue")
const e404 = () => import('@/pages/E404.vue')
const UsrInfo = () => import('@/pages/UsrInfo.vue')

Vue.use(Router)

const router = new Router({
    routes: [
        {
            path: "/",
            redirect: "/home"
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
                    path: "quickops",
                    component: QuickOps
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
                {
                    path: "issue",
                    component: Issue
                }
            ]
        },
        {
            path: "/visitor",
            component: Visitor,
            children: [
                {
                    path: "breakdep",
                    component: BreakDep
                },
                {
                    path: "accounting",
                    component: Accounting
                },
            ]
        },
        {
            path: "/elkstack",
            component: ElkStack,
            children: [
                {
                    path: "kibana",
                    component: Kibana
                },
            ]
        },
        {
            path: "/note",
            component: Note
        },
        {
            path: "/iframe",
            component: IFrame,
            children: [
                // {
                //     path: "basic",
                //     component: BasicFrame
                // },
                {
                    path: "raft",
                    component: Raft
                },
                {
                    path: "plans",
                    component: Plans
                },
                {
                    path: "gm239",
                    component: GM239
                },
                {
                    path: "gm",
                    component: GM
                },
                {
                    path: "devdoc",
                    component: DevDoc
                },
                {
                    path: "taomeedoc",
                    component: TaomeeDoc
                },
                {
                    path: "clienthub",
                    component: ClientHub
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