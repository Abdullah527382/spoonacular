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

// GetShoppingList200ResponseAislesInnerItemsInnerMeasures struct for GetShoppingList200ResponseAislesInnerItemsInnerMeasures
type GetShoppingList200ResponseAislesInnerItemsInnerMeasures struct {
	Original ParseIngredients200ResponseInnerNutritionWeightPerServing `json:"original"`
	Metric ParseIngredients200ResponseInnerNutritionWeightPerServing `json:"metric"`
	Us ParseIngredients200ResponseInnerNutritionWeightPerServing `json:"us"`
}

// NewGetShoppingList200ResponseAislesInnerItemsInnerMeasures instantiates a new GetShoppingList200ResponseAislesInnerItemsInnerMeasures object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetShoppingList200ResponseAislesInnerItemsInnerMeasures(original ParseIngredients200ResponseInnerNutritionWeightPerServing, metric ParseIngredients200ResponseInnerNutritionWeightPerServing, us ParseIngredients200ResponseInnerNutritionWeightPerServing) *GetShoppingList200ResponseAislesInnerItemsInnerMeasures {
	this := GetShoppingList200ResponseAislesInnerItemsInnerMeasures{}
	this.Original = original
	this.Metric = metric
	this.Us = us
	return &this
}

// NewGetShoppingList200ResponseAislesInnerItemsInnerMeasuresWithDefaults instantiates a new GetShoppingList200ResponseAislesInnerItemsInnerMeasures object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetShoppingList200ResponseAislesInnerItemsInnerMeasuresWithDefaults() *GetShoppingList200ResponseAislesInnerItemsInnerMeasures {
	this := GetShoppingList200ResponseAislesInnerItemsInnerMeasures{}
	return &this
}

// GetOriginal returns the Original field value
func (o *GetShoppingList200ResponseAislesInnerItemsInnerMeasures) GetOriginal() ParseIngredients200ResponseInnerNutritionWeightPerServing {
	if o == nil {
		var ret ParseIngredients200ResponseInnerNutritionWeightPerServing
		return ret
	}

	return o.Original
}

// GetOriginalOk returns a tuple with the Original field value
// and a boolean to check if the value has been set.
func (o *GetShoppingList200ResponseAislesInnerItemsInnerMeasures) GetOriginalOk() (*ParseIngredients200ResponseInnerNutritionWeightPerServing, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Original, true
}

// SetOriginal sets field value
func (o *GetShoppingList200ResponseAislesInnerItemsInnerMeasures) SetOriginal(v ParseIngredients200ResponseInnerNutritionWeightPerServing) {
	o.Original = v
}

// GetMetric returns the Metric field value
func (o *GetShoppingList200ResponseAislesInnerItemsInnerMeasures) GetMetric() ParseIngredients200ResponseInnerNutritionWeightPerServing {
	if o == nil {
		var ret ParseIngredients200ResponseInnerNutritionWeightPerServing
		return ret
	}

	return o.Metric
}

// GetMetricOk returns a tuple with the Metric field value
// and a boolean to check if the value has been set.
func (o *GetShoppingList200ResponseAislesInnerItemsInnerMeasures) GetMetricOk() (*ParseIngredients200ResponseInnerNutritionWeightPerServing, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Metric, true
}

// SetMetric sets field value
func (o *GetShoppingList200ResponseAislesInnerItemsInnerMeasures) SetMetric(v ParseIngredients200ResponseInnerNutritionWeightPerServing) {
	o.Metric = v
}

// GetUs returns the Us field value
func (o *GetShoppingList200ResponseAislesInnerItemsInnerMeasures) GetUs() ParseIngredients200ResponseInnerNutritionWeightPerServing {
	if o == nil {
		var ret ParseIngredients200ResponseInnerNutritionWeightPerServing
		return ret
	}

	return o.Us
}

// GetUsOk returns a tuple with the Us field value
// and a boolean to check if the value has been set.
func (o *GetShoppingList200ResponseAislesInnerItemsInnerMeasures) GetUsOk() (*ParseIngredients200ResponseInnerNutritionWeightPerServing, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Us, true
}

// SetUs sets field value
func (o *GetShoppingList200ResponseAislesInnerItemsInnerMeasures) SetUs(v ParseIngredients200ResponseInnerNutritionWeightPerServing) {
	o.Us = v
}

func (o GetShoppingList200ResponseAislesInnerItemsInnerMeasures) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["original"] = o.Original
	}
	if true {
		toSerialize["metric"] = o.Metric
	}
	if true {
		toSerialize["us"] = o.Us
	}
	return json.Marshal(toSerialize)
}

type NullableGetShoppingList200ResponseAislesInnerItemsInnerMeasures struct {
	value *GetShoppingList200ResponseAislesInnerItemsInnerMeasures
	isSet bool
}

func (v NullableGetShoppingList200ResponseAislesInnerItemsInnerMeasures) Get() *GetShoppingList200ResponseAislesInnerItemsInnerMeasures {
	return v.value
}

func (v *NullableGetShoppingList200ResponseAislesInnerItemsInnerMeasures) Set(val *GetShoppingList200ResponseAislesInnerItemsInnerMeasures) {
	v.value = val
	v.isSet = true
}

func (v NullableGetShoppingList200ResponseAislesInnerItemsInnerMeasures) IsSet() bool {
	return v.isSet
}

func (v *NullableGetShoppingList200ResponseAislesInnerItemsInnerMeasures) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetShoppingList200ResponseAislesInnerItemsInnerMeasures(val *GetShoppingList200ResponseAislesInnerItemsInnerMeasures) *NullableGetShoppingList200ResponseAislesInnerItemsInnerMeasures {
	return &NullableGetShoppingList200ResponseAislesInnerItemsInnerMeasures{value: val, isSet: true}
}

func (v NullableGetShoppingList200ResponseAislesInnerItemsInnerMeasures) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetShoppingList200ResponseAislesInnerItemsInnerMeasures) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


