//**********************************************************
//
// Copyright (C) 2018 - 2021 J&J Ideenschmiede UG (haftungsbeschränkt) <info@jj-ideenschmiede.de>
//
// This file is part of gotrello.
// All code may be used. Feel free and maybe code something better.
//
// Author: Jonas Kwiedor
//
//**********************************************************

package gotrello

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

// To decode the trello card response
type Cart struct {
	ID                    string        `json:"id"`
	CheckItemStates       []interface{} `json:"checkItemStates"`
	Closed                bool          `json:"closed"`
	DateLastActivity      string        `json:"dateLastActivity"`
	Desc                  string        `json:"desc"`
	DescData              interface{}   `json:"descData"`
	DueReminder           interface{}   `json:"dueReminder"`
	IDBoard               string        `json:"idBoard"`
	IDList                string        `json:"idList"`
	IDMembersVoted        []interface{} `json:"idMembersVoted"`
	IDShort               int           `json:"idShort"`
	IDAttachmentCover     interface{}   `json:"idAttachmentCover"`
	IDLabels              []interface{} `json:"idLabels"`
	ManualCoverAttachment bool          `json:"manualCoverAttachment"`
	Name                  string        `json:"name"`
	Pos                   int           `json:"pos"`
	ShortLink             string        `json:"shortLink"`
	IsTemplate            bool          `json:"isTemplate"`
	CardRole              interface{}   `json:"cardRole"`
	DueComplete           bool          `json:"dueComplete"`
	Due                   string        `json:"due"`
	Email                 interface{}   `json:"email"`
	Labels                []interface{} `json:"labels"`
	ShortUrl              string        `json:"shortUrl"`
	Start                 interface{}   `json:"start"`
	Url                   string        `json:"url"`
	Cover                 Cover         `json:"cover"`
	IDMembers             []interface{} `json:"idMembers"`
	Attachments           []interface{} `json:"attachments"`
	Badges                interface{}   `json:"badges"`
	Subscribed            bool          `json:"subscribed"`
	IDChecklists          []interface{} `json:"idChecklists"`
	Stickers              []interface{} `json:"stickers"`
	Limits                interface{}   `json:"limits"`
}

type Cover struct {
	IDAttachment         interface{} `json:"idAttachment"`
	Color                interface{} `json:"color"`
	IDUploadedBackground interface{} `json:"idUploadedBackground"`
	Size                 string      `json:"size"`
	Brightness           string      `json:"brightness"`
	IDPlugin             interface{} `json:"idPlugin"`
}

type Badges struct {
	AttachmentsByType     AttachmentsByType `json:"attachmentsByType"`
	Location              bool              `json:"location"`
	Votes                 int               `json:"votes"`
	ViewingMemberVoted    bool              `json:"viewingMemberVoted"`
	Subscribed            bool              `json:"subscribed"`
	Fogbugz               string            `json:"fogbugz"`
	CheckItems            int               `json:"checkItems"`
	CheckItemsChecked     int               `json:"checkItemsChecked"`
	CheckItemsEarliestDue interface{}       `json:"checkItemsEarliestDue"`
	Comments              int               `json:"comments"`
	Attachments           int               `json:"attachments"`
	Description           bool              `json:"description"`
	Due                   string            `json:"due"`
	DueComplete           bool              `json:"dueComplete"`
	Start                 interface{}       `json:"start"`
}

type AttachmentsByType struct {
	Trello Trello `json:"trello"`
}

type Trello struct {
	Board int `json:"board"`
	Card  int `json:"card"`
}

type CartMember struct {
	ID                 string      `json:"id"`
	Username           string      `json:"username"`
	ActivityBlocked    bool        `json:"activityBlocked"`
	AvatarHash         string      `json:"avatarHash"`
	AvatarUrl          string      `json:"avatarUrl"`
	FullName           string      `json:"fullName"`
	IDMemberReferrer   interface{} `json:"idMemberReferrer"`
	Initials           string      `json:"initials"`
	NonPublic          interface{} `json:"nonPublic"`
	NonPublicAvailable bool        `Ajson:"nonPublicAvailable"`
}

// To create an cart
func CreateCart(key, token, list, name, description, date string) (Cart, error) {

	// Get request
	request, err := http.NewRequest("POST", fmt.Sprintf("https://api.trello.com/1/cards?key=%s&token=%s&idList=%s&name=%s&desc=%s&due=%s", key, token, list, url.QueryEscape(name), url.QueryEscape(description), date), nil)
	if err != nil {
		return Cart{}, err
	}

	// Create client
	client := &http.Client{}

	// Send request
	response, err := client.Do(request)
	if err != nil {
		return Cart{}, err
	}

	// Close response body
	defer response.Body.Close()

	// Decode json data
	var decode Cart

	err = json.NewDecoder(response.Body).Decode(&decode)
	if err != nil {
		return Cart{}, err
	}

	// Return data
	return decode, nil

}

// Create new card member
func CreateCardMember(key, token, card, member string) ([]CartMember, error) {

	// Get request
	request, err := http.NewRequest("POST", fmt.Sprintf("https://api.trello.com/1/cards/%s/idMembers?key=%s&token=%s&value=%s", card, key, token, member), nil)
	if err != nil {
		return nil, err
	}

	// Create client
	client := &http.Client{}

	// Send request
	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}

	// Close response body
	defer response.Body.Close()

	// Decode json data
	var decode []CartMember

	err = json.NewDecoder(response.Body).Decode(&decode)
	if err != nil {
		return nil, err
	}

	// Return data
	return decode, nil

}
