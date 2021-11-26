/* Requests and Responses

In Go kit, the primary messaging pattern is RPC. 
This means that every method in our interface should be modeled as a client-server interaction. 
The requesting program is a client, and the service is the server. 
This allows us to specify the parameters and return types of each method.  */

package mystr

type IsPalRequest struct {
Word string `json:"word"`
}

type ReverseRequest struct {
Word string `json:"word"`
}