// Copyright 2023 sigma
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

package daemon

import (
	"context"
	"sync"

	"github.com/rs/zerolog/log"
	"github.com/tidwall/gjson"

	"github.com/go-sigma/sigma/pkg/dal/dao"
	"github.com/go-sigma/sigma/pkg/dal/models"
	"github.com/go-sigma/sigma/pkg/dal/query"
	"github.com/go-sigma/sigma/pkg/types/enums"
)

// DecoratorArtifactStatus is a status for decorator
type DecoratorArtifactStatus struct {
	Daemon  enums.Daemon
	Status  enums.TaskCommonStatus
	Raw     []byte
	Result  []byte
	Stdout  []byte
	Stderr  []byte
	Message string
}

// DecoratorArtifact is a decorator for daemon task runners
func DecoratorArtifact(runner func(context.Context, *models.Artifact, chan DecoratorArtifactStatus) error) func(context.Context, []byte) error {
	return func(ctx context.Context, payload []byte) error {
		log.Info().Msg("got a task")
		ctx = log.Logger.WithContext(ctx)

		artifactServiceFactory := dao.NewArtifactServiceFactory()
		artifactService := artifactServiceFactory.New()

		id := gjson.GetBytes(payload, "artifact_id").Int()
		artifact, err := artifactService.Get(ctx, id)
		if err != nil {
			return err
		}

		var waitAllEvents = &sync.WaitGroup{}
		waitAllEvents.Add(1)

		var statusChan = make(chan DecoratorArtifactStatus, 1)
		go func() {
			defer waitAllEvents.Done()
			ctx := log.Logger.WithContext(ctx)
			for status := range statusChan {
				switch status.Daemon {
				case enums.DaemonVulnerability:
					err = artifactService.UpdateVulnerability(ctx, id,
						map[string]any{
							query.ArtifactVulnerability.Raw.ColumnName().String():     status.Raw,
							query.ArtifactVulnerability.Result.ColumnName().String():  status.Result,
							query.ArtifactVulnerability.Status.ColumnName().String():  status.Status,
							query.ArtifactVulnerability.Stdout.ColumnName().String():  status.Stdout,
							query.ArtifactVulnerability.Stderr.ColumnName().String():  status.Stderr,
							query.ArtifactVulnerability.Message.ColumnName().String(): status.Message,
						},
					)
				case enums.DaemonSbom:
					err = artifactService.UpdateSbom(ctx,
						id,
						map[string]any{
							query.ArtifactSbom.Raw.ColumnName().String():     status.Raw,
							query.ArtifactSbom.Result.ColumnName().String():  status.Result,
							query.ArtifactSbom.Status.ColumnName().String():  status.Status,
							query.ArtifactSbom.Stdout.ColumnName().String():  status.Stdout,
							query.ArtifactSbom.Stderr.ColumnName().String():  status.Stderr,
							query.ArtifactSbom.Message.ColumnName().String(): status.Message,
						},
					)
				default:
					continue
				}
				if err != nil {
					log.Error().Err(err).Msg("Update artifact status failed")
				}
			}
		}()

		err = runner(ctx, artifact, statusChan)
		if err != nil {
			return err
		}

		waitAllEvents.Wait()

		return nil
	}
}
