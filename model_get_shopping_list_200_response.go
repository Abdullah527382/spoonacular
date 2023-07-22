/*
spoonacular API

The spoonacular Nutrition, Recipe, and Food API allows you to access over thousands of recipes, thousands of ingredients, 800,000 food products, over 100,000 menu items, and restaurants. Our food ontology and semantic recipe search engine makes it possible to search for recipes using natural language queries, such as \"gluten free brownies without sugar\" or \"low fat vegan cupcakes.\" You can automatically calculate the nutritional information for any recipe, analyze recipe costs, visualize ingredient lists, find recipes for what's in your fridge, find recipes based on special diets, nutritional requirements, or favorite ingredients, classify recipes into types and cuisines, convert ingredient amounts, or even compute an entire meal plan. With our powerful API, you can create many kinds of food and especially nutrition apps.  Special diets/dietary requirements currently available include: vegan, vegetarian, pescetarian, gluten free, grain free, dairy free, high protein, whole 30, low sodium, low carb, Paleo, ketogenic, FODMAP, and Primal.

API version: 1.1
Contact: mail@spoonacular.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// GetShoppingList200Response 
type GetShoppingList200Response struct {
	Aisles []GetShoppingList200ResponseAislesInner `json:"aisles"`
	Cost float32 `json:"cost"`
	StartDate float32 `json:"startDate"`
	EndDate float32 `json:"endDate"`
}

// NewGetShoppingList200Response instantiates a new GetShoppingList200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetShoppingList200Response(aisles []GetShoppingList200ResponseAislesInner, cost float32, startDate float32, endDate float32) *GetShoppingList200Response {
	this := GetShoppingList200Response{}
	this.Aisles = aisles
	this.Cost = cost
	this.StartDate = startDate
	this.EndDate = endDate
	return &this
}

// NewGetShoppingList200ResponseWithDefaults instantiates a new GetShoppingList200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetShoppingList200ResponseWithDefaults() *GetShoppingList200Response {
	this := GetShoppingList200Response{}
	return &this
}

// GetAisles returns the Aisles field value
func (o *GetShoppingList200Response) GetAisles() []GetShoppingList200ResponseAislesInner {
	if o == nil {
		var ret []GetShoppingList200ResponseAislesInner
		return ret
	}

	return o.Aisles
}

// GetAislesOk returns a tuple with the Aisles field value
// and a boolean to check if the value has been set.
func (o *GetShoppingList200Response) GetAislesOk() ([]GetShoppingList200ResponseAislesInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.Aisles, true
}

// SetAisles sets field value
func (o *GetShoppingList200Response) SetAisles(v []GetShoppingList200ResponseAislesInner) {
	o.Aisles = v
}

// GetCost returns the Cost field value
func (o *GetShoppingList200Response) GetCost() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Cost
}

// GetCostOk returns a tuple with the Cost field value
// and a boolean to check if the value has been set.
func (o *GetShoppingList200Response) GetCostOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Cost, true
}

// SetCost sets field value
func (o *GetShoppingList200Response) SetCost(v float32) {
	o.Cost = v
}

// GetStartDate returns the StartDate field value
func (o *GetShoppingList200Response) GetStartDate() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.StartDate
}

// GetStartDateOk returns a tuple with the StartDate field value
// and a boolean to check if the value has been set.
func (o *GetShoppingList200Response) GetStartDateOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StartDate, true
}

// SetStartDate sets field value
func (o *GetShoppingList200Response) SetStartDate(v float32) {
	o.StartDate = v
}

// GetEndDate returns the EndDate field value
func (o *GetShoppingList200Response) GetEndDate() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.EndDate
}

// GetEndDateOk returns a tuple with the EndDate field value
// and a boolean to check if the value has been set.
func (o *GetShoppingList200Response) GetEndDateOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EndDate, true
}

// SetEndDate sets field value
func (o *GetShoppingList200Response) SetEndDate(v float32) {
	o.EndDate = v
}

func (o GetShoppingList200Response) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["aisles"] = o.Aisles
	}
	if true {
		toSerialize["cost"] = o.Cost
	}
	if true {
		toSerialize["startDate"] = o.StartDate
	}
	if true {
		toSerialize["endDate"] = o.EndDate
	}
	return json.Marshal(toSerialize)
}

type NullableGetShoppingList200Response struct {
	value *GetShoppingList200Response
	isSet bool
}

func (v NullableGetShoppingList200Response) Get() *GetShoppingList200Response {
	return v.value
}

func (v *NullableGetShoppingList200Response) Set(val *GetShoppingList200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGetShoppingList200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGetShoppingList200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetShoppingList200Response(val *GetShoppingList200Response) *NullableGetShoppingList200Response {
	return &NullableGetShoppingList200Response{value: val, isSet: true}
}

func (v NullableGetShoppingList200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetShoppingList200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


