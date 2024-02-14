package requests

type Create struct {
	Name    string `json:"name"`
	Age     int    `json:"age"`
	Friends []int  `json:"friends"`
}

type MakeFriends struct {
	Source_id int `json:"source_id"`
	Target_id int `json:"target_id"`
}

type Delete struct {
	Target_id int `json:"target_id"`
}

type Update struct {
	NewAge int `json:"new age"`
}
