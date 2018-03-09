package mts

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

// invoke QueryFacerecogJobList api with *QueryFacerecogJobListRequest synchronously
// api document: https://help.aliyun.com/api/mts/queryfacerecogjoblist.html
func (client *Client) QueryFacerecogJobList(request *QueryFacerecogJobListRequest) (response *QueryFacerecogJobListResponse, err error) {
	response = CreateQueryFacerecogJobListResponse()
	err = client.DoAction(request, response)
	return
}

// invoke QueryFacerecogJobList api with *QueryFacerecogJobListRequest asynchronously
// api document: https://help.aliyun.com/api/mts/queryfacerecogjoblist.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) QueryFacerecogJobListWithChan(request *QueryFacerecogJobListRequest) (<-chan *QueryFacerecogJobListResponse, <-chan error) {
	responseChan := make(chan *QueryFacerecogJobListResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.QueryFacerecogJobList(request)
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

// invoke QueryFacerecogJobList api with *QueryFacerecogJobListRequest asynchronously
// api document: https://help.aliyun.com/api/mts/queryfacerecogjoblist.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) QueryFacerecogJobListWithCallback(request *QueryFacerecogJobListRequest, callback func(response *QueryFacerecogJobListResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *QueryFacerecogJobListResponse
		var err error
		defer close(result)
		response, err = client.QueryFacerecogJobList(request)
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

type QueryFacerecogJobListRequest struct {
	*requests.RpcRequest
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	FacerecogJobIds      string           `position:"Query" name:"FacerecogJobIds"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
}

type QueryFacerecogJobListResponse struct {
	*responses.BaseResponse
	RequestId        string                             `json:"RequestId" xml:"RequestId"`
	NonExistIds      NonExistIdsInQueryFacerecogJobList `json:"NonExistIds" xml:"NonExistIds"`
	FacerecogJobList FacerecogJobList                   `json:"FacerecogJobList" xml:"FacerecogJobList"`
}

// create a request to invoke QueryFacerecogJobList API
func CreateQueryFacerecogJobListRequest() (request *QueryFacerecogJobListRequest) {
	request = &QueryFacerecogJobListRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Mts", "2014-06-18", "QueryFacerecogJobList", "mts", "openAPI")
	return
}

// create a response to parse from QueryFacerecogJobList response
func CreateQueryFacerecogJobListResponse() (response *QueryFacerecogJobListResponse) {
	response = &QueryFacerecogJobListResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
