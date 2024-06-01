watch_tailwind:
	npx tailwindcss -i ./style/index.css -o ./assets/app.css --watch

setup_env:
	export GOOS=darwin
	export GOARCH=arm64
	export CGO_ENABLED=1

run:
	go run main.go

watch:
	reflex -s -r '.*\.(go|html|css|js|tmpl)' -R '^node_modules/' -- sh -c 'go run main.go'

build:
	go build -o finy main.go
