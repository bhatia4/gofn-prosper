package thin

import (
	"fmt"
	"net/http"
	"reflect"
	"testing"

	"github.com/bhatia4/gofn-prosper/prosper/auth"
)

type mockTokenManager struct {
}

func (m mockTokenManager) Token() (auth.OAuthToken, error) {
	return auth.OAuthToken{}, nil
}

func TestSearchSuccessfulResponse(t *testing.T) {
	setUp()
	defer tearDown()

	mux.HandleFunc("/search/listings/",
		func(w http.ResponseWriter, r *http.Request) {
			testMethod(t, r, "GET")
			testFormValues(t, r, values{
				"limit":                     "50",
				"exclude_listings_invested": "true",
			})
			fmt.Fprint(w, `{
  "result": [
    {
      "member_key": "AF001946468823A105AE",
      "listing_number": 4247229,
      "credit_pull_date": "2015-12-04 01:03:03 +0000",
      "listing_start_date": "2015-12-04 17:02:28 +0000",
      "listing_creation_date": "2015-12-04 00:31:34 +0000",
      "listing_status": 2,
      "listing_status_reason": "Active",
      "verification_stage": 2,
      "listing_amount": 8000,
      "amount_funded": 833.91,
      "amount_remaining": 7166.09,
      "percent_funded": 0.1042,
      "partial_funding_indicator": true,
      "funding_threshold": 0.7,
      "prosper_rating": "C",
      "estimated_return": 0.0758,
      "estimated_loss_rate": 0.0824,
      "lender_yield": 0.1606,
      "effective_yield": 0.1582,
      "borrower_rate": 0.1706,
      "borrower_apr": 0.20777,
      "listing_term": 36,
      "listing_monthly_payment": 285.46,
      "scorex": "702-723",
      "fico_score": "660-679",
      "prosper_score": 4,
      "listing_category_id": 14,
      "listing_title": "Large Purchases",
      "income_range": 3,
      "income_range_description": "$25,000-49,999",
      "stated_monthly_income": 3000,
      "income_verifiable": true,
      "dti_wprosper_loan": 0.14,
      "employment_status_description": "Other",
      "months_employed": 72,
      "borrower_state": "MD",
      "borrower_city": "PASADENA",
      "borrower_metropolitan_area": "(Not Implemented)",
      "prior_prosper_loans_active": 0,
      "prior_prosper_loans": 0,
      "lender_indicator": 0,
      "group_indicator": false,
      "channel_code": "90000",
      "amount_participation": 0,
      "monthly_debt": 144,
      "current_delinquencies": 0,
      "delinquencies_last7_years": 14,
      "public_records_last10_years": 0,
      "public_records_last12_months": 0,
      "first_recorded_credit_line": "1991-03-22 08:00:00 +0000",
      "credit_lines_last7_years": 28,
      "inquiries_last6_months": 3,
      "amount_delinquent": 0,
      "current_credit_lines": 2,
      "open_credit_lines": 2,
      "bankcard_utilization": 0.32,
      "total_open_revolving_accounts": 3,
      "installment_balance": 0,
      "real_estate_balance": 0,
      "revolving_balance": 978,
      "real_estate_payment": 0,
      "revolving_available_percent": 72,
      "total_inquiries": 6,
      "total_trade_items": 28,
      "satisfactory_accounts": 23,
      "now_delinquent_derog": 0,
      "was_delinquent_derog": 5,
      "oldest_trade_open_date": "03221991",
      "delinquencies_over30_days": 5,
      "delinquencies_over60_days": 3,
      "delinquencies_over90_days": 14,
      "investment_typeid": 1,
      "investment_type_description": "Fractional",
      "is_homeowner": false
    },
    {
      "member_key": "8E36142078174706D82B",
      "listing_number": 4247097,
      "credit_pull_date": "2015-11-22 21:42:00 +0000",
      "listing_start_date": "2015-12-04 17:02:24 +0000",
      "listing_creation_date": "2015-12-04 00:15:33 +0000",
      "listing_status": 2,
      "listing_status_reason": "Active",
      "verification_stage": 1,
      "listing_amount": 20000,
      "amount_funded": 1530,
      "amount_remaining": 18470,
      "percent_funded": 0.0765,
      "partial_funding_indicator": true,
      "funding_threshold": 0.7,
      "prosper_rating": "C",
      "estimated_return": 0.0705,
      "estimated_loss_rate": 0.0724,
      "lender_yield": 0.1443,
      "effective_yield": 0.1429,
      "borrower_rate": 0.1543,
      "borrower_apr": 0.17792,
      "listing_term": 60,
      "listing_monthly_payment": 480.32,
      "scorex": "650-664",
      "fico_score": "680-699",
      "prosper_score": 8,
      "listing_category_id": 7,
      "listing_title": "Other",
      "income_range": 4,
      "income_range_description": "$50,000-74,999",
      "stated_monthly_income": 5694.33,
      "income_verifiable": true,
      "dt_iwprosper_loan": 0.22,
      "employment_status_description": "Other",
      "occupation": "Other",
      "months_employed": 615,
      "borrower_state": "HI",
      "borrower_city": "HILO",
      "borrower_metropolitan_area": "(Not Implemented)",
      "prior_prosper_loans_active": 0,
      "prior_prosper_loans": 0,
      "lender_indicator": 0,
      "group_indicator": false,
      "channel_code": "40000",
      "amount_participation": 0,
      "monthly_debt": 751,
      "current_delinquencies": 0,
      "delinquencies_last7_years": 0,
      "public_records_last10_years": 0,
      "public_records_last12_months": 0,
      "first_recorded_credit_line": "2001-10-24 07:00:00 +0000",
      "credit_lines_last7_years": 22,
      "inquiries_last6_months": 2,
      "amount_delinquent": 0,
      "current_credit_lines": 10,
      "open_credit_lines": 10,
      "bankcard_utilization": 0.76,
      "total_open_revolving_accounts": 9,
      "installment_balance": 0,
      "real_estate_balance": 52177,
      "revolving_balance": 34949,
      "real_estate_payment": 476,
      "revolving_available_percent": 29,
      "total_inquiries": 3,
      "total_trade_items": 22,
      "satisfactory_accounts": 22,
      "now_delinquent_derog": 0,
      "was_delinquent_derog": 0,
      "oldest_trade_open_date": "10242001",
      "delinquencies_over30_days": 0,
      "delinquencies_over60_days": 0,
      "delinquencies_over90_days": 0,
      "investment_typeid": 1,
      "investment_type_description": "Fractional",
      "is_homeowner": true
    },
    {
      "member_key": "34D036969548861949ADDA2",
      "listing_number": 4245951,
      "credit_pull_date": "2015-12-03 21:47:05 +0000",
      "listing_start_date": "2015-12-04 17:02:20 +0000",
      "listing_creation_date": "2015-12-03 21:47:04 +0000",
      "listing_status": 2,
      "listing_status_reason": "Active",
      "verification_stage": 3,
      "listing_amount": 10000,
      "amount_funded": 7601.92,
      "amount_remaining": 2398.08,
      "percent_funded": 0.7601,
      "partial_funding_indicator": true,
      "funding_threshold": 0.7,
      "prosper_rating": "E",
      "estimated_return": 0.0992,
      "estimated_loss_rate": 0.1425,
      "lender_yield": 0.2531,
      "effective_yield": 0.2417,
      "borrower_rate": 0.2631,
      "borrower_apr": 0.30244,
      "listing_term": 36,
      "listing_monthly_payment": 404.56,
      "scorex": "724-747",
      "fico_score": "680-699",
      "prosper_score": 3,
      "listing_category_id": 1,
      "listing_title": "Debt Consolidation",
      "income_range": 3,
      "income_range_description": "$25,000-49,999",
      "stated_monthly_income": 2944.75,
      "income_verifiable": true,
      "dt_iwprosper_loan": 0.6,
      "employment_status_description": "Employed",
      "occupation": "Professional",
      "months_employed": 206,
      "borrower_state": "UT",
      "borrower_city": "LOGAN",
      "borrower_metropolitan_area": "(Not Implemented)",
      "prior_prosper_loans_active": 0,
      "prior_prosper_loans": 0,
      "lender_indicator": 0,
      "group_indicator": false,
      "channel_code": "70000",
      "amount_participation": 0,
      "monthly_debt": 1354,
      "current_delinquencies": 0,
      "delinquencies_last7_years": 0,
      "public_records_last10_years": 0,
      "public_records_last12_months": 0,
      "first_recorded_credit_line": "1983-04-14 08:00:00 +0000",
      "credit_lines_last7_years": 23,
      "inquiries_last6_months": 0,
      "amount_delinquent": 0,
      "current_credit_lines": 10,
      "open_credit_lines": 10,
      "bankcard_utilization": 0.91,
      "total_open_revolving_accounts": 8,
      "installment_balance": 29479,
      "real_estate_balance": 62896,
      "revolving_balance": 13465,
      "real_estate_payment": 857,
      "revolving_available_percent": 30,
      "total_inquiries": 7,
      "total_trade_items": 23,
      "satisfactory_accounts": 22,
      "now_delinquent_derog": 0,
      "was_delinquent_derog": 1,
      "oldest_trade_open_date": "04141983",
      "delinquencies_over30_days": 1,
      "delinquencies_over60_days": 0,
      "delinquencies_over90_days": 0,
      "investment_typeid": 1,
      "investment_type_description": "Fractional",
      "is_homeowner": true
    }
  ],
  "result_count": 3,
  "total_count": 3
}`)
		},
	)

	client := defaultClient{
		baseURL:      server.URL,
		tokenManager: mockTokenManager{},
	}
	got, err := client.Search(SearchParams{
		Offset: 0,
		Limit:  50,
		ExcludeListingsInvested: true,
	})
	if err != nil {
		t.Fatalf("client.Search failed: %v", err)
	}

	want := SearchResponse{
		Results: []SearchResult{
			{
				PriorProsperLoans:                         0,
				AmountDelinquent:                          0,
				AmountParticipation:                       0,
				DelinquenciesOver60Days:                   3,
				GroupIndicator:                            false,
				IncomeRange:                               3,
				ListingMonthlyPayment:                     285.46,
				OldestTradeOpenDate:                       "03221991",
				PriorProsperLoansPrincipalOutstanding:     0,
				PublicRecordsLast12Months:                 0,
				TotalOpenRevolvingAccounts:                3,
				VerificationStage:                         2,
				ListingStatus:                             2,
				ListingTitle:                              "Large Purchases",
				ScorexChange:                              "",
				BorrowerListingDescription:                "",
				DelinquenciesLast7Years:                   14,
				EmploymentStatusDescription:               "Other",
				ListingStartDate:                          "2015-12-04 17:02:28 +0000",
				TotalTradeItems:                           28,
				BorrowerRate:                              0.1706,
				IsHomeowner:                               false,
				LastUpdatedDate:                           "",
				ListingAmount:                             8000,
				ListingNumber:                             4247229,
				PriorProsperLoans61dpd:                    0,
				PriorProsperLoansPrincipalBorrowed:        0,
				WasDelinquentDerog:                        5,
				BankcardUtilization:                       0.32,
				InstallmentBalance:                        0,
				InvestmentTypeid:                          1,
				ListingCategoryID:                         14,
				BorrowerCity:                              "PASADENA",
				BorrowerState:                             "MD",
				IncomeRangeDescription:                    "$25,000-49,999",
				ProsperScore:                              4,
				RevolvingAvailablePercent:                 72,
				WholeLoanStartDate:                        "",
				CurrentCreditLines:                        2,
				DtiWprosperLoan:                           0.14,
				FicoScore:                                 "660-679",
				FirstRecordedCreditLine:                   "1991-03-22 08:00:00 +0000",
				RealEstateBalance:                         0,
				SatisfactoryAccounts:                      23,
				ChannelCode:                               "90000",
				FundingThreshold:                          0.7,
				InquiriesLast6Months:                      3,
				LenderYield:                               0.1606,
				MemberKey:                                 "AF001946468823A105AE",
				PriorProsperLoanEarliestPayOff:            0,
				PriorProsperLoansCyclesBilled:             0,
				CurrentDelinquencies:                      0,
				DelinquenciesOver30Days:                   5,
				InvestmentTypeDescription:                 "Fractional",
				ListingStatusReason:                       "Active",
				MonthlyDebt:                               144,
				MonthsEmployed:                            72,
				PartialFundingIndicator:                   true,
				Rating:                                    "C",
				BorrowerApr:                               0.20777,
				GroupName:                                 "",
				ListingPurpose:                            "",
				PriorProsperLoans31dpd:                    0,
				PriorProsperLoansLatePaymentsOneMonthPlus: 0,
				RealEstatePayment:                         0,
				CreditLinesLast7Years:                     28,
				CreditPullDate:                            "2015-12-04 01:03:03 +0000",
				PublicRecordsLast10Years:                  0,
				RevolvingBalance:                          978,
				Scorex:                                    "702-723",
				BorrowerMetropolitanArea:            "(Not Implemented)",
				MinPriorProsperLoan:                 0,
				WholeLoanEndDate:                    "",
				AmountFunded:                        833.91,
				EffectiveYield:                      0.1582,
				EstimatedLossRate:                   0.0824,
				ListingCreationDate:                 "2015-12-04 00:31:34 +0000",
				Occupation:                          "",
				PercentFunded:                       0.1042,
				PriorProsperLoansBalanceOutstanding: 0,
				AmountRemaining:                     7166.09,
				DelinquenciesOver90Days:             14,
				ListingEndDate:                      "",
				OpenCreditLines:                     2,
				PriorProsperLoansActive:             0,
				PriorProsperLoansLateCycles:         0,
				ListingTerm:                         36,
				PriorProsperLoansOntimePayments:     0,
				EstimatedReturn:                     0.0758,
				IncomeVerifiable:                    true,
				LenderIndicator:                     0,
				MaxPriorProsperLoan:                 0,
				NowDelinquentDerog:                  0,
				StatedMonthlyIncome:                 3000,
				TotalInquiries:                      6,
			},
			{
				PriorProsperLoans:                         0,
				AmountDelinquent:                          0,
				AmountParticipation:                       0,
				DelinquenciesOver60Days:                   0,
				GroupIndicator:                            false,
				IncomeRange:                               4,
				ListingMonthlyPayment:                     480.32,
				OldestTradeOpenDate:                       "10242001",
				PriorProsperLoansPrincipalOutstanding:     0,
				PublicRecordsLast12Months:                 0,
				TotalOpenRevolvingAccounts:                9,
				VerificationStage:                         1,
				ListingStatus:                             2,
				ListingTitle:                              "Other",
				ScorexChange:                              "",
				BorrowerListingDescription:                "",
				DelinquenciesLast7Years:                   0,
				EmploymentStatusDescription:               "Other",
				ListingStartDate:                          "2015-12-04 17:02:24 +0000",
				TotalTradeItems:                           22,
				BorrowerRate:                              0.1543,
				IsHomeowner:                               true,
				LastUpdatedDate:                           "",
				ListingAmount:                             20000,
				ListingNumber:                             4247097,
				PriorProsperLoans61dpd:                    0,
				PriorProsperLoansPrincipalBorrowed:        0,
				WasDelinquentDerog:                        0,
				BankcardUtilization:                       0.76,
				InstallmentBalance:                        0,
				InvestmentTypeid:                          1,
				ListingCategoryID:                         7,
				BorrowerCity:                              "HILO",
				BorrowerState:                             "HI",
				IncomeRangeDescription:                    "$50,000-74,999",
				ProsperScore:                              8,
				RevolvingAvailablePercent:                 29,
				WholeLoanStartDate:                        "",
				CurrentCreditLines:                        10,
				DtiWprosperLoan:                           0,
				FicoScore:                                 "680-699",
				FirstRecordedCreditLine:                   "2001-10-24 07:00:00 +0000",
				RealEstateBalance:                         52177,
				SatisfactoryAccounts:                      22,
				ChannelCode:                               "40000",
				FundingThreshold:                          0.7,
				InquiriesLast6Months:                      2,
				LenderYield:                               0.1443,
				MemberKey:                                 "8E36142078174706D82B",
				PriorProsperLoanEarliestPayOff:            0,
				PriorProsperLoansCyclesBilled:             0,
				CurrentDelinquencies:                      0,
				DelinquenciesOver30Days:                   0,
				InvestmentTypeDescription:                 "Fractional",
				ListingStatusReason:                       "Active",
				MonthlyDebt:                               751,
				MonthsEmployed:                            615,
				PartialFundingIndicator:                   true,
				Rating:                                    "C",
				BorrowerApr:                               0.17792,
				GroupName:                                 "",
				ListingPurpose:                            "",
				PriorProsperLoans31dpd:                    0,
				PriorProsperLoansLatePaymentsOneMonthPlus: 0,
				RealEstatePayment:                         476,
				CreditLinesLast7Years:                     22,
				CreditPullDate:                            "2015-11-22 21:42:00 +0000",
				PublicRecordsLast10Years:                  0,
				RevolvingBalance:                          34949,
				Scorex:                                    "650-664",
				BorrowerMetropolitanArea:            "(Not Implemented)",
				MinPriorProsperLoan:                 0,
				WholeLoanEndDate:                    "",
				AmountFunded:                        1530,
				EffectiveYield:                      0.1429,
				EstimatedLossRate:                   0.0724,
				ListingCreationDate:                 "2015-12-04 00:15:33 +0000",
				Occupation:                          "Other",
				PercentFunded:                       0.0765,
				PriorProsperLoansBalanceOutstanding: 0,
				AmountRemaining:                     18470,
				DelinquenciesOver90Days:             0,
				ListingEndDate:                      "",
				OpenCreditLines:                     10,
				PriorProsperLoansActive:             0,
				PriorProsperLoansLateCycles:         0,
				ListingTerm:                         60,
				PriorProsperLoansOntimePayments:     0,
				EstimatedReturn:                     0.0705,
				IncomeVerifiable:                    true,
				LenderIndicator:                     0,
				MaxPriorProsperLoan:                 0,
				NowDelinquentDerog:                  0,
				StatedMonthlyIncome:                 5694.33,
				TotalInquiries:                      3,
			},
			{
				PriorProsperLoans:                         0,
				AmountDelinquent:                          0,
				AmountParticipation:                       0,
				DelinquenciesOver60Days:                   0,
				GroupIndicator:                            false,
				IncomeRange:                               3,
				ListingMonthlyPayment:                     404.56,
				OldestTradeOpenDate:                       "04141983",
				PriorProsperLoansPrincipalOutstanding:     0,
				PublicRecordsLast12Months:                 0,
				TotalOpenRevolvingAccounts:                8,
				VerificationStage:                         3,
				ListingStatus:                             2,
				ListingTitle:                              "Debt Consolidation",
				ScorexChange:                              "",
				BorrowerListingDescription:                "",
				DelinquenciesLast7Years:                   0,
				EmploymentStatusDescription:               "Employed",
				ListingStartDate:                          "2015-12-04 17:02:20 +0000",
				TotalTradeItems:                           23,
				BorrowerRate:                              0.2631,
				IsHomeowner:                               true,
				LastUpdatedDate:                           "",
				ListingAmount:                             10000,
				ListingNumber:                             4245951,
				PriorProsperLoans61dpd:                    0,
				PriorProsperLoansPrincipalBorrowed:        0,
				WasDelinquentDerog:                        1,
				BankcardUtilization:                       0.91,
				InstallmentBalance:                        29479,
				InvestmentTypeid:                          1,
				ListingCategoryID:                         1,
				BorrowerCity:                              "LOGAN",
				BorrowerState:                             "UT",
				IncomeRangeDescription:                    "$25,000-49,999",
				ProsperScore:                              3,
				RevolvingAvailablePercent:                 30,
				WholeLoanStartDate:                        "",
				CurrentCreditLines:                        10,
				DtiWprosperLoan:                           0,
				FicoScore:                                 "680-699",
				FirstRecordedCreditLine:                   "1983-04-14 08:00:00 +0000",
				RealEstateBalance:                         62896,
				SatisfactoryAccounts:                      22,
				ChannelCode:                               "70000",
				FundingThreshold:                          0.7,
				InquiriesLast6Months:                      0,
				LenderYield:                               0.2531,
				MemberKey:                                 "34D036969548861949ADDA2",
				PriorProsperLoanEarliestPayOff:            0,
				PriorProsperLoansCyclesBilled:             0,
				CurrentDelinquencies:                      0,
				DelinquenciesOver30Days:                   1,
				InvestmentTypeDescription:                 "Fractional",
				ListingStatusReason:                       "Active",
				MonthlyDebt:                               1354,
				MonthsEmployed:                            206,
				PartialFundingIndicator:                   true,
				Rating:                                    "E",
				BorrowerApr:                               0.30244,
				GroupName:                                 "",
				ListingPurpose:                            "",
				PriorProsperLoans31dpd:                    0,
				PriorProsperLoansLatePaymentsOneMonthPlus: 0,
				RealEstatePayment:                         857,
				CreditLinesLast7Years:                     23,
				CreditPullDate:                            "2015-12-03 21:47:05 +0000",
				PublicRecordsLast10Years:                  0,
				RevolvingBalance:                          13465,
				Scorex:                                    "724-747",
				BorrowerMetropolitanArea:            "(Not Implemented)",
				MinPriorProsperLoan:                 0,
				WholeLoanEndDate:                    "",
				AmountFunded:                        7601.92,
				EffectiveYield:                      0.2417,
				EstimatedLossRate:                   0.1425,
				ListingCreationDate:                 "2015-12-03 21:47:04 +0000",
				Occupation:                          "Professional",
				PercentFunded:                       0.7601,
				PriorProsperLoansBalanceOutstanding: 0,
				AmountRemaining:                     2398.08,
				DelinquenciesOver90Days:             0,
				ListingEndDate:                      "",
				OpenCreditLines:                     10,
				PriorProsperLoansActive:             0,
				PriorProsperLoansLateCycles:         0,
				ListingTerm:                         36,
				PriorProsperLoansOntimePayments:     0,
				EstimatedReturn:                     0.0992,
				IncomeVerifiable:                    true,
				LenderIndicator:                     0,
				MaxPriorProsperLoan:                 0,
				NowDelinquentDerog:                  0,
				StatedMonthlyIncome:                 2944.75,
				TotalInquiries:                      7,
			},
		},
		ResultCount: 3,
		TotalCount:  3,
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("client.Search returned %#v, want %#v", got, want)
	}
}

func TestSearchFailedResponse(t *testing.T) {
	setUp()
	defer tearDown()

	mux.HandleFunc("/search/listings/",
		func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprint(w, "mock server error: request failed")
		},
	)
	client := defaultClient{
		baseURL:      server.URL,
		tokenManager: mockTokenManager{},
	}
	_, err := client.Search(SearchParams{
		Offset: 0,
		Limit:  50,
		ExcludeListingsInvested: true,
	})
	if err == nil {
		t.Fatal("client.Search should fail when server returns error")
	}
}
