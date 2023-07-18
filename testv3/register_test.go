// Copyright © 2023 OpenIM SDK. All rights reserved.
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

// @Author BanTanger 2023/7/9 15:30:00
package testv3

import (
	"open_im_sdk/pkg/utils"
	"open_im_sdk/testv3/funcation"
	"strconv"
	"testing"
	"time"
)

func Test_RegisterOne(t *testing.T) {
	uid := "123456"
	nickname := "123456"
	faceUrl := ""
	register, err := funcation.RegisterOne(uid, nickname, faceUrl)
	if err != nil {
		t.Fatal(err)
	}
	if register != true {
		t.Errorf("uid [%v] register expected be successful, but fail got", uid)
	}
	t.Log(register)
}

func Test_RegisterBatch(t *testing.T) {
	count := 100
	var users []funcation.Users
	for i := 0; i < count; i++ {
		users = append(users, funcation.Users{
			Uid:      funcation.GenUid(i, "register_test_"+utils.Int64ToString(time.Now().Unix())),
			Nickname: "register_test_" + strconv.FormatInt(int64(i), 10),
			FaceUrl:  "",
		})
	}
	success, fail := funcation.RegisterBatch(users)
	t.Log(success)
	t.Log(fail)
}
