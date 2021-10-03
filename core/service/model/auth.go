package model

import "octopus/core/helper"

type Auth struct {
	Model
	Media       string `json:"media"`        // 平台
	AccessToken string `json:"access_token"` // 授权
}

// MarshalBinary for writing to redis
func (a Auth) MarshalBinary() ([]byte, error) {
	return helper.NewJson().Marshal(a)
}

// UnmarshalBinary for reading from redis
func (a Auth) UnmarshalBinary(data []byte) error {
	return helper.NewJson().Unmarshal(data, &a)
}
