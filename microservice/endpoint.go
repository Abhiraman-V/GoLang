/* Endpoints

This is where we introduce one of the main functionalities of Go kit: endpoints. 
These are abstractions provided by Go kit that work pretty much like an action or handler on a controller. 
It’s also the place to add safety and security logic. Each endpoint represents a single method in our service interface. 

One of the structures provided by Go kit to ensure security is the middlewares. 
Middlewares work the same way they do in any other language. 
They are methods executed over the request before it reaches its handler. 
Here you can add functionality such as logging, load balancing, tracing, etc.*/


package mystr

import (
"context"
"errors"

"github.com/go-kit/kit/endpoint"
"github.com/go-kit/kit/log"
"github.com/go-kit/kit/log/level"
)

type Endpoints struct {
GetIsPalindrome endpoint.Endpoint
GetReverse      endpoint.Endpoint
}

func MakeEndpoints(svc Service, logger log.Logger,
					middlewares []endpoint.Middleware) Endpoints {
return Endpoints{
  GetIsPalindrome: wrapEndpoint(makeGetIsPalindromeEndpoint(svc, logger), middlewares),
  GetReverse:      wrapEndpoint(makeGetReverseEndpoint(svc, logger), middlewares),
}
}

func makeGetIsPalindromeEndpoint(svc Service, logger log.Logger) endpoint.Endpoint {
return func(ctx context.Context, request interface{}) (interface{}, error) {
  req, ok := request.(*IsPalRequest)
  if !ok {
    level.Error(logger).Log("message", "invalid request")
    return nil, errors.New("invalid request")
  }
  msg := svc.IsPal(ctx, req.Word)
  return &IsPalResponse{
    Message: msg,
  }, nil
}

}

func makeGetReverseEndpoint(svc Service, logger log.Logger) endpoint.Endpoint {
return func(ctx context.Context, request interface{}) (interface{}, error) {
  req, ok := request.(*ReverseRequest)
  if !ok {
    level.Error(logger).Log("message", "invalid request")
    return nil, errors.New("invalid request")
  }
  reverseString := svc.IsPal(ctx, req.Word)
  return &ReverseResponse{
    Word: reverseString,
  }, nil
}
}

func wrapEndpoint(e endpoint.Endpoint, middlewares []endpoint.Middleware) endpoint.Endpoint {
for _, m := range middlewares {
  e = m(e)
}
return e
}