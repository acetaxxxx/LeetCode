class Solution:
    def moveZeroes(self, nums):
        zeroCount = 0
        for i in range(len(nums)):
            if( nums[i] == 0):
                    zeroCount+=1
            else :
                tmp = nums[i-zeroCount]
                nums[i-zeroCount] = nums[i]
                nums[i] = tmp                    
        return nums
        """
        :type nums: List[int]
        :rtype: void Do not return anything, modify nums in-place instead.
        """
        