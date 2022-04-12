# Property

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**Value** | [**NullableOneOfstringnumberbooleanobjectarray**](oneOf&lt;string,number,boolean,object,array&gt;.md) |  | 
**ObservedAt** | Pointer to **time.Time** |  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**ModifiedAt** | Pointer to **time.Time** |  | [optional] 
**DatasetId** | Pointer to **string** |  | [optional] 
**InstanceId** | Pointer to **string** |  | [optional] 
**UnitCode** | Pointer to **string** |  | [optional] 

## Methods

### NewProperty

`func NewProperty(type_ string, value NullableOneOfstringnumberbooleanobjectarray, ) *Property`

NewProperty instantiates a new Property object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPropertyWithDefaults

`func NewPropertyWithDefaults() *Property`

NewPropertyWithDefaults instantiates a new Property object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *Property) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Property) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Property) SetType(v string)`

SetType sets Type field to given value.


### GetValue

`func (o *Property) GetValue() OneOfstringnumberbooleanobjectarray`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *Property) GetValueOk() (*OneOfstringnumberbooleanobjectarray, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *Property) SetValue(v OneOfstringnumberbooleanobjectarray)`

SetValue sets Value field to given value.


### SetValueNil

`func (o *Property) SetValueNil(b bool)`

 SetValueNil sets the value for Value to be an explicit nil

### UnsetValue
`func (o *Property) UnsetValue()`

UnsetValue ensures that no value is present for Value, not even an explicit nil
### GetObservedAt

`func (o *Property) GetObservedAt() time.Time`

GetObservedAt returns the ObservedAt field if non-nil, zero value otherwise.

### GetObservedAtOk

`func (o *Property) GetObservedAtOk() (*time.Time, bool)`

GetObservedAtOk returns a tuple with the ObservedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObservedAt

`func (o *Property) SetObservedAt(v time.Time)`

SetObservedAt sets ObservedAt field to given value.

### HasObservedAt

`func (o *Property) HasObservedAt() bool`

HasObservedAt returns a boolean if a field has been set.

### GetCreatedAt

`func (o *Property) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Property) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Property) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *Property) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetModifiedAt

`func (o *Property) GetModifiedAt() time.Time`

GetModifiedAt returns the ModifiedAt field if non-nil, zero value otherwise.

### GetModifiedAtOk

`func (o *Property) GetModifiedAtOk() (*time.Time, bool)`

GetModifiedAtOk returns a tuple with the ModifiedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedAt

`func (o *Property) SetModifiedAt(v time.Time)`

SetModifiedAt sets ModifiedAt field to given value.

### HasModifiedAt

`func (o *Property) HasModifiedAt() bool`

HasModifiedAt returns a boolean if a field has been set.

### GetDatasetId

`func (o *Property) GetDatasetId() string`

GetDatasetId returns the DatasetId field if non-nil, zero value otherwise.

### GetDatasetIdOk

`func (o *Property) GetDatasetIdOk() (*string, bool)`

GetDatasetIdOk returns a tuple with the DatasetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatasetId

`func (o *Property) SetDatasetId(v string)`

SetDatasetId sets DatasetId field to given value.

### HasDatasetId

`func (o *Property) HasDatasetId() bool`

HasDatasetId returns a boolean if a field has been set.

### GetInstanceId

`func (o *Property) GetInstanceId() string`

GetInstanceId returns the InstanceId field if non-nil, zero value otherwise.

### GetInstanceIdOk

`func (o *Property) GetInstanceIdOk() (*string, bool)`

GetInstanceIdOk returns a tuple with the InstanceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceId

`func (o *Property) SetInstanceId(v string)`

SetInstanceId sets InstanceId field to given value.

### HasInstanceId

`func (o *Property) HasInstanceId() bool`

HasInstanceId returns a boolean if a field has been set.

### GetUnitCode

`func (o *Property) GetUnitCode() string`

GetUnitCode returns the UnitCode field if non-nil, zero value otherwise.

### GetUnitCodeOk

`func (o *Property) GetUnitCodeOk() (*string, bool)`

GetUnitCodeOk returns a tuple with the UnitCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnitCode

`func (o *Property) SetUnitCode(v string)`

SetUnitCode sets UnitCode field to given value.

### HasUnitCode

`func (o *Property) HasUnitCode() bool`

HasUnitCode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


