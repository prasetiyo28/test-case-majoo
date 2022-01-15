package response

import (
	"log"
	"os"
	"sort"
	"test-case-majoo/entity/responses"
	"time"

	"github.com/gin-gonic/gin"
)

func Response(ctx *gin.Context, res *responses.Response) {
	var (
		APP_TIMEZONE, _ = time.LoadLocation(os.Getenv("APP_TIMEZONE"))
		TIME_NOW        = time.Now().In(APP_TIMEZONE)
		TIME_UNIX       = TIME_NOW.Unix()
	)
	if res.Code == 0 && res.Message == "" {
		res.Code = 204
		res.Message = "No Content"
	}

	if res.Meta.UnixTime == 0 {
		res.Meta.UnixTime = TIME_UNIX
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
