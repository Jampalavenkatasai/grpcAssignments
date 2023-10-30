// pkg/user/user.go
package user

import (
	"encoding/json"
)

type User struct {
	Id      int32   `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Fname   string  `protobuf:"bytes,2,opt,name=fname,proto3" json:"fname,omitempty"`
	City    string  `protobuf:"bytes,3,opt,name=city,proto3" json:"city,omitempty"`
	Phone   int64   `protobuf:"varint,4,opt,name=phone,proto3" json:"phone,omitempty"`
	Height  float64 `protobuf:"fixed64,5,opt,name=height,proto3" json:"height,omitempty"`
	Married bool    `protobuf:"varint,6,opt,name=married,proto3" json:"married,omitempty"`
}

// NewUser creates a new User instance with the provided data.
func NewUser(id int32, fname, city string, phone int64, height float64, married bool) *User {
	return &User{
		Id:      id,
		Fname:   fname,
		City:    city,
		Phone:   phone,
		Height:  height,
		Married: married,
	}
}

// SerializeToJSON serializes the User to a JSON string.
func (u *User) SerializeToJSON() (string, error) {
	userJSON, err := json.Marshal(u)
	if err != nil {
		return "", err
	}
	return string(userJSON), nil
}

// DeserializeFromJSON deserializes a JSON string into a User instance.
func (u *User) DeserializeFromJSON(jsonStr string) error {
	return json.Unmarshal([]byte(jsonStr), u)
}
