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

// Package schema provides definitions for the JSON schema of the different
// manifests in the App Registry specification.
package schema

import "errors"

const (
	// ManifestListMimeType is the MIME type that should be used for content type
	// negotiation for a ManifestList.
	ManifestListMimeType = "application/vnd.appr.manifest.list.v0+json"

	// ManifestMimeType is the MIME type that should be used for content type
	// negotiation for a Manifest.
	ManifestMimeType = "application/vnd.appr.manifest.v0+json"
)

// Release represents an immutable state of application that has been assigned
// a version.
type Release struct {
	Namespace  string `json:"namespace"`
	Repository string `json:"repository"`
	Platform   string `json:"platform"`
	Version    string `json:"version"`
}

// Channel represents the head of a stream of Releases.
type Channel struct {
	Name string `json:"name"`
	Release
}

// SHA256Digest is a SHA256 digest of an application artifact.
type SHA256Digest string

// Blob represents the metadata of an application artifact.
type Blob struct {
	MediaType string       `json:"mediaType"` // The MIME type of the referenced object.
	Size      uint64       `json:"size"`      // The size in bytes of the object.
	Digest    SHA256Digest `json:"digest"`    // The SHA256 hash of the object.
	URLs      []string     `json:"urls"`      // The list of URLs from which the content may be fetched.
}

// ManifestList represents a list Manifests for the same object differentiated
// by Platform.
type ManifestList struct {
	SchemaVersion int                    `json:"schemaVersion"` // The version of schema being used for this object.
	MediaType     string                 `json:"mediaType"`     // The MIME type of the referenced object.
	Manifests     []ManifestListManifest `json:"manifests"`     // The list of Manifests for this reference.
}

// FindManifest attempts to find the ManifestListManifest in a ManifestList for
// the provided platform.
func (ml ManifestList) FindManifest(platform string) (ManifestListManifest, error) {
	for _, mlm := range ml.Manifests {
		if mlm.Platform.Name == platform {
			return mlm, nil
		}
	}
	return ManifestListManifest{}, errors.New("no matching Manifest in ManifestList")
}

// ManifestListManifest is the representation of a Manifest embedded in a
// ManifestList.
type ManifestListManifest struct {
	SchemaVersion int          `json:"schemaVersion"` // The version of the schema being used for this object.
	MediaType     string       `json:"mediaType"`     // The MIME type of the refernced object.
	Size          uint64       `json:"size"`          // The size in bytes of the object.
	Digest        SHA256Digest `json:"digest"`        // The SHA256 hash of the object.
	Platform      Platform     `json:"platform"`      // The platform for which this object belongs.
}

// Platform represents the unique metadata that differentiates between versions
// of ManifestListManifests.
type Platform struct {
	Name string `json:"name"` // The name of the platform.
	// TODO(jzelinskie): maybe include fields like OS/arch
}

// Manifest represents the list of blobs that compose a particular release of
// an application.
type Manifest struct {
	SchemaVersion int    `json:"schemaVersion"`
	MediaType     string `json:"mediaType"`
	Blobs         []Blob `json:"blobs"`
}
