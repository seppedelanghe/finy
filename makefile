watch_tailwind:
	npx tailwindcss -i ./style/index.css -o ./assets/app.css --watch

watch:
	export GOOS=darwin
	export GOARCH=arm64
	export CGO_ENABLED=1
	reflex -s -r '.*\.(go|html|css|js|tmpl)' -- go run main.go
