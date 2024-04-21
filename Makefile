build:
	templ generate
	@go build -o bin/dbproject.exe ./cmd
run: build
	npx tailwindcss -i ./view/css/app.css -o ./view/assets/css/styles.css
	@./bin/dbproject.exe
docker:
	docker-compose down
	docker-compose up --build
