# Go Test

Go has a built-in [testing] command `go test` and a package testing

## Why Unit test need

Unit tests are crucial to long-term project. We are expected to learn by perceiving, but often we end up dooming ourselves from the start, due to misconceptions or gaps in knowledge. I hope to fill in some of those gaps and provide a broader way of ideas to tackle go unit tests.

Key benefits of unit tests:

- Provide a safety when refactoring
- Can help identify dead code
- Provide a measure of confidence for management
- Can sometimes find missed use cases
- Helps produce higher quality code

There are costs associated with writing unit tests as well:

- Time and effort to write and maintain
- False sense of security (poor coverage, duplicate tests, testing the wrong thing, poorly written tests, the code is not fully covered)

## Writing go unit test

Unit testing in Go is just as opinionated as any other aspect of the language like formatting or naming. The syntax deliberately avoids the use of assertions and leaves the responsibility for checking values and behaviour to the developer.

The requirements for a valid go test file

- File name ends with _test.go the test file should be `user_test.go`.
- Include a package declaration in the file should be real function `package handler` file test `package handler_test`
- function should begins with the word Test followed by a word or phrase starting with a capital letter and should have only one parameter `t *testing.T`.
`go func TestGetUserDetail(t *testing.T) {}`

- `t.Error` or `t.Fail` to indicate a failure.
- `t.Log` can be used to provide non-failing debug information.
- Mocking function by [mockery].
- Using the [assert package] provides some helpful methods that allow you to write better test code in Go.
- Using [Data Faker] for mock dummy data.
- http mock [httpmock] if need interaction with external api


### Executing command

| Command | Description |
| :-----: | :---------: |
| `go test` | picks up any files matching packagename_test.go                                                   |
|                      `go tool cover -func=coverage.out`                      | show coverage lits of function                                                                       |
|                                  `go test ./...`                                  | picks up any files matching \*\_test.go all the packages from directory                           |
|                                   `go test -v`                                    | verbose output with PASS/FAIL result of each test including any extra logging produced by _t.Log_ |
|                                 `go test -cover`                                  | verbose output with code-coverage                                                                 |
| `go test -cover -coverprofile=c.out` `go tool cover -html=c.out -o coverage.html` | generating an HTML coverage report                                                                |



in this project use this command to get coverage and get result.

```
godotenv -f .env go test -v -coverprofile=coverage.out ./...
```


or use this command to see coverage in any function.

```
go tool cover -func=coverage.out
```
## Example basic of unit test

Here is an example of a method we want to test in the main package. We have defined an exported function called Sum which takes in two integers and adds them together.

```
package entity

func Sum(x int, y int) int {
    return x + y
}

func main() {
    Sum(5, 5)
}
```
We then write our test in a separate file. The test file can be in a different package (and folder) or the same one (main). Here's a unit test to check addition:

```
package entity_test

import "testing"

func TestSum(t *testing.T) {
    total := Sum(5, 5)
    if total != 10 {
       t.Errorf("Sum was incorrect, got: %d, want: %d.", total, 10)
    }
}
```

execute test using command go test, you can read from previos table.

```
tidigital@TIs-MacBook-Pro entity % go test -run TestSum
PASS
ok      go-oauth2-kompasid/entity       1.631s
tidigital@TIs-MacBook-Pro entity %
```

print coverage

```
tidigital@TIs-MacBook-Pro entity % go test -run TestSum -cover
PASS
coverage: 1.3% of statements
ok      go-oauth2-kompasid/entity       1.140s
tidigital@TIs-MacBook-Pro entity % 
```
## Table Test
Table Test used to convenience unit test with too many scenarios. Table driven testing is not a tool, package or anything else, it's just a way and perspective to write cleaner tests.

## Template Table Test
```

func TestTable(t *testing.T) {
    type Unit struct {
        Name   string
        Result auths.OauthClients
        Error  *responses.Response
    }

	tests := []struct {
		name    string
        wantCode    int
		unit        Unit
        ... //more argument
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.handler.EncryptPassword(tt.args.c)
		})
	}
}

```

## Example table test

Here is an example of a implementation table test
```
func TestGetAccessList(t *testing.T) {
    type Unit struct {
        Name   string
        Result auths.OauthClients
        Error  *responses.Response
    }

	tests := []struct {
		name        string
		wantCode    int
		accessToken string
		unit        Unit
	}{
		{
			name:        "success#200",
			wantCode:    http.StatusOK,
			accessToken: accessToken,
			unit: Unit{
				Name: "OauthClientList",
				Result: auths.OauthClients{
					{
						GuID:          "abcdes-123-12as",
						OauthClientId: "2",
						Status:        1,
						Token:         "randomtoken",
						Name:          "Google SWG",
						Scope:         "nama lengkap, alamat, email",
						Description:   "untuk SWG",
					},
				},
				Error: nil,
			},
		},
		{
			name:        "unauthorized#401",
			wantCode:    http.StatusUnauthorized,
			accessToken: accessTokenExpired,
			unit: Unit{
				Name:   "OauthClientList",
				Result: nil,
				Error: &responses.Response{
					Code: http.StatusUnauthorized,
				},
			},
		},
		{
			name:        "badRequest#400",
			wantCode:    http.StatusBadRequest,
			accessToken: atWithoutGuid,
			unit: Unit{
				Name:   "OauthClientList",
				Result: nil,
				Error: &responses.Response{
					Code: http.StatusBadRequest,
				},
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			authMock := new(MOCK_AUTH.Repository)
			userMock := new(MOCK_USER.Repository)
			redisMock := new(MOCK_REDIS.Repository)

			// call mock function TokenByEmail
			authMock.On("OauthClientList", mock.AnythingOfType("string")).Return(test.unit.Result, test.unit.Error)
			ssoClientList, errSsoClientList := authMock.OauthClientList(userGuid)

			authService := auth.NewService(authMock)
			redisService := redis.NewService(redisMock)
			userService := user.NewService(userMock)
			controller := &SsoController{
				at: authService,
				us: userService,
				rd: redisService,
			}

			c, _ := gin.CreateTestContext(httptest.NewRecorder())
			if test.wantCode != http.StatusUnauthorized {
				c.Set("userGuid", userGuid)
			}

			// Main function Access List
			controller.GetAccessList(c)

			if test.unit.Error != nil {
				assert.NotNil(t, errSsoClientList)
				assert.Nil(t, ssoClientList)
			} else {
				assert.Nil(t, errSsoClientList)
				assert.NotNil(t, ssoClientList)
			}

			authMock.AssertExpectations(t)
			userMock.AssertExpectations(t)
			redisMock.AssertExpectations(t)
		})
	}
}

```
## Example Structure test
```
.example
├── api
│   ├── api.go
│   └── handler // integration test, test endpoint make sure work all together include sub function
│         ├── user.go
│         └── user_test.go // integration test with endpoint
│ 
├── entity // unit test logic & external api
│   ├── auth.go
│   ├── auth_test.go // unit test logic
│   ├── apimy.go
│   └── apimy_test.go // unit test with external api
│ 
├── usecase // mocks and interface/services
│   ├── mocks // auto generate by mockery based on interface
│   │     └── Repository.go
│   └── interface.go
│
├── infrastructure
│   └── repository
│         ├── user.go
│         └── user_test.go
│ 
└── README.md
```

# [Link to next section](go-test-tools-&-standard-library.md)

## *References*

*https://medium.com/rungo/unit-testing-made-easy-in-go-25077669318*

*https://blog.alexellis.io/golang-writing-unit-tests/*

*https://github.com/gin-gonic/gin#testing*

*https://github.com/bxcodec/go-clean-arch*

*https://medium.com/@matryer/5-simple-tips-and-tricks-for-writing-unit-tests-in-golang-619653f90742*

*https://github.com/andanhm/gounittest*


[testing]: https://golang.org/pkg/testing/
[assert package]: https://github.com/gsamokovarov/assert
[mockery]: https://github.com/vektra/mockery
[Data Faker]: https://github.com/bxcodec/faker
[httpmock]: https://github.com/jarcoal/httpmock
