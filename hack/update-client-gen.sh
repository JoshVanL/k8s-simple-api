#!/bin/bash

# The only argument this script should ever be called with is '--verify-only'

set -o errexit
set -o nounset
set -o pipefail

REPO_ROOT=$GOPATH/src/github.com/joshvanl/k8s-simple-api
BINDIR=${REPO_ROOT}/bin

# Generate the internal clientset (pkg/client/clientset_generated/internalclientset)
${BINDIR}/client-gen "$@" \
	      --input-base "github.com/joshvanl/k8s-simple-api/pkg/apis/" \
	      --input "simple/" \
	      --clientset-path "github.com/joshvanl/k8s-simple-api/pkg/client/" \
	      --clientset-name internalclientset \
	      --go-header-file "${GOPATH}/src/github.com/kubernetes/repo-infra/verify/boilerplate/boilerplate.go.txt"
# Generate the versioned clientset (pkg/client/clientset_generated/clientset)
${BINDIR}/client-gen "$@" \
		  --input-base "github.com/joshvanl/k8s-simple-api/pkg/apis/" \
		  --input "simple/v1alpha1" \
	      --clientset-path "github.com/joshvanl/k8s-simple-api/pkg/" \
	      --clientset-name "client" \
	      --go-header-file "${GOPATH}/src/github.com/kubernetes/repo-infra/verify/boilerplate/boilerplate.go.txt"
# generate lister
${BINDIR}/lister-gen "$@" \
		  --input-dirs="github.com/joshvanl/k8s-simple-api/pkg/apis/simple" \
	      --input-dirs="github.com/joshvanl/k8s-simple-api/pkg/apis/simple/v1alpha1" \
	      --output-package "github.com/joshvanl/k8s-simple-api/pkg/listers" \
	      --go-header-file "${GOPATH}/src/github.com/kubernetes/repo-infra/verify/boilerplate/boilerplate.go.txt"
# generate informer
${BINDIR}/informer-gen "$@" \
	      --go-header-file "${GOPATH}/src/github.com/kubernetes/repo-infra/verify/boilerplate/boilerplate.go.txt" \
	      --input-dirs "github.com/joshvanl/k8s-simple-api/pkg/apis/simple" \
	      --input-dirs "github.com/joshvanl/k8s-simple-api/pkg/apis/simple/v1alpha1" \
	      --internal-clientset-package "github.com/joshvanl/k8s-simple-api/pkg/client/internalclientset" \
	      --versioned-clientset-package "github.com/joshvanl/k8s-simple-api/pkg/client" \
	      --listers-package "github.com/joshvanl/k8s-simple-api/pkg/listers" \
	      --output-package "github.com/joshvanl/k8s-simple-api/pkg/informers"
