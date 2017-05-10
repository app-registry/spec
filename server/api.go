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

package server

import (
	"io"
	"net/http"
	"strconv"
	"strings"

	"github.com/julienschmidt/httprouter"
)

// API represents the required HTTP routes for an App Registry server.
type API interface {
	HeadManifest(w http.ResponseWriter, r *http.Request, p httprouter.Params)
	GetManifest(w http.ResponseWriter, r *http.Request, p httprouter.Params)
	PutManifest(w http.ResponseWriter, r *http.Request, p httprouter.Params)
	DeleteManifest(w http.ResponseWriter, r *http.Request, p httprouter.Params)

	HeadBlob(w http.ResponseWriter, r *http.Request, p httprouter.Params)
	GetBlob(w http.ResponseWriter, r *http.Request, p httprouter.Params)
	DeleteBlob(w http.ResponseWriter, r *http.Request, p httprouter.Params)
	PostBlobUpload(w http.ResponseWriter, r *http.Request, p httprouter.Params)
	PutBlobUpload(w http.ResponseWriter, r *http.Request, p httprouter.Params)
	DeleteBlobUpload(w http.ResponseWriter, r *http.Request, p httprouter.Params)

	GetTags(w http.ResponseWriter, r *http.Request, p httprouter.Params)
}

// This forces the compiler to ensure that SimpleAPI implements the API
// interface.
var _ API = &SimpleAPI{}

// SimpleAPI implements an API using the Storage and Authorization interfaces.
type SimpleAPI struct {
	Storage
	Authorization
}

// writeError reduces the boilerplate of writing the correct HTTP Status Code
// given an error.
func writeError(w http.ResponseWriter, r *http.Request, err error) {
	switch err {
	case ErrResourceNotFound:
		w.WriteHeader(http.StatusNotFound)
	case ErrUnauthorized:
		w.WriteHeader(http.StatusUnauthorized)
	default:
		w.WriteHeader(http.StatusInternalServerError)
	}
}

// authToken parses an Authorization token out of the Authorization HTTP
// header.
func authToken(r *http.Request) string {
	header := r.Header.Get("Authorization")
	return strings.TrimPrefix("Bearer ", header)
}

// redirectOrStreamResource checks the Read permission and attempts to redirect
// the client to the resource.
//
// If the Storage does not implement this functionality, this function
// fallsback to streaming the resource directly to the client.
func (s *SimpleAPI) redirectOrStreamResource(w http.ResponseWriter, r *http.Request, res ResourceIdentifier) {
	ctx := r.Context()
	if err := s.ReadPermission(ctx, authToken(r), res); err != nil {
		writeError(w, r, err)
		return
	}

	url, err := s.ResourceDownloadURL(ctx, res)
	if err == nil {
		http.Redirect(w, r, url, 302)
		return
	} else if err == ErrUnsupported {
		reader, err := s.ReadResource(ctx, res)
		if err != nil {
			writeError(w, r, err)
			return
		}
		defer reader.Close()
		w.WriteHeader(http.StatusOK)
		io.Copy(w, reader)
		return
	}
	writeError(w, r, err)
}

// setContentLengthOfResource checks the Read permission and stats the resource
// in order to set the Content-Length HTTP header to the size of the resource
// in bytes.
func (s *SimpleAPI) setContentLengthOfResource(w http.ResponseWriter, r *http.Request, res ResourceIdentifier) {
	ctx := r.Context()
	if err := s.ReadPermission(ctx, authToken(r), res); err != nil {
		writeError(w, r, err)
		return
	}

	size, err := s.StatResource(ctx, res)
	if err != nil {
		writeError(w, r, err)
		return
	}

	w.Header().Add("Content-Length", strconv.FormatUint(size, 10))
	w.WriteHeader(http.StatusOK)
}

// deleteResource checks the Write permission and deletes the source.
func (s *SimpleAPI) deleteResource(w http.ResponseWriter, r *http.Request, res ResourceIdentifier) {
	ctx := r.Context()
	if err := s.WritePermission(ctx, authToken(r), res); err != nil {
		writeError(w, r, err)
		return
	}

	if err := s.DeleteResource(ctx, res); err != nil {
		writeError(w, r, err)
	}

	w.WriteHeader(http.StatusNoContent)
}

func (s *SimpleAPI) HeadManifest(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	s.setContentLengthOfResource(w, r, ResourceIdentifier{
		ManifestKind,
		p.ByName("namespace"),
		p.ByName("repository"),
		p.ByName("ref"),
	})
}

func (s *SimpleAPI) GetManifest(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	s.redirectOrStreamResource(w, r, ResourceIdentifier{
		ManifestKind,
		p.ByName("namespace"),
		p.ByName("repository"),
		p.ByName("ref"),
	})
}

func (s *SimpleAPI) PutManifest(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
}

func (s *SimpleAPI) DeleteManifest(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	s.deleteResource(w, r, ResourceIdentifier{
		ManifestKind,
		p.ByName("namespace"),
		p.ByName("repository"),
		p.ByName("ref"),
	})
}

func (s *SimpleAPI) HeadBlob(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	s.setContentLengthOfResource(w, r, ResourceIdentifier{
		BlobKind,
		p.ByName("namespace"),
		p.ByName("repository"),
		p.ByName("digest"),
	})
}

func (s *SimpleAPI) GetBlob(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	s.redirectOrStreamResource(w, r, ResourceIdentifier{
		BlobKind,
		p.ByName("namespace"),
		p.ByName("repository"),
		p.ByName("digest"),
	})
}

func (s *SimpleAPI) DeleteBlob(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	s.deleteResource(w, r, ResourceIdentifier{
		BlobKind,
		p.ByName("namespace"),
		p.ByName("repository"),
		p.ByName("digest"),
	})
}

func (s *SimpleAPI) PostBlobUpload(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
}

func (s *SimpleAPI) PutBlobUpload(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
}

func (s *SimpleAPI) DeleteBlobUpload(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
}

func (s *SimpleAPI) GetTags(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
}
