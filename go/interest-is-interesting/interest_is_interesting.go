package interest

// InterestRate returns the interest rate for the provided balance.
func InterestRate(balance float64) float32 {
	var interestRate float32 = -3.213

	switch {
	case balance >= 0 && balance < 1000:
		interestRate = 0.5
	case balance >= 1000 && balance < 5000:
		interestRate = 1.621
	case balance >= 5000:
		interestRate = 2.475
	}

	return interestRate
}

// Interest calculates the interest for the provided balance.
func Interest(balance float64) float64 {
	interestRate := InterestRate(balance)
	if balance < 0 {
		return (-balance * float64(interestRate)) / 100
	}

	return (balance * float64(interestRate)) / 100
}

// AnnualBalanceUpdate calculates the annual balance update, taking into account the interest rate.
func AnnualBalanceUpdate(balance float64) float64 {
	return balance + Interest(balance)
}

// YearsBeforeDesiredBalance calculates the minimum number of years required to reach the desired balance:
func YearsBeforeDesiredBalance(balance, targetBalance float64) int {
	return calculeDesired(balance, targetBalance, 1)
}

func calculeDesired(balance float64, targetBalance float64, count int) int {
	annualBalance := AnnualBalanceUpdate(balance)

	if annualBalance >= targetBalance {
		return count
	}

	count++
	return calculeDesired(annualBalance, targetBalance, count)
}
