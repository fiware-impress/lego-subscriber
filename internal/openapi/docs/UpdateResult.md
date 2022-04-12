# UpdateResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Updated** | Pointer to **[]string** |  | [optional] 
**NotUpdated** | Pointer to [**[]NotUpdatedDetails**](NotUpdatedDetails.md) |  | [optional] 

## Methods

### NewUpdateResult

`func NewUpdateResult() *UpdateResult`

NewUpdateResult instantiates a new UpdateResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateResultWithDefaults

`func NewUpdateResultWithDefaults() *UpdateResult`

NewUpdateResultWithDefaults instantiates a new UpdateResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUpdated

`func (o *UpdateResult) GetUpdated() []string`

GetUpdated returns the Updated field if non-nil, zero value otherwise.

### GetUpdatedOk

`func (o *UpdateResult) GetUpdatedOk() (*[]string, bool)`

GetUpdatedOk returns a tuple with the Updated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdated

`func (o *UpdateResult) SetUpdated(v []string)`

SetUpdated sets Updated field to given value.

### HasUpdated

`func (o *UpdateResult) HasUpdated() bool`

HasUpdated returns a boolean if a field has been set.

### GetNotUpdated

`func (o *UpdateResult) GetNotUpdated() []NotUpdatedDetails`

GetNotUpdated returns the NotUpdated field if non-nil, zero value otherwise.

### GetNotUpdatedOk

`func (o *UpdateResult) GetNotUpdatedOk() (*[]NotUpdatedDetails, bool)`

GetNotUpdatedOk returns a tuple with the NotUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotUpdated

`func (o *UpdateResult) SetNotUpdated(v []NotUpdatedDetails)`

SetNotUpdated sets NotUpdated field to given value.

### HasNotUpdated

`func (o *UpdateResult) HasNotUpdated() bool`

HasNotUpdated returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


