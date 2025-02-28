/*
Copyright 2016 Medcl (m AT medcl.net)

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

   http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package main

type ESAPIV5 struct {
	ESAPIV0
}

func NewEsApiV5(host string, auth *Auth, httpProxy string, compress bool, version *ClusterVersion) ESAPI {
	apiV5 := &ESAPIV5{}

	apiV5.Host = host
	apiV5.Auth = auth
	apiV5.HttpProxy = httpProxy
	apiV5.Compress = compress
	apiV5.Version = version

	return apiV5
}
