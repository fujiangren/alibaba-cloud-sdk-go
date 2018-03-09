package rds

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

// invoke AllocateReadWriteSplittingConnection api with *AllocateReadWriteSplittingConnectionRequest synchronously
// api document: https://help.aliyun.com/api/rds/allocatereadwritesplittingconnection.html
func (client *Client) AllocateReadWriteSplittingConnection(request *AllocateReadWriteSplittingConnectionRequest) (response *AllocateReadWriteSplittingConnectionResponse, err error) {
	response = CreateAllocateReadWriteSplittingConnectionResponse()
	err = client.DoAction(request, response)
	return
}

// invoke AllocateReadWriteSplittingConnection api with *AllocateReadWriteSplittingConnectionRequest asynchronously
// api document: https://help.aliyun.com/api/rds/allocatereadwritesplittingconnection.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) AllocateReadWriteSplittingConnectionWithChan(request *AllocateReadWriteSplittingConnectionRequest) (<-chan *AllocateReadWriteSplittingConnectionResponse, <-chan error) {
	responseChan := make(chan *AllocateReadWriteSplittingConnectionResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.AllocateReadWriteSplittingConnection(request)
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

// invoke AllocateReadWriteSplittingConnection api with *AllocateReadWriteSplittingConnectionRequest asynchronously
// api document: https://help.aliyun.com/api/rds/allocatereadwritesplittingconnection.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) AllocateReadWriteSplittingConnectionWithCallback(request *AllocateReadWriteSplittingConnectionRequest, callback func(response *AllocateReadWriteSplittingConnectionResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *AllocateReadWriteSplittingConnectionResponse
		var err error
		defer close(result)
		response, err = client.AllocateReadWriteSplittingConnection(request)
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

type AllocateReadWriteSplittingConnectionRequest struct {
	*requests.RpcRequest
	OwnerId                requests.Integer `position:"Query" name:"OwnerId"`
	ResourceOwnerAccount   string           `position:"Query" name:"ResourceOwnerAccount"`
	ResourceOwnerId        requests.Integer `position:"Query" name:"ResourceOwnerId"`
	OwnerAccount           string           `position:"Query" name:"OwnerAccount"`
	DBInstanceId           string           `position:"Query" name:"DBInstanceId"`
	ConnectionStringPrefix string           `position:"Query" name:"ConnectionStringPrefix"`
	Port                   string           `position:"Query" name:"Port"`
	MaxDelayTime           string           `position:"Query" name:"MaxDelayTime"`
	IPType                 string           `position:"Query" name:"IPType"`
	DistributionType       string           `position:"Query" name:"DistributionType"`
	Weight                 string           `position:"Query" name:"Weight"`
}

type AllocateReadWriteSplittingConnectionResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// create a request to invoke AllocateReadWriteSplittingConnection API
func CreateAllocateReadWriteSplittingConnectionRequest() (request *AllocateReadWriteSplittingConnectionRequest) {
	request = &AllocateReadWriteSplittingConnectionRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Rds", "2014-08-15", "AllocateReadWriteSplittingConnection", "rds", "openAPI")
	return
}

// create a response to parse from AllocateReadWriteSplittingConnection response
func CreateAllocateReadWriteSplittingConnectionResponse() (response *AllocateReadWriteSplittingConnectionResponse) {
	response = &AllocateReadWriteSplittingConnectionResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
