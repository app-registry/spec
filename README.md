# ⚠️️ DISCLAIMER ⚠️️

This specification is currently under active development and has not yet had a stable release.
Processes described by this set of documents is subject to change drastically in this time.
Now is a great time to voice opinions that will shape the future of the specification.
See [CONTRIBUTING.md](CONTRIBUTING.md) for more details.

# Container Native Registry

This repository contains the written specification for the Container Native Registry (CNR) specification.
This includes technical details on how containerized applications' configuration artifacts should be programatically discovered and distributed.
See [SPEC.md](SPEC.md) for details of the specification itself.

## What is the CNR spec?

Container Native Registry (CNR) is a well-specified and community developed specification for the distribution of configuration artifacts for containerized applications.
The core of CNR is an HTTP API specification defining the process for common operations such as search, upload, and download.

#### What is a containerized application configuration artifact?

TODO(jzelinskie): write this section

Examples of artifacts include:

* [Kubernetes Object Files](https://kubernetes.io/docs/concepts/abstractions/overview/)
* [Helm Charts](https://github.com/kubernetes/helm/blob/master/docs/charts.md)
* [KPM Packages](https://github.com/coreos/kpm/blob/master/Documentation/create_packages.md)
* [Docker Compose Files](https://docs.docker.com/compose/)
* [Docker Distributed Application Bundles](https://docs.docker.com/compose/bundles/)

## What is the promise of the CNR spec?

TODO(jzelinskie): write this section

## Who is using the spec?

TODO(jzelinskie): write this section

## What are some of the implementations of the spec?

### Server

#### Reference Implementations:

* [cnr-server](https://github.com/cn-app-registry/cnr-server) - open source Python implementation of a CNR server supporting storage via redis, etcd, or local filesystem.

#### Third Party Implementations:

* [Quay](https://quay.io) - commercial registry product from CoreOS (Python/Closed Source)

### Client

#### Reference Implementations:

* [cnr-python-cli](https://github.com/cn-app-registry/cnr-python-cli) - open source Python implementation of a CNR client
* [cnr-go-lib](https://github.com/cn-app-registry/cnr-go-lib) - open source Go implementation for a CNR client.


## Who controls the spec?

CNR is an open-source, community-driven project, developed under the Apache 2.0 license.
For information on governance and contribution policies, see [POLICY.md](POLICY.md)
