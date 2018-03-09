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

// invoke ModifyDBDescription api with *ModifyDBDescriptionRequest synchronously
// api document: https://help.aliyun.com/api/rds/modifydbdescription.html
func (client *Client) ModifyDBDescription(request *ModifyDBDescriptionRequest) (response *ModifyDBDescriptionResponse, err error) {
	response = CreateModifyDBDescriptionResponse()
	err = client.DoAction(request, response)
	return
}

// invoke ModifyDBDescription api with *ModifyDBDescriptionRequest asynchronously
// api document: https://help.aliyun.com/api/rds/modifydbdescription.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ModifyDBDescriptionWithChan(request *ModifyDBDescriptionRequest) (<-chan *ModifyDBDescriptionResponse, <-chan error) {
	responseChan := make(chan *ModifyDBDescriptionResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ModifyDBDescription(request)
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

// invoke ModifyDBDescription api with *ModifyDBDescriptionRequest asynchronously
// api document: https://help.aliyun.com/api/rds/modifydbdescription.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ModifyDBDescriptionWithCallback(request *ModifyDBDescriptionRequest, callback func(response *ModifyDBDescriptionResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ModifyDBDescriptionResponse
		var err error
		defer close(result)
		response, err = client.ModifyDBDescription(request)
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

type ModifyDBDescriptionRequest struct {
	*requests.RpcRequest
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	DBInstanceId         string           `position:"Query" name:"DBInstanceId"`
	DBName               string           `position:"Query" name:"DBName"`
	DBDescription        string           `position:"Query" name:"DBDescription"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
}

type ModifyDBDescriptionResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// create a request to invoke ModifyDBDescription API
func CreateModifyDBDescriptionRequest() (request *ModifyDBDescriptionRequest) {
	request = &ModifyDBDescriptionRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Rds", "2014-08-15", "ModifyDBDescription", "rds", "openAPI")
	return
}

// create a response to parse from ModifyDBDescription response
func CreateModifyDBDescriptionResponse() (response *ModifyDBDescriptionResponse) {
	response = &ModifyDBDescriptionResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
