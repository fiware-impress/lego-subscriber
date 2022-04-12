# \CSourceSubscriptionsApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateCSourceSubscription**](CSourceSubscriptionsApi.md#CreateCSourceSubscription) | **Post** /csourceSubscriptions/ | 
[**RemoveCSourceSubscription**](CSourceSubscriptionsApi.md#RemoveCSourceSubscription) | **Delete** /csourceSubscriptions/{subscriptionId} | 
[**RetrieveCSourceSubscriptions**](CSourceSubscriptionsApi.md#RetrieveCSourceSubscriptions) | **Get** /csourceSubscriptions/ | 
[**RetrieveCSourceSubscriptionsById**](CSourceSubscriptionsApi.md#RetrieveCSourceSubscriptionsById) | **Get** /csourceSubscriptions/{subscriptionId} | 
[**UpdateCSourceSubscription**](CSourceSubscriptionsApi.md#UpdateCSourceSubscription) | **Patch** /csourceSubscriptions/{subscriptionId} | 



## CreateCSourceSubscription

> CreateCSourceSubscription(ctx).Subscription(subscription).Execute()





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
    subscription := *openapiclient.NewSubscription() // Subscription | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CSourceSubscriptionsApi.CreateCSourceSubscription(context.Background()).Subscription(subscription).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CSourceSubscriptionsApi.CreateCSourceSubscription``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateCSourceSubscriptionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **subscription** | [**Subscription**](Subscription.md) |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/ld+json
- **Accept**: application/ld+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RemoveCSourceSubscription

> RemoveCSourceSubscription(ctx, subscriptionId).Execute()





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
    subscriptionId := "subscriptionId_example" // string | Subscription Id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CSourceSubscriptionsApi.RemoveCSourceSubscription(context.Background(), subscriptionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CSourceSubscriptionsApi.RemoveCSourceSubscription``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**subscriptionId** | **string** | Subscription Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemoveCSourceSubscriptionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/ld+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RetrieveCSourceSubscriptions

> []Subscription RetrieveCSourceSubscriptions(ctx).Limit(limit).Execute()





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
    limit := int32(56) // int32 | Pagination limit (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CSourceSubscriptionsApi.RetrieveCSourceSubscriptions(context.Background()).Limit(limit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CSourceSubscriptionsApi.RetrieveCSourceSubscriptions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RetrieveCSourceSubscriptions`: []Subscription
    fmt.Fprintf(os.Stdout, "Response from `CSourceSubscriptionsApi.RetrieveCSourceSubscriptions`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRetrieveCSourceSubscriptionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** | Pagination limit | 

### Return type

[**[]Subscription**](Subscription.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/ld+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RetrieveCSourceSubscriptionsById

> Subscription RetrieveCSourceSubscriptionsById(ctx, subscriptionId).Execute()





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
    subscriptionId := "subscriptionId_example" // string | Subscription Id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CSourceSubscriptionsApi.RetrieveCSourceSubscriptionsById(context.Background(), subscriptionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CSourceSubscriptionsApi.RetrieveCSourceSubscriptionsById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RetrieveCSourceSubscriptionsById`: Subscription
    fmt.Fprintf(os.Stdout, "Response from `CSourceSubscriptionsApi.RetrieveCSourceSubscriptionsById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**subscriptionId** | **string** | Subscription Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiRetrieveCSourceSubscriptionsByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Subscription**](Subscription.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/ld+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateCSourceSubscription

> UpdateCSourceSubscription(ctx, subscriptionId).SubscriptionFragment(subscriptionFragment).Execute()





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
    subscriptionId := "subscriptionId_example" // string | Subscription Id
    subscriptionFragment := *openapiclient.NewSubscriptionFragment() // SubscriptionFragment | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CSourceSubscriptionsApi.UpdateCSourceSubscription(context.Background(), subscriptionId).SubscriptionFragment(subscriptionFragment).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CSourceSubscriptionsApi.UpdateCSourceSubscription``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**subscriptionId** | **string** | Subscription Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateCSourceSubscriptionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **subscriptionFragment** | [**SubscriptionFragment**](SubscriptionFragment.md) |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/ld+json
- **Accept**: application/ld+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

