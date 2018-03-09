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

// invoke DeleteCACertificate api with *DeleteCACertificateRequest synchronously
// api document: https://help.aliyun.com/api/slb/deletecacertificate.html
func (client *Client) DeleteCACertificate(request *DeleteCACertificateRequest) (response *DeleteCACertificateResponse, err error) {
	response = CreateDeleteCACertificateResponse()
	err = client.DoAction(request, response)
	return
}

// invoke DeleteCACertificate api with *DeleteCACertificateRequest asynchronously
// api document: https://help.aliyun.com/api/slb/deletecacertificate.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DeleteCACertificateWithChan(request *DeleteCACertificateRequest) (<-chan *DeleteCACertificateResponse, <-chan error) {
	responseChan := make(chan *DeleteCACertificateResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DeleteCACertificate(request)
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

// invoke DeleteCACertificate api with *DeleteCACertificateRequest asynchronously
// api document: https://help.aliyun.com/api/slb/deletecacertificate.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DeleteCACertificateWithCallback(request *DeleteCACertificateRequest, callback func(response *DeleteCACertificateResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DeleteCACertificateResponse
		var err error
		defer close(result)
		response, err = client.DeleteCACertificate(request)
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

type DeleteCACertificateRequest struct {
	*requests.RpcRequest
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	AccessKeyId          string           `position:"Query" name:"access_key_id"`
	CACertificateId      string           `position:"Query" name:"CACertificateId"`
}

type DeleteCACertificateResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// create a request to invoke DeleteCACertificate API
func CreateDeleteCACertificateRequest() (request *DeleteCACertificateRequest) {
	request = &DeleteCACertificateRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Slb", "2014-05-15", "DeleteCACertificate", "slb", "openAPI")
	return
}

// create a response to parse from DeleteCACertificate response
func CreateDeleteCACertificateResponse() (response *DeleteCACertificateResponse) {
	response = &DeleteCACertificateResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
