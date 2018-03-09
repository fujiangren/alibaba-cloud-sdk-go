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

// invoke DescribeMasterSlaveServerGroups api with *DescribeMasterSlaveServerGroupsRequest synchronously
// api document: https://help.aliyun.com/api/slb/describemasterslaveservergroups.html
func (client *Client) DescribeMasterSlaveServerGroups(request *DescribeMasterSlaveServerGroupsRequest) (response *DescribeMasterSlaveServerGroupsResponse, err error) {
	response = CreateDescribeMasterSlaveServerGroupsResponse()
	err = client.DoAction(request, response)
	return
}

// invoke DescribeMasterSlaveServerGroups api with *DescribeMasterSlaveServerGroupsRequest asynchronously
// api document: https://help.aliyun.com/api/slb/describemasterslaveservergroups.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeMasterSlaveServerGroupsWithChan(request *DescribeMasterSlaveServerGroupsRequest) (<-chan *DescribeMasterSlaveServerGroupsResponse, <-chan error) {
	responseChan := make(chan *DescribeMasterSlaveServerGroupsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeMasterSlaveServerGroups(request)
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

// invoke DescribeMasterSlaveServerGroups api with *DescribeMasterSlaveServerGroupsRequest asynchronously
// api document: https://help.aliyun.com/api/slb/describemasterslaveservergroups.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeMasterSlaveServerGroupsWithCallback(request *DescribeMasterSlaveServerGroupsRequest, callback func(response *DescribeMasterSlaveServerGroupsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeMasterSlaveServerGroupsResponse
		var err error
		defer close(result)
		response, err = client.DescribeMasterSlaveServerGroups(request)
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

type DescribeMasterSlaveServerGroupsRequest struct {
	*requests.RpcRequest
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	AccessKeyId          string           `position:"Query" name:"access_key_id"`
	Tags                 string           `position:"Query" name:"Tags"`
	LoadBalancerId       string           `position:"Query" name:"LoadBalancerId"`
}

type DescribeMasterSlaveServerGroupsResponse struct {
	*responses.BaseResponse
	RequestId               string                  `json:"RequestId" xml:"RequestId"`
	MasterSlaveServerGroups MasterSlaveServerGroups `json:"MasterSlaveServerGroups" xml:"MasterSlaveServerGroups"`
}

// create a request to invoke DescribeMasterSlaveServerGroups API
func CreateDescribeMasterSlaveServerGroupsRequest() (request *DescribeMasterSlaveServerGroupsRequest) {
	request = &DescribeMasterSlaveServerGroupsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Slb", "2014-05-15", "DescribeMasterSlaveServerGroups", "slb", "openAPI")
	return
}

// create a response to parse from DescribeMasterSlaveServerGroups response
func CreateDescribeMasterSlaveServerGroupsResponse() (response *DescribeMasterSlaveServerGroupsResponse) {
	response = &DescribeMasterSlaveServerGroupsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
