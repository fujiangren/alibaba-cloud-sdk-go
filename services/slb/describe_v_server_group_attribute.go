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

// invoke DescribeVServerGroupAttribute api with *DescribeVServerGroupAttributeRequest synchronously
// api document: https://help.aliyun.com/api/slb/describevservergroupattribute.html
func (client *Client) DescribeVServerGroupAttribute(request *DescribeVServerGroupAttributeRequest) (response *DescribeVServerGroupAttributeResponse, err error) {
	response = CreateDescribeVServerGroupAttributeResponse()
	err = client.DoAction(request, response)
	return
}

// invoke DescribeVServerGroupAttribute api with *DescribeVServerGroupAttributeRequest asynchronously
// api document: https://help.aliyun.com/api/slb/describevservergroupattribute.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeVServerGroupAttributeWithChan(request *DescribeVServerGroupAttributeRequest) (<-chan *DescribeVServerGroupAttributeResponse, <-chan error) {
	responseChan := make(chan *DescribeVServerGroupAttributeResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeVServerGroupAttribute(request)
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

// invoke DescribeVServerGroupAttribute api with *DescribeVServerGroupAttributeRequest asynchronously
// api document: https://help.aliyun.com/api/slb/describevservergroupattribute.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeVServerGroupAttributeWithCallback(request *DescribeVServerGroupAttributeRequest, callback func(response *DescribeVServerGroupAttributeResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeVServerGroupAttributeResponse
		var err error
		defer close(result)
		response, err = client.DescribeVServerGroupAttribute(request)
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

type DescribeVServerGroupAttributeRequest struct {
	*requests.RpcRequest
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	AccessKeyId          string           `position:"Query" name:"access_key_id"`
	Tags                 string           `position:"Query" name:"Tags"`
	VServerGroupId       string           `position:"Query" name:"VServerGroupId"`
}

type DescribeVServerGroupAttributeResponse struct {
	*responses.BaseResponse
	RequestId        string                                        `json:"RequestId" xml:"RequestId"`
	VServerGroupId   string                                        `json:"VServerGroupId" xml:"VServerGroupId"`
	VServerGroupName string                                        `json:"VServerGroupName" xml:"VServerGroupName"`
	BackendServers   BackendServersInDescribeVServerGroupAttribute `json:"BackendServers" xml:"BackendServers"`
}

// create a request to invoke DescribeVServerGroupAttribute API
func CreateDescribeVServerGroupAttributeRequest() (request *DescribeVServerGroupAttributeRequest) {
	request = &DescribeVServerGroupAttributeRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Slb", "2014-05-15", "DescribeVServerGroupAttribute", "slb", "openAPI")
	return
}

// create a response to parse from DescribeVServerGroupAttribute response
func CreateDescribeVServerGroupAttributeResponse() (response *DescribeVServerGroupAttributeResponse) {
	response = &DescribeVServerGroupAttributeResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
