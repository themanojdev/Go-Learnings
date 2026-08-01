package Arrays

func BestStockToBuy(nums []int) int {

	bestBuy := nums[0]
	maxProfit := 0

	for i := 0 ; i < len(nums) ; i++ {

		if nums[i] < bestBuy {
			bestBuy = nums[i]
		}

		profit := nums[i] - bestBuy

		if profit > maxProfit {
			maxProfit = profit
		}
	}
	return maxProfit
}