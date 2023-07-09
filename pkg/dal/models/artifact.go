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

package models

import (
	"database/sql"
	"errors"
	"time"

	"gorm.io/gorm"
	"gorm.io/plugin/soft_delete"

	"github.com/ximager/ximager/pkg/types/enums"
)

// Artifact represents an artifact
type Artifact struct {
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt soft_delete.DeletedAt `gorm:"softDelete:milli"`
	ID        int64                 `gorm:"primaryKey"`

	RepositoryID int64
	Digest       string
	Size         int64 `gorm:"default:0"`
	BlobsSize    int64 `gorm:"default:0"`
	ContentType  string
	Raw          []byte

	LastPull  sql.NullTime
	PushedAt  time.Time `gorm:"autoCreateTime"`
	PullTimes int64     `gorm:"default:0"`

	Repository Repository

	ArtifactIndexes []*Artifact `gorm:"many2many:artifact_artifacts;"`
	Blobs           []*Blob     `gorm:"many2many:artifact_blobs;"`
	Tags            []*Tag      `gorm:"foreignKey:ArtifactID;"`
}

// AfterCreate ...
// if something err occurs, the create will be aborted
func (a *Artifact) BeforeCreate(tx *gorm.DB) error {
	if a == nil {
		return nil
	}
	var repositoryObj Repository
	err := tx.Model(&Repository{}).Where(&Repository{ID: a.RepositoryID}).First(&repositoryObj).Error
	if err != nil {
		return err
	}
	var namespaceObj Namespace
	err = tx.Model(&Namespace{}).Where(&Namespace{ID: repositoryObj.NamespaceID}).First(&namespaceObj).Error
	if err != nil {
		return err
	}
	if namespaceObj.SizeLimit > 0 && namespaceObj.Size+a.BlobsSize > namespaceObj.SizeLimit {
		return errors.New("namespace's size quota exceeded")
	}
	if repositoryObj.SizeLimit > 0 && repositoryObj.Size+a.BlobsSize > repositoryObj.SizeLimit {
		return errors.New("repository's size quota exceeded")
	}

	// we should check all the checker here, and update the size and tag count
	err = tx.Model(&Namespace{}).Where(&Namespace{ID: repositoryObj.NamespaceID}).UpdateColumns(
		map[string]any{
			"size": namespaceObj.Size + a.BlobsSize,
		}).Error
	if err != nil {
		return err
	}
	err = tx.Model(&Repository{}).Where(&Repository{ID: repositoryObj.ID}).UpdateColumns(map[string]any{
		"size": repositoryObj.Size + a.BlobsSize,
	}).Error
	if err != nil {
		return err
	}
	return nil
}

// BeforeDelete ...
// if something err occurs, the delete will be aborted
func (a *Artifact) BeforeUpdate(tx *gorm.DB) error {
	if a == nil {
		return nil
	}
	var repositoryObj Repository
	err := tx.Model(&Repository{}).Where("id = ?", a.RepositoryID).First(&repositoryObj).Error
	if err != nil {
		return err
	}

	err = tx.Exec(`UPDATE
  namespaces
SET
  size = (
    SELECT
      SUM(artifacts.blobs_size)
    FROM
      repositories
      INNER JOIN artifacts ON repositories.id = artifacts.repository_id
    WHERE
      repositories.namespace_id = ?)
WHERE
  id = ?`, repositoryObj.NamespaceID, repositoryObj.NamespaceID).Error
	if err != nil {
		return err
	}
	err = tx.Exec(`UPDATE
  repositories
SET
  size = (
    SELECT
      SUM(size)
    FROM
      artifacts
    WHERE
		  artifacts.repository_id = ?)
WHERE
  id = ?`, repositoryObj.ID, repositoryObj.ID).Error
	if err != nil {
		return err
	}
	return nil
}

// ArtifactSbom represents an artifact sbom
type ArtifactSbom struct {
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt soft_delete.DeletedAt `gorm:"softDelete:milli"`
	ID        int64                 `gorm:"primaryKey"`

	ArtifactID int64
	Raw        []byte
	Status     enums.TaskCommonStatus
	Stdout     []byte
	Stderr     []byte
	Message    string

	Artifact *Artifact
}

// ArtifactVulnerability represents an artifact vulnerability
type ArtifactVulnerability struct {
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt soft_delete.DeletedAt `gorm:"softDelete:milli"`
	ID        int64                 `gorm:"primaryKey"`

	ArtifactID int64
	Metadata   []byte // is the trivy db metadata
	Raw        []byte
	Status     enums.TaskCommonStatus
	Stdout     []byte
	Stderr     []byte
	Message    string

	Artifact *Artifact
}
