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

// invoke ModifyPostpaidDBInstanceSpec api with *ModifyPostpaidDBInstanceSpecRequest synchronously
// api document: https://help.aliyun.com/api/rds/modifypostpaiddbinstancespec.html
func (client *Client) ModifyPostpaidDBInstanceSpec(request *ModifyPostpaidDBInstanceSpecRequest) (response *ModifyPostpaidDBInstanceSpecResponse, err error) {
	response = CreateModifyPostpaidDBInstanceSpecResponse()
	err = client.DoAction(request, response)
	return
}

// invoke ModifyPostpaidDBInstanceSpec api with *ModifyPostpaidDBInstanceSpecRequest asynchronously
// api document: https://help.aliyun.com/api/rds/modifypostpaiddbinstancespec.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ModifyPostpaidDBInstanceSpecWithChan(request *ModifyPostpaidDBInstanceSpecRequest) (<-chan *ModifyPostpaidDBInstanceSpecResponse, <-chan error) {
	responseChan := make(chan *ModifyPostpaidDBInstanceSpecResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ModifyPostpaidDBInstanceSpec(request)
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

// invoke ModifyPostpaidDBInstanceSpec api with *ModifyPostpaidDBInstanceSpecRequest asynchronously
// api document: https://help.aliyun.com/api/rds/modifypostpaiddbinstancespec.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ModifyPostpaidDBInstanceSpecWithCallback(request *ModifyPostpaidDBInstanceSpecRequest, callback func(response *ModifyPostpaidDBInstanceSpecResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ModifyPostpaidDBInstanceSpecResponse
		var err error
		defer close(result)
		response, err = client.ModifyPostpaidDBInstanceSpec(request)
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

type ModifyPostpaidDBInstanceSpecRequest struct {
	*requests.RpcRequest
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	ClientToken          string           `position:"Query" name:"ClientToken"`
	DBInstanceId         string           `position:"Query" name:"DBInstanceId"`
	DBInstanceClass      string           `position:"Query" name:"DBInstanceClass"`
	DBInstanceStorage    requests.Integer `position:"Query" name:"DBInstanceStorage"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
}

type ModifyPostpaidDBInstanceSpecResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// create a request to invoke ModifyPostpaidDBInstanceSpec API
func CreateModifyPostpaidDBInstanceSpecRequest() (request *ModifyPostpaidDBInstanceSpecRequest) {
	request = &ModifyPostpaidDBInstanceSpecRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Rds", "2014-08-15", "ModifyPostpaidDBInstanceSpec", "rds", "openAPI")
	return
}

// create a response to parse from ModifyPostpaidDBInstanceSpec response
func CreateModifyPostpaidDBInstanceSpecResponse() (response *ModifyPostpaidDBInstanceSpecResponse) {
	response = &ModifyPostpaidDBInstanceSpecResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
