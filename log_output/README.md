# Log output

Generates a random UUID on startup and prints it with a timestamp every 5 seconds.

## Run

```sh
docker build -t xabim1/log-output:1.1 .
docker push xabim1/log-output:1.1

k3d cluster create -a 2
kubectl apply -f manifests/deployment.yaml
kubectl logs -f deploy/log-output-dep
```
