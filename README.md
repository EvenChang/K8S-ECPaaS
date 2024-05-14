# How to develop a controller in ECPaaS Project.

## Add CRD resource definition
1. Define CRD resource at staging/src/kubesphere.io/api/ path

## Add Controller behavior
1. Add the reconcile code at pkg/controller

## Modify Makefile
1. Add scheme and your module path in global variable **GV** and **MANIFESTS**

## Generate relate code automatically
1. make deepcopy
2. make clientset
3. make manifests
4. make ks-controller-manager

## Test Controller behavior
1. Install telepresence
```
sudo curl -fL https://app.getambassador.io/download/tel2/linux/amd64/latest/telepresence -o /usr/local/bin/telepresence

sudo chmod a+x /usr/local/bin/telepresence

telepresence helm install
```

2. Run k8s-controller-manager
```
telepresence --namespace  kubesphere-system --swap-deployment ks-controller-manager --run go run controller-manager.go
```

If you encounter this issues 
> F0514 15:23:22.817881   32404 server.go:254] unable to run the manager: open /tmp/k8s-webhook-server/serving-certs/tls.crt: no such file or directory

Enter these commands to get relative secure files.
```
kubectl get secret ks-controller-manager-webhook-cert -n kubesphere-system -o jsonpath="{.data.ca\.crt}" | base64 --decode > /tmp/k8s-webhook-server/serving-certs/ca.crt
kubectl get secret ks-controller-manager-webhook-cert -n kubesphere-system -o jsonpath="{.data.tls\.crt}" | base64 --decode > /tmp/k8s-webhook-server/serving-certs/tls.crt
kubectl get secret ks-controller-manager-webhook-cert -n kubesphere-system -o jsonpath="{.data.tls\.key}" | base64 --decode > /tmp/k8s-webhook-server/serving-certs/tls.key
```
