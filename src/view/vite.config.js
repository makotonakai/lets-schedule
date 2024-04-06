import { fileURLToPath, URL } from 'url'
import { defineConfig, loadEnv} from 'vite'
import vue from '@vitejs/plugin-vue'

// https://vitejs.dev/config/
export default defineConfig(({mode}) => {
  const env = loadEnv(mode, process.cwd(), '');
  
  return {
  define: {
    'process.env.HOST': JSON.stringify(env.HOST),
    'process.env.PORT': JSON.stringify(env.PORT),
  },
  plugins: [vue()],
  resolve: {
    alias: {
      '@': fileURLToPath(new URL('./src', import.meta.url))
    }
  },
  server: {
    proxy: {
      '/api': {
        target: `${process.env.HOST}:${process.env.PORT}`,
        changeOrigin: true
      }
    }
  }
}
})
