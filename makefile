watch_tailwind:
	npx tailwindcss -i ./style/index.css -o ./assets/app.css --watch

run:
	export GOOS=darwin
	export GOARCH=arm64
	export CGO_ENABLED=1
	go run main.go

watch:
	export GOOS=darwin && \
	export GOARCH=arm64 && \
	export CGO_ENABLED=1 && \
	ulimit -n 1024 && \
	reflex -s -r '.*\.(go|html|css|js|tmpl)' -R '^node_modules/' -- sh -c 'go run -v main.go'

# watch:
# 	export GOOS=darwin
# 	export GOARCH=arm64
# 	export CGO_ENABLED=1
# 	reflex -s -r '.*\.(go|html|css|js|tmpl)' -- go run -v main.go

build:
	export GOOS=darwin
	export GOARCH=arm64
	export CGO_ENABLED=1
	go build main.go
