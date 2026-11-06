FROM golang:1.24-alpine AS build
WORKDIR /src
COPY go.mod main.go ./
RUN CGO_ENABLED=0 go build -o /ping_pong .

FROM scratch
COPY --from=build /ping_pong /ping_pong
CMD ["/ping_pong"]
