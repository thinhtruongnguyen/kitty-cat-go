package aiozaiplatformgosdk

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"strconv"
)

func ParseFailResponse(err error) (*ResponseFailResponse, bool) {
	if err == nil {
		return nil, false
	}

	if apiErr, ok := err.(*ErrorResponse); ok {
		var resp ResponseFailResponse
		if unmarshalErr := json.Unmarshal(apiErr.Body, &resp); unmarshalErr == nil {
			return &resp, true
		}
	}

	return nil, false
}

func PrettyPrint(v interface{}) {
	b, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		log.Printf("failed to marshal json: %v", err)
		return
	}
	fmt.Println(string(b))
}

func ToPtr[T any](v T) *T {
	return &v
}
func CheckBalanceDataField(data map[string]interface{}, key string) bool {
	val, exist := data[key]
	if !exist {
		return false
	}

	_, ok := val.(string)
	return ok
}

func CheckPriceDataField(data map[string]interface{}, key string) bool {
	val, exist := data[key]
	if !exist {
		return false
	}

	_, ok := val.(float64)
	return ok
}
func ConverToUSD(balance string, debt string, freeBalance string, tokenPrice float64) (float64, error) {
	balanceFloat, err1 := strconv.ParseFloat(balance, 64)
	debtFloat, err2 := strconv.ParseFloat(debt, 64)
	freeBalanceFloat, err3 := strconv.ParseFloat(freeBalance, 64)

	if err1 != nil || err2 != nil || err3 != nil {
		return 0, errors.New("invalid number format")
	}

	totalFloat := balanceFloat - debtFloat + freeBalanceFloat
	usdBalance := totalFloat * tokenPrice
	return usdBalance, nil
}

func AddonUsdBalance(balance map[string]interface{}, price map[string]interface{}) (map[string]interface{}, string) {
	isBalanceFieldOk := CheckBalanceDataField(balance, "balance")
	isDebtFieldOk := CheckBalanceDataField(balance, "debt")
	isFreeBalanceFieldOk := CheckBalanceDataField(balance, "free_balance")
	isPriceFieldOk := CheckPriceDataField(price, "current_price")

	if !isBalanceFieldOk || !isDebtFieldOk || !isFreeBalanceFieldOk || !isPriceFieldOk {
		return nil, "response format is incorrect"
	}

	usdBalance, err := ConverToUSD(
		balance["balance"].(string),
		balance["debt"].(string),
		balance["free_balance"].(string),
		price["current_price"].(float64),
	)

	if err != nil {
		return nil, "convert data fail"
	}

	balance["usd_balance"] = strconv.FormatFloat(usdBalance, 'f', 6, 64)
	return balance, ""
}
