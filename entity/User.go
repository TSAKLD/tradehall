package entity

type Status string

const (
	StatusBlocked    Status = "Blocked"
	StatusCommonUser Status = "CommonUser"
	StatusPremium    Status = "Premium"
	StatusAdmin      Status = "Admin"
)

type User struct {
	ID       int     `json:"id"`
	Nickname string  `json:"nickname"`
	Password string  `json:"password"`
	Email    string  `json:"email"`
	Age      int     `json:"age"`
	Inviter  *string `json:"inviter,omitempty"`
	Status   Status  `json:"status"`
}
type UserForFind struct {
	ID       *int    `json:"id"`
	Nickname *string `json:"nickname"`
	Email    *string `json:"email" json:"email,omitempty"`
	Status   *Status `json:"status" json:"status,omitempty"`
}
