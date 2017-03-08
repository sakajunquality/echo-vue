import Vue from 'vue'
import Router from 'vue-router'
import Todo from '@/components/Todo'
import Hello from '@/components/Hello'
import NotFound from '@/components/NotFound'

Vue.use(Router)

export default new Router({
  routes: [
    {
      path: '/',
      name: 'Todo',
      component: Todo
    },
    {
      path: '/hello',
      name: 'Hello',
      component: Hello
    },
    {
      path: '*',
      name: 'NotFound',
      component: NotFound
    }
  ]
})
