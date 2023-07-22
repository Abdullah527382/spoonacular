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

// ParseIngredients200ResponseInnerNutritionPropertiesInner struct for ParseIngredients200ResponseInnerNutritionPropertiesInner
type ParseIngredients200ResponseInnerNutritionPropertiesInner struct {
	Name string `json:"name"`
	Amount float32 `json:"amount"`
	Unit string `json:"unit"`
}

// NewParseIngredients200ResponseInnerNutritionPropertiesInner instantiates a new ParseIngredients200ResponseInnerNutritionPropertiesInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewParseIngredients200ResponseInnerNutritionPropertiesInner(name string, amount float32, unit string) *ParseIngredients200ResponseInnerNutritionPropertiesInner {
	this := ParseIngredients200ResponseInnerNutritionPropertiesInner{}
	this.Name = name
	this.Amount = amount
	this.Unit = unit
	return &this
}

// NewParseIngredients200ResponseInnerNutritionPropertiesInnerWithDefaults instantiates a new ParseIngredients200ResponseInnerNutritionPropertiesInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewParseIngredients200ResponseInnerNutritionPropertiesInnerWithDefaults() *ParseIngredients200ResponseInnerNutritionPropertiesInner {
	this := ParseIngredients200ResponseInnerNutritionPropertiesInner{}
	return &this
}

// GetName returns the Name field value
func (o *ParseIngredients200ResponseInnerNutritionPropertiesInner) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ParseIngredients200ResponseInnerNutritionPropertiesInner) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ParseIngredients200ResponseInnerNutritionPropertiesInner) SetName(v string) {
	o.Name = v
}

// GetAmount returns the Amount field value
func (o *ParseIngredients200ResponseInnerNutritionPropertiesInner) GetAmount() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Amount
}

// GetAmountOk returns a tuple with the Amount field value
// and a boolean to check if the value has been set.
func (o *ParseIngredients200ResponseInnerNutritionPropertiesInner) GetAmountOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Amount, true
}

// SetAmount sets field value
func (o *ParseIngredients200ResponseInnerNutritionPropertiesInner) SetAmount(v float32) {
	o.Amount = v
}

// GetUnit returns the Unit field value
func (o *ParseIngredients200ResponseInnerNutritionPropertiesInner) GetUnit() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Unit
}

// GetUnitOk returns a tuple with the Unit field value
// and a boolean to check if the value has been set.
func (o *ParseIngredients200ResponseInnerNutritionPropertiesInner) GetUnitOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Unit, true
}

// SetUnit sets field value
func (o *ParseIngredients200ResponseInnerNutritionPropertiesInner) SetUnit(v string) {
	o.Unit = v
}

func (o ParseIngredients200ResponseInnerNutritionPropertiesInner) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["amount"] = o.Amount
	}
	if true {
		toSerialize["unit"] = o.Unit
	}
	return json.Marshal(toSerialize)
}

type NullableParseIngredients200ResponseInnerNutritionPropertiesInner struct {
	value *ParseIngredients200ResponseInnerNutritionPropertiesInner
	isSet bool
}

func (v NullableParseIngredients200ResponseInnerNutritionPropertiesInner) Get() *ParseIngredients200ResponseInnerNutritionPropertiesInner {
	return v.value
}

func (v *NullableParseIngredients200ResponseInnerNutritionPropertiesInner) Set(val *ParseIngredients200ResponseInnerNutritionPropertiesInner) {
	v.value = val
	v.isSet = true
}

func (v NullableParseIngredients200ResponseInnerNutritionPropertiesInner) IsSet() bool {
	return v.isSet
}

func (v *NullableParseIngredients200ResponseInnerNutritionPropertiesInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableParseIngredients200ResponseInnerNutritionPropertiesInner(val *ParseIngredients200ResponseInnerNutritionPropertiesInner) *NullableParseIngredients200ResponseInnerNutritionPropertiesInner {
	return &NullableParseIngredients200ResponseInnerNutritionPropertiesInner{value: val, isSet: true}
}

func (v NullableParseIngredients200ResponseInnerNutritionPropertiesInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableParseIngredients200ResponseInnerNutritionPropertiesInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

