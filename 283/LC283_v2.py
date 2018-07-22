class Solution:
    def moveZeroes(self, nums):
        zeroCount = 0
        length = len(nums)
        for i in range(length):
            if( nums[i] == 0):
                zeroCount += 1
            else:
                nums[i-zeroCount] = nums[i]
        
        for i in range(zeroCount):
            nums[length-zeroCount+i] = 0 
        return nums
        
        