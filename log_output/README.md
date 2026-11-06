# Log output

Generates a random string on startup and prints it with a timestamp every 5 seconds. The same line, with the current timestamp, is served over HTTP on `/`. The port is set with the `PORT` environment variable (default 3000).

## Run

```sh
docker build -t xabim1/log-output:1.7 .
docker push xabim1/log-output:1.7

k3d cluster create --port 8082:30080@agent:0 -p 8081:80@loadbalancer --agents 2
kubectl apply -f manifests/
kubectl logs -f deploy/log-output-dep
```

The Ingress serves the app at http://localhost:8081.
