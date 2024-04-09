package student

type Student struct {
	ID         string `json:"id"`
	Name       string `json:"name"`
	Department string `json:"department"`
	Age        int    `json:"age"`
}

type Sport struct {
	Sportname string `json:"sportname"`
	SportType string `json:"sporttype"`
}
