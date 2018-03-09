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

// invoke QueryLoginEvent api with *QueryLoginEventRequest synchronously
// api document: https://help.aliyun.com/api/aegis/queryloginevent.html
func (client *Client) QueryLoginEvent(request *QueryLoginEventRequest) (response *QueryLoginEventResponse, err error) {
	response = CreateQueryLoginEventResponse()
	err = client.DoAction(request, response)
	return
}

// invoke QueryLoginEvent api with *QueryLoginEventRequest asynchronously
// api document: https://help.aliyun.com/api/aegis/queryloginevent.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) QueryLoginEventWithChan(request *QueryLoginEventRequest) (<-chan *QueryLoginEventResponse, <-chan error) {
	responseChan := make(chan *QueryLoginEventResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.QueryLoginEvent(request)
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

// invoke QueryLoginEvent api with *QueryLoginEventRequest asynchronously
// api document: https://help.aliyun.com/api/aegis/queryloginevent.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) QueryLoginEventWithCallback(request *QueryLoginEventRequest, callback func(response *QueryLoginEventResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *QueryLoginEventResponse
		var err error
		defer close(result)
		response, err = client.QueryLoginEvent(request)
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

type QueryLoginEventRequest struct {
	*requests.RpcRequest
	Uuid        string           `position:"Query" name:"Uuid"`
	Status      requests.Integer `position:"Query" name:"Status"`
	CurrentPage requests.Integer `position:"Query" name:"CurrentPage"`
	StartTime   string           `position:"Query" name:"StartTime"`
	EndTime     string           `position:"Query" name:"EndTime"`
}

type QueryLoginEventResponse struct {
	*responses.BaseResponse
	RequestId string `json:"requestId" xml:"requestId"`
	Code      string `json:"Code" xml:"Code"`
	Success   bool   `json:"Success" xml:"Success"`
	Message   string `json:"Message" xml:"Message"`
	Data      Data   `json:"Data" xml:"Data"`
}

// create a request to invoke QueryLoginEvent API
func CreateQueryLoginEventRequest() (request *QueryLoginEventRequest) {
	request = &QueryLoginEventRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("aegis", "2016-11-11", "QueryLoginEvent", "vipaegis", "openAPI")
	return
}

// create a response to parse from QueryLoginEvent response
func CreateQueryLoginEventResponse() (response *QueryLoginEventResponse) {
	response = &QueryLoginEventResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
