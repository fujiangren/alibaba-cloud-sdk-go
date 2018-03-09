package aegis

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

// invoke GetStatisticsByUuid api with *GetStatisticsByUuidRequest synchronously
// api document: https://help.aliyun.com/api/aegis/getstatisticsbyuuid.html
func (client *Client) GetStatisticsByUuid(request *GetStatisticsByUuidRequest) (response *GetStatisticsByUuidResponse, err error) {
	response = CreateGetStatisticsByUuidResponse()
	err = client.DoAction(request, response)
	return
}

// invoke GetStatisticsByUuid api with *GetStatisticsByUuidRequest asynchronously
// api document: https://help.aliyun.com/api/aegis/getstatisticsbyuuid.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) GetStatisticsByUuidWithChan(request *GetStatisticsByUuidRequest) (<-chan *GetStatisticsByUuidResponse, <-chan error) {
	responseChan := make(chan *GetStatisticsByUuidResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetStatisticsByUuid(request)
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

// invoke GetStatisticsByUuid api with *GetStatisticsByUuidRequest asynchronously
// api document: https://help.aliyun.com/api/aegis/getstatisticsbyuuid.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) GetStatisticsByUuidWithCallback(request *GetStatisticsByUuidRequest, callback func(response *GetStatisticsByUuidResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetStatisticsByUuidResponse
		var err error
		defer close(result)
		response, err = client.GetStatisticsByUuid(request)
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

type GetStatisticsByUuidRequest struct {
	*requests.RpcRequest
	Uuid string `position:"Query" name:"Uuid"`
}

type GetStatisticsByUuidResponse struct {
	*responses.BaseResponse
	RequestId string                    `json:"requestId" xml:"requestId"`
	Code      string                    `json:"Code" xml:"Code"`
	Success   bool                      `json:"Success" xml:"Success"`
	Message   string                    `json:"Message" xml:"Message"`
	Data      DataInGetStatisticsByUuid `json:"Data" xml:"Data"`
}

// create a request to invoke GetStatisticsByUuid API
func CreateGetStatisticsByUuidRequest() (request *GetStatisticsByUuidRequest) {
	request = &GetStatisticsByUuidRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("aegis", "2016-11-11", "GetStatisticsByUuid", "vipaegis", "openAPI")
	return
}

// create a response to parse from GetStatisticsByUuid response
func CreateGetStatisticsByUuidResponse() (response *GetStatisticsByUuidResponse) {
	response = &GetStatisticsByUuidResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
