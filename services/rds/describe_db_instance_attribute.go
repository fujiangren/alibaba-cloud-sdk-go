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

// invoke DescribeDBInstanceAttribute api with *DescribeDBInstanceAttributeRequest synchronously
// api document: https://help.aliyun.com/api/rds/describedbinstanceattribute.html
func (client *Client) DescribeDBInstanceAttribute(request *DescribeDBInstanceAttributeRequest) (response *DescribeDBInstanceAttributeResponse, err error) {
	response = CreateDescribeDBInstanceAttributeResponse()
	err = client.DoAction(request, response)
	return
}

// invoke DescribeDBInstanceAttribute api with *DescribeDBInstanceAttributeRequest asynchronously
// api document: https://help.aliyun.com/api/rds/describedbinstanceattribute.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeDBInstanceAttributeWithChan(request *DescribeDBInstanceAttributeRequest) (<-chan *DescribeDBInstanceAttributeResponse, <-chan error) {
	responseChan := make(chan *DescribeDBInstanceAttributeResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeDBInstanceAttribute(request)
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

// invoke DescribeDBInstanceAttribute api with *DescribeDBInstanceAttributeRequest asynchronously
// api document: https://help.aliyun.com/api/rds/describedbinstanceattribute.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeDBInstanceAttributeWithCallback(request *DescribeDBInstanceAttributeRequest, callback func(response *DescribeDBInstanceAttributeResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeDBInstanceAttributeResponse
		var err error
		defer close(result)
		response, err = client.DescribeDBInstanceAttribute(request)
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

type DescribeDBInstanceAttributeRequest struct {
	*requests.RpcRequest
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	DBInstanceId         string           `position:"Query" name:"DBInstanceId"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	ResourceGroupId      string           `position:"Query" name:"ResourceGroupId"`
}

type DescribeDBInstanceAttributeResponse struct {
	*responses.BaseResponse
	RequestId string                             `json:"RequestId" xml:"RequestId"`
	Items     ItemsInDescribeDBInstanceAttribute `json:"Items" xml:"Items"`
}

// create a request to invoke DescribeDBInstanceAttribute API
func CreateDescribeDBInstanceAttributeRequest() (request *DescribeDBInstanceAttributeRequest) {
	request = &DescribeDBInstanceAttributeRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Rds", "2014-08-15", "DescribeDBInstanceAttribute", "rds", "openAPI")
	return
}

// create a response to parse from DescribeDBInstanceAttribute response
func CreateDescribeDBInstanceAttributeResponse() (response *DescribeDBInstanceAttributeResponse) {
	response = &DescribeDBInstanceAttributeResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
