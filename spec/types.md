## Types

Several schema types shared between different parts of the specification are defined below.

### Media Types

The following media types are used to identify resources referenced by manifests and manifests themselves:

* `application/vnd.cnr.blob.v1.tar+gzip`
* `application/vnd.cnr.manifest.v1+json`
* `application/vnd.cnr.manifest.list.v1+json`

### Manifest Lists

#### Fields

#### Example

```json
{
  "schemaVersion": 1,
  "mediaType": "application/vnd.cnr.manifest.list.v1+json",
  "manifests": [
    {
      "mediaType": "application/vnd.kubernetes.helm.manifest.v1+json",
      "size": 1234,
      "digest": "sha256:a948904f2f0f479b8f8197694b30184b0d2ed1c1cd2a1ec0fb85d299a192a447",
      "platform": {
        "architecture": "amd64",
        "os": "linux"
      }
    },
    {
      "mediaType": "application/vnd.coreos.kpm.manifest.v1+json",
      "size": 1234,
      "digest": "sha256:a948904f2f0f479b8f8197694b30184b0d2ed1c1cd2a1ec0fb85d299a192a447",
      "platform": {
        "architecture": "amd64",
        "os": "linux",
      }
    }
  ]
}
```

### Manifests

#### Fields

#### Example
