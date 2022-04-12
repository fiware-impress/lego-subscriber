# GeoProperty

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**Value** | **map[string]interface{}** |  | 
**ObservedAt** | Pointer to **time.Time** |  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**ModifiedAt** | Pointer to **time.Time** |  | [optional] 
**DatasetId** | Pointer to **string** |  | [optional] 
**InstanceId** | Pointer to **string** |  | [optional] 
**UnitCode** | Pointer to **string** |  | [optional] 

## Methods

### NewGeoProperty

`func NewGeoProperty(type_ string, value map[string]interface{}, ) *GeoProperty`

NewGeoProperty instantiates a new GeoProperty object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGeoPropertyWithDefaults

`func NewGeoPropertyWithDefaults() *GeoProperty`

NewGeoPropertyWithDefaults instantiates a new GeoProperty object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *GeoProperty) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GeoProperty) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GeoProperty) SetType(v string)`

SetType sets Type field to given value.


### GetValue

`func (o *GeoProperty) GetValue() map[string]interface{}`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *GeoProperty) GetValueOk() (*map[string]interface{}, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *GeoProperty) SetValue(v map[string]interface{})`

SetValue sets Value field to given value.


### GetObservedAt

`func (o *GeoProperty) GetObservedAt() time.Time`

GetObservedAt returns the ObservedAt field if non-nil, zero value otherwise.

### GetObservedAtOk

`func (o *GeoProperty) GetObservedAtOk() (*time.Time, bool)`

GetObservedAtOk returns a tuple with the ObservedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObservedAt

`func (o *GeoProperty) SetObservedAt(v time.Time)`

SetObservedAt sets ObservedAt field to given value.

### HasObservedAt

`func (o *GeoProperty) HasObservedAt() bool`

HasObservedAt returns a boolean if a field has been set.

### GetCreatedAt

`func (o *GeoProperty) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *GeoProperty) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *GeoProperty) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *GeoProperty) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetModifiedAt

`func (o *GeoProperty) GetModifiedAt() time.Time`

GetModifiedAt returns the ModifiedAt field if non-nil, zero value otherwise.

### GetModifiedAtOk

`func (o *GeoProperty) GetModifiedAtOk() (*time.Time, bool)`

GetModifiedAtOk returns a tuple with the ModifiedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedAt

`func (o *GeoProperty) SetModifiedAt(v time.Time)`

SetModifiedAt sets ModifiedAt field to given value.

### HasModifiedAt

`func (o *GeoProperty) HasModifiedAt() bool`

HasModifiedAt returns a boolean if a field has been set.

### GetDatasetId

`func (o *GeoProperty) GetDatasetId() string`

GetDatasetId returns the DatasetId field if non-nil, zero value otherwise.

### GetDatasetIdOk

`func (o *GeoProperty) GetDatasetIdOk() (*string, bool)`

GetDatasetIdOk returns a tuple with the DatasetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatasetId

`func (o *GeoProperty) SetDatasetId(v string)`

SetDatasetId sets DatasetId field to given value.

### HasDatasetId

`func (o *GeoProperty) HasDatasetId() bool`

HasDatasetId returns a boolean if a field has been set.

### GetInstanceId

`func (o *GeoProperty) GetInstanceId() string`

GetInstanceId returns the InstanceId field if non-nil, zero value otherwise.

### GetInstanceIdOk

`func (o *GeoProperty) GetInstanceIdOk() (*string, bool)`

GetInstanceIdOk returns a tuple with the InstanceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceId

`func (o *GeoProperty) SetInstanceId(v string)`

SetInstanceId sets InstanceId field to given value.

### HasInstanceId

`func (o *GeoProperty) HasInstanceId() bool`

HasInstanceId returns a boolean if a field has been set.

### GetUnitCode

`func (o *GeoProperty) GetUnitCode() string`

GetUnitCode returns the UnitCode field if non-nil, zero value otherwise.

### GetUnitCodeOk

`func (o *GeoProperty) GetUnitCodeOk() (*string, bool)`

GetUnitCodeOk returns a tuple with the UnitCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnitCode

`func (o *GeoProperty) SetUnitCode(v string)`

SetUnitCode sets UnitCode field to given value.

### HasUnitCode

`func (o *GeoProperty) HasUnitCode() bool`

HasUnitCode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


