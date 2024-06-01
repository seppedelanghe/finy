watch_tailwind:
	npx tailwindcss -i ./style/index.css -o ./assets/app.css --watch

watch_api:
	air --build.exclude_dir "node_modules"
