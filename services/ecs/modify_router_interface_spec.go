package ecs

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

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

// invoke ModifyRouterInterfaceSpec api with *ModifyRouterInterfaceSpecRequest synchronously
// api document: https://help.aliyun.com/api/ecs/modifyrouterinterfacespec.html
func (client *Client) ModifyRouterInterfaceSpec(request *ModifyRouterInterfaceSpecRequest) (response *ModifyRouterInterfaceSpecResponse, err error) {
	response = CreateModifyRouterInterfaceSpecResponse()
	err = client.DoAction(request, response)
	return
}

// invoke ModifyRouterInterfaceSpec api with *ModifyRouterInterfaceSpecRequest asynchronously
// api document: https://help.aliyun.com/api/ecs/modifyrouterinterfacespec.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ModifyRouterInterfaceSpecWithChan(request *ModifyRouterInterfaceSpecRequest) (<-chan *ModifyRouterInterfaceSpecResponse, <-chan error) {
	responseChan := make(chan *ModifyRouterInterfaceSpecResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ModifyRouterInterfaceSpec(request)
		if err != nil {
			errChan <- err
		} else {
			responseChan <- response
		}
	})
	if err != nil {
		errChan <- err
		close(responseChan)
		close(errChan)
	}
	return responseChan, errChan
}

// invoke ModifyRouterInterfaceSpec api with *ModifyRouterInterfaceSpecRequest asynchronously
// api document: https://help.aliyun.com/api/ecs/modifyrouterinterfacespec.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ModifyRouterInterfaceSpecWithCallback(request *ModifyRouterInterfaceSpecRequest, callback func(response *ModifyRouterInterfaceSpecResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ModifyRouterInterfaceSpecResponse
		var err error
		defer close(result)
		response, err = client.ModifyRouterInterfaceSpec(request)
		callback(response, err)
		result <- 1
	})
	if err != nil {
		defer close(result)
		callback(nil, err)
		result <- 0
	}
	return result
}

type ModifyRouterInterfaceSpecRequest struct {
	*requests.RpcRequest
	RouterInterfaceId    string           `position:"Query" name:"RouterInterfaceId"`
	Spec                 string           `position:"Query" name:"Spec"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	ClientToken          string           `position:"Query" name:"ClientToken"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	UserCidr             string           `position:"Query" name:"UserCidr"`
}

type ModifyRouterInterfaceSpecResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Spec      string `json:"Spec" xml:"Spec"`
}

// create a request to invoke ModifyRouterInterfaceSpec API
func CreateModifyRouterInterfaceSpecRequest() (request *ModifyRouterInterfaceSpecRequest) {
	request = &ModifyRouterInterfaceSpecRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Ecs", "2014-05-26", "ModifyRouterInterfaceSpec", "ecs", "openAPI")
	return
}

// create a response to parse from ModifyRouterInterfaceSpec response
func CreateModifyRouterInterfaceSpecResponse() (response *ModifyRouterInterfaceSpecResponse) {
	response = &ModifyRouterInterfaceSpecResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
