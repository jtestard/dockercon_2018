## Install Knative

To install Knative on Docker EE, follow installation instructions from [the knative documentation](https://github.com/knative/docs/blob/master/install/Knative-with-any-k8s.md) but replace the
`istio.yaml` and `release.yaml` files used in the document by those from this directory. Make sure to install
both the serving and build components if you want to deploy the functions showed in the demo.

The `istio.yaml` and `release.yaml` have been modified in this directory to be compatible with docker ee.
