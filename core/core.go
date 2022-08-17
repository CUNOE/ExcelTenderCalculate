package core

type Result struct {
	Combination []int
	AvgPrice    float64
	Companies   []InRandomPriceWinTheBidding
	Count       map[string]int
}

type InRandomPriceWinTheBidding struct {
	RandomPrice             float64
	TheCompanyWinTheBidding Company
}

func CalculateResult(combination [][]int, selectCompanies []Company, fixedCompanies []Company, randomValue []float64) []Result {

	var r []Result

	for _, z := range combination {
		var prices []float64
		var cs []Company
		var companiesWinTheBidding []InRandomPriceWinTheBidding
		for _, c := range z {
			prices = append(prices, selectCompanies[c].Price)
			cs = append(cs, selectCompanies[c])
		}
		for _, c := range fixedCompanies {
			prices = append(prices, c.Price)
			cs = append(cs, c)
		}
		randomPrices := CalculateRandomPrices(AvgPrice(prices), randomValue)

		for _, rp := range randomPrices {
			companiesWinTheBidding = append(companiesWinTheBidding, InRandomPriceWinTheBidding{
				RandomPrice:             rp,
				TheCompanyWinTheBidding: ReturnCompanyWhoWinTheBidding(rp, cs),
			})
		}
		r = append(r, Result{
			Combination: z,
			AvgPrice:    AvgPrice(prices),
			Companies:   companiesWinTheBidding,
			Count:       CountTheResult(companiesWinTheBidding),
		})
	}

	return r
}

func CalculateOneResult(Companies []Company, randomValue []float64) []Result {

	var r []Result

	var prices []float64
	var companiesWinTheBidding []InRandomPriceWinTheBidding
	//for _, c := range z {
	//	prices = append(prices, selectCompanies[c].Price)
	//	cs = append(cs, selectCompanies[c])
	//}

	for _, c := range Companies {
		prices = append(prices, c.Price)
	}

	randomPrices := CalculateRandomPrices(AvgPrice(prices), randomValue)

	for _, rp := range randomPrices {
		companiesWinTheBidding = append(companiesWinTheBidding, InRandomPriceWinTheBidding{
			RandomPrice:             rp,
			TheCompanyWinTheBidding: ReturnCompanyWhoWinTheBidding(rp, Companies),
		})
	}
	r = append(r, Result{
		AvgPrice:  AvgPrice(prices),
		Companies: companiesWinTheBidding,
		Count:     CountTheResult(companiesWinTheBidding),
	})

	return r
}

func CountTheResult(zbc []InRandomPriceWinTheBidding) map[string]int {
	var count map[string]int
	count = make(map[string]int)

	for _, d := range zbc {
		name := d.TheCompanyWinTheBidding.Name
		if _, ok := count[name]; ok {
			count[name]++
		} else {
			count[name] = 1
		}
	}
	return count
}
