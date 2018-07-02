#!/bin/bash

kubectl create -f fs1-deploy.yaml
kubectl create -f fs2-deploy.yaml

kubectl expose deployment fs1 --type=ClusterIP --name=fs1
kubectl expose deployment fs2 --type=ClusterIP --name=fs2

