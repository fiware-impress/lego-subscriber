# GeoQuery

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Georel** | **string** |  | 
**Coordinates** | [**[]OneOfintegercoordinates**](OneOfintegercoordinates.md) |  | 
**Geometry** | **string** |  | 
**Geoproperty** | Pointer to **string** |  | [optional] 

## Methods

### NewGeoQuery

`func NewGeoQuery(georel string, coordinates []OneOfintegercoordinates, geometry string, ) *GeoQuery`

NewGeoQuery instantiates a new GeoQuery object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGeoQueryWithDefaults

`func NewGeoQueryWithDefaults() *GeoQuery`

NewGeoQueryWithDefaults instantiates a new GeoQuery object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGeorel

`func (o *GeoQuery) GetGeorel() string`

GetGeorel returns the Georel field if non-nil, zero value otherwise.

### GetGeorelOk

`func (o *GeoQuery) GetGeorelOk() (*string, bool)`

GetGeorelOk returns a tuple with the Georel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGeorel

`func (o *GeoQuery) SetGeorel(v string)`

SetGeorel sets Georel field to given value.


### GetCoordinates

`func (o *GeoQuery) GetCoordinates() []OneOfintegercoordinates`

GetCoordinates returns the Coordinates field if non-nil, zero value otherwise.

### GetCoordinatesOk

`func (o *GeoQuery) GetCoordinatesOk() (*[]OneOfintegercoordinates, bool)`

GetCoordinatesOk returns a tuple with the Coordinates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoordinates

`func (o *GeoQuery) SetCoordinates(v []OneOfintegercoordinates)`

SetCoordinates sets Coordinates field to given value.


### GetGeometry

`func (o *GeoQuery) GetGeometry() string`

GetGeometry returns the Geometry field if non-nil, zero value otherwise.

### GetGeometryOk

`func (o *GeoQuery) GetGeometryOk() (*string, bool)`

GetGeometryOk returns a tuple with the Geometry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGeometry

`func (o *GeoQuery) SetGeometry(v string)`

SetGeometry sets Geometry field to given value.


### GetGeoproperty

`func (o *GeoQuery) GetGeoproperty() string`

GetGeoproperty returns the Geoproperty field if non-nil, zero value otherwise.

### GetGeopropertyOk

`func (o *GeoQuery) GetGeopropertyOk() (*string, bool)`

GetGeopropertyOk returns a tuple with the Geoproperty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGeoproperty

`func (o *GeoQuery) SetGeoproperty(v string)`

SetGeoproperty sets Geoproperty field to given value.

### HasGeoproperty

`func (o *GeoQuery) HasGeoproperty() bool`

HasGeoproperty returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


