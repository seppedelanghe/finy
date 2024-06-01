module.exports = {
	content: ["./templates/**/*.{html,js,tmpl}"],
	theme: {
		extend: {
			colors: { 
				'primary':  '#FAFAFA',
				'secondary': '#e8e8e8',
				'jade': { DEFAULT: '#0DB265', 100: '#032314', 200: '#054628', 300: '#08693c', 400: '#0b8c50', 500: '#0db265', 600: '#12eb86', 700: '#4cf2a4', 800: '#87f6c2', 900: '#c3fbe1' }, 
				'poppy': { DEFAULT: '#E41134', 100: '#2e030a', 200: '#5b0715', 300: '#890a1f', 400: '#b60e2a', 500: '#e41134', 600: '#f03a58', 700: '#f46b82', 800: '#f89cac', 900: '#fbced5' }, 
				'night': { DEFAULT: '#0C0C0C', 100: '#030303', 200: '#050505', 300: '#080808', 400: '#0a0a0a', 500: '#0c0c0c', 600: '#3d3d3d', 700: '#6e6e6e', 800: '#9e9e9e', 900: '#cfcfcf' } 
			}
		},
	},
	plugins: [],
}
