# Nuclio

This file contains the nuclio sample used during the DockerCon 2018 demo.

## Requirements

 - A Docker EE cluster with admin credentials (see [here]() for how to install one)
 - A configured [Kubectl](https://kubernetes.io/docs/tasks/tools/install-kubectl/) (follow [this guide](https://docs.docker.com/ee/ucp/user-access/kubectl/) to configure it on your cluster)
 - The [nuclio cli tool `nuctl`](https://github.com/nuclio/nuclio/releases)

## Hello World Example

```
nuctl deploy simple-auth -v --path simple-auth --registry julestestard
nuctl invoke -n nuclio simple-auth --content-type "application/json" -b '{
    "specversion" : "0.1",
    "type" : "com.example.authLogin",
    "source" : "/mycontext",
    "id" : "C234-1234-1234",
    "time" : "2018-04-05T17:31:00Z",
    "contenttype" : "application/json",
    "data" : {
        "user" : "alice",
        "authToken" : "0af23ed13"
    }
}'
```


