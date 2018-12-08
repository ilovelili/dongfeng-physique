package models

// User user profile entity
type User struct {
	ID      string `json:"id" dapper:"id,primarykey,table=users"`
	Email   string `json:"email" dapper:"email"`
	Name    string `json:"name" dapper:"name"`
	Avatar  string `json:"picture,omitempty" dapper:"avatar"`
	Setting int32  `dapper:"settings"`
	Role    string `json:"role" dapper:"role"`
}
