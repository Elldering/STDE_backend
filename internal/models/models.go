package models

import "time"

type ModelDb struct {
	Data    DataModels    `json:"data"`
	Request RequestModels `json:"request"`
	//Response ResponseModels `json:"response"`
}

type DataModels struct {
	AuthUser             AuthUser             `json:"authUser"`
	AuthGroupPermissions AuthGroupPermissions `json:"authGroupPermissions"`
	AuthUserGroups       AuthUserGroups       `json:"authUserGroups"`
	Basket               Basket               `json:"basket"`
	DocumentAuthUser     DocumentAuthUser     `json:"documentAuthUser"`
	Menu                 Menu                 `json:"menu"`
	Order                Order                `json:"order"`
	OrderPosition        OrderPosition        `json:"orderPosition"`
	Permission           Permission           `json:"permission"`
	Position             Position             `json:"position"`
	VerifyCode           VerifyCode           `json:"verifyCode"`
	UserProfile          UserProfile          `json:"userProfile"`
	Reviews              Reviews              `json:"reviews"`
	UserDocument         UserDocument         `json:"userDocument"`
	UserGroup            UserGroup            `json:"userGroup"`
}

type RequestModels struct {
	LogoutRequest   LogoutRequest   `json:"logoutRequest"`
	Register        Register        `json:"register"`
	AuthUserRequest AuthUserRequest `json:"authUserRequest"`
}

type AuthUser struct {
	ID             int       `json:"id"`
	Password       string    `json:"password"`
	LastLogin      time.Time `json:"last_login"`
	Email          string    `json:"email"`
	PhoneNumber    string    `json:"phone_number"`
	IsSuperUser    bool      `json:"is_super_user"`
	IsEmail        bool      `json:"is_email"`
	IsNumberVerify bool      `json:"is_number_verify"`
	IsBlocked      bool      `json:"is_blocked"`
	IsFrozen       bool      `json:"is_frozen"`
	DateJoined     time.Time `json:"date_joined"`
	UserProfileID  int       `json:"user_profile_id"`
	Balance        float64   `json:"balance"`
	UserRID        int       `json:"user_r_id"`
}

type AuthUserRequest struct {
	ID             int       `json:"id"`
	Password       string    `json:"password"`
	Email          string    `json:"email"`
	Login          string    `json:"login"`
	TypeLogin      string    `json:"type_login"`
	IsNumberVerify bool      `json:"is_number_verify"`
	RefreshToken   string    `json:"refresh_token"`
	UsernameClaim  string    `json:"username_claim"`
	LastLogin      time.Time `json:"last_login"`
	IsEmail        bool      `json:"is_email"`
}

type AuthGroupPermissions struct {
	ID           int64 `json:"id"`            // bigserial в базе данных
	GroupID      int   `json:"group_id"`      // integer в базе данных
	PermissionID int   `json:"permission_id"` // integer в базе данных
}

type AuthUserGroups struct {
	ID      int64 `json:"id"`       // bigserial в базе данных
	GroupID int   `json:"group_id"` // integer в базе данных
	UserID  int   `json:"user_id"`  // integer в базе данных
}

type Basket struct {
	ID         int64 `json:"id"`           // bigserial в базе данных
	AuthUserID int   `json:"auth_user_id"` // integer в базе данных
	PositionID int   `json:"position_id"`  // integer в базе данных
}

type DocumentAuthUser struct {
	ID             int64 `json:"id"`               // bigserial в базе данных
	AuthUserID     int   `json:"auth_user_id"`     // integer в базе данных
	UserDocumentID int   `json:"user_document_id"` // integer в базе данных
}

type Menu struct {
	ID         int64  `json:"id"`
	Name       string `json:"name"`
	PositionID int    `json:"position_id"`
	AuthUserID int    `json:"auth_user_id"`
}

type Order struct {
	ID                int       `json:"id"`
	TotalPrice        float64   `json:"total_price"`
	Status            string    `json:"status"`
	DeliveryAddress   string    `json:"delivery_address"`
	CreatedAt         time.Time `json:"created_at"`
	ClosedAt          time.Time `json:"closed_at"`
	AuthUserID        int       `json:"auth_user_id"`
	TrackNumber       string    `json:"track_number"`
	DelivererAccepted bool      `json:"deliverer_accepted"`
}

type OrderPosition struct {
	ID         int64 `json:"id"`
	OrderID    int   `json:"order_id"`
	PositionID int   `json:"position_id"`
}

type Permission struct {
	ID          int    `json:"id"`
	Codename    string `json:"codename"`
	Description string `json:"description"`
}

type Position struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	Available   bool    `json:"available"`
	Image       string  `json:"image"`
}

type LogoutRequest struct {
	Refresh string `json:"refresh"`
}

type Register struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

type Reviews struct {
	ID                  int       `json:"id"`
	AuthUserSenderID    int       `json:"auth_user_sender_id"`
	AuthUserRecipientID int       `json:"auth_user_recipient_id"`
	Grade               int16     `json:"grade"`
	Comment             string    `json:"comment"`
	CreatedAt           time.Time `json:"created_at"`
}

type UserDocument struct {
	ID         int    `json:"id"`
	Name       string `json:"name"`
	Image      string `json:"image"`
	IsAccepted bool   `json:"is_accepted"`
	Type       string `json:"type"`
}

type UserGroup struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type UserProfile struct {
	ID           int     `json:"id"`
	FirstName    string  `json:"first_name"`
	LastName     string  `json:"last_name"`
	MiddleName   string  `json:"middle_name"`
	Rating       float32 `json:"rating"`
	ProfileImage string  `json:"profile_image"`
}

type VerifyCode struct {
	Code int    `json:"code"`
	Type string `json:"type"`
}
