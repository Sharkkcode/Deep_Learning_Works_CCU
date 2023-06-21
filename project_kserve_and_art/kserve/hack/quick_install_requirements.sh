curl -L https://istio.io/downloadIstio | sh -

cat <<EOF | kubectl apply -f -
apiVersion: v1
kind: Namespace
metadata:
  name: istio-system
  labels:
    istio-injection: disabled
EOF

cat << EOF > ./istio-minimal-operator.yaml
apiVersion: install.istio.io/v1beta1
kind: IstioOperator
spec:
  values:
    global:
      proxy:
        autoInject: disabled
      useMCP: false

  meshConfig:
    accessLogFile: /dev/stdout

  components:
    ingressGateways:
      - name: istio-ingressgateway
        enabled: true
        k8s:
          podAnnotations:
            cluster-autoscaler.kubernetes.io/safe-to-evict: "true"
    pilot:
      enabled: true
      k8s:
        resources:
          requests:
            cpu: 200m
            memory: 200Mi
        podAnnotations:
          cluster-autoscaler.kubernetes.io/safe-to-evict: "true"
        env:
        - name: PILOT_ENABLE_CONFIG_DISTRIBUTION_TRACKING
          value: "false"
EOF

istio-1.17.2/bin/istioctl manifest apply -f istio-minimal-operator.yaml -y

rm -rf istio-1.17.2
rm istio-minimal-operator.yaml

kubectl apply -f https://github.com/knative/serving/releases/download/knative-v1.10.1/serving-crds.yaml
kubectl apply -f https://github.com/knative/serving/releases/download/knative-v1.10.1/serving-core.yaml
kubectl apply -f https://github.com/knative/net-istio/releases/download/knative-v1.10.0/release.yaml

kubectl apply -f https://github.com/cert-manager/cert-manager/releases/download/v1.11.0/cert-manager.yaml
kubectl wait --for=condition=available --timeout=600s deployment/cert-manager-webhook -n cert-manager

echo "Succeed 😀"
