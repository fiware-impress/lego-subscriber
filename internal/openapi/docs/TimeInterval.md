# TimeInterval

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Start** | **time.Time** |  | 
**End** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewTimeInterval

`func NewTimeInterval(start time.Time, ) *TimeInterval`

NewTimeInterval instantiates a new TimeInterval object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTimeIntervalWithDefaults

`func NewTimeIntervalWithDefaults() *TimeInterval`

NewTimeIntervalWithDefaults instantiates a new TimeInterval object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStart

`func (o *TimeInterval) GetStart() time.Time`

GetStart returns the Start field if non-nil, zero value otherwise.

### GetStartOk

`func (o *TimeInterval) GetStartOk() (*time.Time, bool)`

GetStartOk returns a tuple with the Start field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStart

`func (o *TimeInterval) SetStart(v time.Time)`

SetStart sets Start field to given value.


### GetEnd

`func (o *TimeInterval) GetEnd() time.Time`

GetEnd returns the End field if non-nil, zero value otherwise.

### GetEndOk

`func (o *TimeInterval) GetEndOk() (*time.Time, bool)`

GetEndOk returns a tuple with the End field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnd

`func (o *TimeInterval) SetEnd(v time.Time)`

SetEnd sets End field to given value.

### HasEnd

`func (o *TimeInterval) HasEnd() bool`

HasEnd returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


