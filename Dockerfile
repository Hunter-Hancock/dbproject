FROM golang:alpine

RUN apk add --no-cache nodejs npm

# Install tailwindcss using npm
RUN npm i -D tailwindcss

WORKDIR /dbproject

COPY . .

RUN go mod download

RUN go install github.com/a-h/templ/cmd/templ@latest

RUN templ generate

RUN go build -o bin/dbproject ./cmd
RUN go build -o bin/migrate ./cmd/migrate 

# Run tailwindcss command
RUN npx tailwindcss -i ./view/css/app.css -o ./view/assets/css/styles.css

CMD ["./bin/dbproject"]
EXPOSE 3000