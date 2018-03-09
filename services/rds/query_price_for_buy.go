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

// invoke QueryPriceForBuy api with *QueryPriceForBuyRequest synchronously
// api document: https://help.aliyun.com/api/rds/querypriceforbuy.html
func (client *Client) QueryPriceForBuy(request *QueryPriceForBuyRequest) (response *QueryPriceForBuyResponse, err error) {
	response = CreateQueryPriceForBuyResponse()
	err = client.DoAction(request, response)
	return
}

// invoke QueryPriceForBuy api with *QueryPriceForBuyRequest asynchronously
// api document: https://help.aliyun.com/api/rds/querypriceforbuy.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) QueryPriceForBuyWithChan(request *QueryPriceForBuyRequest) (<-chan *QueryPriceForBuyResponse, <-chan error) {
	responseChan := make(chan *QueryPriceForBuyResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.QueryPriceForBuy(request)
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

// invoke QueryPriceForBuy api with *QueryPriceForBuyRequest asynchronously
// api document: https://help.aliyun.com/api/rds/querypriceforbuy.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) QueryPriceForBuyWithCallback(request *QueryPriceForBuyRequest, callback func(response *QueryPriceForBuyResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *QueryPriceForBuyResponse
		var err error
		defer close(result)
		response, err = client.QueryPriceForBuy(request)
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

type QueryPriceForBuyRequest struct {
	*requests.RpcRequest
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	ClientToken          string           `position:"Query" name:"ClientToken"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	CommodityCode        string           `position:"Query" name:"CommodityCode"`
	Engine               string           `position:"Query" name:"Engine"`
	EngineVersion        string           `position:"Query" name:"EngineVersion"`
	DBInstanceClass      string           `position:"Query" name:"DBInstanceClass"`
	DBInstanceStorage    requests.Integer `position:"Query" name:"DBInstanceStorage"`
	PayType              string           `position:"Query" name:"PayType"`
	ZoneId               string           `position:"Query" name:"ZoneId"`
	UsedTime             string           `position:"Query" name:"UsedTime"`
	TimeType             string           `position:"Query" name:"TimeType"`
	Quantity             requests.Integer `position:"Query" name:"Quantity"`
	InstanceUsedType     requests.Integer `position:"Query" name:"InstanceUsedType"`
	OrderType            string           `position:"Query" name:"OrderType"`
}

type QueryPriceForBuyResponse struct {
	*responses.BaseResponse
	RequestId string                  `json:"RequestId" xml:"RequestId"`
	PriceInfo PriceInfo               `json:"PriceInfo" xml:"PriceInfo"`
	Rules     RulesInQueryPriceForBuy `json:"Rules" xml:"Rules"`
}

// create a request to invoke QueryPriceForBuy API
func CreateQueryPriceForBuyRequest() (request *QueryPriceForBuyRequest) {
	request = &QueryPriceForBuyRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Rds", "2014-08-15", "QueryPriceForBuy", "rds", "openAPI")
	return
}

// create a response to parse from QueryPriceForBuy response
func CreateQueryPriceForBuyResponse() (response *QueryPriceForBuyResponse) {
	response = &QueryPriceForBuyResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
