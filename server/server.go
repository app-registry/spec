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
	"crypto/tls"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// Server wraps an instance of an API and an *http.Server to server App
// Registry requests.
type Server struct {
	API
	*http.Server
}

// NewServer creates a new instance of a Server configured to serve requests
// using the provided API.
func NewServer(addr string, tlsCfg *tls.Config, api API) *Server {
	return &Server{
		Server: &http.Server{
			Addr:      addr,
			TLSConfig: tlsCfg,
			Handler:   handler(api),
		},
		API: api,
	}
}

func handler(api API) http.Handler {
	router := httprouter.New()
	router.HEAD("/v0/:namespace/:repository/manifests/:ref", api.HeadManifest)
	router.GET("/v0/:namespace/:repository/manifests/:ref", api.GetManifest)

	router.HEAD("/v0/:namespace/:repository/blobs/:digest", api.HeadBlob)
	router.GET("/v0/:namespace/:repository/blobs/:digest", api.GetBlob)
	router.POST("/v0/:namespace/:repository/blobs/uploads", api.PostBlobUpload)
	router.PUT("/v0/:namespace/:repository/blobs/uploads/:uuid", api.PutBlobUpload)

	router.GET("/v0/:namespace/:repository/tags/list", api.GetTags)
	return router
}
