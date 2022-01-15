# Golang Naming Convention
## Files
1. Go follows a convention where source files are all lower case with underscore separating multiple words.
2. Compound file names are separated with _
3. File names that begin with “.” or “_” are ignored by the go tool
4. Files with the suffix `_test.go` are only compiled and run by the `go test` tool.
```go
// examples
1. main.go
2. healt_check.go
3. .gitignore
4. main_test.go
```

## Function and Methods
- A name must begin with a letter, and can have any number of additional letters and numbers.
- A function name cannot start with a number.
- Use camel case, exported functions should start with uppercase:
```go
func writeToDB // unexported, only visible within the package
func WriteToDB // exported, visible within the other package
```
If the functions with names that start with an uppercase letter will be exported to other packages. If the function name starts with a lowercase letter, it won't be exported to other packages, but you can call this function within the same package, this is similar to the public/private features in most OO languages. Here's an example or two:  
```go
package a

func ThisFunctionIsExported() {}

func thisOneIsNot() {}

// = = = = = = = = = = = = = = = = =

package b

import a

a.ThisFunctionIsExported() // works

a.thisOneIsNot() // compiler error
```

## Constants
Constant should use all capital letters and use underscore `_` to separate words.
```go
const PRODUCT string = "Laptop"
const PRICE = 500
const PAYMENT_STATUS = true
```

## Variables
- Use `camelCase`.
- Acronyms should be all capitals, as in `ServeHTTP`.
- Generally, use relatively simple (short) name but descriptive names: `cust` not `customer`.
- Consistent naming style should be used the entire source code.
```go
user to u
userID to uid
```
If variable type is `bool`, its name should start with `Has`, `Is`, `Can` or `Allow`, etc.
```go
var isExist bool
allowNull := true
```
Single letter represents index: `i, j, k`
```go
sample := "ab£c"
for i := 0; i < 4; i++ {
    fmt.Printf("%c\n", sample[i])
}
fmt.Printf("Length is %d\n", len(sample))

//or 

names := []string{"Mary", "John", "Bob", "Anna"}
for i, n := range names {
    fmt.Printf("index: %d = %q\n", i, n)
}
```

## Packages
Good package names are short and clear. They are lower case, with no under_scores or mixedCaps. They are often simple nouns, such as:
```go
time (provides functionality for measuring and displaying time)
list (implements a doubly linked list)
http (provides HTTP client and server implementations)
```
Abbreviate judiciously. Package names may be abbreviated when the abbreviation is familiar to the programmer. Widely-used packages often have compressed names:
```go
strconv (string conversion)
syscall (system call)
fmt (formatted I/O)
On the other hand, if abbreviating a package name makes it ambiguous or unclear, don`t do it.
``` 
avoid repeating package name:
```go
log.Info()    // good
log.LogInfo() // bad
```
Don’t name like getters or setters :
```go
custSvc.cust()    // good
custSvc.getCust() // bad
```

***Don't steal good names from the user***. Avoid giving a package a name that is commonly used in client code. For example, the buffered I/O package is called bufio, not buf, since buf is a good variable name for a buffer.

References : 
- [Kassim Damilola](https://medium.com/@kdnotes/golang-naming-rules-and-conventions-8efeecd23b68)
- [Dave Cheney](https://dave.cheney.net/practical-go/presentations/qcon-china.html)
- [Lindsay](https://medium.com/@lynzt/variable-naming-conventions-in-go-89fe1ef17b0a)
- [Blog Golang](https://blog.golang.org/package-names)
- [golangprograms.com](https://www.golangprograms.com/naming-conventions-for-golang-functions.html)
- [digitalocean.com](https://www.digitalocean.com/community/tutorials/how-to-use-variables-and-constants-in-go)
- [Learn Go Programming](https://www.youtube.com/watch?v=yQUAHpEvb9A)
- [GopherCon Russia](https://www.youtube.com/watch?v=MzTcsI6tn-0)

## Endpoint Naming
### URIs as resources as nouns
RESTful URIs should refer to a resource that is a thing (noun) instead of referring to an action (verb) because nouns have properties which verbs do not have – similar to resources have attributes.  
Example: `/users/{id}` instead of `/getUser`
### Pluralized resources
Next up is the question of whether resource names should be pluralized. Admittedly, this is a matter of preference; however, most API design experts would suggest you pluralize all resources unless they are singleton resources.  
Example: `/users` (typical resource) or `/users/{id}/address` (singleton resource)
### Forward slashes for hierarchy
As shown in the examples above, forward slashes are conventionally used to show the hierarchy between individual resources and collections.  
Example: `/users/{id}/address` clearly falls under the `/users/{id}` resource which falls under the `/users` collection.  
### Lowercase letters and dashes
By convention, resource names should use exclusively lowercase letters. Similarly, dashes `(-)` are conventionally used in place of underscores `(_)`.
Example: `/users/{id}/pending-orders` instead of `/users/{id}/Pending_Orders`  
### No trailing forward slash
Similarly, in the interests of keeping URIs clean, do not add a trailing forward slash to the end of URIs.  
Example: `/users/{id}/pending-orders` instead of `/users/{id}/pending-orders/`
### Do not use CRUD function names in URIs
URIs should not be used to indicate that a CRUD function is performed. URIs should be used to uniquely identify resources and not any action upon them. HTTP request methods should be used to indicate which CRUD function is performed.  
Example: `router.POST("/users", handler.FuncName)` instead of `router.POST("/users/create", handler.FuncName)`
```go
router.GET("/users", handler.FuncName)
router.GET("/users/:id", handler.FuncName)
router.POST("/users", handler.FuncName)
router.PUT("/users", handler.FuncName)
router.Delete("/users", handler.FuncName)
```
References : 
- [Thomas Bush](https://nordicapis.com/10-best-practices-for-naming-api-endpoints/)
- [restfulapi.net](https://restfulapi.net/resource-naming/)
- [John Au-Yeung and Ryan Donovan](https://stackoverflow.blog/2020/03/02/best-practices-for-rest-api-design/)

## sample structure folder 
```text
api
├── handler
│   ├── auths
│   │   ├── login
│   │   └── register
│   ├── check
│   ├── masters
│   └── users
├── middleware
├── presenter
│   ├── auths
│   └── masters
└── routers
    └── v1
```