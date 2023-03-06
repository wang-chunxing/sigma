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

package tests

import (
	"fmt"
	"os"

	"github.com/google/uuid"
	"github.com/spf13/viper"

	"github.com/ximager/ximager/pkg/dal"
)

func init() {
	err := RegisterCIDatabaseFactory("sqlite3", &sqlite3Factory{})
	if err != nil {
		panic(err)
	}
}

type sqlite3Factory struct{}

var _ Factory = &sqlite3Factory{}

func (sqlite3Factory) New() CIDatabase {
	return &sqlite3CIDatabase{}
}

type sqlite3CIDatabase struct {
	path string
}

var _ CIDatabase = &sqlite3CIDatabase{}

// Init sets the default values for the database configuration in ci tests
func (d *sqlite3CIDatabase) Init() error {
	d.path = fmt.Sprintf("%s.db", uuid.New().String())
	viper.SetDefault("database.type", "sqlite")
	viper.SetDefault("database.sqlite.path", d.path)
	err := dal.Initialize()
	if err != nil {
		return err
	}
	return nil
}

// DeInit remove the database or database file for ci tests
func (d *sqlite3CIDatabase) DeInit() error {
	err := os.Remove(d.path)
	if err != nil {
		return err
	}
	return nil
}