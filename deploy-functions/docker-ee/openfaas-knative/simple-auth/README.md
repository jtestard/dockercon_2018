# OpenFaaS on Knative

## Requirements

 - A Docker EE cluster with admin credentials (see [here]() for how to install one)
 - A configured [Kubectl](https://kubernetes.io/docs/tasks/tools/install-kubectl/) (follow [this guide](https://docs.docker.com/ee/ucp/user-access/kubectl/) to configure it on your cluster)
 - The [faas-cli](https://github.com/openfaas/faas-cli)
 - OpenFaaS and Knative installed

### Simple Auth example

Now edit the `simple-auth.yml` file as follows:

```
provider:
  name: faas

functions:
  hello-world:
    lang: golang-http
    handler: ./simple-auth
    image: $DOCKER_HUB_USERNAME/simple-auth
```

where `$DOCKER_HUB_USERNAME` is your Docker Hub username.

Here is how to build the function. `$PASSWORD` refers to your openfaas password which you specified when you install the openfaas platform:

```
gateway=http://$UCP_HOST:33000
echo $PASSWORD | faas-cli login -u admin --password-stdin --gateway $gateway
faas-cli build -f simple-auth.yml
faas-cli push -f simple-auth.yml
```

### Deploy using Knative

Apply the following yaml:

```
kubectl apply -f - << EOF
apiVersion: serving.knative.dev/v1alpha1
kind: Service
metadata:
  name: knative-simple-auth
  namespace: default
spec:
  runLatest:
    configuration:
      revisionTemplate:
        spec:
          container:
            image: julestestard/simpleauth:latest
EOF
```

Find the knative host used:

```
kubectl get services.serving.knative.dev knative-hello-world -o=custom-columns=NAME:.metadata.name,DOMAIN:.status.domain
```

Finally, invoke your function, where `$host` is the knative function hostname and `$UCP_HOST` is the IP of a host on your Docker EE cluster:

```
curl -X POST \
    -v \
    -H "Content-Type: application/json" \
    -d @payload.json \
    -H "Host: $host" \
    http://$UCP_HOST:34380 \
```
