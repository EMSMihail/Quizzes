FROM golang:1.20

RUN mkdir /app
WORKDIR /app

COPY go.mod go.sum /app/
RUN go mod download && go mod verify

COPY . .

RUN chmod +x /app/docker/web_app/check.sh

# ENV DATABASE_URL=postgres://aizek:1234@quizzes_db/quizzes?sslmode=disable

ENV STATIC_DIR=/app/web/static

ENV TEMPLATES_DIR=/app/web/templates

ENV PORT=5000

RUN go build -v -o /app/cmd/app/main /app/cmd/app/main.go

EXPOSE 5000

CMD ["/app/docker/web_app/check.sh", "/app/cmd/app/main"]