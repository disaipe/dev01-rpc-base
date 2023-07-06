package rpc

import "io"

type Response interface{}
type ResultResponse interface{}

type ActionResponse struct {
	Response

	Status bool
	Data   string
}

type ActionFunction func(rpc *Rpc, body io.ReadCloser, appAuth string) (Response, error)
