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

// Package client implements an App Registry client.
package client

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"net/url"
	"path"

	"github.com/app-registry/spec/schema"
)

// RegistryEndpoint represents the location where an App Registry API can be
// found.
type RegistryEndpoint struct {
	*url.URL
	Version string
}

// DiscoverRegistryEndpoints uses the default http.Client to request the
// well-known endpoint for finding RegistryEndpoints given a domain.
func DiscoverRegistryEndpoints(discoveryURL string) ([]RegistryEndpoint, error) {
	// TODO(jzelinskie): implement discovery according to the spec
	return nil, nil
}

// NewRegistryEndpoint creates and validates a RegistryEndpoint.
func NewRegistryEndpoint(baseURL, version string) (RegistryEndpoint, error) {
	parsedBaseURL, err := url.Parse(baseURL)
	if err != nil {
		return RegistryEndpoint{}, err
	}

	return RegistryEndpoint{
		URL:     parsedBaseURL,
		Version: version,
	}, nil
}

// RegistrySession represents the state used to engage with a Registry.
type RegistrySession struct {
	Endpoint   RegistryEndpoint
	Credential string
	Client     *http.Client
}

// NewRegistrySession creates the necessary state to begin engaging with a
// Registry.
func NewRegistrySession(credential string, r RegistryEndpoint, c *http.Client) *RegistrySession {
	if c == nil {
		c = &http.Client{}
	}

	return &RegistrySession{
		Endpoint:   r,
		Credential: credential,
		Client:     c,
	}
}

// FetchRelease is an alias for calling Fetch() with all of the parameters
// provided from a Release.
func (rs *RegistrySession) FetchRelease(r schema.Release) ([]io.ReadCloser, error) {
	return rs.Fetch(r.Namespace, r.Repository, r.Version, r.Platform)
}

// Fetch performs the full registry flow that resolves the provided arguments
// ultimately into a slice of io.ReadCloser containing the raw bytes for each
// blob that composes a release of an application.
func (rs *RegistrySession) Fetch(namespace, repository, ref, platform string) ([]io.ReadCloser, error) {
	ml, err := rs.FetchManifestList(namespace, repository, ref)
	if err != nil {
		return nil, err
	}

	mlm, err := ml.FindManifest(platform)
	if err != nil {
		return nil, err
	}

	manifest, err := rs.FetchManifest(namespace, repository, mlm.Digest)
	if err != nil {
		return nil, err
	}

	var blobReaders []io.ReadCloser
	for _, blob := range manifest.Blobs {
		blobReader, err := rs.blobRequest(namespace, repository, blob.Digest)
		if err != nil {
			return nil, err
		}
		blobReaders = append(blobReaders, blobReader)
	}

	return blobReaders, nil
}

func (rs *RegistrySession) blobRequest(namespace, repository string, digest schema.SHA256Digest) (io.ReadCloser, error) {
	blobURL := &url.URL{
		Host:   rs.Endpoint.Host,
		Scheme: rs.Endpoint.Scheme,
		Path:   path.Join(rs.Endpoint.Version, namespace, repository, "blobs", string(digest)),
	}

	req, err := http.NewRequest("GET", blobURL.String(), nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Authorization", "Bearer "+rs.Credential)

	resp, err := rs.Client.Do(req)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		// TODO(jzelinskie): handle different errors for different failures
		return nil, errors.New("non 200 status code")
	}

	return resp.Body, nil
}

func (rs *RegistrySession) manifestRequest(namespace, repository, ref string, manifestOnly bool) (io.ReadCloser, error) {
	manifestURL := &url.URL{
		Host:   rs.Endpoint.Host,
		Scheme: rs.Endpoint.Scheme,
		Path:   path.Join(rs.Endpoint.Version, namespace, repository, "manifests", ref),
	}

	req, err := http.NewRequest("GET", manifestURL.String(), nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Authorization", "Bearer "+rs.Credential)
	if manifestOnly {
		req.Header.Add("Accept", schema.ManifestMimeType)
	} else {
		req.Header.Add("Accept", schema.ManifestListMimeType)
	}

	resp, err := rs.Client.Do(req)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		// TODO(jzelinskie): handle different errors for different failures
		return nil, errors.New("non 200 status code")
	}

	return resp.Body, nil
}

// FetchManifest performs the registry flow to provide a Manifest.
func (rs *RegistrySession) FetchManifest(namespace, repository string, digest schema.SHA256Digest) (schema.Manifest, error) {
	manifestReader, err := rs.manifestRequest(namespace, repository, string(digest), true)
	if err != nil {
		return schema.Manifest{}, err
	}
	defer manifestReader.Close()

	// TODO(jzelinskie): check SchemaVersion before fully decoding.
	var m schema.Manifest
	err = json.NewDecoder(manifestReader).Decode(&m)
	if err != nil {
		return schema.Manifest{}, err
	}

	return m, nil
}

// FetchManifestList performs the registry flow to provide a ManifestList.
func (rs *RegistrySession) FetchManifestList(namespace, repository, ref string) (schema.ManifestList, error) {
	manifestListReader, err := rs.manifestRequest(namespace, repository, ref, false)
	if err != nil {
		return schema.ManifestList{}, err
	}
	defer manifestListReader.Close()

	// TODO(jzelinskie): check SchemaVersion before fully decoding.
	var ml schema.ManifestList
	err = json.NewDecoder(manifestListReader).Decode(&ml)
	if err != nil {
		return schema.ManifestList{}, err
	}

	return ml, nil
}
