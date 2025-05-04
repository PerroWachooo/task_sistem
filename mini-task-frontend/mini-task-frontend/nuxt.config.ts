// https://nuxt.com/docs/api/configuration/nuxt-config


export default defineNuxtConfig({
  axios: {
    baseURL: 'http://localhost:8080',  // URL del backend
  },

  typescript: {
    strict: true
  },
  
  compatibilityDate: '2024-11-01',
  devtools: { enabled: true }
})
