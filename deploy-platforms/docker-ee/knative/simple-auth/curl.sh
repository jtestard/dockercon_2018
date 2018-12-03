set -x
curl -X POST \
    -v \
    -H "Content-Type: application/json" \
    -d @payload.json \
    -H "Host: knative-simple-auth.default.example.com" \
    http://35.159.24.132:34380 \
