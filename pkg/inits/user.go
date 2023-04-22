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

package inits

import (
	"context"
	"fmt"

	"github.com/spf13/viper"

	"github.com/ximager/ximager/pkg/dal/dao"
	"github.com/ximager/ximager/pkg/dal/models"
	"github.com/ximager/ximager/pkg/utils/password"
)

func init() {
	inits["user"] = initUser
}

func initUser() error {
	passwordService := password.New()
	userService := dao.NewUserService()
	userCount, err := userService.Count(context.Background())
	if err != nil {
		return err
	}
	if userCount > 0 {
		return nil
	}
	internalUserPassword := viper.GetString("auth.internalUser.password")
	if internalUserPassword == "" {
		return fmt.Errorf("the internal user password is not set")
	}
	internalUserUsername := viper.GetString("auth.internalUser.username")
	if internalUserUsername == "" {
		return fmt.Errorf("the internal user username is not set")
	}
	internalUserPasswordHashed, err := passwordService.Hash(internalUserPassword)
	if err != nil {
		return err
	}
	internalUser := &models.User{
		Username: internalUserUsername,
		Password: internalUserPasswordHashed,
		Email:    "internal-fake@gmail.com",
		Role:     "root", // TODO: change to read-only role
	}
	err = userService.Create(context.Background(), internalUser)
	if err != nil {
		return err
	}

	adminUserPassword := viper.GetString("auth.admin.password")
	if adminUserPassword == "" {
		return fmt.Errorf("the admin user password is not set")
	}
	adminUserUsername := viper.GetString("auth.admin.username")
	if adminUserUsername == "" {
		return fmt.Errorf("the admin user username is not set")
	}
	adminUserPasswordHashed, err := passwordService.Hash(adminUserPassword)
	if err != nil {
		return err
	}
	adminUserEmail := viper.GetString("auth.admin.email")
	if adminUserEmail == "" {
		adminUserEmail = "fake@gmail.com"
	}
	adminUser := &models.User{
		Username: adminUserUsername,
		Password: adminUserPasswordHashed,
		Email:    adminUserEmail,
		Role:     "root",
	}
	err = userService.Create(context.Background(), adminUser)
	if err != nil {
		return err
	}

	return nil
}