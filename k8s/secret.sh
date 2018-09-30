#!/usr/bin/env bash

cat <<-EOF
apiVersion: v1
kind: Secret
metadata:
  name: betterdev
  namespace: yeo
type: Opaque
data:
  AWS_ACCESS_KEY_ID: `echo -n $AWS_ACCESS_KEY_ID | base64 -`
  AWS_SECRET_ACCESS_KEY: `echo -n $AWS_SECRET_ACCESS_KEY | base64 -`
EOF
