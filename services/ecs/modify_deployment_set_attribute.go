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

// invoke ModifyDeploymentSetAttribute api with *ModifyDeploymentSetAttributeRequest synchronously
// api document: https://help.aliyun.com/api/ecs/modifydeploymentsetattribute.html
func (client *Client) ModifyDeploymentSetAttribute(request *ModifyDeploymentSetAttributeRequest) (response *ModifyDeploymentSetAttributeResponse, err error) {
	response = CreateModifyDeploymentSetAttributeResponse()
	err = client.DoAction(request, response)
	return
}

// invoke ModifyDeploymentSetAttribute api with *ModifyDeploymentSetAttributeRequest asynchronously
// api document: https://help.aliyun.com/api/ecs/modifydeploymentsetattribute.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ModifyDeploymentSetAttributeWithChan(request *ModifyDeploymentSetAttributeRequest) (<-chan *ModifyDeploymentSetAttributeResponse, <-chan error) {
	responseChan := make(chan *ModifyDeploymentSetAttributeResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ModifyDeploymentSetAttribute(request)
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

// invoke ModifyDeploymentSetAttribute api with *ModifyDeploymentSetAttributeRequest asynchronously
// api document: https://help.aliyun.com/api/ecs/modifydeploymentsetattribute.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ModifyDeploymentSetAttributeWithCallback(request *ModifyDeploymentSetAttributeRequest, callback func(response *ModifyDeploymentSetAttributeResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ModifyDeploymentSetAttributeResponse
		var err error
		defer close(result)
		response, err = client.ModifyDeploymentSetAttribute(request)
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

type ModifyDeploymentSetAttributeRequest struct {
	*requests.RpcRequest
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	DeploymentSetId      string           `position:"Query" name:"DeploymentSetId"`
	Description          string           `position:"Query" name:"Description"`
	DeploymentSetName    string           `position:"Query" name:"DeploymentSetName"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
}

type ModifyDeploymentSetAttributeResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// create a request to invoke ModifyDeploymentSetAttribute API
func CreateModifyDeploymentSetAttributeRequest() (request *ModifyDeploymentSetAttributeRequest) {
	request = &ModifyDeploymentSetAttributeRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Ecs", "2014-05-26", "ModifyDeploymentSetAttribute", "ecs", "openAPI")
	return
}

// create a response to parse from ModifyDeploymentSetAttribute response
func CreateModifyDeploymentSetAttributeResponse() (response *ModifyDeploymentSetAttributeResponse) {
	response = &ModifyDeploymentSetAttributeResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
