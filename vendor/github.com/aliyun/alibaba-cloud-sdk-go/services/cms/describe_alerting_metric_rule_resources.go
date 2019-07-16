package cms

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

// DescribeAlertingMetricRuleResources invokes the cms.DescribeAlertingMetricRuleResources API synchronously
// api document: https://help.aliyun.com/api/cms/describealertingmetricruleresources.html
func (client *Client) DescribeAlertingMetricRuleResources(request *DescribeAlertingMetricRuleResourcesRequest) (response *DescribeAlertingMetricRuleResourcesResponse, err error) {
	response = CreateDescribeAlertingMetricRuleResourcesResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeAlertingMetricRuleResourcesWithChan invokes the cms.DescribeAlertingMetricRuleResources API asynchronously
// api document: https://help.aliyun.com/api/cms/describealertingmetricruleresources.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeAlertingMetricRuleResourcesWithChan(request *DescribeAlertingMetricRuleResourcesRequest) (<-chan *DescribeAlertingMetricRuleResourcesResponse, <-chan error) {
	responseChan := make(chan *DescribeAlertingMetricRuleResourcesResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeAlertingMetricRuleResources(request)
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

// DescribeAlertingMetricRuleResourcesWithCallback invokes the cms.DescribeAlertingMetricRuleResources API asynchronously
// api document: https://help.aliyun.com/api/cms/describealertingmetricruleresources.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeAlertingMetricRuleResourcesWithCallback(request *DescribeAlertingMetricRuleResourcesRequest, callback func(response *DescribeAlertingMetricRuleResourcesResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeAlertingMetricRuleResourcesResponse
		var err error
		defer close(result)
		response, err = client.DescribeAlertingMetricRuleResources(request)
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

// DescribeAlertingMetricRuleResourcesRequest is the request struct for api DescribeAlertingMetricRuleResources
type DescribeAlertingMetricRuleResourcesRequest struct {
	*requests.RpcRequest
	GroupId string `position:"Query" name:"GroupId"`
	RuleId  string `position:"Query" name:"RuleId"`
}

// DescribeAlertingMetricRuleResourcesResponse is the response struct for api DescribeAlertingMetricRuleResources
type DescribeAlertingMetricRuleResourcesResponse struct {
	*responses.BaseResponse
	RequestId string                                         `json:"RequestId" xml:"RequestId"`
	Success   bool                                           `json:"Success" xml:"Success"`
	Code      int                                            `json:"Code" xml:"Code"`
	Message   string                                         `json:"Message" xml:"Message"`
	Total     int                                            `json:"Total" xml:"Total"`
	Resources ResourcesInDescribeAlertingMetricRuleResources `json:"Resources" xml:"Resources"`
}

// CreateDescribeAlertingMetricRuleResourcesRequest creates a request to invoke DescribeAlertingMetricRuleResources API
func CreateDescribeAlertingMetricRuleResourcesRequest() (request *DescribeAlertingMetricRuleResourcesRequest) {
	request = &DescribeAlertingMetricRuleResourcesRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Cms", "2019-01-01", "DescribeAlertingMetricRuleResources", "cms", "openAPI")
	return
}

// CreateDescribeAlertingMetricRuleResourcesResponse creates a response to parse from DescribeAlertingMetricRuleResources response
func CreateDescribeAlertingMetricRuleResourcesResponse() (response *DescribeAlertingMetricRuleResourcesResponse) {
	response = &DescribeAlertingMetricRuleResourcesResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}