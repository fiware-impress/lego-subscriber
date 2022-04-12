# EntityTemporal

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Context** | Pointer to **map[string]interface{}** |  | [optional] 
**Location** | Pointer to [**[]GeoProperty**](GeoProperty.md) |  | [optional] 
**ObservationSpace** | Pointer to [**[]GeoProperty**](GeoProperty.md) |  | [optional] 
**OperationSpace** | Pointer to [**[]GeoProperty**](GeoProperty.md) |  | [optional] 
**Id** | **string** |  | 
**Type** | **string** | NGSI-LD Name | 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**ModifiedAt** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewEntityTemporal

`func NewEntityTemporal(id string, type_ string, ) *EntityTemporal`

NewEntityTemporal instantiates a new EntityTemporal object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEntityTemporalWithDefaults

`func NewEntityTemporalWithDefaults() *EntityTemporal`

NewEntityTemporalWithDefaults instantiates a new EntityTemporal object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContext

`func (o *EntityTemporal) GetContext() map[string]interface{}`

GetContext returns the Context field if non-nil, zero value otherwise.

### GetContextOk

`func (o *EntityTemporal) GetContextOk() (*map[string]interface{}, bool)`

GetContextOk returns a tuple with the Context field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContext

`func (o *EntityTemporal) SetContext(v map[string]interface{})`

SetContext sets Context field to given value.

### HasContext

`func (o *EntityTemporal) HasContext() bool`

HasContext returns a boolean if a field has been set.

### GetLocation

`func (o *EntityTemporal) GetLocation() []GeoProperty`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *EntityTemporal) GetLocationOk() (*[]GeoProperty, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *EntityTemporal) SetLocation(v []GeoProperty)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *EntityTemporal) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### GetObservationSpace

`func (o *EntityTemporal) GetObservationSpace() []GeoProperty`

GetObservationSpace returns the ObservationSpace field if non-nil, zero value otherwise.

### GetObservationSpaceOk

`func (o *EntityTemporal) GetObservationSpaceOk() (*[]GeoProperty, bool)`

GetObservationSpaceOk returns a tuple with the ObservationSpace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObservationSpace

`func (o *EntityTemporal) SetObservationSpace(v []GeoProperty)`

SetObservationSpace sets ObservationSpace field to given value.

### HasObservationSpace

`func (o *EntityTemporal) HasObservationSpace() bool`

HasObservationSpace returns a boolean if a field has been set.

### GetOperationSpace

`func (o *EntityTemporal) GetOperationSpace() []GeoProperty`

GetOperationSpace returns the OperationSpace field if non-nil, zero value otherwise.

### GetOperationSpaceOk

`func (o *EntityTemporal) GetOperationSpaceOk() (*[]GeoProperty, bool)`

GetOperationSpaceOk returns a tuple with the OperationSpace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperationSpace

`func (o *EntityTemporal) SetOperationSpace(v []GeoProperty)`

SetOperationSpace sets OperationSpace field to given value.

### HasOperationSpace

`func (o *EntityTemporal) HasOperationSpace() bool`

HasOperationSpace returns a boolean if a field has been set.

### GetId

`func (o *EntityTemporal) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *EntityTemporal) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *EntityTemporal) SetId(v string)`

SetId sets Id field to given value.


### GetType

`func (o *EntityTemporal) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *EntityTemporal) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *EntityTemporal) SetType(v string)`

SetType sets Type field to given value.


### GetCreatedAt

`func (o *EntityTemporal) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *EntityTemporal) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *EntityTemporal) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *EntityTemporal) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetModifiedAt

`func (o *EntityTemporal) GetModifiedAt() time.Time`

GetModifiedAt returns the ModifiedAt field if non-nil, zero value otherwise.

### GetModifiedAtOk

`func (o *EntityTemporal) GetModifiedAtOk() (*time.Time, bool)`

GetModifiedAtOk returns a tuple with the ModifiedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedAt

`func (o *EntityTemporal) SetModifiedAt(v time.Time)`

SetModifiedAt sets ModifiedAt field to given value.

### HasModifiedAt

`func (o *EntityTemporal) HasModifiedAt() bool`

HasModifiedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


