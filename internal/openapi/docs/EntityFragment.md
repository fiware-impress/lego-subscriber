# EntityFragment

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

### NewEntityFragment

`func NewEntityFragment() *EntityFragment`

NewEntityFragment instantiates a new EntityFragment object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEntityFragmentWithDefaults

`func NewEntityFragmentWithDefaults() *EntityFragment`

NewEntityFragmentWithDefaults instantiates a new EntityFragment object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContext

`func (o *EntityFragment) GetContext() map[string]interface{}`

GetContext returns the Context field if non-nil, zero value otherwise.

### GetContextOk

`func (o *EntityFragment) GetContextOk() (*map[string]interface{}, bool)`

GetContextOk returns a tuple with the Context field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContext

`func (o *EntityFragment) SetContext(v map[string]interface{})`

SetContext sets Context field to given value.

### HasContext

`func (o *EntityFragment) HasContext() bool`

HasContext returns a boolean if a field has been set.

### GetLocation

`func (o *EntityFragment) GetLocation() GeoProperty`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *EntityFragment) GetLocationOk() (*GeoProperty, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *EntityFragment) SetLocation(v GeoProperty)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *EntityFragment) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### GetObservationSpace

`func (o *EntityFragment) GetObservationSpace() GeoProperty`

GetObservationSpace returns the ObservationSpace field if non-nil, zero value otherwise.

### GetObservationSpaceOk

`func (o *EntityFragment) GetObservationSpaceOk() (*GeoProperty, bool)`

GetObservationSpaceOk returns a tuple with the ObservationSpace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObservationSpace

`func (o *EntityFragment) SetObservationSpace(v GeoProperty)`

SetObservationSpace sets ObservationSpace field to given value.

### HasObservationSpace

`func (o *EntityFragment) HasObservationSpace() bool`

HasObservationSpace returns a boolean if a field has been set.

### GetOperationSpace

`func (o *EntityFragment) GetOperationSpace() GeoProperty`

GetOperationSpace returns the OperationSpace field if non-nil, zero value otherwise.

### GetOperationSpaceOk

`func (o *EntityFragment) GetOperationSpaceOk() (*GeoProperty, bool)`

GetOperationSpaceOk returns a tuple with the OperationSpace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperationSpace

`func (o *EntityFragment) SetOperationSpace(v GeoProperty)`

SetOperationSpace sets OperationSpace field to given value.

### HasOperationSpace

`func (o *EntityFragment) HasOperationSpace() bool`

HasOperationSpace returns a boolean if a field has been set.

### GetId

`func (o *EntityFragment) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *EntityFragment) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *EntityFragment) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *EntityFragment) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *EntityFragment) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *EntityFragment) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *EntityFragment) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *EntityFragment) HasType() bool`

HasType returns a boolean if a field has been set.

### GetCreatedAt

`func (o *EntityFragment) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *EntityFragment) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *EntityFragment) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *EntityFragment) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetModifiedAt

`func (o *EntityFragment) GetModifiedAt() time.Time`

GetModifiedAt returns the ModifiedAt field if non-nil, zero value otherwise.

### GetModifiedAtOk

`func (o *EntityFragment) GetModifiedAtOk() (*time.Time, bool)`

GetModifiedAtOk returns a tuple with the ModifiedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedAt

`func (o *EntityFragment) SetModifiedAt(v time.Time)`

SetModifiedAt sets ModifiedAt field to given value.

### HasModifiedAt

`func (o *EntityFragment) HasModifiedAt() bool`

HasModifiedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


