package models

type UserResponse struct {
	ID              string   `json:"id,omitempty" bson:"_id,omitempty"`
	Role            string   `json:"role"`
	Email           []string `json:"email"`
	Phone           []string `json:"phone"`
	ApproveStatus   bool     `json:"approve_status"`
	AccountID       string   `json:"account_id" bson:"account_id,omitempty"`
	StaffId         string   `json:"staff_id,omitempty" bson:"staff_id,omitempty"`
	FirstNameTh     string   `json:"first_name_th" bson:"first_name_th,omitempty"`
	LastNameTh      string   `json:"last_name_th"  bson:"last_name_th,omitempty"`
	FirstNameEng    string   `json:"first_name_eng"  bson:"first_name_eng,omitempty"`
	LastNameEng     string   `json:"last_name_eng"  bson:"last_name_eng,omitempty"`
	AccountTitleTh  string   `json:"account_title_th"  bson:"account_title_th,omitempty"`
	AccountTitleEng string   `json:"account_title_eng"  bson:"account_title_eng,omitempty"`
}
