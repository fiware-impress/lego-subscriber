# \TemporalRetrievalApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**QueryTemporalEntities**](TemporalRetrievalApi.md#QueryTemporalEntities) | **Get** /temporal/entities/ | 
[**QueryTemporalEntitiesOnPost**](TemporalRetrievalApi.md#QueryTemporalEntitiesOnPost) | **Post** /temporal/entityOperations/query | 
[**RetrieveEntityTemporalById**](TemporalRetrievalApi.md#RetrieveEntityTemporalById) | **Get** /temporal/entities/{entityId} | 



## QueryTemporalEntities

> []EntityTemporal QueryTemporalEntities(ctx).Link(link).Id(id).IdPattern(idPattern).Type_(type_).Attrs(attrs).Q(q).Georel(georel).Geometry(geometry).Coordinates(coordinates).Geoproperty(geoproperty).Timerel(timerel).Timeproperty(timeproperty).TimeAt(timeAt).EndTimeAt(endTimeAt).Csf(csf).PageSize(pageSize).PageAnchor(pageAnchor).Limit(limit).Options(options).LastN(lastN).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "time"
    openapiclient "./openapi"
)

func main() {
    link := "link_example" // string | Link header to be used as described in the json-ld spec. (optional)
    id := "id_example" // string | Comma separated list of URIs to be retrieved (optional)
    idPattern := "idPattern_example" // string | Regular expression that must be matched by Entity ids (optional)
    type_ := "type__example" // string | Comma separated list of Entity type names to be retrieved (optional)
    attrs := "attrs_example" // string | Comma separated list of attribute names (properties or relationships) to be retrieved (optional)
    q := "q_example" // string | Query (optional)
    georel := "georel_example" // string | Geo-relationship (optional)
    geometry := "geometry_example" // string | Geometry (optional)
    coordinates := "coordinates_example" // string | Coordinates serialized as a string (optional)
    geoproperty := "geoproperty_example" // string | The name of the property that contains the geo-spatial data that will be used to resolve the geoquery (optional)
    timerel := openapiclient.timerel("before") // Timerel | Time relationship (optional)
    timeproperty := "timeproperty_example" // string | The name of the property that contains the temporal data that will be used to resolve the temporal query (optional)
    timeAt := time.Now() // time.Time | start time for temporal query (optional)
    endTimeAt := time.Now() // time.Time | end time for temporal query (optional)
    csf := "csf_example" // string | Context Source Filter (optional)
    pageSize := int32(56) // int32 | Size of the page to be returned (optional)
    pageAnchor := "pageAnchor_example" // string | Size of the page to be returned (optional)
    limit := int32(56) // int32 | Pagination limit (optional)
    options := "options_example" // string | Options dictionary (optional)
    lastN := int32(56) // int32 | Only retrieve last N instances (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TemporalRetrievalApi.QueryTemporalEntities(context.Background()).Link(link).Id(id).IdPattern(idPattern).Type_(type_).Attrs(attrs).Q(q).Georel(georel).Geometry(geometry).Coordinates(coordinates).Geoproperty(geoproperty).Timerel(timerel).Timeproperty(timeproperty).TimeAt(timeAt).EndTimeAt(endTimeAt).Csf(csf).PageSize(pageSize).PageAnchor(pageAnchor).Limit(limit).Options(options).LastN(lastN).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TemporalRetrievalApi.QueryTemporalEntities``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `QueryTemporalEntities`: []EntityTemporal
    fmt.Fprintf(os.Stdout, "Response from `TemporalRetrievalApi.QueryTemporalEntities`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiQueryTemporalEntitiesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **link** | **string** | Link header to be used as described in the json-ld spec. | 
 **id** | **string** | Comma separated list of URIs to be retrieved | 
 **idPattern** | **string** | Regular expression that must be matched by Entity ids | 
 **type_** | **string** | Comma separated list of Entity type names to be retrieved | 
 **attrs** | **string** | Comma separated list of attribute names (properties or relationships) to be retrieved | 
 **q** | **string** | Query | 
 **georel** | **string** | Geo-relationship | 
 **geometry** | **string** | Geometry | 
 **coordinates** | **string** | Coordinates serialized as a string | 
 **geoproperty** | **string** | The name of the property that contains the geo-spatial data that will be used to resolve the geoquery | 
 **timerel** | [**Timerel**](Timerel.md) | Time relationship | 
 **timeproperty** | **string** | The name of the property that contains the temporal data that will be used to resolve the temporal query | 
 **timeAt** | **time.Time** | start time for temporal query | 
 **endTimeAt** | **time.Time** | end time for temporal query | 
 **csf** | **string** | Context Source Filter | 
 **pageSize** | **int32** | Size of the page to be returned | 
 **pageAnchor** | **string** | Size of the page to be returned | 
 **limit** | **int32** | Pagination limit | 
 **options** | **string** | Options dictionary | 
 **lastN** | **int32** | Only retrieve last N instances | 

### Return type

[**[]EntityTemporal**](EntityTemporal.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/ld+json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## QueryTemporalEntitiesOnPost

> []EntityTemporal QueryTemporalEntitiesOnPost(ctx).Query(query).Link(link).PageSize(pageSize).PageAnchor(pageAnchor).Limit(limit).Options(options).LastN(lastN).Execute()





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
    query := *openapiclient.NewQuery() // Query | 
    link := "link_example" // string | Link header to be used as described in the json-ld spec. (optional)
    pageSize := int32(56) // int32 | Size of the page to be returned (optional)
    pageAnchor := "pageAnchor_example" // string | Size of the page to be returned (optional)
    limit := int32(56) // int32 | Pagination limit (optional)
    options := "options_example" // string | Options dictionary (optional)
    lastN := int32(56) // int32 | Only retrieve last N instances (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TemporalRetrievalApi.QueryTemporalEntitiesOnPost(context.Background()).Query(query).Link(link).PageSize(pageSize).PageAnchor(pageAnchor).Limit(limit).Options(options).LastN(lastN).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TemporalRetrievalApi.QueryTemporalEntitiesOnPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `QueryTemporalEntitiesOnPost`: []EntityTemporal
    fmt.Fprintf(os.Stdout, "Response from `TemporalRetrievalApi.QueryTemporalEntitiesOnPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiQueryTemporalEntitiesOnPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **query** | [**Query**](Query.md) |  | 
 **link** | **string** | Link header to be used as described in the json-ld spec. | 
 **pageSize** | **int32** | Size of the page to be returned | 
 **pageAnchor** | **string** | Size of the page to be returned | 
 **limit** | **int32** | Pagination limit | 
 **options** | **string** | Options dictionary | 
 **lastN** | **int32** | Only retrieve last N instances | 

### Return type

[**[]EntityTemporal**](EntityTemporal.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/ld+json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RetrieveEntityTemporalById

> EntityTemporal RetrieveEntityTemporalById(ctx, entityId).Link(link).Attrs(attrs).Options(options).Timerel(timerel).Timeproperty(timeproperty).TimeAt(timeAt).EndTimeAt(endTimeAt).LastN(lastN).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "time"
    openapiclient "./openapi"
)

func main() {
    entityId := "entityId_example" // string | Entity Id
    link := "link_example" // string | Link header to be used as described in the json-ld spec. (optional)
    attrs := "attrs_example" // string | Comma separated list of attribute names (properties or relationships) to be retrieved (optional)
    options := "options_example" // string | Options dictionary (optional)
    timerel := openapiclient.timerel("before") // Timerel | Time relationship (optional)
    timeproperty := "timeproperty_example" // string | The name of the property that contains the temporal data that will be used to resolve the temporal query (optional)
    timeAt := time.Now() // time.Time | start time for temporal query (optional)
    endTimeAt := time.Now() // time.Time | end time for temporal query (optional)
    lastN := int32(56) // int32 | Only retrieve last N instances (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TemporalRetrievalApi.RetrieveEntityTemporalById(context.Background(), entityId).Link(link).Attrs(attrs).Options(options).Timerel(timerel).Timeproperty(timeproperty).TimeAt(timeAt).EndTimeAt(endTimeAt).LastN(lastN).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TemporalRetrievalApi.RetrieveEntityTemporalById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RetrieveEntityTemporalById`: EntityTemporal
    fmt.Fprintf(os.Stdout, "Response from `TemporalRetrievalApi.RetrieveEntityTemporalById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**entityId** | **string** | Entity Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiRetrieveEntityTemporalByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **link** | **string** | Link header to be used as described in the json-ld spec. | 
 **attrs** | **string** | Comma separated list of attribute names (properties or relationships) to be retrieved | 
 **options** | **string** | Options dictionary | 
 **timerel** | [**Timerel**](Timerel.md) | Time relationship | 
 **timeproperty** | **string** | The name of the property that contains the temporal data that will be used to resolve the temporal query | 
 **timeAt** | **time.Time** | start time for temporal query | 
 **endTimeAt** | **time.Time** | end time for temporal query | 
 **lastN** | **int32** | Only retrieve last N instances | 

### Return type

[**EntityTemporal**](EntityTemporal.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/ld+json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

