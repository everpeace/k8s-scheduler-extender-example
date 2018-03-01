# k8s-scheduler-extender-example
This is an example of [Kubernetes Scheduler Extender](https://github.com/kubernetes/community/blob/master/contributors/design-proposals/scheduling/scheduler_extender.md)

## How to

### 1. buid a docker image

```
$ docker build . -t YOUR_ORG/YOUR_IMAGE:YOUR_TAG
$ docker push YOUR_ORG/YOUR_IMAGE:YOUR_TAG
```

### 2. deploy `my-scheduler` in `kube-system` namespace
please see ConfigMap in [extender.yaml](extender.yaml) for scheduler policy json which includes scheduler extender config.

```
# edit extender.yaml to specify your docker image
$ $EDITOR extender.yaml

# deploy it.
$ kubectl create -f extender.yaml
```

### 3. schedule test pod
you will see `test-pod` will be scheduled by `my-scheduler`.

```
$ kubectl create -f test-pod.yaml

$ kubectl desceribe test-pod
Name:         test-pod
...
Events:
  Type    Reason                 Age   From               Message
  ----    ------                 ----  ----               -------
  Normal  Scheduled              25s   my-scheduler       Successfully assigned test-pod to minikube
  Normal  SuccessfulMountVolume  25s   kubelet, minikube  MountVolume.SetUp succeeded for volume "default-token-wrk5s"
  Normal  Pulling                24s   kubelet, minikube  pulling image "nginx"
  Normal  Pulled                 8s    kubelet, minikube  Successfully pulled image "nginx"
  Normal  Created                8s    kubelet, minikube  Created container
  Normal  Started                8s    kubelet, minikube  Started container
```
