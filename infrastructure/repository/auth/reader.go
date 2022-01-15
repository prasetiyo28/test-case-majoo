package auth

import (
	"fmt"
	"net/http"
	"test-case-majoo/entity/auths"
	"test-case-majoo/entity/responses"
)

func (auth *AuthRepository) GetUserByID(userID string) (auths.User, *responses.Response) {
	var user auths.User
	fmt.Println("userID", userID)
	err := auth.db.Table("users").Where("user_name = ?", userID).First(&user).Error
	if err != nil {
		return user, &responses.Response{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		}
	}
	return user, nil
}

func (auth *AuthRepository) GetMonthlyReport(userID, Month string) (auths.MonthlyReport, *responses.Response) {
	var report auths.MonthlyReport
	rows, err := auth.db.Table("users").Select("merchants.name,transactions.id,transaction.bill").Joins("left join merchants on users.id = merchants.user_id").Joins("left join transactions on merchants.id = transactions.merchants_id").Where("users.id = ?", userID).Rows()
	if err != nil {
		return report, &responses.Response{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		}
	}

	for rows.Next() {
		println("rows", rows)
	}
	return report, nil
}
