package Model

type Message map[string]interface {
	//code int
	//message string
	//data string
}

type LoginMessage struct {
	Code string `json:"code"`
	Data struct {
		UserId string `json:"userId"`
		Grade  string `json:"grade"`
	}
}
