FROM golang:alpine

WORKDIR /dbproject

COPY go.mod go.sum ./
RUN go mod download

RUN go install github.com/a-h/templ/cmd/templ@latest

COPY . .

RUN templ generate

RUN go build -o bin/dbproject ./cmd && go build -o bin/migrate ./cmd/migrate

CMD ["/dbproject/bin/dbproject"]
EXPOSE 3000
