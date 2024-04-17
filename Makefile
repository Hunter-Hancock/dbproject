build:
	templ generate
	@go build -o bin/dbproject ./cmd
run: build
	npx tailwindcss -i ./view/css/app.css -o ./view/assets/css/styles.css
	@./bin/dbproject
docker:
	docker-compose up --build
