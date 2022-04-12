# TemporalQuery

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Timerel** | Pointer to **string** |  | [optional] 
**TimeAt** | Pointer to **time.Time** |  | [optional] 
**EndTimeAt** | Pointer to **time.Time** |  | [optional] 
**Timeproperty** | Pointer to **string** |  | [optional] 

## Methods

### NewTemporalQuery

`func NewTemporalQuery() *TemporalQuery`

NewTemporalQuery instantiates a new TemporalQuery object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTemporalQueryWithDefaults

`func NewTemporalQueryWithDefaults() *TemporalQuery`

NewTemporalQueryWithDefaults instantiates a new TemporalQuery object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTimerel

`func (o *TemporalQuery) GetTimerel() string`

GetTimerel returns the Timerel field if non-nil, zero value otherwise.

### GetTimerelOk

`func (o *TemporalQuery) GetTimerelOk() (*string, bool)`

GetTimerelOk returns a tuple with the Timerel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimerel

`func (o *TemporalQuery) SetTimerel(v string)`

SetTimerel sets Timerel field to given value.

### HasTimerel

`func (o *TemporalQuery) HasTimerel() bool`

HasTimerel returns a boolean if a field has been set.

### GetTimeAt

`func (o *TemporalQuery) GetTimeAt() time.Time`

GetTimeAt returns the TimeAt field if non-nil, zero value otherwise.

### GetTimeAtOk

`func (o *TemporalQuery) GetTimeAtOk() (*time.Time, bool)`

GetTimeAtOk returns a tuple with the TimeAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeAt

`func (o *TemporalQuery) SetTimeAt(v time.Time)`

SetTimeAt sets TimeAt field to given value.

### HasTimeAt

`func (o *TemporalQuery) HasTimeAt() bool`

HasTimeAt returns a boolean if a field has been set.

### GetEndTimeAt

`func (o *TemporalQuery) GetEndTimeAt() time.Time`

GetEndTimeAt returns the EndTimeAt field if non-nil, zero value otherwise.

### GetEndTimeAtOk

`func (o *TemporalQuery) GetEndTimeAtOk() (*time.Time, bool)`

GetEndTimeAtOk returns a tuple with the EndTimeAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTimeAt

`func (o *TemporalQuery) SetEndTimeAt(v time.Time)`

SetEndTimeAt sets EndTimeAt field to given value.

### HasEndTimeAt

`func (o *TemporalQuery) HasEndTimeAt() bool`

HasEndTimeAt returns a boolean if a field has been set.

### GetTimeproperty

`func (o *TemporalQuery) GetTimeproperty() string`

GetTimeproperty returns the Timeproperty field if non-nil, zero value otherwise.

### GetTimepropertyOk

`func (o *TemporalQuery) GetTimepropertyOk() (*string, bool)`

GetTimepropertyOk returns a tuple with the Timeproperty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeproperty

`func (o *TemporalQuery) SetTimeproperty(v string)`

SetTimeproperty sets Timeproperty field to given value.

### HasTimeproperty

`func (o *TemporalQuery) HasTimeproperty() bool`

HasTimeproperty returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


