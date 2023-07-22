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

// GetRecipeInformation200ResponseWinePairingProductMatchesInner struct for GetRecipeInformation200ResponseWinePairingProductMatchesInner
type GetRecipeInformation200ResponseWinePairingProductMatchesInner struct {
	Id int32 `json:"id"`
	Title string `json:"title"`
	Description string `json:"description"`
	Price string `json:"price"`
	ImageUrl string `json:"imageUrl"`
	AverageRating float32 `json:"averageRating"`
	RatingCount int32 `json:"ratingCount"`
	Score float32 `json:"score"`
	Link string `json:"link"`
}

// NewGetRecipeInformation200ResponseWinePairingProductMatchesInner instantiates a new GetRecipeInformation200ResponseWinePairingProductMatchesInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetRecipeInformation200ResponseWinePairingProductMatchesInner(id int32, title string, description string, price string, imageUrl string, averageRating float32, ratingCount int32, score float32, link string) *GetRecipeInformation200ResponseWinePairingProductMatchesInner {
	this := GetRecipeInformation200ResponseWinePairingProductMatchesInner{}
	this.Id = id
	this.Title = title
	this.Description = description
	this.Price = price
	this.ImageUrl = imageUrl
	this.AverageRating = averageRating
	this.RatingCount = ratingCount
	this.Score = score
	this.Link = link
	return &this
}

// NewGetRecipeInformation200ResponseWinePairingProductMatchesInnerWithDefaults instantiates a new GetRecipeInformation200ResponseWinePairingProductMatchesInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetRecipeInformation200ResponseWinePairingProductMatchesInnerWithDefaults() *GetRecipeInformation200ResponseWinePairingProductMatchesInner {
	this := GetRecipeInformation200ResponseWinePairingProductMatchesInner{}
	return &this
}

// GetId returns the Id field value
func (o *GetRecipeInformation200ResponseWinePairingProductMatchesInner) GetId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *GetRecipeInformation200ResponseWinePairingProductMatchesInner) GetIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *GetRecipeInformation200ResponseWinePairingProductMatchesInner) SetId(v int32) {
	o.Id = v
}

// GetTitle returns the Title field value
func (o *GetRecipeInformation200ResponseWinePairingProductMatchesInner) GetTitle() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Title
}

// GetTitleOk returns a tuple with the Title field value
// and a boolean to check if the value has been set.
func (o *GetRecipeInformation200ResponseWinePairingProductMatchesInner) GetTitleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Title, true
}

// SetTitle sets field value
func (o *GetRecipeInformation200ResponseWinePairingProductMatchesInner) SetTitle(v string) {
	o.Title = v
}

// GetDescription returns the Description field value
func (o *GetRecipeInformation200ResponseWinePairingProductMatchesInner) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *GetRecipeInformation200ResponseWinePairingProductMatchesInner) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value
func (o *GetRecipeInformation200ResponseWinePairingProductMatchesInner) SetDescription(v string) {
	o.Description = v
}

// GetPrice returns the Price field value
func (o *GetRecipeInformation200ResponseWinePairingProductMatchesInner) GetPrice() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Price
}

// GetPriceOk returns a tuple with the Price field value
// and a boolean to check if the value has been set.
func (o *GetRecipeInformation200ResponseWinePairingProductMatchesInner) GetPriceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Price, true
}

// SetPrice sets field value
func (o *GetRecipeInformation200ResponseWinePairingProductMatchesInner) SetPrice(v string) {
	o.Price = v
}

// GetImageUrl returns the ImageUrl field value
func (o *GetRecipeInformation200ResponseWinePairingProductMatchesInner) GetImageUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ImageUrl
}

// GetImageUrlOk returns a tuple with the ImageUrl field value
// and a boolean to check if the value has been set.
func (o *GetRecipeInformation200ResponseWinePairingProductMatchesInner) GetImageUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ImageUrl, true
}

// SetImageUrl sets field value
func (o *GetRecipeInformation200ResponseWinePairingProductMatchesInner) SetImageUrl(v string) {
	o.ImageUrl = v
}

// GetAverageRating returns the AverageRating field value
func (o *GetRecipeInformation200ResponseWinePairingProductMatchesInner) GetAverageRating() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.AverageRating
}

// GetAverageRatingOk returns a tuple with the AverageRating field value
// and a boolean to check if the value has been set.
func (o *GetRecipeInformation200ResponseWinePairingProductMatchesInner) GetAverageRatingOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AverageRating, true
}

// SetAverageRating sets field value
func (o *GetRecipeInformation200ResponseWinePairingProductMatchesInner) SetAverageRating(v float32) {
	o.AverageRating = v
}

// GetRatingCount returns the RatingCount field value
func (o *GetRecipeInformation200ResponseWinePairingProductMatchesInner) GetRatingCount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.RatingCount
}

// GetRatingCountOk returns a tuple with the RatingCount field value
// and a boolean to check if the value has been set.
func (o *GetRecipeInformation200ResponseWinePairingProductMatchesInner) GetRatingCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RatingCount, true
}

// SetRatingCount sets field value
func (o *GetRecipeInformation200ResponseWinePairingProductMatchesInner) SetRatingCount(v int32) {
	o.RatingCount = v
}

// GetScore returns the Score field value
func (o *GetRecipeInformation200ResponseWinePairingProductMatchesInner) GetScore() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Score
}

// GetScoreOk returns a tuple with the Score field value
// and a boolean to check if the value has been set.
func (o *GetRecipeInformation200ResponseWinePairingProductMatchesInner) GetScoreOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Score, true
}

// SetScore sets field value
func (o *GetRecipeInformation200ResponseWinePairingProductMatchesInner) SetScore(v float32) {
	o.Score = v
}

// GetLink returns the Link field value
func (o *GetRecipeInformation200ResponseWinePairingProductMatchesInner) GetLink() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Link
}

// GetLinkOk returns a tuple with the Link field value
// and a boolean to check if the value has been set.
func (o *GetRecipeInformation200ResponseWinePairingProductMatchesInner) GetLinkOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Link, true
}

// SetLink sets field value
func (o *GetRecipeInformation200ResponseWinePairingProductMatchesInner) SetLink(v string) {
	o.Link = v
}

func (o GetRecipeInformation200ResponseWinePairingProductMatchesInner) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["title"] = o.Title
	}
	if true {
		toSerialize["description"] = o.Description
	}
	if true {
		toSerialize["price"] = o.Price
	}
	if true {
		toSerialize["imageUrl"] = o.ImageUrl
	}
	if true {
		toSerialize["averageRating"] = o.AverageRating
	}
	if true {
		toSerialize["ratingCount"] = o.RatingCount
	}
	if true {
		toSerialize["score"] = o.Score
	}
	if true {
		toSerialize["link"] = o.Link
	}
	return json.Marshal(toSerialize)
}

type NullableGetRecipeInformation200ResponseWinePairingProductMatchesInner struct {
	value *GetRecipeInformation200ResponseWinePairingProductMatchesInner
	isSet bool
}

func (v NullableGetRecipeInformation200ResponseWinePairingProductMatchesInner) Get() *GetRecipeInformation200ResponseWinePairingProductMatchesInner {
	return v.value
}

func (v *NullableGetRecipeInformation200ResponseWinePairingProductMatchesInner) Set(val *GetRecipeInformation200ResponseWinePairingProductMatchesInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetRecipeInformation200ResponseWinePairingProductMatchesInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetRecipeInformation200ResponseWinePairingProductMatchesInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetRecipeInformation200ResponseWinePairingProductMatchesInner(val *GetRecipeInformation200ResponseWinePairingProductMatchesInner) *NullableGetRecipeInformation200ResponseWinePairingProductMatchesInner {
	return &NullableGetRecipeInformation200ResponseWinePairingProductMatchesInner{value: val, isSet: true}
}

func (v NullableGetRecipeInformation200ResponseWinePairingProductMatchesInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetRecipeInformation200ResponseWinePairingProductMatchesInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


