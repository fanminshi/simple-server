# simple-server

The simple-server serves a static Hello World page via TLS.

## Quick Start

The first, build and push image to a public registry:

`$ docker build -t quay.io/fanminshi/simple-server:latest .`

`$ docker push quay.io/fanminshi/simple-server:latest`

The second, before deploy the simpler server, we need to create the tls asset that the server needs:

`kubectl create secret tls simple-server-tls --key tls/server-key.pem --cert tls/server.pem`

Next, deploy the simple server deployment and its corresponding service:

`$ kubectl create -f deploy.yaml`
`$ kubectl create -f service.yaml`

Verify that the deployment and service are succeeded:

```sh
$ kubectl get deploy
NAME                       DESIRED   CURRENT   UP-TO-DATE   AVAILABLE   AGE
simple-server-deployment   1         1         1            1           1m


$ kubectl get svc
NAME                    TYPE        CLUSTER-IP     EXTERNAL-IP   PORT(S)   AGE
simple-server-service   ClusterIP   10.108.1.231   <none>        443/TCP   1m
```

Download and run a busybox-curl pod:

`$ kubectl run curl --image=radial/busyboxplus:curl -i --tty`


Retrieve the `hello world` page:

```sh
$ curl -k https://simple-server-service:443/
<!DOCTYPE html>
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