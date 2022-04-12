# \EntitiesApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AppendEntityAttrs**](EntitiesApi.md#AppendEntityAttrs) | **Post** /entities/{entityId}/attrs/ | 
[**BatchEntityCreation**](EntitiesApi.md#BatchEntityCreation) | **Post** /entityOperations/create | 
[**BatchEntityDelete**](EntitiesApi.md#BatchEntityDelete) | **Post** /entityOperations/delete | 
[**BatchEntityUpdate**](EntitiesApi.md#BatchEntityUpdate) | **Post** /entityOperations/update | 
[**BatchEntityUpsert**](EntitiesApi.md#BatchEntityUpsert) | **Post** /entityOperations/upsert | 
[**CreateEntity**](EntitiesApi.md#CreateEntity) | **Post** /entities/ | 
[**PartialAttrUpdate**](EntitiesApi.md#PartialAttrUpdate) | **Patch** /entities/{entityId}/attrs/{attrId} | 
[**QueryEntities**](EntitiesApi.md#QueryEntities) | **Get** /entities/ | 
[**RemoveEntityAttr**](EntitiesApi.md#RemoveEntityAttr) | **Delete** /entities/{entityId}/attrs/{attrId} | 
[**RemoveEntityById**](EntitiesApi.md#RemoveEntityById) | **Delete** /entities/{entityId} | 
[**RetrieveEntityById**](EntitiesApi.md#RetrieveEntityById) | **Get** /entities/{entityId} | 
[**UpdateEntityAttrs**](EntitiesApi.md#UpdateEntityAttrs) | **Patch** /entities/{entityId}/attrs/ | 



## AppendEntityAttrs

> AppendEntityAttrs(ctx, entityId).NGSILDTenant(nGSILDTenant).RequestBody(requestBody).Options(options).Execute()





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
    nGSILDTenant := "nGSILDTenant_example" // string |  (default to "orion")
    entityId := "entityId_example" // string | Entity Id
    requestBody := map[string]OneOfPropertyRelationshipGeoProperty{"key": "TODO"} // map[string]OneOfPropertyRelationshipGeoProperty | 
    options := "options_example" // string | Indicates that no attribute overwrite shall be performed (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EntitiesApi.AppendEntityAttrs(context.Background(), entityId).NGSILDTenant(nGSILDTenant).RequestBody(requestBody).Options(options).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EntitiesApi.AppendEntityAttrs``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**entityId** | **string** | Entity Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiAppendEntityAttrsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **nGSILDTenant** | **string** |  | [default to &quot;orion&quot;]

 **requestBody** | [**map[string]OneOfPropertyRelationshipGeoProperty**](oneOf&lt;Property,Relationship,GeoProperty&gt;.md) |  | 
 **options** | **string** | Indicates that no attribute overwrite shall be performed | 

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
    resp, r, err := apiClient.EntitiesApi.BatchEntityCreation(context.Background()).Entity(entity).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EntitiesApi.BatchEntityCreation``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BatchEntityCreation`: BatchOperationResult
    fmt.Fprintf(os.Stdout, "Response from `EntitiesApi.BatchEntityCreation`: %v\n", resp)
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
    resp, r, err := apiClient.EntitiesApi.BatchEntityDelete(context.Background()).RequestBody(requestBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EntitiesApi.BatchEntityDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BatchEntityDelete`: BatchOperationResult
    fmt.Fprintf(os.Stdout, "Response from `EntitiesApi.BatchEntityDelete`: %v\n", resp)
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
    resp, r, err := apiClient.EntitiesApi.BatchEntityUpdate(context.Background()).Entity(entity).Options(options).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EntitiesApi.BatchEntityUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BatchEntityUpdate`: BatchOperationResult
    fmt.Fprintf(os.Stdout, "Response from `EntitiesApi.BatchEntityUpdate`: %v\n", resp)
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
    resp, r, err := apiClient.EntitiesApi.BatchEntityUpsert(context.Background()).Entity(entity).Options(options).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EntitiesApi.BatchEntityUpsert``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BatchEntityUpsert`: BatchOperationResult
    fmt.Fprintf(os.Stdout, "Response from `EntitiesApi.BatchEntityUpsert`: %v\n", resp)
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


## CreateEntity

> CreateEntity(ctx).NGSILDTenant(nGSILDTenant).RequestBody(requestBody).Execute()





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
    nGSILDTenant := "nGSILDTenant_example" // string |  (default to "orion")
    requestBody := map[string]OneOfPropertyRelationshipGeoProperty{"key": "TODO"} // map[string]OneOfPropertyRelationshipGeoProperty | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EntitiesApi.CreateEntity(context.Background()).NGSILDTenant(nGSILDTenant).RequestBody(requestBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EntitiesApi.CreateEntity``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateEntityRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **nGSILDTenant** | **string** |  | [default to &quot;orion&quot;]
 **requestBody** | [**map[string]OneOfPropertyRelationshipGeoProperty**](oneOf&lt;Property,Relationship,GeoProperty&gt;.md) |  | 

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


## PartialAttrUpdate

> PartialAttrUpdate(ctx, entityId, attrId).NGSILDTenant(nGSILDTenant).RequestBody(requestBody).Execute()





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
    nGSILDTenant := "nGSILDTenant_example" // string |  (default to "orion")
    entityId := "entityId_example" // string | Entity Id
    attrId := "attrId_example" // string | Attribute Id
    requestBody := map[string]OneOfPropertyRelationshipGeoProperty{"key": "TODO"} // map[string]OneOfPropertyRelationshipGeoProperty | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EntitiesApi.PartialAttrUpdate(context.Background(), entityId, attrId).NGSILDTenant(nGSILDTenant).RequestBody(requestBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EntitiesApi.PartialAttrUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**entityId** | **string** | Entity Id | 
**attrId** | **string** | Attribute Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiPartialAttrUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **nGSILDTenant** | **string** |  | [default to &quot;orion&quot;]


 **requestBody** | [**map[string]OneOfPropertyRelationshipGeoProperty**](oneOf&lt;Property,Relationship,GeoProperty&gt;.md) |  | 

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


## QueryEntities

> []Entity QueryEntities(ctx).NGSILDTenant(nGSILDTenant).Id(id).IdPattern(idPattern).Type_(type_).Attrs(attrs).Q(q).Georel(georel).Geometry(geometry).Coordinates(coordinates).Geoproperty(geoproperty).Csf(csf).Limit(limit).Options(options).Link(link).Execute()





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
    nGSILDTenant := "nGSILDTenant_example" // string |  (default to "orion")
    id := "id_example" // string | Comma separated list of URIs to be retrieved (optional)
    idPattern := "idPattern_example" // string | Regular expression that must be matched by Entity ids (optional)
    type_ := "type__example" // string | Comma separated list of Entity type names to be retrieved (optional)
    attrs := "attrs_example" // string | Comma separated list of attribute names (properties or relationships) to be retrieved (optional)
    q := "q_example" // string | Query (optional)
    georel := "georel_example" // string | Geo-relationship (optional)
    geometry := "geometry_example" // string | Geometry (optional)
    coordinates := "coordinates_example" // string | Coordinates serialized as a string (optional)
    geoproperty := "geoproperty_example" // string | The name of the property that contains the geo-spatial data that will be used to resolve the geoquery (optional)
    csf := "csf_example" // string | Context Source Filter (optional)
    limit := int32(56) // int32 | Pagination limit (optional)
    options := "options_example" // string | Options dictionary (optional)
    link := "link_example" // string | Link header to be used as described in the json-ld spec. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EntitiesApi.QueryEntities(context.Background()).NGSILDTenant(nGSILDTenant).Id(id).IdPattern(idPattern).Type_(type_).Attrs(attrs).Q(q).Georel(georel).Geometry(geometry).Coordinates(coordinates).Geoproperty(geoproperty).Csf(csf).Limit(limit).Options(options).Link(link).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EntitiesApi.QueryEntities``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `QueryEntities`: []Entity
    fmt.Fprintf(os.Stdout, "Response from `EntitiesApi.QueryEntities`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiQueryEntitiesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **nGSILDTenant** | **string** |  | [default to &quot;orion&quot;]
 **id** | **string** | Comma separated list of URIs to be retrieved | 
 **idPattern** | **string** | Regular expression that must be matched by Entity ids | 
 **type_** | **string** | Comma separated list of Entity type names to be retrieved | 
 **attrs** | **string** | Comma separated list of attribute names (properties or relationships) to be retrieved | 
 **q** | **string** | Query | 
 **georel** | **string** | Geo-relationship | 
 **geometry** | **string** | Geometry | 
 **coordinates** | **string** | Coordinates serialized as a string | 
 **geoproperty** | **string** | The name of the property that contains the geo-spatial data that will be used to resolve the geoquery | 
 **csf** | **string** | Context Source Filter | 
 **limit** | **int32** | Pagination limit | 
 **options** | **string** | Options dictionary | 
 **link** | **string** | Link header to be used as described in the json-ld spec. | 

### Return type

[**[]Entity**](Entity.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/ld+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RemoveEntityAttr

> RemoveEntityAttr(ctx, entityId, attrId).NGSILDTenant(nGSILDTenant).Execute()





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
    nGSILDTenant := "nGSILDTenant_example" // string |  (default to "orion")
    entityId := "entityId_example" // string | Entity Id
    attrId := "attrId_example" // string | Attribute Id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EntitiesApi.RemoveEntityAttr(context.Background(), entityId, attrId).NGSILDTenant(nGSILDTenant).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EntitiesApi.RemoveEntityAttr``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**entityId** | **string** | Entity Id | 
**attrId** | **string** | Attribute Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemoveEntityAttrRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **nGSILDTenant** | **string** |  | [default to &quot;orion&quot;]



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


## RemoveEntityById

> RemoveEntityById(ctx, entityId).NGSILDTenant(nGSILDTenant).Type_(type_).Execute()





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
    nGSILDTenant := "nGSILDTenant_example" // string |  (default to "orion")
    entityId := "entityId_example" // string | Entity Id
    type_ := "type__example" // string | Entity Type (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EntitiesApi.RemoveEntityById(context.Background(), entityId).NGSILDTenant(nGSILDTenant).Type_(type_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EntitiesApi.RemoveEntityById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**entityId** | **string** | Entity Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemoveEntityByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **nGSILDTenant** | **string** |  | [default to &quot;orion&quot;]

 **type_** | **string** | Entity Type | 

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


## RetrieveEntityById

> Entity RetrieveEntityById(ctx, entityId).NGSILDTenant(nGSILDTenant).Attrs(attrs).Type_(type_).Options(options).Link(link).Execute()





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
    nGSILDTenant := "nGSILDTenant_example" // string |  (default to "orion")
    entityId := "entityId_example" // string | Entity Id
    attrs := "attrs_example" // string | Comma separated list of attribute names (properties or relationships) to be retrieved (optional)
    type_ := "type__example" // string | Entity Type (optional)
    options := "options_example" // string | Options dictionary (optional)
    link := "link_example" // string | Link header to be used as described in the json-ld spec. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EntitiesApi.RetrieveEntityById(context.Background(), entityId).NGSILDTenant(nGSILDTenant).Attrs(attrs).Type_(type_).Options(options).Link(link).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EntitiesApi.RetrieveEntityById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RetrieveEntityById`: Entity
    fmt.Fprintf(os.Stdout, "Response from `EntitiesApi.RetrieveEntityById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**entityId** | **string** | Entity Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiRetrieveEntityByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **nGSILDTenant** | **string** |  | [default to &quot;orion&quot;]

 **attrs** | **string** | Comma separated list of attribute names (properties or relationships) to be retrieved | 
 **type_** | **string** | Entity Type | 
 **options** | **string** | Options dictionary | 
 **link** | **string** | Link header to be used as described in the json-ld spec. | 

### Return type

[**Entity**](Entity.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/ld+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateEntityAttrs

> UpdateEntityAttrs(ctx, entityId).NGSILDTenant(nGSILDTenant).RequestBody(requestBody).Execute()





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
    nGSILDTenant := "nGSILDTenant_example" // string |  (default to "orion")
    entityId := "entityId_example" // string | Entity Id
    requestBody := map[string]OneOfPropertyRelationshipGeoProperty{"key": "TODO"} // map[string]OneOfPropertyRelationshipGeoProperty | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EntitiesApi.UpdateEntityAttrs(context.Background(), entityId).NGSILDTenant(nGSILDTenant).RequestBody(requestBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EntitiesApi.UpdateEntityAttrs``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**entityId** | **string** | Entity Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateEntityAttrsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **nGSILDTenant** | **string** |  | [default to &quot;orion&quot;]

 **requestBody** | [**map[string]OneOfPropertyRelationshipGeoProperty**](oneOf&lt;Property,Relationship,GeoProperty&gt;.md) |  | 

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

