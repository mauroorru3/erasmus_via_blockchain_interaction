package utilfunc

import (
	"encoding/json"
	"fmt"
	"os"
	"university_chain_it/x/universitychainit/types"
)

type TaxesStruct struct {
	Year            uint16 `json:"year"`
	Amount          uint32 `json:"amount"`
	Income_bracket  uint32 `json:"income_bracket"`
	Payment_made    bool   `json:"payment_made"`
	Date_of_payment string `json:"date_of_payment"`
}

func CheckTaxPayment(student types.StoredStudent) (ok bool, err error) {

	taxesDataBytes := []byte(student.TaxesData.TaxesHistory)
	var taxesData []TaxesStruct
	err = json.Unmarshal(taxesDataBytes, &taxesData)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error "+err.Error())
		return ok, err
	}

	return taxesData[0].Payment_made, err
}
