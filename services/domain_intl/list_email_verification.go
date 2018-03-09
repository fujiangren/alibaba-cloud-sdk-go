package domain_intl

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

// invoke ListEmailVerification api with *ListEmailVerificationRequest synchronously
// api document: https://help.aliyun.com/api/domain-intl/listemailverification.html
func (client *Client) ListEmailVerification(request *ListEmailVerificationRequest) (response *ListEmailVerificationResponse, err error) {
	response = CreateListEmailVerificationResponse()
	err = client.DoAction(request, response)
	return
}

// invoke ListEmailVerification api with *ListEmailVerificationRequest asynchronously
// api document: https://help.aliyun.com/api/domain-intl/listemailverification.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ListEmailVerificationWithChan(request *ListEmailVerificationRequest) (<-chan *ListEmailVerificationResponse, <-chan error) {
	responseChan := make(chan *ListEmailVerificationResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListEmailVerification(request)
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

// invoke ListEmailVerification api with *ListEmailVerificationRequest asynchronously
// api document: https://help.aliyun.com/api/domain-intl/listemailverification.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ListEmailVerificationWithCallback(request *ListEmailVerificationRequest, callback func(response *ListEmailVerificationResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListEmailVerificationResponse
		var err error
		defer close(result)
		response, err = client.ListEmailVerification(request)
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

type ListEmailVerificationRequest struct {
	*requests.RpcRequest
	Lang               string           `position:"Query" name:"Lang"`
	BeginCreateTime    requests.Integer `position:"Query" name:"BeginCreateTime"`
	EndCreateTime      requests.Integer `position:"Query" name:"EndCreateTime"`
	Email              string           `position:"Query" name:"Email"`
	VerificationStatus requests.Integer `position:"Query" name:"VerificationStatus"`
	PageNum            requests.Integer `position:"Query" name:"PageNum"`
	PageSize           requests.Integer `position:"Query" name:"PageSize"`
}

type ListEmailVerificationResponse struct {
	*responses.BaseResponse
	RequestId      string              `json:"RequestId" xml:"RequestId"`
	TotalItemNum   int                 `json:"TotalItemNum" xml:"TotalItemNum"`
	CurrentPageNum int                 `json:"CurrentPageNum" xml:"CurrentPageNum"`
	TotalPageNum   int                 `json:"TotalPageNum" xml:"TotalPageNum"`
	PageSize       int                 `json:"PageSize" xml:"PageSize"`
	PrePage        bool                `json:"PrePage" xml:"PrePage"`
	NextPage       bool                `json:"NextPage" xml:"NextPage"`
	Data           []EmailVerification `json:"Data" xml:"Data"`
}

// create a request to invoke ListEmailVerification API
func CreateListEmailVerificationRequest() (request *ListEmailVerificationRequest) {
	request = &ListEmailVerificationRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Domain-intl", "2017-12-18", "ListEmailVerification", "", "")
	return
}

// create a response to parse from ListEmailVerification response
func CreateListEmailVerificationResponse() (response *ListEmailVerificationResponse) {
	response = &ListEmailVerificationResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
