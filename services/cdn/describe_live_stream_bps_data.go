package cdn

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

// invoke DescribeLiveStreamBpsData api with *DescribeLiveStreamBpsDataRequest synchronously
// api document: https://help.aliyun.com/api/cdn/describelivestreambpsdata.html
func (client *Client) DescribeLiveStreamBpsData(request *DescribeLiveStreamBpsDataRequest) (response *DescribeLiveStreamBpsDataResponse, err error) {
	response = CreateDescribeLiveStreamBpsDataResponse()
	err = client.DoAction(request, response)
	return
}

// invoke DescribeLiveStreamBpsData api with *DescribeLiveStreamBpsDataRequest asynchronously
// api document: https://help.aliyun.com/api/cdn/describelivestreambpsdata.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeLiveStreamBpsDataWithChan(request *DescribeLiveStreamBpsDataRequest) (<-chan *DescribeLiveStreamBpsDataResponse, <-chan error) {
	responseChan := make(chan *DescribeLiveStreamBpsDataResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeLiveStreamBpsData(request)
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

// invoke DescribeLiveStreamBpsData api with *DescribeLiveStreamBpsDataRequest asynchronously
// api document: https://help.aliyun.com/api/cdn/describelivestreambpsdata.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeLiveStreamBpsDataWithCallback(request *DescribeLiveStreamBpsDataRequest, callback func(response *DescribeLiveStreamBpsDataResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeLiveStreamBpsDataResponse
		var err error
		defer close(result)
		response, err = client.DescribeLiveStreamBpsData(request)
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

type DescribeLiveStreamBpsDataRequest struct {
	*requests.RpcRequest
	OwnerId       requests.Integer `position:"Query" name:"OwnerId"`
	SecurityToken string           `position:"Query" name:"SecurityToken"`
	DomainName    string           `position:"Query" name:"DomainName"`
	AppName       string           `position:"Query" name:"AppName"`
	StreamName    string           `position:"Query" name:"StreamName"`
	StartTime     string           `position:"Query" name:"StartTime"`
	EndTime       string           `position:"Query" name:"EndTime"`
}

type DescribeLiveStreamBpsDataResponse struct {
	*responses.BaseResponse
	RequestId string                              `json:"RequestId" xml:"RequestId"`
	BpsDatas  BpsDatasInDescribeLiveStreamBpsData `json:"BpsDatas" xml:"BpsDatas"`
}

// create a request to invoke DescribeLiveStreamBpsData API
func CreateDescribeLiveStreamBpsDataRequest() (request *DescribeLiveStreamBpsDataRequest) {
	request = &DescribeLiveStreamBpsDataRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Cdn", "2014-11-11", "DescribeLiveStreamBpsData", "", "")
	return
}

// create a response to parse from DescribeLiveStreamBpsData response
func CreateDescribeLiveStreamBpsDataResponse() (response *DescribeLiveStreamBpsDataResponse) {
	response = &DescribeLiveStreamBpsDataResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
