import Application from '@/app'

const application = new Application()

application.mount('#app')

export const getApplication = (): Application => application
