package prosper

import "github.com/bhatia4/gofn-prosper/prosper/thin"

type accountParser interface {
	Parse(thin.AccountResponse) (AccountInformation, error)
}

type defaultAccountParser struct{}

func (a defaultAccountParser) Parse(r thin.AccountResponse) (AccountInformation, error) {
	lastDepositDate, err := parseProsperDate(r.LastDepositDate)
	if err != nil {
		return AccountInformation{}, err
	}
	lastWithdrawDate, err := parseProsperDate(r.LastWithdrawDate)
	if err != nil {
		return AccountInformation{}, err
	}
	return AccountInformation{
		AvailableCashBalance:                r.AvailableCashBalance,
		TotalPrincipalReceivedOnActiveNotes: r.TotalPrincipalReceivedOnActiveNotes,
		OutstandingPrincipalOnActiveNotes:   r.OutstandingPrincipalOnActiveNotes,
		LastWithdrawAmount:                  r.LastWithdrawAmount,
		LastDepositAmount:                   r.LastDepositAmount,
		LastDepositDate:                     lastDepositDate,
		PendingInvestmentsPrimaryMarket:     r.PendingInvestmentsPrimaryMarket,
		PendingInvestmentsSecondaryMarket:   r.PendingInvestmentsSecondaryMarket,
		PendingQuickInvestOrders:            r.PendingQuickInvestOrders,
		TotalAmountInvestedOnActiveNotes:    r.TotalAmountInvestedOnActiveNotes,
		TotalAccountValue:                   r.TotalAccountValue,
		InflightGross:                       r.InflightGross,
		LastWithdrawDate:                    lastWithdrawDate,
	}, nil
}
