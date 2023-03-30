[Best practices for developing microservices based on Go/Grpc/kubernetes/Istio - Part01](https://github.com/janrs-io/Jgrpc/blob/master/best-practices-for-developing-microservices-based-on-go-grpc-kubernetes-Istio-part01.md "Best practices for developing microservices based on Go/Grpc/kubernetes/Istio - Part01")

[Best practices for developing microservices based on Go/Grpc/kubernetes/Istio - Part02](https://github.com/janrs-io/Jgrpc/blob/master/best-practices-for-developing-microservices-based-on-go-grpc-kubernetes-Istio-part02.md "Best practices for developing microservices based on Go/Grpc/kubernetes/Istio - Part02")

[Best practices for developing microservices based on Go/Grpc/kubernetes/Istio - Part03](https://github.com/janrs-io/Jgrpc/blob/master/best-practices-for-developing-microservices-based-on-go-grpc-kubernetes-Istio-part03.md "Best practices for developing microservices based on Go/Grpc/kubernetes/Istio - Part03")

***

In the previous two parts, we created two microservices: `pingservice` and `pongservice`. In this part, we will
create `CICD` Pipeline for automatic deployment.

We assume that you have deployed `Jenkins/Gitlab/Harbor` and `Kubenertes/Istio`.

# Project Structure

```bash
devops
├── README.md
├── ping
│   └── dev
│       ├── Deployment.yaml
│       ├── Dockerfile
│       ├── Jenkinsfile
│       └── Service.yaml
└── pong
    └── dev
        ├── Deployment.yaml
        ├── Dockerfile
        ├── Jenkinsfile
        └── Service.yaml

4 directories, 9 files
```

# Usage

On Jenkinfs, create a directory for each microservice project, and then create a `dev/test/prod` pipeline under the
directory.

On Gitlab, set up three branch protection branches: `dev/test/prod`. These three branches are used
for `dev/test/production` environments.These three branches can only be merged but not submitted.

If there is a new microservice to be developed, create a new branch based on the dev branch, and the name format
is: `dev-*`. For example: dev-ping, dev-pong.
Then set up a webhook for each branch to automatically trigger the Jenkins pipeline to automatically deploy to the
kubernetes cluster.

Local development of microservices requires debugging. You can use
the [kubefwd](https://github.com/txn2/kubefwd "kubefwd") tool
or [telepresence](https://kubernetes.io/zh-cn/docs/tasks/debug/debug-cluster/local-debugging/ "telepresence") officially
recommended by kubernetes.

# Advanced usage

If your company has grown to a group size and needs collaborative development in different places, you can
separate `devops`, `istio-manifests`, `kubernetes-manifests` and create an independent git-repo for management.

And it is also possible to separate different microservices under the `src/` directory and create different git-repos
for management.

Different teams need to document the developed grpc interface and publish it online, and all personnel develop and debug
according to the online interface document.

# Related projects and materials

Thanks to the contributors of the following resources:

- [GoogleCloudPlatform/microservices-demo](https://github.com/GoogleCloudPlatform/microservices-demo "GoogleCloudPlatform / microservices-demo")

- [GitOps Decisions: Creating a Mature Pipeline](https://blog.container-solutions.com/gitops-decisions "GitOps Decisions: Creating a Mature Pipeline")

- [buf build](https://buf.build/ "buf build")

- [wire](https://github.com/google/wire "wire")

- [gRPC Ecosystem](https://github.com/grpc-ecosystem "gRPC Ecosystem")









