watch_tailwind:
	npx tailwindcss -i ./style/index.css -o ./assets/app.css --watch

watch:
	reflex -s -r '.*\.(go|html|css|js|tmpl)' -- go run main.go
