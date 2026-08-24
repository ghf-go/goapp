import Vue from 'vue'
import VueRouter from 'vue-router'
import Layout from '../views/Layout.vue'
import Login from '../views/Login.vue'
import Project from '../views/Project.vue'
import BuildRecord from '../views/BuildRecord.vue'

Vue.use(VueRouter)

const routes = [
  { path: '/login', name: 'Login', component: Login },
  {
    path: '/',
    component: Layout,
    redirect: '/project',
    children: [
      { path: 'project', name: 'Project', component: Project },
      { path: 'build-record', name: 'BuildRecord', component: BuildRecord }
    ]
  }
]

const router = new VueRouter({
  mode: 'hash',
  routes
})

router.beforeEach((to, from, next) => {
  const token = localStorage.getItem('token')
  if (to.path !== '/login' && !token) {
    next('/login')
  } else if (to.path === '/login' && token) {
    next('/')
  } else {
    next()
  }
})

export default router
