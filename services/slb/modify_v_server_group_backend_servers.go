package slb

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

// invoke ModifyVServerGroupBackendServers api with *ModifyVServerGroupBackendServersRequest synchronously
// api document: https://help.aliyun.com/api/slb/modifyvservergroupbackendservers.html
func (client *Client) ModifyVServerGroupBackendServers(request *ModifyVServerGroupBackendServersRequest) (response *ModifyVServerGroupBackendServersResponse, err error) {
	response = CreateModifyVServerGroupBackendServersResponse()
	err = client.DoAction(request, response)
	return
}

// invoke ModifyVServerGroupBackendServers api with *ModifyVServerGroupBackendServersRequest asynchronously
// api document: https://help.aliyun.com/api/slb/modifyvservergroupbackendservers.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ModifyVServerGroupBackendServersWithChan(request *ModifyVServerGroupBackendServersRequest) (<-chan *ModifyVServerGroupBackendServersResponse, <-chan error) {
	responseChan := make(chan *ModifyVServerGroupBackendServersResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ModifyVServerGroupBackendServers(request)
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

// invoke ModifyVServerGroupBackendServers api with *ModifyVServerGroupBackendServersRequest asynchronously
// api document: https://help.aliyun.com/api/slb/modifyvservergroupbackendservers.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ModifyVServerGroupBackendServersWithCallback(request *ModifyVServerGroupBackendServersRequest, callback func(response *ModifyVServerGroupBackendServersResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ModifyVServerGroupBackendServersResponse
		var err error
		defer close(result)
		response, err = client.ModifyVServerGroupBackendServers(request)
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

type ModifyVServerGroupBackendServersRequest struct {
	*requests.RpcRequest
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	AccessKeyId          string           `position:"Query" name:"access_key_id"`
	Tags                 string           `position:"Query" name:"Tags"`
	VServerGroupId       string           `position:"Query" name:"VServerGroupId"`
	OldBackendServers    string           `position:"Query" name:"OldBackendServers"`
	NewBackendServers    string           `position:"Query" name:"NewBackendServers"`
}

type ModifyVServerGroupBackendServersResponse struct {
	*responses.BaseResponse
	RequestId      string                                           `json:"RequestId" xml:"RequestId"`
	VServerGroupId string                                           `json:"VServerGroupId" xml:"VServerGroupId"`
	BackendServers BackendServersInModifyVServerGroupBackendServers `json:"BackendServers" xml:"BackendServers"`
}

// create a request to invoke ModifyVServerGroupBackendServers API
func CreateModifyVServerGroupBackendServersRequest() (request *ModifyVServerGroupBackendServersRequest) {
	request = &ModifyVServerGroupBackendServersRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Slb", "2014-05-15", "ModifyVServerGroupBackendServers", "slb", "openAPI")
	return
}

// create a response to parse from ModifyVServerGroupBackendServers response
func CreateModifyVServerGroupBackendServersResponse() (response *ModifyVServerGroupBackendServersResponse) {
	response = &ModifyVServerGroupBackendServersResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
