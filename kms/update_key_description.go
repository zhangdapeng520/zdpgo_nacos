package kms

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
	"github.com/zhangdapeng520/zdpgo_nacos/sdk/requests"
	"github.com/zhangdapeng520/zdpgo_nacos/sdk/responses"
)

// UpdateKeyDescription invokes the kms.UpdateKeyDescription API synchronously
// api document: https://help.aliyun.com/api/kms/updatekeydescription.html
func (client *Client) UpdateKeyDescription(request *UpdateKeyDescriptionRequest) (response *UpdateKeyDescriptionResponse, err error) {
	response = CreateUpdateKeyDescriptionResponse()
	err = client.DoAction(request, response)
	return
}

// UpdateKeyDescriptionWithChan invokes the kms.UpdateKeyDescription API asynchronously
// api document: https://help.aliyun.com/api/kms/updatekeydescription.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) UpdateKeyDescriptionWithChan(request *UpdateKeyDescriptionRequest) (<-chan *UpdateKeyDescriptionResponse, <-chan error) {
	responseChan := make(chan *UpdateKeyDescriptionResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.UpdateKeyDescription(request)
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

// UpdateKeyDescriptionWithCallback invokes the kms.UpdateKeyDescription API asynchronously
// api document: https://help.aliyun.com/api/kms/updatekeydescription.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) UpdateKeyDescriptionWithCallback(request *UpdateKeyDescriptionRequest, callback func(response *UpdateKeyDescriptionResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *UpdateKeyDescriptionResponse
		var err error
		defer close(result)
		response, err = client.UpdateKeyDescription(request)
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

// UpdateKeyDescriptionRequest is the request struct for api UpdateKeyDescription
type UpdateKeyDescriptionRequest struct {
	*requests.RpcRequest
	KeyId       string `position:"Query" name:"KeyId"`
	Description string `position:"Query" name:"Description"`
}

// UpdateKeyDescriptionResponse is the response struct for api UpdateKeyDescription
type UpdateKeyDescriptionResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateUpdateKeyDescriptionRequest creates a request to invoke UpdateKeyDescription API
func CreateUpdateKeyDescriptionRequest() (request *UpdateKeyDescriptionRequest) {
	request = &UpdateKeyDescriptionRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Kms", "2016-01-20", "UpdateKeyDescription", "kms", "openAPI")
	return
}

// CreateUpdateKeyDescriptionResponse creates a response to parse from UpdateKeyDescription response
func CreateUpdateKeyDescriptionResponse() (response *UpdateKeyDescriptionResponse) {
	response = &UpdateKeyDescriptionResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
