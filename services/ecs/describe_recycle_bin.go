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

// invoke DescribeRecycleBin api with *DescribeRecycleBinRequest synchronously
// api document: https://help.aliyun.com/api/ecs/describerecyclebin.html
func (client *Client) DescribeRecycleBin(request *DescribeRecycleBinRequest) (response *DescribeRecycleBinResponse, err error) {
	response = CreateDescribeRecycleBinResponse()
	err = client.DoAction(request, response)
	return
}

// invoke DescribeRecycleBin api with *DescribeRecycleBinRequest asynchronously
// api document: https://help.aliyun.com/api/ecs/describerecyclebin.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeRecycleBinWithChan(request *DescribeRecycleBinRequest) (<-chan *DescribeRecycleBinResponse, <-chan error) {
	responseChan := make(chan *DescribeRecycleBinResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeRecycleBin(request)
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

// invoke DescribeRecycleBin api with *DescribeRecycleBinRequest asynchronously
// api document: https://help.aliyun.com/api/ecs/describerecyclebin.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeRecycleBinWithCallback(request *DescribeRecycleBinRequest, callback func(response *DescribeRecycleBinResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeRecycleBinResponse
		var err error
		defer close(result)
		response, err = client.DescribeRecycleBin(request)
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

type DescribeRecycleBinRequest struct {
	*requests.RpcRequest
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	ResourceId           string           `position:"Query" name:"ResourceId"`
	Status               string           `position:"Query" name:"Status"`
	PageNumber           requests.Integer `position:"Query" name:"PageNumber"`
	PageSize             requests.Integer `position:"Query" name:"PageSize"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
}

type DescribeRecycleBinResponse struct {
	*responses.BaseResponse
	RequestId        string           `json:"RequestId" xml:"RequestId"`
	TotalCount       int              `json:"TotalCount" xml:"TotalCount"`
	RecycleBinModels RecycleBinModels `json:"RecycleBinModels" xml:"RecycleBinModels"`
}

// create a request to invoke DescribeRecycleBin API
func CreateDescribeRecycleBinRequest() (request *DescribeRecycleBinRequest) {
	request = &DescribeRecycleBinRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Ecs", "2014-05-26", "DescribeRecycleBin", "ecs", "openAPI")
	return
}

// create a response to parse from DescribeRecycleBin response
func CreateDescribeRecycleBinResponse() (response *DescribeRecycleBinResponse) {
	response = &DescribeRecycleBinResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
