const isDev = process.env.NODE_ENV === 'development'

// https://nuxt.com/docs/api/configuration/nuxt-config
export default defineNuxtConfig({
  modules: [
    // https://vueuse.org/guide/
    '@vueuse/nuxt',
  ],
  vite: {
    server: {
      hmr: {
        clientPort: 80,
        path: 'hmr/',
      },
    },
  },
})
