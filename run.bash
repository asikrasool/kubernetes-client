# to run locally
docker build --build-arg USE_KUBECONFIG=True -t asikrasool/client-go:0.1 .
docker run  -v $HOME/.kube/config:/root/.kube/config asikrasool/client-go:0.1

# to run in kubernetes
docker build -t asikrasool/client-go:latest .
docker push asikrasool/client-go:latest
kubectl apply -f  k8s/
