
class Solution2:
    def __init__(self):
        self.factorialArray = [1,1] # 0! 1!

    def factorial(self, n):
        if len(self.factorialArray) < n+1 :
            self.factorialArray.append( n * self.factorial(n-1))
            return self.factorialArray[n]
        return self.factorialArray[n]
       
    def getNthRow(self,n):
        array = []
        for i in range(0,n+1):
            array.append(int( self.factorial(n) / (self.factorial(i)*self.factorial(n-i)) ))
        return array
    
    def generate(self, numRows):
        result = []
        for i in range(numRows-1):           
            result.append(self.getNthRow(i))
        return result
        """
        :type numRows: int
        :rtype: List[List[int]]
        """
        