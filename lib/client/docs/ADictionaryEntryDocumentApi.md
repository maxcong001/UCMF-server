# \ADictionaryEntryDocumentApi

All URIs are relative to *https://example.com/nucmf-uecm/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateDictionaryEntry**](ADictionaryEntryDocumentApi.md#CreateDictionaryEntry) | **Post** /dic-entries | Create a dictionary entry in the UCMF



## CreateDictionaryEntry

> DicEntryCreatedData CreateDictionaryEntry(ctx).JsonData(jsonData).BinaryDataUeRadioCapability5GS(binaryDataUeRadioCapability5GS).BinaryDataUeRadioCapabilityEPS(binaryDataUeRadioCapabilityEPS).Execute()

Create a dictionary entry in the UCMF

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    jsonData := *openapiclient.NewDicEntryCreateData("TypeAllocationCode_example") // DicEntryCreateData |  (optional)
    binaryDataUeRadioCapability5GS := os.NewFile(1234, "some_file") // *os.File |  (optional)
    binaryDataUeRadioCapabilityEPS := os.NewFile(1234, "some_file") // *os.File |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ADictionaryEntryDocumentApi.CreateDictionaryEntry(context.Background()).JsonData(jsonData).BinaryDataUeRadioCapability5GS(binaryDataUeRadioCapability5GS).BinaryDataUeRadioCapabilityEPS(binaryDataUeRadioCapabilityEPS).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ADictionaryEntryDocumentApi.CreateDictionaryEntry``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateDictionaryEntry`: DicEntryCreatedData
    fmt.Fprintf(os.Stdout, "Response from `ADictionaryEntryDocumentApi.CreateDictionaryEntry`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateDictionaryEntryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **jsonData** | [**DicEntryCreateData**](DicEntryCreateData.md) |  | 
 **binaryDataUeRadioCapability5GS** | ***os.File** |  | 
 **binaryDataUeRadioCapabilityEPS** | ***os.File** |  | 

### Return type

[**DicEntryCreatedData**](DicEntryCreatedData.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: multipart/related
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

