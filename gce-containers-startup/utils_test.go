// Copyright 2017 Google Inc. All Rights Reserved.
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

package main

import (
	"fmt"
	"testing"

	utils "github.com/konlet/utils"
)

func TestDefaultRegistry_default(t *testing.T) {
	assertRegistryGetsToken(t, "tomcat", false)
	assertRegistryGetsToken(t, "tomcat:1.1", false)
}

func TestDefaultRegistry_dockerIo(t *testing.T) {
	assertRegistryGetsToken(t, "docker.io/tomcat", false)
	assertRegistryGetsToken(t, "index.docker.io/tomcat", false)
	assertRegistryGetsToken(t, "docker.io/tomcat:1.1", false)
}

func TestDefaultRegistry_localRegistry(t *testing.T) {
	assertRegistryGetsToken(t, "localhost.localdomain:5000/ubuntu", false)
}

func TestDefaultRegistry_gcr(t *testing.T) {
	assertRegistryGetsToken(t, "gcr.io/google-containers/nginx", true)
	assertRegistryGetsToken(t, "gcr.io/google-containers/nginx:1.2", true)
	assertRegistryGetsToken(t, "asia.gcr.io/other-containers/busybox", true)
}

func assertRegistryGetsToken(t *testing.T, image string, expectedToken bool) {
	assertEqual(t,
		utils.UseGcpTokenForImage(image),
		expectedToken,
		fmt.Sprintf("registry for %s: unexpected use token: %t", image, !expectedToken))
}
