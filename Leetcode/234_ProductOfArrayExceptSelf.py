class Solution:
    def productExceptSelf(self, nums: List[int]) -> List[int]:
        x=1
        zeros = 0
        for num in nums:
            if num:
                x*=num
            else: zeros+=1

        if zeros > 1: return [0]*len(nums)

        res = [0] * len(nums)
        for i,c in enumerate(nums):
            if zeros == 0:
                res[i] = x // c

            elif c == 0:
                res[i] = x
            
        return res
