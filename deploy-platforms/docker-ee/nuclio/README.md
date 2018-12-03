# Nuclio Platform

**In this file, we detail how to deploy [Nuclio](https://nuclio.io/) as shown in the DockerCon demo.**

## Requirements

 - [Docker CLI](https://www.docker.com/get-docker)
 - A Docker EE cluster with admin credentials (see [here]() for how to install one)
 - A configured [Kubectl](https://kubernetes.io/docs/tasks/tools/install-kubectl/) (follow [this guide](https://docs.docker.com/ee/ucp/user-access/kubectl/) to configure it on your cluster)

## Deploy  Nuclio

Instructions shown here are largely inspired from [nuclio's own instructions](https://nuclio.io/docs/latest/setup/k8s/getting-started-k8s/).

Create the namespace and docker hub credentials:

```
kubectl create namespace nuclio
kubectl create secret docker-registry registry-credentials --namespace nuclio \
    --docker-username $username \
    --docker-password $mypassword \
    --docker-server registry.hub.docker.com \
    --docker-email ignored@nuclio.io
```

You can then do:

```
kubectl -n kube-system create sa tiller \
 && kubectl create clusterrolebinding tiller \
      --clusterrole cluster-admin \
      --serviceaccount=kube-system:tiller

helm init --skip-refresh --upgrade --service-account tiller
```

This is will initialize the Tiller component. Before continuing, wait until the output of `helm version` includes a server version in addition to the client's.

Now install nuclio:

```
kubectl apply -f nuclio.yaml
```

## TODO

 - Add support for service of type LoadBalancer for benchmarking purposes.
