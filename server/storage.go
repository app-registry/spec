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
	"context"
	"errors"
	"io"
)

var (
	// ManifestKind is the ResourceKind for a Manifest or ManifestList.
	ManifestKind ResourceKind = "manifest"

	// BlobKind is the ResourceKind for a Blob.
	BlobKind ResourceKind = "blob"
)

// ResourceKind represents the type of the object a ResourceIdentifier
// represents.
type ResourceKind string

// ResourceIdentifier is a triple used to uniquely identify a resource.
type ResourceIdentifier struct {
	Kind       ResourceKind
	Namespace  string
	Repository string
	Reference  string
}

var (
	// ErrResourceNotFound is the error value returned when a resource is not
	// present.
	ErrResourceNotFound = errors.New("resource could not be found")

	// ErrUnsupported is the error value returned from a Storage method that is
	// not implemented by an instance of Storage.
	ErrUnsupported = errors.New("method is not supported")
)

// Storage represents the methods required for the storage and retrieval of
// resources.
type Storage interface {
	// (Stat|Read|Write|Delete)Resource are used to serve resources to App
	// Registry requests.
	//
	// These methods should return ErrResourceNotFound if a particular resource
	// is not within the Storage instance.
	StatResource(ctx context.Context, b ResourceIdentifier) (size uint64, err error)
	ReadResource(ctx context.Context, b ResourceIdentifier) (blobBytes io.ReadCloser, err error)
	WriteResource(ctx context.Context, b ResourceIdentifier, blobBytes io.ReadCloser) error
	DeleteResource(ctx context.Context, b ResourceIdentifier) error

	// Resource(Up|Download)URL are used to redirect the client to another
	// endpoint for (up|down)loads.
	//
	// These methods are optional and should return ErrUnsupported if they are
	// not implemented.
	ResourceDownloadURL(ctx context.Context, b ResourceIdentifier) (url string, err error)
	ResourceUploadURL(ctx context.Context, b ResourceIdentifier) (url string, err error)
}
