# \ContextSourcesApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateCSourceSubscription**](ContextSourcesApi.md#CreateCSourceSubscription) | **Post** /csourceSubscriptions/ | 
[**QueryCsources**](ContextSourcesApi.md#QueryCsources) | **Get** /csourceRegistrations/ | 
[**RegisterCsource**](ContextSourcesApi.md#RegisterCsource) | **Post** /csourceRegistrations/ | 
[**RemoveCSourceSubscription**](ContextSourcesApi.md#RemoveCSourceSubscription) | **Delete** /csourceSubscriptions/{subscriptionId} | 
[**RemoveCsource**](ContextSourcesApi.md#RemoveCsource) | **Delete** /csourceRegistrations/{registrationId} | 
[**RetrieveCSourceSubscriptions**](ContextSourcesApi.md#RetrieveCSourceSubscriptions) | **Get** /csourceSubscriptions/ | 
[**RetrieveCSourceSubscriptionsById**](ContextSourcesApi.md#RetrieveCSourceSubscriptionsById) | **Get** /csourceSubscriptions/{subscriptionId} | 
[**RetrieveCsource**](ContextSourcesApi.md#RetrieveCsource) | **Get** /csourceRegistrations/{registrationId} | 
[**UpdateCSourceSubscription**](ContextSourcesApi.md#UpdateCSourceSubscription) | **Patch** /csourceSubscriptions/{subscriptionId} | 



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
    resp, r, err := apiClient.ContextSourcesApi.CreateCSourceSubscription(context.Background()).Subscription(subscription).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContextSourcesApi.CreateCSourceSubscription``: %v\n", err)
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


## QueryCsources

> []ContextSourceRegistration QueryCsources(ctx).Id(id).IdPattern(idPattern).Type_(type_).Attrs(attrs).Q(q).Georel(georel).Geometry(geometry).Coordinates(coordinates).Geoproperty(geoproperty).Limit(limit).Execute()





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
    id := "id_example" // string | Comma separated list of URIs to be retrieved (optional)
    idPattern := "idPattern_example" // string | Regular expression that must be matched by Entity ids (optional)
    type_ := "type__example" // string | Comma separated list of Entity type names to be retrieved (optional)
    attrs := "attrs_example" // string | Comma separated list of attribute names (properties or relationships) to be retrieved (optional)
    q := "q_example" // string | Query (optional)
    georel := "georel_example" // string | Geo-relationship (optional)
    geometry := "geometry_example" // string | Geometry (optional)
    coordinates := "coordinates_example" // string | Coordinates serialized as a string (optional)
    geoproperty := "geoproperty_example" // string | The name of the property that contains the geo-spatial data that will be used to resolve the geoquery (optional)
    limit := int32(56) // int32 | Pagination limit (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ContextSourcesApi.QueryCsources(context.Background()).Id(id).IdPattern(idPattern).Type_(type_).Attrs(attrs).Q(q).Georel(georel).Geometry(geometry).Coordinates(coordinates).Geoproperty(geoproperty).Limit(limit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContextSourcesApi.QueryCsources``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `QueryCsources`: []ContextSourceRegistration
    fmt.Fprintf(os.Stdout, "Response from `ContextSourcesApi.QueryCsources`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiQueryCsourcesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string** | Comma separated list of URIs to be retrieved | 
 **idPattern** | **string** | Regular expression that must be matched by Entity ids | 
 **type_** | **string** | Comma separated list of Entity type names to be retrieved | 
 **attrs** | **string** | Comma separated list of attribute names (properties or relationships) to be retrieved | 
 **q** | **string** | Query | 
 **georel** | **string** | Geo-relationship | 
 **geometry** | **string** | Geometry | 
 **coordinates** | **string** | Coordinates serialized as a string | 
 **geoproperty** | **string** | The name of the property that contains the geo-spatial data that will be used to resolve the geoquery | 
 **limit** | **int32** | Pagination limit | 

### Return type

[**[]ContextSourceRegistration**](ContextSourceRegistration.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/ld+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RegisterCsource

> RegisterCsource(ctx).ContextSourceRegistration(contextSourceRegistration).Execute()





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
    contextSourceRegistration := *openapiclient.NewContextSourceRegistration([]openapiclient.RegistrationInfo{*openapiclient.NewRegistrationInfo()}, "Endpoint_example", "Id_example", "Type_example") // ContextSourceRegistration | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ContextSourcesApi.RegisterCsource(context.Background()).ContextSourceRegistration(contextSourceRegistration).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContextSourcesApi.RegisterCsource``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRegisterCsourceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **contextSourceRegistration** | [**ContextSourceRegistration**](ContextSourceRegistration.md) |  | 

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
    resp, r, err := apiClient.ContextSourcesApi.RemoveCSourceSubscription(context.Background(), subscriptionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContextSourcesApi.RemoveCSourceSubscription``: %v\n", err)
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


## RemoveCsource

> RemoveCsource(ctx, registrationId).Execute()





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
    registrationId := "registrationId_example" // string | Registration Id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ContextSourcesApi.RemoveCsource(context.Background(), registrationId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContextSourcesApi.RemoveCsource``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**registrationId** | **string** | Registration Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemoveCsourceRequest struct via the builder pattern


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
    resp, r, err := apiClient.ContextSourcesApi.RetrieveCSourceSubscriptions(context.Background()).Limit(limit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContextSourcesApi.RetrieveCSourceSubscriptions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RetrieveCSourceSubscriptions`: []Subscription
    fmt.Fprintf(os.Stdout, "Response from `ContextSourcesApi.RetrieveCSourceSubscriptions`: %v\n", resp)
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
    resp, r, err := apiClient.ContextSourcesApi.RetrieveCSourceSubscriptionsById(context.Background(), subscriptionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContextSourcesApi.RetrieveCSourceSubscriptionsById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RetrieveCSourceSubscriptionsById`: Subscription
    fmt.Fprintf(os.Stdout, "Response from `ContextSourcesApi.RetrieveCSourceSubscriptionsById`: %v\n", resp)
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


## RetrieveCsource

> ContextSourceRegistration RetrieveCsource(ctx, registrationId).Execute()





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
    registrationId := "registrationId_example" // string | Registration Id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ContextSourcesApi.RetrieveCsource(context.Background(), registrationId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContextSourcesApi.RetrieveCsource``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RetrieveCsource`: ContextSourceRegistration
    fmt.Fprintf(os.Stdout, "Response from `ContextSourcesApi.RetrieveCsource`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**registrationId** | **string** | Registration Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiRetrieveCsourceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ContextSourceRegistration**](ContextSourceRegistration.md)

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
    resp, r, err := apiClient.ContextSourcesApi.UpdateCSourceSubscription(context.Background(), subscriptionId).SubscriptionFragment(subscriptionFragment).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContextSourcesApi.UpdateCSourceSubscription``: %v\n", err)
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

