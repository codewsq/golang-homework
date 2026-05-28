import { createApp } from 'vue'
import { createPinia } from 'pinia'
import ArcoVue from '@arco-design/web-vue'
import ArcoVueIcon from '@arco-design/web-vue/es/icon'
import '@arco-design/web-vue/dist/arco.css'

import App from './App.vue'
import router from './router'
import './style.css'
const app = createApp(App)

app.use(createPinia())
app.use(router)
app.use(ArcoVue, {
  // 全局 Arco 组件默认尺寸，可按需改为 'mini' | 'small' | 'medium' | 'large'
  size: 'medium',
})
app.use(ArcoVueIcon)

app.mount('#app')
