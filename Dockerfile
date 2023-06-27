FROM golang:1.20

WORKDIR /

COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .

# RUN go mod init Quizzes
# RUN go mod tidy

# RUN echo "nameserver 8.8.8.8" >> /etc/resolv.conf
ENV DATABASE_URL=postgres://aizek:1234@quizzes_db/quizzes?sslmode=disable


ENV STATIC_DIR=/app/web/static

ENV TEMPLATES_DIR=/app/web/templates

ENV PORT=5000

RUN go build -v -o /app/main /cmd/app/main.go

EXPOSE 5000

CMD ["/app/main"]