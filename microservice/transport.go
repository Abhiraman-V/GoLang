/*Transport
The transport file is meant to represent the matching layer of the OSI model. It’s the file that contains every piece of code responsible for the end to end communication strategy. Go kit supports many transports out of the box. For simplicity, we will work with JSON over HTTP. 

The transport consists of a series of Server objects, one for each endpoint of our service, and it receives four parameters. These are:
An endpoint: The request handler. Given the separation of structures and responsibilities, we can use the same structure across multiple transport implementations.
A decode function: This function receives the external requests and translates it to Golang code.
An encode function: The exact opposite of the decode function. This function translates the Golang response to the corresponding output for the selected transport.
A set of server options: A set of options that could be credentials, codec, keep parameters alive, etc. These provide extra capabilities to our transport layer. */
package mystr

import (
"context"
"encoding/json"
"errors"
"net/http"

"github.com/go-kit/kit/endpoint"
httptransport "github.com/go-kit/kit/transport/http"
)

func GetIsPalHandler(ep endpoint.Endpoint, options []httptransport.ServerOption) *httptransport.Server {
return httptransport.NewServer(
  ep,
  decodeGetIsPalRequest,
  encodeGetIsPalResponse,
  options...,
)
}

func decodeGetIsPalRequest(_ context.Context, r *http.Request) (interface{}, error) {
var req IsPalRequest
if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
  return nil, err
}
return req, nil
}

func encodeGetIsPalResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
resp, ok := response.(*IsPalResponse)
if !ok {
  return errors.New("error decoding")
}
return json.NewEncoder(w).Encode(resp)
}

func GetReverseHandler(ep endpoint.Endpoint, options []httptransport.ServerOption) *httptransport.Server {
return httptransport.NewServer(
  ep,
  decodeGetReverseRequest,
  encodeGetReverseResponse,
  options...,
)
}

func decodeGetReverseRequest(_ context.Context, r *http.Request) (interface{}, error) {
var req ReverseRequest
if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
  return nil, err
}
return req, nil
}

func encodeGetReverseResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
resp, ok := response.(*ReverseResponse)
if !ok {
  return errors.New("error decoding")
}

return json.NewEncoder(w).Encode(resp)
}