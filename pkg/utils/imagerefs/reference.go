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

// Package imagerefs ...
package imagerefs

import (
	"fmt"
	"strings"

	"github.com/distribution/distribution/v3/reference"
)

// Parse ...
func Parse(name string) (string, string, string, string, error) {
	if !strings.Contains(name, "/") {
		return "", "", "", "", fmt.Errorf("invalid reference: %s", name)
	}

	named, err := reference.ParseNormalizedNamed(name)
	if err != nil {
		return "", "", "", "", fmt.Errorf("failed to parse reference: %v, %s", err, name)
	}
	named = reference.TagNameOnly(named)
	domain := reference.Domain(named)
	path := reference.Path(named)
	tagged, ok := named.(reference.Tagged)
	if !ok {
		return "", "", "", "", fmt.Errorf("reference is not tagged: %v, %s", named, name)
	}
	tag := tagged.Tag()
	if !strings.Contains(path, "/") {
		return "", "", "", "", fmt.Errorf("invalid reference: %s", name)
	}
	parts := strings.Split(path, "/")
	ns := parts[0]
	repo := path
	return domain, ns, repo, tag, nil
}
