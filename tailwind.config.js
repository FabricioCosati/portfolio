/** @type {import('tailwindcss').Config} */
import daisyUi from 'daisyui'
import flowbite from 'flowbite/plugin'
import tailgrid from 'tailgrids/plugin'

module.exports = {
  content: ['./internal/**/*.go', './internal/**/*.html', './internal/**/*.templ', './node_modules/tailgrids/**/*.js', "./internal/**/*.{html,js}"],
  darkMode: 'class',
  theme: {
    extend: {
      fontFamily: {
        title: ['Prompt', 'sans-serif'],
        text: ['Space Mono', 'monospace']
      },
      keyframes: {
        fadeUpTitle: {
          '0%': {
            opacity: "0",
            transform: "translateY(30px) scale(0.9)",
          },
          '100%': {
            opacity: "1",
            transform: "translateY(0px) scale(1)",
          },
        },
        fadeUp: {
          from: {
            opacity: "0",
            scale: "0.5",
          },
          to: {
            opacity: "1",
            scale: "1",
          },
        },
        blink: {
          '50%': { opacity: '0' },
        }
      },
      animation: {
        fadeUpTitle: 'fadeUpTitle 0.5s',
        fadeUp: 'fadeUp linear',
        blink: 'blink 700ms cubic-bezier(0.18, 0.89, 0.32, 1.28) infinite',
      },
      supports: {
        "no-scroll-driven-animations": "not(animation-timeline: scroll())",
      },
      boxShadow: {
        'journeyBox': '0 0 0 2px #fff, inset 0 2px 0 rgba(0, 0, 0, .08), 0 3px 0 4px rgba(0, 0, 0, .05)',
      },
      screens: {
        '2xl': '1536px',
        '2sm': '701px'
      }
    },
  },
  corePlugins: {
    preflight: true
  },
  plugins: [
    tailgrid,
    flowbite,
    daisyUi,
  ],
}