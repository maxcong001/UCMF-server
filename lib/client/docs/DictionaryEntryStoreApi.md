# \DictionaryEntryStoreApi

All URIs are relative to *https://example.com/nucmf-uecm/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**RetrieveDictionaryEntry**](DictionaryEntryStoreApi.md#RetrieveDictionaryEntry) | **Get** /dic-entries | retrieve a dictionary entry matching query parameters



## RetrieveDictionaryEntry

> InlineResponse200 RetrieveDictionaryEntry(ctx).UeRadioCapaId(ueRadioCapaId).RacFormat(racFormat).Execute()

retrieve a dictionary entry matching query parameters

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
    ueRadioCapaId := *openapiclient.NewUeRadioCapaId() // UeRadioCapaId | UE Radio Capability ID, either PLMN Assigned or Manufacturer Assigned
    racFormat := *openapiclient.NewRacFormat() // RacFormat | Encoding format of RAC Info (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DictionaryEntryStoreApi.RetrieveDictionaryEntry(context.Background()).UeRadioCapaId(ueRadioCapaId).RacFormat(racFormat).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DictionaryEntryStoreApi.RetrieveDictionaryEntry``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RetrieveDictionaryEntry`: InlineResponse200
    fmt.Fprintf(os.Stdout, "Response from `DictionaryEntryStoreApi.RetrieveDictionaryEntry`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRetrieveDictionaryEntryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ueRadioCapaId** | [**UeRadioCapaId**](UeRadioCapaId.md) | UE Radio Capability ID, either PLMN Assigned or Manufacturer Assigned | 
 **racFormat** | [**RacFormat**](RacFormat.md) | Encoding format of RAC Info | 

### Return type

[**InlineResponse200**](inline_response_200.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: multipart/related, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

