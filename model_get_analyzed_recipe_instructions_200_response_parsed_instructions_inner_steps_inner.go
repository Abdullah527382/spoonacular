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

// GetAnalyzedRecipeInstructions200ResponseParsedInstructionsInnerStepsInner struct for GetAnalyzedRecipeInstructions200ResponseParsedInstructionsInnerStepsInner
type GetAnalyzedRecipeInstructions200ResponseParsedInstructionsInnerStepsInner struct {
	Number float32 `json:"number"`
	Step string `json:"step"`
	Ingredients []GetAnalyzedRecipeInstructions200ResponseParsedInstructionsInnerStepsInnerIngredientsInner `json:"ingredients,omitempty"`
	Equipment []GetAnalyzedRecipeInstructions200ResponseParsedInstructionsInnerStepsInnerIngredientsInner `json:"equipment,omitempty"`
}

// NewGetAnalyzedRecipeInstructions200ResponseParsedInstructionsInnerStepsInner instantiates a new GetAnalyzedRecipeInstructions200ResponseParsedInstructionsInnerStepsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetAnalyzedRecipeInstructions200ResponseParsedInstructionsInnerStepsInner(number float32, step string) *GetAnalyzedRecipeInstructions200ResponseParsedInstructionsInnerStepsInner {
	this := GetAnalyzedRecipeInstructions200ResponseParsedInstructionsInnerStepsInner{}
	this.Number = number
	this.Step = step
	return &this
}

// NewGetAnalyzedRecipeInstructions200ResponseParsedInstructionsInnerStepsInnerWithDefaults instantiates a new GetAnalyzedRecipeInstructions200ResponseParsedInstructionsInnerStepsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetAnalyzedRecipeInstructions200ResponseParsedInstructionsInnerStepsInnerWithDefaults() *GetAnalyzedRecipeInstructions200ResponseParsedInstructionsInnerStepsInner {
	this := GetAnalyzedRecipeInstructions200ResponseParsedInstructionsInnerStepsInner{}
	return &this
}

// GetNumber returns the Number field value
func (o *GetAnalyzedRecipeInstructions200ResponseParsedInstructionsInnerStepsInner) GetNumber() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Number
}

// GetNumberOk returns a tuple with the Number field value
// and a boolean to check if the value has been set.
func (o *GetAnalyzedRecipeInstructions200ResponseParsedInstructionsInnerStepsInner) GetNumberOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Number, true
}

// SetNumber sets field value
func (o *GetAnalyzedRecipeInstructions200ResponseParsedInstructionsInnerStepsInner) SetNumber(v float32) {
	o.Number = v
}

// GetStep returns the Step field value
func (o *GetAnalyzedRecipeInstructions200ResponseParsedInstructionsInnerStepsInner) GetStep() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Step
}

// GetStepOk returns a tuple with the Step field value
// and a boolean to check if the value has been set.
func (o *GetAnalyzedRecipeInstructions200ResponseParsedInstructionsInnerStepsInner) GetStepOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Step, true
}

// SetStep sets field value
func (o *GetAnalyzedRecipeInstructions200ResponseParsedInstructionsInnerStepsInner) SetStep(v string) {
	o.Step = v
}

// GetIngredients returns the Ingredients field value if set, zero value otherwise.
func (o *GetAnalyzedRecipeInstructions200ResponseParsedInstructionsInnerStepsInner) GetIngredients() []GetAnalyzedRecipeInstructions200ResponseParsedInstructionsInnerStepsInnerIngredientsInner {
	if o == nil || o.Ingredients == nil {
		var ret []GetAnalyzedRecipeInstructions200ResponseParsedInstructionsInnerStepsInnerIngredientsInner
		return ret
	}
	return o.Ingredients
}

// GetIngredientsOk returns a tuple with the Ingredients field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAnalyzedRecipeInstructions200ResponseParsedInstructionsInnerStepsInner) GetIngredientsOk() ([]GetAnalyzedRecipeInstructions200ResponseParsedInstructionsInnerStepsInnerIngredientsInner, bool) {
	if o == nil || o.Ingredients == nil {
		return nil, false
	}
	return o.Ingredients, true
}

// HasIngredients returns a boolean if a field has been set.
func (o *GetAnalyzedRecipeInstructions200ResponseParsedInstructionsInnerStepsInner) HasIngredients() bool {
	if o != nil && o.Ingredients != nil {
		return true
	}

	return false
}

// SetIngredients gets a reference to the given []GetAnalyzedRecipeInstructions200ResponseParsedInstructionsInnerStepsInnerIngredientsInner and assigns it to the Ingredients field.
func (o *GetAnalyzedRecipeInstructions200ResponseParsedInstructionsInnerStepsInner) SetIngredients(v []GetAnalyzedRecipeInstructions200ResponseParsedInstructionsInnerStepsInnerIngredientsInner) {
	o.Ingredients = v
}

// GetEquipment returns the Equipment field value if set, zero value otherwise.
func (o *GetAnalyzedRecipeInstructions200ResponseParsedInstructionsInnerStepsInner) GetEquipment() []GetAnalyzedRecipeInstructions200ResponseParsedInstructionsInnerStepsInnerIngredientsInner {
	if o == nil || o.Equipment == nil {
		var ret []GetAnalyzedRecipeInstructions200ResponseParsedInstructionsInnerStepsInnerIngredientsInner
		return ret
	}
	return o.Equipment
}

// GetEquipmentOk returns a tuple with the Equipment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAnalyzedRecipeInstructions200ResponseParsedInstructionsInnerStepsInner) GetEquipmentOk() ([]GetAnalyzedRecipeInstructions200ResponseParsedInstructionsInnerStepsInnerIngredientsInner, bool) {
	if o == nil || o.Equipment == nil {
		return nil, false
	}
	return o.Equipment, true
}

// HasEquipment returns a boolean if a field has been set.
func (o *GetAnalyzedRecipeInstructions200ResponseParsedInstructionsInnerStepsInner) HasEquipment() bool {
	if o != nil && o.Equipment != nil {
		return true
	}

	return false
}

// SetEquipment gets a reference to the given []GetAnalyzedRecipeInstructions200ResponseParsedInstructionsInnerStepsInnerIngredientsInner and assigns it to the Equipment field.
func (o *GetAnalyzedRecipeInstructions200ResponseParsedInstructionsInnerStepsInner) SetEquipment(v []GetAnalyzedRecipeInstructions200ResponseParsedInstructionsInnerStepsInnerIngredientsInner) {
	o.Equipment = v
}

func (o GetAnalyzedRecipeInstructions200ResponseParsedInstructionsInnerStepsInner) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["number"] = o.Number
	}
	if true {
		toSerialize["step"] = o.Step
	}
	if o.Ingredients != nil {
		toSerialize["ingredients"] = o.Ingredients
	}
	if o.Equipment != nil {
		toSerialize["equipment"] = o.Equipment
	}
	return json.Marshal(toSerialize)
}

type NullableGetAnalyzedRecipeInstructions200ResponseParsedInstructionsInnerStepsInner struct {
	value *GetAnalyzedRecipeInstructions200ResponseParsedInstructionsInnerStepsInner
	isSet bool
}

func (v NullableGetAnalyzedRecipeInstructions200ResponseParsedInstructionsInnerStepsInner) Get() *GetAnalyzedRecipeInstructions200ResponseParsedInstructionsInnerStepsInner {
	return v.value
}

func (v *NullableGetAnalyzedRecipeInstructions200ResponseParsedInstructionsInnerStepsInner) Set(val *GetAnalyzedRecipeInstructions200ResponseParsedInstructionsInnerStepsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetAnalyzedRecipeInstructions200ResponseParsedInstructionsInnerStepsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetAnalyzedRecipeInstructions200ResponseParsedInstructionsInnerStepsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetAnalyzedRecipeInstructions200ResponseParsedInstructionsInnerStepsInner(val *GetAnalyzedRecipeInstructions200ResponseParsedInstructionsInnerStepsInner) *NullableGetAnalyzedRecipeInstructions200ResponseParsedInstructionsInnerStepsInner {
	return &NullableGetAnalyzedRecipeInstructions200ResponseParsedInstructionsInnerStepsInner{value: val, isSet: true}
}

func (v NullableGetAnalyzedRecipeInstructions200ResponseParsedInstructionsInnerStepsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetAnalyzedRecipeInstructions200ResponseParsedInstructionsInnerStepsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

