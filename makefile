SHELL := /bin/bash

run:
	go run main.go

build: 
	go build -ldflags "-X main.build=local"

VERSION = 1.0

all: service

service:
	docker build \
		-f zarf/docker/dockerfile \
		-t service-amd64:${VERSION} \
		--build-arg BUILD_REF=${VERSION} \
		--build-arg BUILD_DATE=`date -u +"%Y-%m-%dT%H:%M:%SZ"` \
		.

# ==============================================================================
# Running from within k8s/kind

KIND_CLUSTER := ardan-starter-cluster

kind-up:
	kind create cluster \
		--image kindest/node:v1.23.4@sha256:0e34f0d0fd448aa2f2819cfd74e99fe5793a6e4938b328f657c8e3f81ee0dfb9 \
		--name $(KIND_CLUSTER) \
		--config zarf/k8s/kind/kind-config.yaml

kind-down:
	kind delete cluster --name $(KIND_CLUSTER)

kind-load:
	kind load docker-image service-amd64:$(VERSION) --name $(KIND_CLUSTER)

kind-apply:
	cat zarf/k8s/base/service-pod/base-service.yaml | kubectl apply -f -

kind-logs:
	kubectl logs -l app=service --all-containers=true -f --tail=100 --namespace=service-system

kind-restart:
	kubectl rollout restart deployment service-pod --namespace=service-system

kind-update: service kind-load kind-restart

kind-status:
	kubectl get nodes -o wide
	kubectl get svc -o wide
	kubectl get pods -o wide --watch --all-namespaces

kind-status-service:
	kubectl get pods -o wide --watch --namespace=service-system

kind-describe:
	kubectl describe pod -l app=service