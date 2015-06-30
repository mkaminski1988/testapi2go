# testapi2go
test repo

Test project to test multiple refernce documents.

Problem... 
1. go build the project  
2. Run project. 

Program Panics ..stack trace as below 

panic: a handle is already registered for path ''/v1/authors/:id/links/:name'

goroutine 1 [running]:
github.com/julienschmidt/httprouter.(*node).addRoute(0xc208046960, 0xc2080a3756, 0x5, 0xc2080a3720)
	/home/bellrock1/gocode/src/github.com/julienschmidt/httprouter/tree.go:190 +0xc7d
github.com/julienschmidt/httprouter.(*Router).Handle(0xc2080cc150, 0x73b0f0, 0x4, 0xc2080a3740, 0x1b, 0xc2080a3720)
	/home/bellrock1/gocode/src/github.com/julienschmidt/httprouter/router.go:229 +0x1fd
github.com/julienschmidt/httprouter.(*Router).POST(0xc2080cc150, 0xc2080a3740, 0x1b, 0xc2080a3720)
	/home/bellrock1/gocode/src/github.com/julienschmidt/httprouter/router.go:188 +0x59
github.com/manyminds/api2go.(*API).addResource(0xc208046dc0, 0x7fb3672faea8, 0xc208058240, 0x7fb3672fae68, 0xc208038258, 0xc20807de30)
	/home/bellrock1/gocode/src/github.com/manyminds/api2go/api.go:329 +0xe42
github.com/manyminds/api2go.(*API).AddResource(0xc208046dc0, 0x7fb3672faea8, 0xc208058240, 0x7fb3672fae68, 0xc208038258)
	/home/bellrock1/gocode/src/github.com/manyminds/api2go/api.go:373 +0x50
main.main()
	/home/bellrock1/gocode/src/github.com/testapi2go/main.go:15 +0x113


====================================
If you remove `func (author *Author) DeleteToManyIDs(name string, IDs []string) error`
from file controller/author.go line 120
 Program Runs without any errors.





