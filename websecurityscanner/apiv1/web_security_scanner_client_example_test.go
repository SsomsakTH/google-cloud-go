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

package websecurityscanner_test

import (
	"context"

	websecurityscanner "cloud.google.com/go/websecurityscanner/apiv1"
	"google.golang.org/api/iterator"
	websecurityscannerpb "google.golang.org/genproto/googleapis/cloud/websecurityscanner/v1"
)

func ExampleNewClient() {
	ctx := context.Background()
	c, err := websecurityscanner.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	// TODO: Use client.
	_ = c
}

func ExampleClient_CreateScanConfig() {
	ctx := context.Background()
	c, err := websecurityscanner.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &websecurityscannerpb.CreateScanConfigRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/google.golang.org/genproto/googleapis/cloud/websecurityscanner/v1#CreateScanConfigRequest.
	}
	resp, err := c.CreateScanConfig(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleClient_DeleteScanConfig() {
	ctx := context.Background()
	c, err := websecurityscanner.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &websecurityscannerpb.DeleteScanConfigRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/google.golang.org/genproto/googleapis/cloud/websecurityscanner/v1#DeleteScanConfigRequest.
	}
	err = c.DeleteScanConfig(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
}

func ExampleClient_GetScanConfig() {
	ctx := context.Background()
	c, err := websecurityscanner.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &websecurityscannerpb.GetScanConfigRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/google.golang.org/genproto/googleapis/cloud/websecurityscanner/v1#GetScanConfigRequest.
	}
	resp, err := c.GetScanConfig(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleClient_ListScanConfigs() {
	ctx := context.Background()
	c, err := websecurityscanner.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &websecurityscannerpb.ListScanConfigsRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/google.golang.org/genproto/googleapis/cloud/websecurityscanner/v1#ListScanConfigsRequest.
	}
	it := c.ListScanConfigs(ctx, req)
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

func ExampleClient_UpdateScanConfig() {
	ctx := context.Background()
	c, err := websecurityscanner.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &websecurityscannerpb.UpdateScanConfigRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/google.golang.org/genproto/googleapis/cloud/websecurityscanner/v1#UpdateScanConfigRequest.
	}
	resp, err := c.UpdateScanConfig(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleClient_StartScanRun() {
	ctx := context.Background()
	c, err := websecurityscanner.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &websecurityscannerpb.StartScanRunRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/google.golang.org/genproto/googleapis/cloud/websecurityscanner/v1#StartScanRunRequest.
	}
	resp, err := c.StartScanRun(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleClient_GetScanRun() {
	ctx := context.Background()
	c, err := websecurityscanner.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &websecurityscannerpb.GetScanRunRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/google.golang.org/genproto/googleapis/cloud/websecurityscanner/v1#GetScanRunRequest.
	}
	resp, err := c.GetScanRun(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleClient_ListScanRuns() {
	ctx := context.Background()
	c, err := websecurityscanner.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &websecurityscannerpb.ListScanRunsRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/google.golang.org/genproto/googleapis/cloud/websecurityscanner/v1#ListScanRunsRequest.
	}
	it := c.ListScanRuns(ctx, req)
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

func ExampleClient_StopScanRun() {
	ctx := context.Background()
	c, err := websecurityscanner.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &websecurityscannerpb.StopScanRunRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/google.golang.org/genproto/googleapis/cloud/websecurityscanner/v1#StopScanRunRequest.
	}
	resp, err := c.StopScanRun(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleClient_ListCrawledUrls() {
	ctx := context.Background()
	c, err := websecurityscanner.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &websecurityscannerpb.ListCrawledUrlsRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/google.golang.org/genproto/googleapis/cloud/websecurityscanner/v1#ListCrawledUrlsRequest.
	}
	it := c.ListCrawledUrls(ctx, req)
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

func ExampleClient_GetFinding() {
	ctx := context.Background()
	c, err := websecurityscanner.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &websecurityscannerpb.GetFindingRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/google.golang.org/genproto/googleapis/cloud/websecurityscanner/v1#GetFindingRequest.
	}
	resp, err := c.GetFinding(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleClient_ListFindings() {
	ctx := context.Background()
	c, err := websecurityscanner.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &websecurityscannerpb.ListFindingsRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/google.golang.org/genproto/googleapis/cloud/websecurityscanner/v1#ListFindingsRequest.
	}
	it := c.ListFindings(ctx, req)
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

func ExampleClient_ListFindingTypeStats() {
	ctx := context.Background()
	c, err := websecurityscanner.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &websecurityscannerpb.ListFindingTypeStatsRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/google.golang.org/genproto/googleapis/cloud/websecurityscanner/v1#ListFindingTypeStatsRequest.
	}
	resp, err := c.ListFindingTypeStats(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}