// Copyright 2014 Google Inc. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// +build integration

package logging_test

import (
	"testing"

	"github.com/golang/glog"

	"google.golang.org/cloud/internal/testutil"
	"google.golang.org/cloud/logging"
)

func TestAll(t *testing.T) {
	var (
		projID = testutil.ProjID()
		ctx    = testutil.Context()

		c   *logging.Client
		err error
	)

	if c, err = logging.NewClient(ctx, projID, "logging-integration-test"); err != nil {
		t.Fatalf("error creating client: %v", err)
	}

	if err := c.Ping(); err != nil {
		glog.Fatalf("error pinging logging api: %v", err)
	}

	payload := struct{ Message string }{Message: "test log"}

	if err := c.LogSync(logging.Entry{Payload: payload}); err != nil {
		t.Fatalf("error writing log: %v", err)
	}

	if err := c.LogSync(logging.Entry{Payload: payload}); err != nil {
		t.Fatalf("error writing log: %v", err)
	}

	if _, err := c.Writer(logging.Debug).Write([]byte("test log\n")); err != nil {
		t.Fatalf("error writing log using io.Writer: %v", err)
	}

	c.Logger(logging.Debug).Println("test log")

	if err := c.Flush(); err != nil {
		t.Fatalf("error flushing logs: %v", err)
	}
}
