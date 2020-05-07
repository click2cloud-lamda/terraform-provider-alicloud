package waf_openapi

//Licensed under the Apache License, Version 2.0 (the "License");
//you may not use this file except in compliance with the License.
//You may obtain a copy of the License at
//
//http://www.apache.org/licenses/LICENSE-2.0
//
//Unless required by applicable law or agreed to in writing, software
//distributed under the License is distributed on an "AS IS" BASIS,
//WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//See the License for the specific language governing permissions and
//limitations under the License.
//
// Code generated by Alibaba Cloud SDK Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

// Rule is a nested struct in waf_openapi response
type Rule struct {
	Name    string `json:"Name" xml:"Name"`
	Version int64  `json:"Version" xml:"Version"`
	Time    int64  `json:"Time" xml:"Time"`
	Content string `json:"Content" xml:"Content"`
	Status  int64  `json:"Status" xml:"Status"`
	Scene   string `json:"Scene" xml:"Scene"`
	RuleId  int64  `json:"RuleId" xml:"RuleId"`
	Enabled int    `json:"Enabled" xml:"Enabled"`
	Origin  string `json:"Origin" xml:"Origin"`
}
