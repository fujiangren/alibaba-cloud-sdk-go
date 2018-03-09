package cloudphoto

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

// invoke GetLibrary api with *GetLibraryRequest synchronously
// api document: https://help.aliyun.com/api/cloudphoto/getlibrary.html
func (client *Client) GetLibrary(request *GetLibraryRequest) (response *GetLibraryResponse, err error) {
	response = CreateGetLibraryResponse()
	err = client.DoAction(request, response)
	return
}

// invoke GetLibrary api with *GetLibraryRequest asynchronously
// api document: https://help.aliyun.com/api/cloudphoto/getlibrary.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) GetLibraryWithChan(request *GetLibraryRequest) (<-chan *GetLibraryResponse, <-chan error) {
	responseChan := make(chan *GetLibraryResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetLibrary(request)
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

// invoke GetLibrary api with *GetLibraryRequest asynchronously
// api document: https://help.aliyun.com/api/cloudphoto/getlibrary.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) GetLibraryWithCallback(request *GetLibraryRequest, callback func(response *GetLibraryResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetLibraryResponse
		var err error
		defer close(result)
		response, err = client.GetLibrary(request)
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

type GetLibraryRequest struct {
	*requests.RpcRequest
	StoreName string `position:"Query" name:"StoreName"`
	LibraryId string `position:"Query" name:"LibraryId"`
}

type GetLibraryResponse struct {
	*responses.BaseResponse
	Code      string  `json:"Code" xml:"Code"`
	Message   string  `json:"Message" xml:"Message"`
	RequestId string  `json:"RequestId" xml:"RequestId"`
	Action    string  `json:"Action" xml:"Action"`
	Library   Library `json:"Library" xml:"Library"`
}

// create a request to invoke GetLibrary API
func CreateGetLibraryRequest() (request *GetLibraryRequest) {
	request = &GetLibraryRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("CloudPhoto", "2017-07-11", "GetLibrary", "cloudphoto", "openAPI")
	return
}

// create a response to parse from GetLibrary response
func CreateGetLibraryResponse() (response *GetLibraryResponse) {
	response = &GetLibraryResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
