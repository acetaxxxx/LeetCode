class Solution:
    def twoSum(self, nums, target):
        for i in reversed(range(len(nums))):
            if(target < nums[i]):
                pass
            else :
               try:
                   a = nums.index(target-nums[i])
                   return [a,i]
               except :
                   pass
        return []

            
        """
        :type nums: List[int]
        :type target: int
        :rtype: List[int]
        """
        