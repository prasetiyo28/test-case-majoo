package login

import (
	"fmt"
	"net/http"
	ENTITY_AUTH "test-case-majoo/entity/auths"
	"test-case-majoo/entity/responses"
	"test-case-majoo/pkg/response"
	"test-case-majoo/pkg/success"
	"test-case-majoo/pkg/utils"

	"github.com/gin-gonic/gin"
)

func (handler *LoginController) MonthlyReport(context *gin.Context) {
	var input ENTITY_AUTH.ReportRequest
	guid, _ := context.Get("userGuid")

	guidString := guid.(string)

	err := utils.ValidateRequest(utils.BIND_TYPE_JSON, "Report", context, &input)
	if err != nil {
		result := responses.Response{
			Code:    err.Code,
			Message: err.Message,
		}
		response.Response(context, &result)
		return
	}

	data, erre := handler.at.GetMonthlyReport(guidString, input.Month, input.Limit, input.Page)
	if erre != nil {
		fmt.Println("erre", erre)
	}

	result := responses.Response{
		Code:    http.StatusOK,
		Message: success.SuccessReport,
		Data:    data,
	}
	response.Response(context, &result)
}
