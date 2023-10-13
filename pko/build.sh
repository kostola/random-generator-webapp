#!/usr/bin/env bash

set -e
set -x

kubectl package validate backend
rm -f backend.tar
kubectl package build -t quay.io/alcosta/package-operator-packages/random-generator-webapp/backend -o backend.tar backend
podman image load -i backend.tar
podman push quay.io/alcosta/package-operator-packages/random-generator-webapp/backend:latest

kubectl package validate frontend
rm -f frontend.tar
kubectl package build -t quay.io/alcosta/package-operator-packages/random-generator-webapp/frontend -o frontend.tar frontend
podman image load -i frontend.tar
podman push quay.io/alcosta/package-operator-packages/random-generator-webapp/frontend:latest
