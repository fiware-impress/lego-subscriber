# Point

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** |  | [optional] 
**Coordinates** | Pointer to **[]float32** | A single position | [optional] 

## Methods

### NewPoint

`func NewPoint() *Point`

NewPoint instantiates a new Point object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPointWithDefaults

`func NewPointWithDefaults() *Point`

NewPointWithDefaults instantiates a new Point object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *Point) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Point) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Point) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *Point) HasType() bool`

HasType returns a boolean if a field has been set.

### GetCoordinates

`func (o *Point) GetCoordinates() []float32`

GetCoordinates returns the Coordinates field if non-nil, zero value otherwise.

### GetCoordinatesOk

`func (o *Point) GetCoordinatesOk() (*[]float32, bool)`

GetCoordinatesOk returns a tuple with the Coordinates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoordinates

`func (o *Point) SetCoordinates(v []float32)`

SetCoordinates sets Coordinates field to given value.

### HasCoordinates

`func (o *Point) HasCoordinates() bool`

HasCoordinates returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


