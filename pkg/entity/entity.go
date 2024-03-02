package entity

type User struct {
	ID            int64
	FirstName     string `json:"first_name"`
	Lastname      string `json:"last_name"`
	EncryptedPass string `json:"encrypted_password"`
	AvatarLink    string `json:"avatar_link"`
	Gender        string `json:"gender"`
	Age           int64  `json:"age"`
	PhoneNum      string `json:"phone_number"`
	ResidenceCity string `json:"city_of_residence"`
	Description   string `json:"description"`
}
