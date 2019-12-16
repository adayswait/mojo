import Vue from 'vue'
import Router from 'vue-router'

import e404 from '@/components/E404.vue'
Vue.use(Router)

// 创建路由对象并配置路由规则（嵌套路由）
const router = new Router({
    routes: [
        {
            path: '*',
            component: e404
        }
    ]
})
export default router