package model

type Admin struct {
	ID           string `json:"_id,omitempty" bson:"_id,omitempty"`
	Name         string `json:"name,omitempty" validate:"required" bson:"name,omitempty"`
	Password     string `json:"password,omitempty" validate:"required" bson:"password,omitempty"`
	Role         string `json:"role,omitempty"  bson:"role,omitempty" default:"Admin"`
	Email        string `json:"email,omitempty" validate:"required" bson:"email,omitempty"`
	SchoolName   string `json:"schoolName,omitempty" validate:"required" bson:"schoolName,omitempty"`
	BytePassword []byte `json:"bytepassword,omitempty"  bson:"bytepassword,omitempty"`
	Token        string `json:"token,omitempty"`
}
