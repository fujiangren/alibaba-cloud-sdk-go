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

// invoke DescribeLoadBalancerTCPListenerAttribute api with *DescribeLoadBalancerTCPListenerAttributeRequest synchronously
// api document: https://help.aliyun.com/api/slb/describeloadbalancertcplistenerattribute.html
func (client *Client) DescribeLoadBalancerTCPListenerAttribute(request *DescribeLoadBalancerTCPListenerAttributeRequest) (response *DescribeLoadBalancerTCPListenerAttributeResponse, err error) {
	response = CreateDescribeLoadBalancerTCPListenerAttributeResponse()
	err = client.DoAction(request, response)
	return
}

// invoke DescribeLoadBalancerTCPListenerAttribute api with *DescribeLoadBalancerTCPListenerAttributeRequest asynchronously
// api document: https://help.aliyun.com/api/slb/describeloadbalancertcplistenerattribute.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeLoadBalancerTCPListenerAttributeWithChan(request *DescribeLoadBalancerTCPListenerAttributeRequest) (<-chan *DescribeLoadBalancerTCPListenerAttributeResponse, <-chan error) {
	responseChan := make(chan *DescribeLoadBalancerTCPListenerAttributeResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeLoadBalancerTCPListenerAttribute(request)
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

// invoke DescribeLoadBalancerTCPListenerAttribute api with *DescribeLoadBalancerTCPListenerAttributeRequest asynchronously
// api document: https://help.aliyun.com/api/slb/describeloadbalancertcplistenerattribute.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeLoadBalancerTCPListenerAttributeWithCallback(request *DescribeLoadBalancerTCPListenerAttributeRequest, callback func(response *DescribeLoadBalancerTCPListenerAttributeResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeLoadBalancerTCPListenerAttributeResponse
		var err error
		defer close(result)
		response, err = client.DescribeLoadBalancerTCPListenerAttribute(request)
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

type DescribeLoadBalancerTCPListenerAttributeRequest struct {
	*requests.RpcRequest
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	LoadBalancerId       string           `position:"Query" name:"LoadBalancerId"`
	ListenerPort         requests.Integer `position:"Query" name:"ListenerPort"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	AccessKeyId          string           `position:"Query" name:"access_key_id"`
	Tags                 string           `position:"Query" name:"Tags"`
}

type DescribeLoadBalancerTCPListenerAttributeResponse struct {
	*responses.BaseResponse
	RequestId                 string `json:"RequestId" xml:"RequestId"`
	ListenerPort              int    `json:"ListenerPort" xml:"ListenerPort"`
	BackendServerPort         int    `json:"BackendServerPort" xml:"BackendServerPort"`
	Status                    string `json:"Status" xml:"Status"`
	Bandwidth                 int    `json:"Bandwidth" xml:"Bandwidth"`
	Scheduler                 string `json:"Scheduler" xml:"Scheduler"`
	SynProxy                  string `json:"SynProxy" xml:"SynProxy"`
	PersistenceTimeout        int    `json:"PersistenceTimeout" xml:"PersistenceTimeout"`
	EstablishedTimeout        int    `json:"EstablishedTimeout" xml:"EstablishedTimeout"`
	HealthCheck               string `json:"HealthCheck" xml:"HealthCheck"`
	HealthyThreshold          int    `json:"HealthyThreshold" xml:"HealthyThreshold"`
	UnhealthyThreshold        int    `json:"UnhealthyThreshold" xml:"UnhealthyThreshold"`
	HealthCheckConnectTimeout int    `json:"HealthCheckConnectTimeout" xml:"HealthCheckConnectTimeout"`
	HealthCheckConnectPort    int    `json:"HealthCheckConnectPort" xml:"HealthCheckConnectPort"`
	HealthCheckInterval       int    `json:"HealthCheckInterval" xml:"HealthCheckInterval"`
	HealthCheckHttpCode       string `json:"HealthCheckHttpCode" xml:"HealthCheckHttpCode"`
	HealthCheckDomain         string `json:"HealthCheckDomain" xml:"HealthCheckDomain"`
	HealthCheckURI            string `json:"HealthCheckURI" xml:"HealthCheckURI"`
	HealthCheckType           string `json:"HealthCheckType" xml:"HealthCheckType"`
	MaxConnection             int    `json:"MaxConnection" xml:"MaxConnection"`
	VServerGroupId            string `json:"VServerGroupId" xml:"VServerGroupId"`
	MasterSlaveServerGroupId  string `json:"MasterSlaveServerGroupId" xml:"MasterSlaveServerGroupId"`
}

// create a request to invoke DescribeLoadBalancerTCPListenerAttribute API
func CreateDescribeLoadBalancerTCPListenerAttributeRequest() (request *DescribeLoadBalancerTCPListenerAttributeRequest) {
	request = &DescribeLoadBalancerTCPListenerAttributeRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Slb", "2014-05-15", "DescribeLoadBalancerTCPListenerAttribute", "slb", "openAPI")
	return
}

// create a response to parse from DescribeLoadBalancerTCPListenerAttribute response
func CreateDescribeLoadBalancerTCPListenerAttributeResponse() (response *DescribeLoadBalancerTCPListenerAttributeResponse) {
	response = &DescribeLoadBalancerTCPListenerAttributeResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
