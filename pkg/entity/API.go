package entity

type httpMethod string

const (
	// HTTPGet HTTPGet
	HTTPGet = httpMethod("GET")
	// HTTPPost HTTPPost
	HTTPPost = httpMethod("POST")
	// HTTPDelete HTTPDelete
	HTTPDelete = httpMethod("DELETE")
	// HTTPPut HTTPPut
	HTTPPut = httpMethod("PUT")
	// HTTPHead HTTPHead
	HTTPHead = httpMethod("HEAD")
)

// API API
type API struct {
	ID               string
	Method           httpMethod
	URI              string
	RequestTemplate  string
	ResponseTemplate string
}
