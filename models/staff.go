package models

import "time"

type Staff struct {
	Email          []interface{} `json:"email"`
	Phone          []interface{} `json:"phone"`
	Active         bool          `json:"active"`
	IsTransfer     bool          `json:"isTransfer"`
	LastActiveDate time.Time     `json:"last_active_date"`
	Skill          []struct {
		Skill string `json:"skill"`
		Level int    `json:"level"`
	} `json:"skill"`
	Certificate       []interface{} `json:"certificate"`
	ID                string        `json:"id"`
	Prefix            string        `json:"prefix"`
	Fname             string        `json:"fname"`
	Lname             string        `json:"lname"`
	Nname             string        `json:"nname"`
	Center            string        `json:"center"`
	Team              string        `json:"team"`
	StartDate         time.Time     `json:"start_date"`
	CreatedAt         time.Time     `json:"createdAt"`
	UpdatedAt         time.Time     `json:"updatedAt"`
	AccountID         string        `json:"account_id"`
	OneEmail          string        `json:"one_email"`
	LeaveDetail       interface{}   `json:"leaveDetail"`
	ResignDescription string        `json:"resign_description"`
}

type StaffResponse struct {
	Obj_ID         string        `json:"_id,omitempty" bson:"_id,omitempty"`
	Email          []interface{} `json:"email" bson:"email,omitempty"`
	Phone          []interface{} `json:"phone" bson:"phone,omitempty"`
	Active         bool          `json:"active" bson:"active,omitempty"`
	IsTransfer     bool          `json:"isTransfer" bson:"isTransfer,omitempty"`
	LastActiveDate time.Time     `json:"last_active_date" bson:"last_active_date,omitempty"`
	Skill          []struct {
		Skill string `json:"skill" bson:"skill,omitempty"`
		Level int    `json:"level" bson:"level,omitempty"`
	} `json:"skill"`
	Certificate       []interface{} `json:"certificate" bson:"certificate,omitempty"`
	ID                string        `json:"id" bson:"id,omitempty"`
	Prefix            string        `json:"prefix" bson:"prefix,omitempty"`
	Fname             string        `json:"fname" bson:"fname,omitempty"`
	Lname             string        `json:"lname" bson:"lname,omitempty"`
	Nname             string        `json:"nname" bson:"nname,omitempty"`
	Center            string        `json:"center" bson:"center,omitempty"`
	Team              string        `json:"team" bson:"team,omitempty"`
	AccountID         string        `json:"account_id" bson:"account_id,omitempty"`
	OneEmail          string        `json:"one_email" bson:"one_email,omitempty"`
	LeaveDetail       interface{}   `json:"leaveDetail" bson:"leaveDetail,omitempty"`
	ResignDescription string        `json:"resign_description" bson:"resign_description,omitempty"`
}
