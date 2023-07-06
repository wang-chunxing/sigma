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

package types

// RepositoryItem represents a repository.
type RepositoryItem struct {
	ID   int64  `json:"id" example:"1"`
	Name string `json:"name" example:"busybox"`

	ArtifactCount int64 `json:"artifact_count"`

	CreatedAt string `json:"created_at" example:"2006-01-02 15:04:05"`
	UpdatedAt string `json:"updated_at" example:"2006-01-02 15:04:05"`
}

// ListRepositoryRequest represents the request to list repositories.
type ListRepositoryRequest struct {
	Pagination

	Namespace string `json:"namespace" param:"namespace" validate:"required,min=2,max=20,is_valid_namespace" example:"test"`
}

// GetRepositoryRequest represents the request to get a repository.
type GetRepositoryRequest struct {
	ID int64 `json:"name" param:"id" validate:"required,number" example:"1"`
}

// DeleteRepositoryRequest represents the request to delete a repository.
type DeleteRepositoryRequest struct {
	Namespace string `json:"namespace" param:"namespace" validate:"required,min=2,max=20,is_valid_namespace" example:"test"`
	ID        int64  `json:"id" param:"id" validate:"required,number" example:"1"`
}

// PostRepositoryRequest represents the request to create a repository.
type PostRepositoryRequest struct {
	Namespace string `json:"namespace" param:"namespace" validate:"required,min=2,max=20,is_valid_namespace" example:"test"`
	Name      string `json:"name" validate:"required" example:"test"`
}

// PostRepositoryRequestSwagger represents the request to create a repository.
type PostRepositoryRequestSwagger struct {
	Name string `json:"name" validate:"required" example:"test"`
}

// PostRepositoryResponse represents the response to create a repository.
type PostRepositoryResponse struct {
	ID int64 `json:"id" example:"21911"`
}

// PutRepositoryRequest represents the request to update a repository.
type PutRepositoryRequest struct {
	Namespace   string  `json:"namespace" param:"namespace" validate:"required,min=2,max=20,is_valid_namespace" example:"test"`
	ID          int64   `json:"id" param:"id" validate:"required,number" example:"1"`
	Description *string `json:"description,omitempty" validate:"omitempty,max=300" example:"i am just description"`
	Overview    *string `json:"overview,omitempty" validate:"omitempty,max=3000" example:"i am just overview"`
}

// PutRepositoryRequestSwagger represents the request to update a repository.
type PutRepositoryRequestSwagger struct {
	Description *string `json:"description,omitempty" validate:"omitempty,max=300" example:"i am just description"`
	Overview    *string `json:"overview,omitempty" validate:"omitempty,max=3000" example:"i am just overview"`
}

// PutRepositoryResponse represents the response to update a repository.
type PutRepositoryResponse struct {
	ID int64 `json:"id" example:"21911"`
}
