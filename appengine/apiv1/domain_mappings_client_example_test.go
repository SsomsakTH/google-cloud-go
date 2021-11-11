// Copyright 2021 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go_gapic. DO NOT EDIT.

package appengine_test

import (
	"context"

	appengine "cloud.google.com/go/appengine/apiv1"
	"google.golang.org/api/iterator"
	appenginepb "google.golang.org/genproto/googleapis/appengine/v1"
)

func ExampleNewDomainMappingsClient() {
	ctx := context.Background()
	c, err := appengine.NewDomainMappingsClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	// TODO: Use client.
	_ = c
}

func ExampleDomainMappingsClient_ListDomainMappings() {
	ctx := context.Background()
	c, err := appengine.NewDomainMappingsClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &appenginepb.ListDomainMappingsRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/google.golang.org/genproto/googleapis/appengine/v1#ListDomainMappingsRequest.
	}
	it := c.ListDomainMappings(ctx, req)
	for {
		resp, err := it.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			// TODO: Handle error.
		}
		// TODO: Use resp.
		_ = resp
	}
}

func ExampleDomainMappingsClient_GetDomainMapping() {
	ctx := context.Background()
	c, err := appengine.NewDomainMappingsClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &appenginepb.GetDomainMappingRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/google.golang.org/genproto/googleapis/appengine/v1#GetDomainMappingRequest.
	}
	resp, err := c.GetDomainMapping(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleDomainMappingsClient_CreateDomainMapping() {
	ctx := context.Background()
	c, err := appengine.NewDomainMappingsClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &appenginepb.CreateDomainMappingRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/google.golang.org/genproto/googleapis/appengine/v1#CreateDomainMappingRequest.
	}
	op, err := c.CreateDomainMapping(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}

	resp, err := op.Wait(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleDomainMappingsClient_UpdateDomainMapping() {
	ctx := context.Background()
	c, err := appengine.NewDomainMappingsClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &appenginepb.UpdateDomainMappingRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/google.golang.org/genproto/googleapis/appengine/v1#UpdateDomainMappingRequest.
	}
	op, err := c.UpdateDomainMapping(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}

	resp, err := op.Wait(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleDomainMappingsClient_DeleteDomainMapping() {
	ctx := context.Background()
	c, err := appengine.NewDomainMappingsClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &appenginepb.DeleteDomainMappingRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/google.golang.org/genproto/googleapis/appengine/v1#DeleteDomainMappingRequest.
	}
	op, err := c.DeleteDomainMapping(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}

	err = op.Wait(ctx)
	if err != nil {
		// TODO: Handle error.
	}
}