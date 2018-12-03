# OpenFaaS - Serverless Functions Made Simple

## Requirements

 - A Docker EE cluster with admin credentials (see [here]() for how to install one)
 - A configured [Kubectl](https://kubernetes.io/docs/tasks/tools/install-kubectl/) (follow [this guide](https://docs.docker.com/ee/ucp/user-access/kubectl/) to configure it on your cluster)
 - The [faas-cli](https://github.com/openfaas/faas-cli)

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

Finally, here is how to build, ship and deploy the function. `$PASSWORD` refers to your openfaas password which you specified when you install the openfaas platform:

```
gateway=http://$UCP_HOST:33000
echo $PASSWORD | faas-cli login -u admin --password-stdin --gateway $gateway
faas-cli build -f simple-auth.yml
faas-cli push -f simple-auth.yml
faas-cli deploy -f simple-auth.yml --gateway $gateway
```

Finally, invoke your function (`$UCP_HOST` is your docker ee ip):

```
set -x
curl -X POST \
    -v \
    -H "Content-Type: application/json" \
    -d @payload.json \
    http://$UCP_HOST:33000/function/simple-auth
```
