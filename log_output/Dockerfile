FROM golang:1.24-alpine AS build
WORKDIR /src
COPY go.mod main.go ./
RUN CGO_ENABLED=0 go build -o /log_output .

FROM scratch
COPY --from=build /log_output /log_output
CMD ["/log_output"]
