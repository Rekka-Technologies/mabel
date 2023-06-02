/** @type {import('tailwindcss').Config} */
export default {
  content: [
    "./index.html",
    "./src/**/*.{html,ts,vue}",
    "./node_modules/vue-tailwind-datepicker/**/*.js"
  ],
  theme: {
  },
  plugins: [
    require('@tailwindcss/forms'),
  ],
}

