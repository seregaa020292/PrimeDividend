const isDev = process.env.NODE_ENV === 'development'

// https://nuxt.com/docs/api/configuration/nuxt-config
export default defineNuxtConfig({
  srcDir: 'src/',

  dir: {
    public: '../public/',
  },

  modules: [
    // https://vueuse.org/guide/
    '@vueuse/nuxt',
  ],

  css: [
    '@/assets/styles/main.scss',
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
