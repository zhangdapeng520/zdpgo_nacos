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

// GenerateDataKeyWithoutPlaintext invokes the kms.GenerateDataKeyWithoutPlaintext API synchronously
// api document: https://help.aliyun.com/api/kms/generatedatakeywithoutplaintext.html
func (client *Client) GenerateDataKeyWithoutPlaintext(request *GenerateDataKeyWithoutPlaintextRequest) (response *GenerateDataKeyWithoutPlaintextResponse, err error) {
	response = CreateGenerateDataKeyWithoutPlaintextResponse()
	err = client.DoAction(request, response)
	return
}

// GenerateDataKeyWithoutPlaintextWithChan invokes the kms.GenerateDataKeyWithoutPlaintext API asynchronously
// api document: https://help.aliyun.com/api/kms/generatedatakeywithoutplaintext.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) GenerateDataKeyWithoutPlaintextWithChan(request *GenerateDataKeyWithoutPlaintextRequest) (<-chan *GenerateDataKeyWithoutPlaintextResponse, <-chan error) {
	responseChan := make(chan *GenerateDataKeyWithoutPlaintextResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GenerateDataKeyWithoutPlaintext(request)
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

// GenerateDataKeyWithoutPlaintextWithCallback invokes the kms.GenerateDataKeyWithoutPlaintext API asynchronously
// api document: https://help.aliyun.com/api/kms/generatedatakeywithoutplaintext.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) GenerateDataKeyWithoutPlaintextWithCallback(request *GenerateDataKeyWithoutPlaintextRequest, callback func(response *GenerateDataKeyWithoutPlaintextResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GenerateDataKeyWithoutPlaintextResponse
		var err error
		defer close(result)
		response, err = client.GenerateDataKeyWithoutPlaintext(request)
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

// GenerateDataKeyWithoutPlaintextRequest is the request struct for api GenerateDataKeyWithoutPlaintext
type GenerateDataKeyWithoutPlaintextRequest struct {
	*requests.RpcRequest
	EncryptionContext string           `position:"Query" name:"EncryptionContext"`
	KeyId             string           `position:"Query" name:"KeyId"`
	KeySpec           string           `position:"Query" name:"KeySpec"`
	NumberOfBytes     requests.Integer `position:"Query" name:"NumberOfBytes"`
}

// GenerateDataKeyWithoutPlaintextResponse is the response struct for api GenerateDataKeyWithoutPlaintext
type GenerateDataKeyWithoutPlaintextResponse struct {
	*responses.BaseResponse
	CiphertextBlob string `json:"CiphertextBlob" xml:"CiphertextBlob"`
	KeyId          string `json:"KeyId" xml:"KeyId"`
	RequestId      string `json:"RequestId" xml:"RequestId"`
	KeyVersionId   string `json:"KeyVersionId" xml:"KeyVersionId"`
}

// CreateGenerateDataKeyWithoutPlaintextRequest creates a request to invoke GenerateDataKeyWithoutPlaintext API
func CreateGenerateDataKeyWithoutPlaintextRequest() (request *GenerateDataKeyWithoutPlaintextRequest) {
	request = &GenerateDataKeyWithoutPlaintextRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Kms", "2016-01-20", "GenerateDataKeyWithoutPlaintext", "kms", "openAPI")
	return
}

// CreateGenerateDataKeyWithoutPlaintextResponse creates a response to parse from GenerateDataKeyWithoutPlaintext response
func CreateGenerateDataKeyWithoutPlaintextResponse() (response *GenerateDataKeyWithoutPlaintextResponse) {
	response = &GenerateDataKeyWithoutPlaintextResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
