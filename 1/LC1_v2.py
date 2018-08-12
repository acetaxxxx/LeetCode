class Solution:
    def twoSum(self, nums, target):
        for i in range(len(nums)):
            try:
                a = nums.index(target-nums[i])
                if(a != i):
                    return [a, i]
            except:
                pass
        return []

        """
        :type nums: List[int]
        :type target: int
        :rtype: List[int]
        """
