package vpc

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

// invoke AssociateGlobalAccelerationInstance api with *AssociateGlobalAccelerationInstanceRequest synchronously
// api document: https://help.aliyun.com/api/vpc/associateglobalaccelerationinstance.html
func (client *Client) AssociateGlobalAccelerationInstance(request *AssociateGlobalAccelerationInstanceRequest) (response *AssociateGlobalAccelerationInstanceResponse, err error) {
	response = CreateAssociateGlobalAccelerationInstanceResponse()
	err = client.DoAction(request, response)
	return
}

// invoke AssociateGlobalAccelerationInstance api with *AssociateGlobalAccelerationInstanceRequest asynchronously
// api document: https://help.aliyun.com/api/vpc/associateglobalaccelerationinstance.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) AssociateGlobalAccelerationInstanceWithChan(request *AssociateGlobalAccelerationInstanceRequest) (<-chan *AssociateGlobalAccelerationInstanceResponse, <-chan error) {
	responseChan := make(chan *AssociateGlobalAccelerationInstanceResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.AssociateGlobalAccelerationInstance(request)
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

// invoke AssociateGlobalAccelerationInstance api with *AssociateGlobalAccelerationInstanceRequest asynchronously
// api document: https://help.aliyun.com/api/vpc/associateglobalaccelerationinstance.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) AssociateGlobalAccelerationInstanceWithCallback(request *AssociateGlobalAccelerationInstanceRequest, callback func(response *AssociateGlobalAccelerationInstanceResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *AssociateGlobalAccelerationInstanceResponse
		var err error
		defer close(result)
		response, err = client.AssociateGlobalAccelerationInstance(request)
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

type AssociateGlobalAccelerationInstanceRequest struct {
	*requests.RpcRequest
	OwnerId                      requests.Integer `position:"Query" name:"OwnerId"`
	ResourceOwnerAccount         string           `position:"Query" name:"ResourceOwnerAccount"`
	ResourceOwnerId              requests.Integer `position:"Query" name:"ResourceOwnerId"`
	GlobalAccelerationInstanceId string           `position:"Query" name:"GlobalAccelerationInstanceId"`
	BackendServerId              string           `position:"Query" name:"BackendServerId"`
	BackendServerRegionId        string           `position:"Query" name:"BackendServerRegionId"`
	OwnerAccount                 string           `position:"Query" name:"OwnerAccount"`
	BackendServerType            string           `position:"Query" name:"BackendServerType"`
}

type AssociateGlobalAccelerationInstanceResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// create a request to invoke AssociateGlobalAccelerationInstance API
func CreateAssociateGlobalAccelerationInstanceRequest() (request *AssociateGlobalAccelerationInstanceRequest) {
	request = &AssociateGlobalAccelerationInstanceRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Vpc", "2016-04-28", "AssociateGlobalAccelerationInstance", "vpc", "openAPI")
	return
}

// create a response to parse from AssociateGlobalAccelerationInstance response
func CreateAssociateGlobalAccelerationInstanceResponse() (response *AssociateGlobalAccelerationInstanceResponse) {
	response = &AssociateGlobalAccelerationInstanceResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
