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
)

// ErrUnauthorized is the error value returned from an Authorization method
// when the provided token does not access to that resource.
var ErrUnauthorized = errors.New("resource access is unauthorized")

// Authorization represents the methods required for the authorizing of access
// to resources.
type Authorization interface {
	ReadPermission(ctx context.Context, authToken string, r ResourceIdentifier) error
	WritePermission(ctx context.Context, authToken string, r ResourceIdentifier) error
}

// This forces the compiler to ensure that NoopAuth implements the
// Authorization interface.
var _ Authorization = &NoopAuth{}

// NoopAuth blindly authorizes any request to access a resource.
type NoopAuth struct{}

// ReadPermission blindly authorizes any request to read a resource.
func (n NoopAuth) ReadPermission(_ context.Context, _ string, _ ResourceIdentifier) error {
	return nil
}

// WritePermission blindly authorizes any request to write a resource.
func (n NoopAuth) WritePermission(_ context.Context, _ string, _ ResourceIdentifier) error {
	return nil
}
