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

// invoke CreateAlbum api with *CreateAlbumRequest synchronously
// api document: https://help.aliyun.com/api/cloudphoto/createalbum.html
func (client *Client) CreateAlbum(request *CreateAlbumRequest) (response *CreateAlbumResponse, err error) {
	response = CreateCreateAlbumResponse()
	err = client.DoAction(request, response)
	return
}

// invoke CreateAlbum api with *CreateAlbumRequest asynchronously
// api document: https://help.aliyun.com/api/cloudphoto/createalbum.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) CreateAlbumWithChan(request *CreateAlbumRequest) (<-chan *CreateAlbumResponse, <-chan error) {
	responseChan := make(chan *CreateAlbumResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateAlbum(request)
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

// invoke CreateAlbum api with *CreateAlbumRequest asynchronously
// api document: https://help.aliyun.com/api/cloudphoto/createalbum.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) CreateAlbumWithCallback(request *CreateAlbumRequest, callback func(response *CreateAlbumResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateAlbumResponse
		var err error
		defer close(result)
		response, err = client.CreateAlbum(request)
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

type CreateAlbumRequest struct {
	*requests.RpcRequest
	AlbumName string `position:"Query" name:"AlbumName"`
	StoreName string `position:"Query" name:"StoreName"`
	Remark    string `position:"Query" name:"Remark"`
	LibraryId string `position:"Query" name:"LibraryId"`
}

type CreateAlbumResponse struct {
	*responses.BaseResponse
	Code      string `json:"Code" xml:"Code"`
	Message   string `json:"Message" xml:"Message"`
	RequestId string `json:"RequestId" xml:"RequestId"`
	Action    string `json:"Action" xml:"Action"`
	Album     Album  `json:"Album" xml:"Album"`
}

// create a request to invoke CreateAlbum API
func CreateCreateAlbumRequest() (request *CreateAlbumRequest) {
	request = &CreateAlbumRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("CloudPhoto", "2017-07-11", "CreateAlbum", "cloudphoto", "openAPI")
	return
}

// create a response to parse from CreateAlbum response
func CreateCreateAlbumResponse() (response *CreateAlbumResponse) {
	response = &CreateAlbumResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
