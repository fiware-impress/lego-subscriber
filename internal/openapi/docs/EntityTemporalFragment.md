# EntityTemporalFragment

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Context** | Pointer to **map[string]interface{}** |  | [optional] 
**Location** | Pointer to [**[]GeoProperty**](GeoProperty.md) |  | [optional] 
**ObservationSpace** | Pointer to [**[]GeoProperty**](GeoProperty.md) |  | [optional] 
**OperationSpace** | Pointer to [**[]GeoProperty**](GeoProperty.md) |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** | NGSI-LD Name | [optional] 

## Methods

### NewEntityTemporalFragment

`func NewEntityTemporalFragment() *EntityTemporalFragment`

NewEntityTemporalFragment instantiates a new EntityTemporalFragment object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEntityTemporalFragmentWithDefaults

`func NewEntityTemporalFragmentWithDefaults() *EntityTemporalFragment`

NewEntityTemporalFragmentWithDefaults instantiates a new EntityTemporalFragment object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContext

`func (o *EntityTemporalFragment) GetContext() map[string]interface{}`

GetContext returns the Context field if non-nil, zero value otherwise.

### GetContextOk

`func (o *EntityTemporalFragment) GetContextOk() (*map[string]interface{}, bool)`

GetContextOk returns a tuple with the Context field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContext

`func (o *EntityTemporalFragment) SetContext(v map[string]interface{})`

SetContext sets Context field to given value.

### HasContext

`func (o *EntityTemporalFragment) HasContext() bool`

HasContext returns a boolean if a field has been set.

### GetLocation

`func (o *EntityTemporalFragment) GetLocation() []GeoProperty`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *EntityTemporalFragment) GetLocationOk() (*[]GeoProperty, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *EntityTemporalFragment) SetLocation(v []GeoProperty)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *EntityTemporalFragment) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### GetObservationSpace

`func (o *EntityTemporalFragment) GetObservationSpace() []GeoProperty`

GetObservationSpace returns the ObservationSpace field if non-nil, zero value otherwise.

### GetObservationSpaceOk

`func (o *EntityTemporalFragment) GetObservationSpaceOk() (*[]GeoProperty, bool)`

GetObservationSpaceOk returns a tuple with the ObservationSpace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObservationSpace

`func (o *EntityTemporalFragment) SetObservationSpace(v []GeoProperty)`

SetObservationSpace sets ObservationSpace field to given value.

### HasObservationSpace

`func (o *EntityTemporalFragment) HasObservationSpace() bool`

HasObservationSpace returns a boolean if a field has been set.

### GetOperationSpace

`func (o *EntityTemporalFragment) GetOperationSpace() []GeoProperty`

GetOperationSpace returns the OperationSpace field if non-nil, zero value otherwise.

### GetOperationSpaceOk

`func (o *EntityTemporalFragment) GetOperationSpaceOk() (*[]GeoProperty, bool)`

GetOperationSpaceOk returns a tuple with the OperationSpace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperationSpace

`func (o *EntityTemporalFragment) SetOperationSpace(v []GeoProperty)`

SetOperationSpace sets OperationSpace field to given value.

### HasOperationSpace

`func (o *EntityTemporalFragment) HasOperationSpace() bool`

HasOperationSpace returns a boolean if a field has been set.

### GetId

`func (o *EntityTemporalFragment) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *EntityTemporalFragment) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *EntityTemporalFragment) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *EntityTemporalFragment) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *EntityTemporalFragment) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *EntityTemporalFragment) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *EntityTemporalFragment) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *EntityTemporalFragment) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


