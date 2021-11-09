import Vue from 'vue'
import VueRouter from 'vue-router'
import ContactPage from './pages/Contacts.vue'
import KeysPage from './pages/Keys.vue'
import EncryptPage from './pages/Encrypt.vue'

Vue.use(VueRouter)

const routes = [
    { component: ContactPage, name: 'Contacts', path: '/contacts' },
    { component: KeysPage, name: 'Contacts', path: '/keys' },
    { component: ContactPage, name: 'Contacts', path: '/wizard/key' },
    { component: ContactPage, name: 'Contacts', path: '/wizard/contact' },
    { component: EncryptPage, name: 'Encrypt', path: '/wizard/encrypt' },
    { component: ContactPage, name: 'Contacts', path: '/wizard/decrypt' },
]

const router = new VueRouter({
  mode: 'abstract',
  routes,
})

export default router