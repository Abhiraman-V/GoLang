/* Microservice architecture is a design pattern that structures applications as a collection of services. 
Some of the advantages of microservices are:

Highly maintainable and testable
Loosely coupled
Independently deployable
Organized around business capabilities
Owned by a small group of people

Microservices are easier to maintain/test, scale more efficiently, 
use fewer resources, and are focused on specific processes that produce value to your company. 

Enter Go Kit

Go kit is a programming toolkit for building microservices in Go. 
It was created to solve common problems in distributed systems and applications 
and allow developers to focus on the business part of programming. 
It’s basically a group of co-related packages that, together, form a framework for constructing large SOAs.
Its main transport layer is RPC (Remote Procedure Call), but it can also use JSON over HTTP 
and it’s easy to integrate with the most common infrastructural components, to reduce deployment friction 
and promote interoperability with existing systems. 

*/


package mystr

import (
"errors"
"strings"
"github.com/go-kit/kit/log"
)

type Service interface {
IsPal(string) error
Reverse(string) string
}

type myStringService struct {
log log.Logger
}

func (svc *myStringService) IsPal(s string) error {
reverse := svc.Reverse(s)
if strings.ToLower(s) != reverse {
  return errors.New("Not palindrome")
}
return nil
}

func (svc *myStringService) Reverse(s string) string {
rns := []rune(s) // convert to rune
for i, j := 0, len(rns)-1; i ‹ j; i, j = i+1, j-1 {

  // swap the letters of the string,
  // like first with last and so on.
  rns[i], rns[j] = rns[j], rns[i]
}

// return the reversed string.
return strings.ToLower(string(rns))
}