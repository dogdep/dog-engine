# POC 2. Setup a custom ingress on GKE and test the use case

Use case.
Whenever developer deploys a release we must provide a friendly dns name.
Developer might choose to deploy few releases that should be accessible
through different URL's.

IE. We have a fictional project ACME, upon deploying a release it should get
a dns friendly name such as:

```
acme.project.dogdep.io
```

If the developer chooses to deploy to different endpoint then it would be:
```
beta.acme.project.dogdep.io
```

This will allow to have multiple versions, potentially with different features
for testing or demoing.

Main goal is to setup a GKE cluster with custom nginx ingress (provided by k8s)
and test if we can get away without using google load balancer for this setup.

There are few ways to do this.

1. Using GLBC and for each domain add a forwarding rule
2. Using GLBC and for `*.dogdep.io` have a forwarding rule which would forward
requests to internal nginx service, which then would proxy to the actual service.
3. Using Nginx ingress on google cloud and as mentioned in 1 point add forwarding
rules.

While its OK to try [2], but this incurs additional cost and because we wont
run production loads it makes sense to try out a cheaper alternative.

## Prerequisites

- docker is set up on your local machine
- gcloud cli installed
- gcloud account with billables

## Walkthrough

First step is to clear a cluster within with node pool across multiple av zones.

```
gcloud container clusters create poc-ingress \
--zone europe-west1-d \
--additional-zones europe-west1-b,europe-west1-c \
--machine-type g1-small \
--num-nodes 1
```

Second step after it is created is to get the credentials

```
gcloud container clusters get-credentials poc-ingress \
    --zone europe-west1-d
```

Lets deploy a sample app:
```
kubectl run echoheaders --image=gcr.io/google_containers/echoserver:1.4 --replicas=1 --port=8080
```

Lets expose the app through different services:
```
kubectl expose deployment echoheaders --port=80 --target-port=8080 --name=echoheaders-x
kubectl expose deployment echoheaders --port=80 --target-port=8080 --name=echoheaders-y
```

Now lets deploy nginx ingress:

```
kubectl create -f examples/default-backend.yaml
kubectl expose rc default-http-backend --port=80 --target-port=8080 --name=default-http-backend

kubectl create -f k8s/rc-default.yaml

kubectl run echoheaders --image=gcr.io/google_containers/echoserver:1.4 --replicas=1 --port=8080

kubectl expose deployment echoheaders --port=80 --target-port=8080 --name=echoheaders-x
kubectl expose deployment echoheaders --port=80 --target-port=8080 --name=echoheaders-y

kubectl create -f k8s/ingress-nginx.yaml

kubectl create -f k8s/ingress.yml
```


References.
1. https://github.com/kubernetes/ingress/tree/master/controllers/nginx
2. https://beroux.com/english/articles/kubernetes/?part=3
