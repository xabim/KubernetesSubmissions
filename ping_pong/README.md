# Ping-pong

Responds to `GET /pingpong` with `pong N`, where N is an in-memory counter that starts at 0 and increases on every request. The port is set with the `PORT` environment variable (default 3000).

## Run

```sh
docker build -t xabim1/ping-pong:1.9 .
docker push xabim1/ping-pong:1.9

kubectl apply -f manifests/
```

The Ingress is shared with Log output, see `../log_output/manifests/ingress.yaml`. With the cluster created as in `../log_output/README.md`, the app is at http://localhost:8081/pingpong.
