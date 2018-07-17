
class Solution:
    def __init__(self):
        self.factorialArray = [1,1] # 0!

    def factorial(self, n):
        if len(factorialArray) < n :
             self.factorialArray.append(n*factorial(n-1))
             return self.factorialArray[n]
        return self.factorialArray[n]
       
    
    
    def generate(self, numRows):
        """
        :type numRows: int
        :rtype: List[List[int]]
        """
        