FROM golang:1.24-alpine AS build
WORKDIR /src
COPY go.mod main.go ./
RUN CGO_ENABLED=0 go build -o /todo_app .

FROM scratch
COPY --from=build /todo_app /todo_app
CMD ["/todo_app"]
