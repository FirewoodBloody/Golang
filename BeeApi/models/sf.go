package models

type Response struct {
	Service string `xml:"service,attr"`
	Head    string
}

type Responses struct {
	Success string `xml:"success"`
}

func Sf(id string) Response {

	m := Response{
		Service: "RoutePushService",
	}
	m.Head = "OK"
	if id == "" {
		return m
	}
	return m
}

func SF(id string) Responses {
	m := Responses{}
	m.Success = "true"

	if id == "" {
		return m
	}
	return m
}
