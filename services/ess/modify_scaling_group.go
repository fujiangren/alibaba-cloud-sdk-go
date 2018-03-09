package ess

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

// invoke ModifyScalingGroup api with *ModifyScalingGroupRequest synchronously
// api document: https://help.aliyun.com/api/ess/modifyscalinggroup.html
func (client *Client) ModifyScalingGroup(request *ModifyScalingGroupRequest) (response *ModifyScalingGroupResponse, err error) {
	response = CreateModifyScalingGroupResponse()
	err = client.DoAction(request, response)
	return
}

// invoke ModifyScalingGroup api with *ModifyScalingGroupRequest asynchronously
// api document: https://help.aliyun.com/api/ess/modifyscalinggroup.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ModifyScalingGroupWithChan(request *ModifyScalingGroupRequest) (<-chan *ModifyScalingGroupResponse, <-chan error) {
	responseChan := make(chan *ModifyScalingGroupResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ModifyScalingGroup(request)
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

// invoke ModifyScalingGroup api with *ModifyScalingGroupRequest asynchronously
// api document: https://help.aliyun.com/api/ess/modifyscalinggroup.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ModifyScalingGroupWithCallback(request *ModifyScalingGroupRequest, callback func(response *ModifyScalingGroupResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ModifyScalingGroupResponse
		var err error
		defer close(result)
		response, err = client.ModifyScalingGroup(request)
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

type ModifyScalingGroupRequest struct {
	*requests.RpcRequest
	OwnerId                      requests.Integer `position:"Query" name:"OwnerId"`
	ResourceOwnerAccount         string           `position:"Query" name:"ResourceOwnerAccount"`
	ResourceOwnerId              requests.Integer `position:"Query" name:"ResourceOwnerId"`
	ScalingGroupId               string           `position:"Query" name:"ScalingGroupId"`
	ScalingGroupName             string           `position:"Query" name:"ScalingGroupName"`
	MinSize                      requests.Integer `position:"Query" name:"MinSize"`
	MaxSize                      requests.Integer `position:"Query" name:"MaxSize"`
	DefaultCooldown              requests.Integer `position:"Query" name:"DefaultCooldown"`
	RemovalPolicy1               string           `position:"Query" name:"RemovalPolicy.1"`
	RemovalPolicy2               string           `position:"Query" name:"RemovalPolicy.2"`
	ActiveScalingConfigurationId string           `position:"Query" name:"ActiveScalingConfigurationId"`
	OwnerAccount                 string           `position:"Query" name:"OwnerAccount"`
}

type ModifyScalingGroupResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// create a request to invoke ModifyScalingGroup API
func CreateModifyScalingGroupRequest() (request *ModifyScalingGroupRequest) {
	request = &ModifyScalingGroupRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Ess", "2014-08-28", "ModifyScalingGroup", "ess", "openAPI")
	return
}

// create a response to parse from ModifyScalingGroup response
func CreateModifyScalingGroupResponse() (response *ModifyScalingGroupResponse) {
	response = &ModifyScalingGroupResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
