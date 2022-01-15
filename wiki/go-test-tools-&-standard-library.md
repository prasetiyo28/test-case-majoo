# Go test #2

Integration tests verify that different modules or services used by your application work well together. For example, it can be testing the interaction with the database or making sure that microservices work together as expected. if we created api we have to test endpoint using integration test.

So what happens if you need to test API endpoints? These are a little different from unit test because they need to be tested as HTTP requests rather than functions.

Take the following very unimpressive and simple API:
```
package main

func setupRouter() *gin.Engine {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})
	return r
}

func main() {
	r := setupRouter()
	r.Run(":8080")
}
```

Test for code example above:

```
package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPingRoute(t *testing.T) {
	router := setupRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/ping", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, "pong", w.Body.String())
}
```

## what happen if these endpoint have sub function?

# Mock
Mocking is a process used in unit testing when the unit being tested has external dependencies. The purpose of mocking is to isolate and focus on the code being tested and not on the behaviour or state of external dependencies. In mocking, the dependencies are replaced by closely controlled replacements objects that simulate the behaviour of the real ones.The replacement objects can be of three types : Fakes,Stubs and Mocks.
A Fake is an object that will replace the actual code by implementing the same interface but without interacting with other objects. Usually the Fake is hard-coded to return fixed results.

## go get
```
go get github.com/vektra/mockery/v2/.../
```

use mockery --all for generate interface which is inside interface have function to be generate by mockery.

![mockery](ss-gotest-1.png)

It's common for a big package to have a lot of interfaces, so mockery provides --all. This option will tell mockery to scan all files under the directory named by --dir ("." by default) and generates mocks for any interfaces it finds. This option implies --recursive=true.

Simplest case

```
package usecase

import (
	"go-oauth2-kompasid/entity"
)

type Usecase interface {
	GetUserDetail(email string) (entity.DetailUser, error)
}
```

Run: mockery --all and the following will be output to mocks/Usecase.go:

```
// GetUserDetail provides a mock function with given fields: email
func (_m *Usecase) GetUserDetail(email string) (entity.DetailUser, error) {
	ret := _m.Called(email)

	var r0 entity.DetailUser
	if rf, ok := ret.Get(0).(func(string) entity.DetailUser); ok {
		r0 = rf(email)
	} else {
		r0 = ret.Get(0).(entity.DetailUser)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(email)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
```

## [Example mock](../usecase/user/mocks/Repository.go)


# Data faker

Faker will generate you a fake data based on your Struct. in mock we have to generate fake data for the result we expected based on type in struct.

## go get

```
go get -u github.com/bxcodec/faker/v3
```
The Struct Field must be PUBLIC.<br>
Support Only For :

* `int`, `int8`, `int16`, `int32` & `int64`
* `[]int`, `[]int8`, `[]int16`, `[]int32` & `[]int64`
* `bool` & `[]bool`
* `string` & `[]string`
* `float32`, `float64`, `[]float32` &`[]float64`
* `time.Time` & `[]time.Time`
* Nested Struct Field

# Example

```
import "github.com/bxcodec/faker"

func TestGetUserDetail(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		r := gin.Default()
		mockUCase := new(mocks.Usecase)
		email := "komasid@kompas.com"

		var mockUser entity.DetailUser
		err := faker.FakeData(&mockUser)
		assert.NoError(t, err)

		mockUCase.On("GetUserDetail", string(email)).Return(mockUser, nil)
		user, errc := mockUCase.GetUserDetail(email)
		require.NoError(t, errc)
```
this struct we have to generate fake data

```
type DetailUser struct {
	ID            string        `json:"userId"`
	Name          string        `json:"-" firestore:"name"`
	FirstName     string        `json:"firstName" firestore:"firstName"`
	Email         string        `json:"email" firestore:"email"`
	UserGuid      string        `json:"userGuid" firestore:"userGuid"`
	LastName      string        `json:"lastName" firestore:"lastName"`
	NoTlp         string        `json:"phoneNumber" firestore:"phoneNumber"`
	TglLahir      time.Time     `json:"dateBirth" firestore:"dateBirth"`
	Negara        string     	`json:"country" firestore:"country"`
	JenisKelamin  string     	`json:"gender" firestore:"gender"`
	DomisiliProv  string     	`json:"province" firestore:"province"`
	Kota          string     	`json:"city" firestore:"city"`
	CreateDate    string     	`json:"-" firestore:"createdDate"`
	CreateDateUs  CreateDate 	`json:"createDate"`
	UserCompleted bool       	`json:"userCompleted"`
	Alamat        Alamat     	`json:"address" firestore:"address"`
}
```
sample faker data from struct DetailUser


![data faker](ss-gotest-2.png)

## [Example](../api/handler/user_test.go)

# Subtests and refactoring

When we’re testing a function, there may be instances where we want to group the test items into meaningful subcategories. For example, for a more complicated function, we may want to have a subtest for the happy path, another subtest for an edge-case and maybe another subtest for error-handling. 

to create subgroups for non-negative cases and negative cases. Golang still allows us to do these subtests and it’s also a good way to reuse input values.

```
func TestFunc(t *testing.T) {
  // place common values to all tests
  t.Run("subtest 1", func(t *testing.T) {
    // place common values to all subtest 1
    t.Run("subsubtest a", func(t *testing.T) {
      // place common values to all subsubtest a
    })
    t.Run("subsubtest b", func(t *testing.T) {
      //
    })
  })
  t.Run("subtest 2", func(t *testing.T) {
    //
  })
}
```


This also makes it easier to read the test cases. In our example, we don’t have common values that we can refactor for the simple usecase, but we can group the test items to easily see that the created function can work for both non-negative and negative case.

## [Example test](../api/handler/user_test.go)
## [Example real function](../api/handler/user.go)

# Asssertify

The most basic tasks performed by unit tests are assertions. Assertions are usually used to verify if the actions performed by the test using determined input produce the expected output. They can also be used to check if the components follow the desired design rules. 

# http mock

For Golang, we using HTTPMock. It allows me to configure a particular response to send when a particular request is received. In fact, you can queue up a few responses but I'm writing a really thin API client so nothing needs more than one call so far.

Inside each test, I configure the request to respond to and the response to send, something like this:

```
  httpmock.Activate()
  defer httpmock.DeactivateAndReset()

  // our database of articles
  articles := make([]map[string]interface{}, 0)

  // mock to list out the articles
  httpmock.RegisterResponder("GET", "https://kompas.id/articles",
    func(req *http.Request) (*http.Response, error) {
      resp, err := httpmock.NewJsonResponse(200, articles)
      if err != nil {
        return httpmock.NewStringResponse(500, ""), nil
      }
      return resp, nil
    },
  )
```

In this example we mock standard net/http package response to an example website, which should return us a resp map response.

## [Example test](../entity/apimy_test.go)

## [Example real function](../entity/apimy_test.go)

# Write table driven tests

Anonymous structs and composite literals allow us to write very clear and simple table tests without relying on any external package.
The following code allows us to setup a range of tests for `CheckFirstName()` function:

```
fnameTest := []struct {
		firstname string
		expected  bool
	}{
		{
			firstname: "kompas",
			expected:  true,
		},
		{
			firstname: "abc123",
			expected:  false,
		},
		{
	}
```

Then our test function just ranges over the slice, calling the `CheckFirstName()` method for each `n`, before asserting that the results are correct:

```
	for _, test := range fnameTest {
		t.Run(test.firstname, func(t *testing.T) {
			result := entity.CheckFirstName(test.firstname)
			require.Equal(t, test.expected, result)
		})
	}
```

## [Example test](../entity/auth_test.go)
## [Example real function](../entity/auth.go)

---
## *references*

*https://www.thepolyglotdeveloper.com/2017/02/unit-testing-golang-application-includes-http/*

*https://levelup.gitconnected.com/unit-testing-using-mocking-in-go-f281122f499f*

*https://github.com/vektra/mockery*

*https://github.com/bxcodec/faker*

*https://gowalker.org/github.com/stretchr/testify/assert*

*https://github.com/jarcoal/httpmock*

*https://medium.com/@matryer/5-simple-tips-and-tricks-for-writing-unit-tests-in-golang-619653f90742*
