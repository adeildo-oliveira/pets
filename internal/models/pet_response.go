package models

type PetResponse struct {
	ID       int64  `json:"id"`
	Name     string `json:"name"`
	Age      int    `json:"age"`
	Breed    string `json:"breed"`
	NickName string `json:"nickName"`
	Status   bool   `json:"status"`
}
