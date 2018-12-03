Build

```
docker build . -t julestestard/faas-knative-starter:latest
docker push julestestard/faas-knative-starter:latest
```

Deploy

```
docker service create --name faas-knative-cli \
  -e UCP_USER="admin" \
  -e UCP_HOST="$UCP_FQDN" \
  -e UCP_SESSION_TOKEN="$AUTHTOKEN" \
  julestestard/faas-knative-starter:latest
```

Run Function

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
host=$(kubectl get services.serving.knative.dev knative-hello-world -o=custom-columns=NAME:.metadata.name,DOMAIN:.status.domain)
curl -v -X POST -H "Host: $host" http://$UCP_FQDN:34380/
```
