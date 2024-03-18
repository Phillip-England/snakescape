/** @type {import('tailwindcss').Config} */
module.exports = {
    content: [
        "./internal/**/*.{go,js,templ,html}"
    ],
    theme: {
      extend: {
        screens: {
          'sm': '640px', // Phones
          'md': '768px', // Tablets
          'lg': '1024px', // Laptops/Desktops
          'xl': '1280px', // Large Screens
          '2xl': '1536px' // Extra Large Screens
        }
      },
    },
    plugins: [],
}
