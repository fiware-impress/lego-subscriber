# Entity

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Context** | Pointer to **map[string]interface{}** |  | [optional] 
**Location** | Pointer to [**GeoProperty**](GeoProperty.md) |  | [optional] 
**ObservationSpace** | Pointer to [**GeoProperty**](GeoProperty.md) |  | [optional] 
**OperationSpace** | Pointer to [**GeoProperty**](GeoProperty.md) |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** | NGSI-LD Name | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**ModifiedAt** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewEntity

`func NewEntity() *Entity`

NewEntity instantiates a new Entity object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEntityWithDefaults

`func NewEntityWithDefaults() *Entity`

NewEntityWithDefaults instantiates a new Entity object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContext

`func (o *Entity) GetContext() map[string]interface{}`

GetContext returns the Context field if non-nil, zero value otherwise.

### GetContextOk

`func (o *Entity) GetContextOk() (*map[string]interface{}, bool)`

GetContextOk returns a tuple with the Context field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContext

`func (o *Entity) SetContext(v map[string]interface{})`

SetContext sets Context field to given value.

### HasContext

`func (o *Entity) HasContext() bool`

HasContext returns a boolean if a field has been set.

### GetLocation

`func (o *Entity) GetLocation() GeoProperty`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *Entity) GetLocationOk() (*GeoProperty, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *Entity) SetLocation(v GeoProperty)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *Entity) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### GetObservationSpace

`func (o *Entity) GetObservationSpace() GeoProperty`

GetObservationSpace returns the ObservationSpace field if non-nil, zero value otherwise.

### GetObservationSpaceOk

`func (o *Entity) GetObservationSpaceOk() (*GeoProperty, bool)`

GetObservationSpaceOk returns a tuple with the ObservationSpace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObservationSpace

`func (o *Entity) SetObservationSpace(v GeoProperty)`

SetObservationSpace sets ObservationSpace field to given value.

### HasObservationSpace

`func (o *Entity) HasObservationSpace() bool`

HasObservationSpace returns a boolean if a field has been set.

### GetOperationSpace

`func (o *Entity) GetOperationSpace() GeoProperty`

GetOperationSpace returns the OperationSpace field if non-nil, zero value otherwise.

### GetOperationSpaceOk

`func (o *Entity) GetOperationSpaceOk() (*GeoProperty, bool)`

GetOperationSpaceOk returns a tuple with the OperationSpace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperationSpace

`func (o *Entity) SetOperationSpace(v GeoProperty)`

SetOperationSpace sets OperationSpace field to given value.

### HasOperationSpace

`func (o *Entity) HasOperationSpace() bool`

HasOperationSpace returns a boolean if a field has been set.

### GetId

`func (o *Entity) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Entity) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Entity) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Entity) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *Entity) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Entity) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Entity) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *Entity) HasType() bool`

HasType returns a boolean if a field has been set.

### GetCreatedAt

`func (o *Entity) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Entity) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Entity) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *Entity) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetModifiedAt

`func (o *Entity) GetModifiedAt() time.Time`

GetModifiedAt returns the ModifiedAt field if non-nil, zero value otherwise.

### GetModifiedAtOk

`func (o *Entity) GetModifiedAtOk() (*time.Time, bool)`

GetModifiedAtOk returns a tuple with the ModifiedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedAt

`func (o *Entity) SetModifiedAt(v time.Time)`

SetModifiedAt sets ModifiedAt field to given value.

### HasModifiedAt

`func (o *Entity) HasModifiedAt() bool`

HasModifiedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


