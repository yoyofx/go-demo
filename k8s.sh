kubectl create -f k8sapp.yaml

kubectl expose deployment godemo1 --type="LoadBalancer"