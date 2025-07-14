const defaultTheme = require('tailwindcss/defaultTheme');
const colors = require('tailwindcss/colors');

module.exports = {
  content: [
    "./index.html",
    "./src/**/*.{vue,js,ts,jsx,tsx}",
    // Include Material Tailwind paths for scanning
    "./node_modules/@material-tailwind/html/theme/components/**/*.{js,ts}",
  ],
  theme: {
    extend: {
      fontFamily: {
        // Headings and prominent text
        serif: ['Montserrat', ...defaultTheme.fontFamily.sans],
        // Body text and general UI
        sans: ['Open Sans', ...defaultTheme.fontFamily.sans],
      },
      colors: {
        // Primary brand color
        primary: colors.blue,
        // Accent color for interactive elements
        accent: colors.teal,
        // Neutral shades for backgrounds, borders, subtle text
        neutral: colors.slate,
        // Custom shades for specific use cases
        'light-bg': colors.slate[50],
        'dark-bg': colors.slate[900],
        'light-text': colors.blueGray[900],
        'dark-text': colors.white,
      },
    },
  },
  plugins: [require('@tailwindcss/line-clamp')],
};