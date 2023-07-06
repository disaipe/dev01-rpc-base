package rpc

type Response interface{}
type ResultResponse interface{}

type ActionResponse struct {
	Response

	Status bool
	Data   string
}
