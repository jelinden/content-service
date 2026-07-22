import { defineConfig } from 'vite';
import react from '@vitejs/plugin-react';

export default defineConfig({
  build: {
    rollupOptions: {
      input: './public/index.html'
    },
    outDir: 'build'
  },
  plugins: [react()]
});
