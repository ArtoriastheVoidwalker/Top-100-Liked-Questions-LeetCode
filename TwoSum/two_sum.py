from typing import List


def twoSum(nums: List[int], target: int) -> List[int]:
    for i in range(len(nums)):
        num = target - nums[i]
        if num in nums[i + 1:len(nums)]:

            return [i, nums[i + 1:].index(num) + 1 + i]
