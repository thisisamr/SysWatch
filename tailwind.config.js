/** @type {import('tailwindcss').Config} */
module.exports = {
  darkMode: ["class"],
  content: ["./**/*.html", "./**/*.templ", "./**/*.go"],
  theme: {
    extend: {
      colors: {
        "dark-bg": "#212529",
      },
    },
  },
  plugins: [],
};
