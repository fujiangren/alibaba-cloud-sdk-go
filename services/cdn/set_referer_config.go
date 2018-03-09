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

// invoke SetRefererConfig api with *SetRefererConfigRequest synchronously
// api document: https://help.aliyun.com/api/cdn/setrefererconfig.html
func (client *Client) SetRefererConfig(request *SetRefererConfigRequest) (response *SetRefererConfigResponse, err error) {
	response = CreateSetRefererConfigResponse()
	err = client.DoAction(request, response)
	return
}

// invoke SetRefererConfig api with *SetRefererConfigRequest asynchronously
// api document: https://help.aliyun.com/api/cdn/setrefererconfig.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) SetRefererConfigWithChan(request *SetRefererConfigRequest) (<-chan *SetRefererConfigResponse, <-chan error) {
	responseChan := make(chan *SetRefererConfigResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.SetRefererConfig(request)
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

// invoke SetRefererConfig api with *SetRefererConfigRequest asynchronously
// api document: https://help.aliyun.com/api/cdn/setrefererconfig.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) SetRefererConfigWithCallback(request *SetRefererConfigRequest, callback func(response *SetRefererConfigResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *SetRefererConfigResponse
		var err error
		defer close(result)
		response, err = client.SetRefererConfig(request)
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

type SetRefererConfigRequest struct {
	*requests.RpcRequest
	OwnerId       requests.Integer `position:"Query" name:"OwnerId"`
	SecurityToken string           `position:"Query" name:"SecurityToken"`
	DomainName    string           `position:"Query" name:"DomainName"`
	ReferType     string           `position:"Query" name:"ReferType"`
	ReferList     string           `position:"Query" name:"ReferList"`
	AllowEmpty    string           `position:"Query" name:"AllowEmpty"`
	DisableAst    string           `position:"Query" name:"DisableAst"`
}

type SetRefererConfigResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// create a request to invoke SetRefererConfig API
func CreateSetRefererConfigRequest() (request *SetRefererConfigRequest) {
	request = &SetRefererConfigRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Cdn", "2014-11-11", "SetRefererConfig", "", "")
	return
}

// create a response to parse from SetRefererConfig response
func CreateSetRefererConfigResponse() (response *SetRefererConfigResponse) {
	response = &SetRefererConfigResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
