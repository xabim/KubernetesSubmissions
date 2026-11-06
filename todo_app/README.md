# Todo app

Web server that prints `Server started in port NNNN` on startup. The port is set with the `PORT` environment variable (default 8080).

## Run

```sh
docker build -t xabim1/todo-app:1.8 .
docker push xabim1/todo-app:1.8

kubectl apply -f manifests/
kubectl logs -f deploy/todo-app-dep
```

The Service is a ClusterIP and the Ingress serves the app at http://localhost:8081 when the cluster was created with:

```sh
k3d cluster create --port 8082:30080@agent:0 -p 8081:80@loadbalancer --agents 2
```
