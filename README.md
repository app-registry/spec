# ⚠️️ DISCLAIMER ⚠️️

This specification is currently under active development and has not yet had a stable release.
Processes described by this set of documents is subject to change drastically in this time.
Now is a great time to voice opinions that will shape the future of the specification.
See [CONTRIBUTING.md](CONTRIBUTING.md) for more details.

# App Registry

This repository contains the definition for the App Registry specification.
This includes technical details on how containerized applications' configuration artifacts should be programatically discovered and distributed.
See [SPEC.md](SPEC.md) for details of the specification itself.

## What is the App Registry spec?

App Registry is a well-specified and community developed specification for the distribution of declarative configuration artifacts for containerized applications.
The core of App Registry is an HTTP API specification defining the process for common operations such as search, upload, and download.

#### What is a declarative application configuration (DAC)?

TODO(jzelinskie): write this section

Examples of DAC formats include:

* [Kubernetes Object Files](https://kubernetes.io/docs/concepts/abstractions/overview/)
* [Helm Charts](https://github.com/kubernetes/helm/blob/master/docs/charts.md)
* [KPM Packages](https://github.com/coreos/kpm/blob/master/Documentation/create_packages.md)
* [Docker Compose Files](https://docs.docker.com/compose/)
* [Docker Distributed Application Bundles](https://docs.docker.com/compose/bundles/)

## What is the promise of the spec?

TODO(jzelinskie): write this section

## Who controls the spec?

App Registry is an open-source, community-driven project, developed under the Apache 2.0 license.
For information on governance and contribution policies, see [POLICY.md](POLICY.md)

## Who is using the spec?

* [Quay](https://quay.io) is a commercial registry product from CoreOS that implements an early prototype of the server.
* The [Helm Registry Plugin](https://github.com/app-registry/cnr-cli) implements an early protoype of the client.

Want to get added to this list?
[Open a Pull Request!](https://github.com/app-registry/spec/edit/master/README.md)

## What are some of the implementations of the spec?

#### Reference Implementations:

* [cnr-server](https://github.com/app-registry/cnr-server) - open source Python implementation of a server backed by [redis](https://redis.io/), [etcd](https://github.com/coreos/etcd), or a local filesystem.
* [cnr-cli](https://github.com/app-registry/cnr-cli) - open source Python implementation of a client

#### Third-Party Implementations:

TODO(jzelinskie): write this section
