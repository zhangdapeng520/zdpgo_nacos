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

// AsymmetricDecrypt invokes the kms.AsymmetricDecrypt API synchronously
// api document: https://help.aliyun.com/api/kms/asymmetricdecrypt.html
func (client *Client) AsymmetricDecrypt(request *AsymmetricDecryptRequest) (response *AsymmetricDecryptResponse, err error) {
	response = CreateAsymmetricDecryptResponse()
	err = client.DoAction(request, response)
	return
}

// AsymmetricDecryptWithChan invokes the kms.AsymmetricDecrypt API asynchronously
// api document: https://help.aliyun.com/api/kms/asymmetricdecrypt.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) AsymmetricDecryptWithChan(request *AsymmetricDecryptRequest) (<-chan *AsymmetricDecryptResponse, <-chan error) {
	responseChan := make(chan *AsymmetricDecryptResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.AsymmetricDecrypt(request)
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

// AsymmetricDecryptWithCallback invokes the kms.AsymmetricDecrypt API asynchronously
// api document: https://help.aliyun.com/api/kms/asymmetricdecrypt.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) AsymmetricDecryptWithCallback(request *AsymmetricDecryptRequest, callback func(response *AsymmetricDecryptResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *AsymmetricDecryptResponse
		var err error
		defer close(result)
		response, err = client.AsymmetricDecrypt(request)
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

// AsymmetricDecryptRequest is the request struct for api AsymmetricDecrypt
type AsymmetricDecryptRequest struct {
	*requests.RpcRequest
	KeyVersionId   string `position:"Query" name:"KeyVersionId"`
	KeyId          string `position:"Query" name:"KeyId"`
	CiphertextBlob string `position:"Query" name:"CiphertextBlob"`
	Algorithm      string `position:"Query" name:"Algorithm"`
}

// AsymmetricDecryptResponse is the response struct for api AsymmetricDecrypt
type AsymmetricDecryptResponse struct {
	*responses.BaseResponse
	Plaintext    string `json:"Plaintext" xml:"Plaintext"`
	KeyId        string `json:"KeyId" xml:"KeyId"`
	RequestId    string `json:"RequestId" xml:"RequestId"`
	KeyVersionId string `json:"KeyVersionId" xml:"KeyVersionId"`
}

// CreateAsymmetricDecryptRequest creates a request to invoke AsymmetricDecrypt API
func CreateAsymmetricDecryptRequest() (request *AsymmetricDecryptRequest) {
	request = &AsymmetricDecryptRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Kms", "2016-01-20", "AsymmetricDecrypt", "kms", "openAPI")
	return
}

// CreateAsymmetricDecryptResponse creates a response to parse from AsymmetricDecrypt response
func CreateAsymmetricDecryptResponse() (response *AsymmetricDecryptResponse) {
	response = &AsymmetricDecryptResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
