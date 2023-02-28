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

package redis

import (
	"context"
	"testing"
	"time"

	"github.com/alicebob/miniredis/v2"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"

	"github.com/ximager/ximager/pkg/utils"
	"github.com/ximager/ximager/pkg/utils/leader"
)

func TestNew(t *testing.T) {
	ctx, ctxCancel := context.WithCancel(context.Background())

	utils.SetLevel(0)

	miniRedis := miniredis.RunT(t)
	viper.SetDefault("redis.url", "redis://"+miniRedis.Addr())

	var f = factory{}
	_, err := f.New(ctx, leader.Options{
		Name:          "leader",
		LeaseDuration: time.Second * 15,
		RenewDeadline: time.Second * 3,
		RetryPeriod:   time.Second * 2,
	})
	assert.NoError(t, err)

	time.Sleep(time.Second * 3)

	ctxCancel()

	time.Sleep(time.Second * 5)
}

func TestLeaderChange(t *testing.T) {
	utils.SetLevel(0)

	miniRedis := miniredis.RunT(t)
	viper.SetDefault("redis.url", "redis://"+miniRedis.Addr())

	ctx1, ctxCancel1 := context.WithCancel(context.Background())
	var f = factory{}
	_, err := f.New(ctx1, leader.Options{
		Name:          "leader1",
		LeaseDuration: time.Second * 15,
		RenewDeadline: time.Second * 3,
		RetryPeriod:   time.Second * 2,
	})
	assert.NoError(t, err)

	time.Sleep(time.Second * 1)

	var f1 = factory{}
	leader1, err := f1.New(context.Background(), leader.Options{
		Name:          "leader2",
		LeaseDuration: time.Second * 15,
		RenewDeadline: time.Second * 3,
		RetryPeriod:   time.Second * 2,
	})
	assert.NoError(t, err)

	time.Sleep(time.Second * 3)

	ctxCancel1()

	time.Sleep(time.Second * 3)

	assert.True(t, leader1.IsLeader())

	time.Sleep(time.Second * 5)
}
