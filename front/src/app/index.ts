import '@/assets/styles/style.scss'
import 'element-plus/dist/index.css'

import ElementPlus from 'element-plus'
import { createApp } from 'vue'

import App from '@/app/App.vue'
import router from '@/app/providers/router'

export default class Application {
  public readonly app = createApp(App)

  constructor() {
    this.app.use(ElementPlus)
    this.app.use(router)
  }

  mount(selector) {
    this.app.mount(selector)
  }
}
