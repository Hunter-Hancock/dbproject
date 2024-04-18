FROM golang:alpine

RUN apk add --no-cache nodejs npm

WORKDIR /dbproject

COPY go.mod go.sum ./
RUN go mod download

RUN go install github.com/a-h/templ/cmd/templ@latest

COPY . .

RUN templ generate

RUN go build -o bin/dbproject ./cmd && go build -o bin/migrate ./cmd/migrate

# Install tailwindcss using npm
RUN npm install -D tailwindcss

# Run tailwindcss command
RUN npx tailwindcss -i ./view/css/app.css -o ./view/assets/css/styles.css

CMD ["/dbproject/bin/dbproject"]
EXPOSE 3000
