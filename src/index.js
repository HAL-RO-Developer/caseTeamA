import Vue from 'vue'
import VueRouter from 'vue-router'
import Index from './components/pages/index.vue'
import Login from './components/pages/login.vue'
import Device from './components/pages/device.vue'
import Records from './components/pages/records.vue'
import SolvedList from './components/pages/solvedList.vue'
import Settings from './components/pages/settings.vue'
import Children from './components/pages/children.vue'
import Messages from './components/pages/messages.vue'
import Bocco from './components/pages/bocco.vue'
// import Manual from './components/pages/Manual.vue'
import NotFound from './components/pages/notFound.vue'
import Admin from './components/pages/admin/admin.vue'
import CreateQues from './components/pages/admin/question.vue'
import Buefy from 'buefy'
import 'buefy/lib/buefy.css'

Vue.use(Buefy)
Vue.use(VueRouter)

const routes = [
    { path: "/", component: Index },
    { path: "/pages/login" , component: Login },
    { path: "/pages/device", component: Device },
    { path: "/pages/records", component: Records },
    { path: "/pages/records/date/:date", component: SolvedList },
    { path: "/pages/settings", component: Settings }, 
    { path: "/pages/children", component: Children },
    { path: "/pages/messages", component: Messages },
    { path: "/pages/bocco", component: Bocco},
    // { path: "/pages/manual", component: Manual },
    { path: "/pages/admin", component: Admin},
    { path: "/pages/admin/question", component: CreateQues },
    { path: "*", component: NotFound },
]
const router = new VueRouter({  mode: 'history', routes })

new Vue({
    router
}).$mount("#app")
