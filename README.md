# simple-server

The simple-server contains two binaries, a server and a client, where the server serves a static "Hello World" page and the client retrives the page via mutual TLS.

## Quick Start

The first, build and push image to a public registry:

`$ docker build -t quay.io/fanminshi/simple-server:latest .`

`$ docker push quay.io/fanminshi/simple-server:latest`

The second, before deploy the simpler server and simple client, we need to create the TLS asset for them:

`$ kubectl create secret tls server-secret --key tls/server-key.pem --cert tls/server.pem`
`$ kubectl create secret tls client-secret --key tls/client-key.pem --cert tls/client.pem`
`$ kubectl create configmap ca-cm --from-file=tls/ca.pem`

Next, deploy the simple server and simple client deployments and their corresponding services:

`$ kubectl create -f server_service.yaml`
`$ kubectl create -f deploy_server.yaml`
`$ kubectl create -f client_service.yaml`
`$ kubectl create -f deploy_client.yaml`

Verify that the deployment and service are succeeded:

```sh
$ kubectl get deploy
NAME                       DESIRED   CURRENT   UP-TO-DATE   AVAILABLE   AGE
simple-client-deployment   1         1         1            0           22s
simple-server-deployment   1         1         1            1           4m


$ kubectl get svc
NAME                    TYPE        CLUSTER-IP     EXTERNAL-IP   PORT(S)   AGE
simple-client-service   ClusterIP   None             <none>        8080/TCP    1m
simple-server-service   ClusterIP   10.106.4.202     <none>        8080/TCP    1m
```

Verify that the client can retrives the "Hello World" page from the server:

```sh
$ kubectl logs -f simple-client-deployment-59676cf8d9-p4tp
2018/07/31 20:22:12 <!DOCTYPE html>
<html>
  <head>
    <meta charset="utf-8">
    <title>hello world</title>
  </head>
  <body>
    <h1>hello world</h1>
  </body>
</html>
```