# Todo app

Web server that prints `Server started in port NNNN` on startup. The port is set with the `PORT` environment variable (default 8080).

## Run

```sh
docker build -t xabim1/todo-app:1.2 .
docker push xabim1/todo-app:1.2

kubectl apply -f manifests/deployment.yaml
kubectl logs -f deploy/todo-app-dep
```
