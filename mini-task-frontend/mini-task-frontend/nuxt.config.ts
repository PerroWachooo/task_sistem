// https://nuxt.com/docs/api/configuration/nuxt-config

import tailwindcss from "@tailwindcss/vite";

export default defineNuxtConfig({
  axios: {
    baseURL: 'http://localhost:8080',  // URL del backend
  },

  vite: {
    plugins: [
      tailwindcss(),
    ],
  },
  css: ['~/assets/css/main.css'],

  typescript: {
    strict: true
  },
  
  compatibilityDate: '2024-11-01',
  devtools: { enabled: true }
})
