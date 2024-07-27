run:
	go run ./cmd/web -cache=true

tailwind:
	npx tailwindcss -i ui/static/css/main.css -o ui/static/css/tailwind.css --watch