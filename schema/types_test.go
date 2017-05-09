// Copyright 2017 The App Registry Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package schema

import (
	"encoding/json"
	"testing"
)

func TestManifestUnmarshal(t *testing.T) {
	manifestBytes := []byte(`
{
  "schemaVersion": 1,
	"mediaType": "application/vnd.appr.manifest.v0+json",
  "blobs": [
    {
      "mediaType": "application/vnd.appr.blob.helm.chart.v0.tar.gzip",
      "size": 32654,
      "digest": "sha256:e692418e4cbaf90ca69d05a66403747baa33ee08806650b51fab815ad7fc331f"
    }
	]
}
`)
	var m Manifest
	err := json.Unmarshal(manifestBytes, &m)
	if err != nil {
		t.Errorf("failed to unmarshal manifestBytes: %s", err)
	}
}

func TestManifestListUnmarshal(t *testing.T) {
	manifestListBytes := []byte(`
{
  "schemaVersion": 1,
	"mediaType": "application/vnd.appr.manifest.list.v0+json",
  "manifests": [
    {
      "mediaType": "application/vnd.appr.manifest.v0+json",
      "size": 7143,
      "digest": "sha256:e692418e4cbaf90ca69d05a66403747baa33ee08806650b51fab815ad7fc331f",
      "platform": { "name": "helm" }
    },
    {
      "mediaType": "application/vnd.appr.manifest.v0+json",
      "size": 7682,
      "digest": "sha256:5b0bcabd1ed22e9fb1310cf6c2dec7cdef19f0ad69efa1f392e94a4333501270",
      "platform": { "name": "kpm" }
    }
  ]
}
`)

	var m ManifestList
	err := json.Unmarshal(manifestListBytes, &m)
	if err != nil {
		t.Errorf("failed to unmarshal manifestListBytes: %s", err)
	}
}
