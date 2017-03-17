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

## Who controls the spec?

CNR is an open-source, community-driven project, developed under the Apache 2.0 license.
For information on governance and contribution policies, see [POLICY.md](POLICY.md)

## Who is using the spec?

* [Quay](https://quay.io) is a commercial registry product from CoreOS that implements an early prototype of the CNR specification.

Want to get added to this list?
[Open a Pull Request!](https://github.com/cn-app-registry/spec/edit/master/README.md)

## What are some of the implementations of the spec?

#### Reference Implementations:

* [cnr-server](https://github.com/cn-app-registry/cnr-server) - open source Python implementation of a CNR server backed by redis, etcd, or a local filesystem.
* [cnr-cli](https://github.com/cn-app-registry/cnr-cli) - open source Python implementation of a CNR client

#### Third-Party Implementations:

TODO(jzelinskie): write this section
