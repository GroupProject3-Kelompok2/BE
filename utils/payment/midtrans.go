package payment

import (
	"strings"

	"github.com/GroupProject3-Kelompok2/BE/app/config"
	"github.com/midtrans/midtrans-go"
	"github.com/midtrans/midtrans-go/coreapi"
)

var c = &coreapi.Client{} // Gunakan tanda ampersand (&) untuk membuat pointer ke objek Client.

func CreateChargeRequest(invoice string, grossAmount uint) (*coreapi.ChargeReq, error) {
	c.New(config.MIDTRANS_SERVERKEY, midtrans.Sandbox)

	request := &coreapi.ChargeReq{
		PaymentType: coreapi.PaymentTypeBankTransfer,
		BankTransfer: &coreapi.BankTransferDetails{
			Bank: midtrans.BankBca,
		},
		TransactionDetails: midtrans.TransactionDetails{
			OrderID:  invoice,
			GrossAmt: int64(grossAmount),
		},
	}

	return request, nil
}

func ProcessChargeResponse(resp *coreapi.ChargeResponse) map[string]string {
	result := make(map[string]string)

	banks := make([]string, len(resp.VaNumbers))
	for i, bank := range resp.VaNumbers {
		banks[i] = bank.Bank
	}
	result["banks"] = strings.Join(banks, ",")

	vaNumbers := make([]string, len(resp.VaNumbers))
	for i, vaNumber := range resp.VaNumbers {
		vaNumbers[i] = vaNumber.VANumber
	}
	result["vaNumbers"] = strings.Join(vaNumbers, ",")

	return result
}
