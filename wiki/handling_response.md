# Handling Response
This file explains how to use the standard error or success return format of a function, which will later be given to the user client in the form of a JSON response.

## Package SetResponse

Use the `Response` struct to wrap your result function either error or success.

```go
//entity/responses/response.go
type Response struct {
	Error   string      `json:"-"`
	Success bool        `json:"success"`
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Meta    Meta        `json:"meta"`
	Data    interface{} `json:"data"`
}

type Meta struct {
	Cache    bool  `json:"cache"`
	UnixTime int64 `json:"time"`
}

```
### Field Explanation
1. `Error`, this column is used to accommodate internal error messages that occur when using a library or fail when carrying out a technical process and can be used for `logError`.
2. `Success`, this column is used to determine whether a process succeeded or failed with data type `boolean`(`true/false`).
3. `Code`, this column is used to specify `httpStatus` when using the client.
4. `Message`, this column is used to accommodate messages that will be used by clients that are non-technical and easy to understand for laypeople.
5. `Meta`, this column is used to show cache value is true or false and Unix timestamp. field cache used to explain the response is get from Redis or DB.
6. `Data`, this column is used to accommodate data in the form of `interface{}` if a process fails or is not appropriate and requires providing data from the results of the process, this column can be used to accommodate the data, which will be accessed by the client.
## Example
```go
/* using for error */ 
return &Response{
	Error	: fmt.Sprintf("Error, FunctionName: v%", variableError.Error()),
	Success	: false,
	Code	: http.StatusInternalServerError,
	Meta	: Meta{
		Cache   : false,
		Unixtime: time.now().unix(),
	},
	Message	: "Creating data registration is fail.",
	Data	: entity.DataUser,
}

/* using for success */ 
return &SetResponse{
	Success	: true,
	Code	: http.StatusOK,
	Message	: "Success",
	Meta	: Meta{
		Cache   : true, // this value can be false or true depend on you set the value
		Unixtime: time.now().unix(),
	},
	Data	: entity.DataUser,
}
```
## Package template response `success` and `error`

Use these template responses for your `JSON` client format.
### Response
```go
//pkg/response/response.go
func Response(ctx *gin.Context, res *responses.Response) {

	if res.Code == 0 && res.Message == "" {
		res.Code = 204
		res.Message = "No Content"
	}

	if res.Meta.UnixTime == 0 {
		res.Meta.UnixTime = utils.TIME_UNIX
	}

	if res.Error != "" {
		log.Println(res.Error)
		// sentry-go logging put here
	}

	res.Success = validateDefaultSuccess(res.Code)
	ctx.JSON(res.Code, res)
}

func validateDefaultSuccess(value int) bool {
	defaultSuccess := []int{200, 201, 202, 203, 204, 206}
	i := sort.Search(len(defaultSuccess), func(i int) bool { return value <= defaultSuccess[i] })
	if i < len(defaultSuccess) && defaultSuccess[i] == value {
		return true
	} else {
		return false
	}
}
```
## Usage
### Case 1
function on repository set the result when error with `SetResponse` and this function will called on handler function
```go
import (
	"context"
	"net/http"
	"project/entity/responses"
	"project/pkg/errors"
	"project/pkg/success"
	"project/pkg/response"
)

// repository function
func (ur *UserRepository) ListCountry() (entity.Countries, bool, *responses.Response) {
	var countries entity.Countries
	ctx := context.Background()
	iter := ur.db.Collection("country").Documents(ctx)
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}

		if err != nil {
			return nil, &responses.Response{
				Error:   fmt.Sprintf("Error, ListCountry: %v", err.Error()),
				Code:    http.StatusInternalServerError,
				Message: "Daftar negara gagal ditampilkan",
			}
		}

		getCountry := doc.Data()["name"].(string)
		appendToCountry := entity.Country{
			ID:   doc.Ref.ID,
			Name: getCountry,
		}
		countries = append(countries, appendToCountry)
	}
	return countries, false, nil
}

// handler function
func (handler *UserController) GetCountry(c *gin.Context) {
	countries, metaCache, err := handler.us.ListCountries()
	if err != nil {
		result := responses.Response{
			Error:   err.Error.Error(),
			Code:    err.Code,
			Message: err.Message,
		}
		response.Response(c, &result)
		return
	}

	result := responses.Response{
		Code:    http.StatusOK,
		Message: success.SuccessGetProvince,
		Meta: responses.Meta{
			Cache: metaCache,
		},
		Data: countries,
	}
	response.Response(c, &result)
}
```
### Case 2
function on handler some section need to use response `error` but without `setResponse`
```go
func (handler *UserController) GetUserDataByEmail(c *gin.Context) {
	// start,  use response error without setResonse
	email, exists := c.Get("email")
	if !exists || email == nil {
		result := responses.Response{
			Error:   "Error, GetUserDataByEmail: Email not found on accces token",
			Code:    http.StatusBadRequest,
			Message: "email dibutuhkan untuk melanjutkan proses.",
		}
		response.Response(c, &result)
		return
	}
	// end,  use response error without setResonse

	getEmail := email.(string) // get email from access token
	getUserData, errUserData := handler.us.UserByEmail(getEmail)
	if errUserData != nil {
		result := responses.Response{
			Error:   errUserData.Error.Error(),
			Code:    errUserData.Code,
			Message: errUserData.Message,
		}
		response.Response(c, &result)
		return
	}
	result := responses.Response{
		Code:    http.StatusOK,
		Message: "Sukses",
		Data: getUserData,
	}
	response.Response(c, &result)
}
```
## References REST API Error Handling in Go
- [Zeynel Özdemir](https://medium.com/@ozdemir.zynl/rest-api-error-handling-in-go-behavioral-type-assertion-509d93636afd)
- [Elliot Lings](https://github.com/gin-gonic/gin/issues/274)
- [Dimitri Balios](https://github.com/astaxie/build-web-application-with-golang/blob/master/en/11.1.md)
- [developpaper](https://developpaper.com/series-gin-framework-custom-error-handling/)
- [Santosh Shrestha](https://medium.com/wesionary-team/capture-errors-with-sentry-in-go-gin-framework-2e0c034b986a)
- [juanmanuel.tirado](https://jmtirado.net/338-2/)
