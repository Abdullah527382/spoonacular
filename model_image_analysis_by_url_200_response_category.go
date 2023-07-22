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

// ImageAnalysisByURL200ResponseCategory struct for ImageAnalysisByURL200ResponseCategory
type ImageAnalysisByURL200ResponseCategory struct {
	Name string `json:"name"`
	Probability float32 `json:"probability"`
}

// NewImageAnalysisByURL200ResponseCategory instantiates a new ImageAnalysisByURL200ResponseCategory object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewImageAnalysisByURL200ResponseCategory(name string, probability float32) *ImageAnalysisByURL200ResponseCategory {
	this := ImageAnalysisByURL200ResponseCategory{}
	this.Name = name
	this.Probability = probability
	return &this
}

// NewImageAnalysisByURL200ResponseCategoryWithDefaults instantiates a new ImageAnalysisByURL200ResponseCategory object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewImageAnalysisByURL200ResponseCategoryWithDefaults() *ImageAnalysisByURL200ResponseCategory {
	this := ImageAnalysisByURL200ResponseCategory{}
	return &this
}

// GetName returns the Name field value
func (o *ImageAnalysisByURL200ResponseCategory) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ImageAnalysisByURL200ResponseCategory) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ImageAnalysisByURL200ResponseCategory) SetName(v string) {
	o.Name = v
}

// GetProbability returns the Probability field value
func (o *ImageAnalysisByURL200ResponseCategory) GetProbability() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Probability
}

// GetProbabilityOk returns a tuple with the Probability field value
// and a boolean to check if the value has been set.
func (o *ImageAnalysisByURL200ResponseCategory) GetProbabilityOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Probability, true
}

// SetProbability sets field value
func (o *ImageAnalysisByURL200ResponseCategory) SetProbability(v float32) {
	o.Probability = v
}

func (o ImageAnalysisByURL200ResponseCategory) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["probability"] = o.Probability
	}
	return json.Marshal(toSerialize)
}

type NullableImageAnalysisByURL200ResponseCategory struct {
	value *ImageAnalysisByURL200ResponseCategory
	isSet bool
}

func (v NullableImageAnalysisByURL200ResponseCategory) Get() *ImageAnalysisByURL200ResponseCategory {
	return v.value
}

func (v *NullableImageAnalysisByURL200ResponseCategory) Set(val *ImageAnalysisByURL200ResponseCategory) {
	v.value = val
	v.isSet = true
}

func (v NullableImageAnalysisByURL200ResponseCategory) IsSet() bool {
	return v.isSet
}

func (v *NullableImageAnalysisByURL200ResponseCategory) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableImageAnalysisByURL200ResponseCategory(val *ImageAnalysisByURL200ResponseCategory) *NullableImageAnalysisByURL200ResponseCategory {
	return &NullableImageAnalysisByURL200ResponseCategory{value: val, isSet: true}
}

func (v NullableImageAnalysisByURL200ResponseCategory) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableImageAnalysisByURL200ResponseCategory) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

