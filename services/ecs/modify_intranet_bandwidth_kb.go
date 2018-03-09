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

// invoke ModifyIntranetBandwidthKb api with *ModifyIntranetBandwidthKbRequest synchronously
// api document: https://help.aliyun.com/api/ecs/modifyintranetbandwidthkb.html
func (client *Client) ModifyIntranetBandwidthKb(request *ModifyIntranetBandwidthKbRequest) (response *ModifyIntranetBandwidthKbResponse, err error) {
	response = CreateModifyIntranetBandwidthKbResponse()
	err = client.DoAction(request, response)
	return
}

// invoke ModifyIntranetBandwidthKb api with *ModifyIntranetBandwidthKbRequest asynchronously
// api document: https://help.aliyun.com/api/ecs/modifyintranetbandwidthkb.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ModifyIntranetBandwidthKbWithChan(request *ModifyIntranetBandwidthKbRequest) (<-chan *ModifyIntranetBandwidthKbResponse, <-chan error) {
	responseChan := make(chan *ModifyIntranetBandwidthKbResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ModifyIntranetBandwidthKb(request)
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

// invoke ModifyIntranetBandwidthKb api with *ModifyIntranetBandwidthKbRequest asynchronously
// api document: https://help.aliyun.com/api/ecs/modifyintranetbandwidthkb.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ModifyIntranetBandwidthKbWithCallback(request *ModifyIntranetBandwidthKbRequest, callback func(response *ModifyIntranetBandwidthKbResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ModifyIntranetBandwidthKbResponse
		var err error
		defer close(result)
		response, err = client.ModifyIntranetBandwidthKb(request)
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

type ModifyIntranetBandwidthKbRequest struct {
	*requests.RpcRequest
	OwnerId                 requests.Integer `position:"Query" name:"OwnerId"`
	ResourceOwnerAccount    string           `position:"Query" name:"ResourceOwnerAccount"`
	ResourceOwnerId         requests.Integer `position:"Query" name:"ResourceOwnerId"`
	InstanceId              string           `position:"Query" name:"InstanceId"`
	IntranetMaxBandwidthIn  requests.Integer `position:"Query" name:"IntranetMaxBandwidthIn"`
	IntranetMaxBandwidthOut requests.Integer `position:"Query" name:"IntranetMaxBandwidthOut"`
	OwnerAccount            string           `position:"Query" name:"OwnerAccount"`
}

type ModifyIntranetBandwidthKbResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// create a request to invoke ModifyIntranetBandwidthKb API
func CreateModifyIntranetBandwidthKbRequest() (request *ModifyIntranetBandwidthKbRequest) {
	request = &ModifyIntranetBandwidthKbRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Ecs", "2014-05-26", "ModifyIntranetBandwidthKb", "ecs", "openAPI")
	return
}

// create a response to parse from ModifyIntranetBandwidthKb response
func CreateModifyIntranetBandwidthKbResponse() (response *ModifyIntranetBandwidthKbResponse) {
	response = &ModifyIntranetBandwidthKbResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
