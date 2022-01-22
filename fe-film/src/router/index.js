import {createRouter, createWebHistory} from 'vue-router'
import Home from '../view/Home.vue'
import NotFound from '../components/NotFound.vue'

const routes = [
    {
        path: "/",
        component:Home,
        name: 'Home'
    },
    {
        path: "/:catchAll(.*)",
        component: NotFound,
    }
]

const router = createRouter({
    history: createWebHistory(),
    routes
})

export default router