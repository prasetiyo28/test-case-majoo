package auth

import (
	"fmt"
	"net/http"
	"strconv"
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

func (auth *AuthRepository) GetMonthlyReport(userID, Month, limit, page string) (auths.MonthlyReports, *responses.Response) {
	var report auths.MonthlyReport
	var reports auths.MonthlyReports
	pageInt, _ := strconv.Atoi(page)
	limitInt, _ := strconv.Atoi(limit)
	offset := (pageInt - 1) * limitInt
	rows := auth.db.Table("users").
		Select("merchants.merchant_name,transactions.id,sum(transactions.bill_total) as bill_total").
		Group("transactions.created_at").
		Joins("left join merchants on users.id = merchants.user_id").
		Joins("left join transactions on merchants.id = transactions.merchant_id").
		Where("users.id = ?", userID).Where("month(transactions.updated_at) = ?", Month).
		Limit(limit).
		Offset(offset).
		Scan(&reports)
	// if err != nil {
	// 	return report, &responses.Response{
	// 		Code:    http.StatusInternalServerError,
	// 		Message: err.Error(),
	// 	}
	// }

	// for rows.Next() {
	println("rows", rows)
	fmt.Println("reports", reports)
	fmt.Println("report", report)
	for i, c := range reports {
		fmt.Println(i)
		fmt.Println(c)
	}
	// }
	return reports, nil
}
