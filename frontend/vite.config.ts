import { fileURLToPath, URL } from 'node:url'

import { defineConfig, loadEnv } from 'vite'
import vue from '@vitejs/plugin-vue'
import vueDevTools from 'vite-plugin-vue-devtools'
import path from 'node:path'

export default defineConfig(({ mode }) => {
  const env = loadEnv(mode, path.join(process.cwd(), '../'), '')
  return {
    plugins: [vue(), vueDevTools()],
    resolve: {
      alias: {
        '@': fileURLToPath(new URL('./src', import.meta.url)),
      },
    },
    define: {
      'import.meta.env.RECAPTCHA_KEY': JSON.stringify(env.RECAPTCHA_KEY),
      'import.meta.env.API_PATH': JSON.stringify(env.API_PATH),
    },
  }
})
