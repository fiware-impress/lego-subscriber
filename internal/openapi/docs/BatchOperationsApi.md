# \BatchOperationsApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**BatchEntityCreation**](BatchOperationsApi.md#BatchEntityCreation) | **Post** /entityOperations/create | 
[**BatchEntityDelete**](BatchOperationsApi.md#BatchEntityDelete) | **Post** /entityOperations/delete | 
[**BatchEntityUpdate**](BatchOperationsApi.md#BatchEntityUpdate) | **Post** /entityOperations/update | 
[**BatchEntityUpsert**](BatchOperationsApi.md#BatchEntityUpsert) | **Post** /entityOperations/upsert | 



## BatchEntityCreation

> BatchOperationResult BatchEntityCreation(ctx).Entity(entity).Execute()





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
    entity := []openapiclient.Entity{*openapiclient.NewEntity()} // []Entity | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BatchOperationsApi.BatchEntityCreation(context.Background()).Entity(entity).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BatchOperationsApi.BatchEntityCreation``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BatchEntityCreation`: BatchOperationResult
    fmt.Fprintf(os.Stdout, "Response from `BatchOperationsApi.BatchEntityCreation`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBatchEntityCreationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **entity** | [**[]Entity**](Entity.md) |  | 

### Return type

[**BatchOperationResult**](BatchOperationResult.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/ld+json
- **Accept**: application/ld+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BatchEntityDelete

> BatchOperationResult BatchEntityDelete(ctx).RequestBody(requestBody).Execute()





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
    requestBody := []string{"Property_example"} // []string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BatchOperationsApi.BatchEntityDelete(context.Background()).RequestBody(requestBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BatchOperationsApi.BatchEntityDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BatchEntityDelete`: BatchOperationResult
    fmt.Fprintf(os.Stdout, "Response from `BatchOperationsApi.BatchEntityDelete`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBatchEntityDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **requestBody** | **[]string** |  | 

### Return type

[**BatchOperationResult**](BatchOperationResult.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/ld+json
- **Accept**: application/ld+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BatchEntityUpdate

> BatchOperationResult BatchEntityUpdate(ctx).Entity(entity).Options(options).Execute()





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
    entity := []openapiclient.Entity{*openapiclient.NewEntity()} // []Entity | 
    options := "options_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BatchOperationsApi.BatchEntityUpdate(context.Background()).Entity(entity).Options(options).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BatchOperationsApi.BatchEntityUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BatchEntityUpdate`: BatchOperationResult
    fmt.Fprintf(os.Stdout, "Response from `BatchOperationsApi.BatchEntityUpdate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBatchEntityUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **entity** | [**[]Entity**](Entity.md) |  | 
 **options** | **string** |  | 

### Return type

[**BatchOperationResult**](BatchOperationResult.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/ld+json
- **Accept**: application/ld+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BatchEntityUpsert

> BatchOperationResult BatchEntityUpsert(ctx).Entity(entity).Options(options).Execute()





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
    entity := []openapiclient.Entity{*openapiclient.NewEntity()} // []Entity | 
    options := "options_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BatchOperationsApi.BatchEntityUpsert(context.Background()).Entity(entity).Options(options).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BatchOperationsApi.BatchEntityUpsert``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BatchEntityUpsert`: BatchOperationResult
    fmt.Fprintf(os.Stdout, "Response from `BatchOperationsApi.BatchEntityUpsert`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBatchEntityUpsertRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **entity** | [**[]Entity**](Entity.md) |  | 
 **options** | **string** |  | 

### Return type

[**BatchOperationResult**](BatchOperationResult.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/ld+json
- **Accept**: application/ld+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

