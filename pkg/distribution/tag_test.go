// Copyright 2023 XImager
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

package distribution

import (
	"context"
	"fmt"
	"net/http"
	"testing"

	"github.com/distribution/distribution/v3/reference"
	"github.com/stretchr/testify/assert"
)

func TestAllTags(t *testing.T) {
	var tran = NewTransport(func(req *http.Request) {
		req.SetBasicAuth("tosone", "8541539655")
	})

	named, err := reference.WithName("library/busybox")
	assert.NoError(t, err)

	repository, err := NewRepository(named, "https://hub.tosone.cn", tran)
	assert.NoError(t, err)
	tags, err := repository.Tags().All(context.Background())
	assert.NoError(t, err)
	fmt.Println(tags)
}
