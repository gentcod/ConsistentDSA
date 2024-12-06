# Best time to Buy and Sell Stock

def max_profit(prices: list[int]) -> int:
   #Two pointers: left=buy, right=sell
   l, r = 0, 1
   maxP = 0

   while r < len(prices):
      if prices[l] < prices[r]:
         profit = prices[r] - prices[l]
         maxP = max(maxP, profit)

      else:
         l = r
      r += 1
   
   return maxP

# stock_prices = [7, 1, 2, 4, 5, 6, 3]

# print(max_profit(stock_prices))