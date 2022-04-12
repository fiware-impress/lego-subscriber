# MultiPoint

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** |  | [optional] 
**Coordinates** | Pointer to **[][]float32** | An array of positions | [optional] 

## Methods

### NewMultiPoint

`func NewMultiPoint() *MultiPoint`

NewMultiPoint instantiates a new MultiPoint object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMultiPointWithDefaults

`func NewMultiPointWithDefaults() *MultiPoint`

NewMultiPointWithDefaults instantiates a new MultiPoint object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *MultiPoint) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *MultiPoint) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *MultiPoint) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *MultiPoint) HasType() bool`

HasType returns a boolean if a field has been set.

### GetCoordinates

`func (o *MultiPoint) GetCoordinates() [][]float32`

GetCoordinates returns the Coordinates field if non-nil, zero value otherwise.

### GetCoordinatesOk

`func (o *MultiPoint) GetCoordinatesOk() (*[][]float32, bool)`

GetCoordinatesOk returns a tuple with the Coordinates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoordinates

`func (o *MultiPoint) SetCoordinates(v [][]float32)`

SetCoordinates sets Coordinates field to given value.

### HasCoordinates

`func (o *MultiPoint) HasCoordinates() bool`

HasCoordinates returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


