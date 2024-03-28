/** @type {import('tailwindcss').Config} */
module.exports = {
    darkMode: 'selector',
    content: [
        "./internal/**/*.{go,js,templ,html}"
    ],
    theme: {
      colors: {
        "black": '#111111',
        "darkborder": "#2e2e2e",
        "darkgray": '#222222',
        "gray": '#999999',
        "lightgray": '#DDDDDD',
        "faintgray": '#EEEEEE',
        "white": '#FAFAFA',
        "blue": '#2c6dde',
        "logoyellow": '#FFC107',
        "orange": "#f08d49",
      },
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
