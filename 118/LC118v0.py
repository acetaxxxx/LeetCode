class Solution1:
    
    def createRow(self,beforeArray):
        newArray = [1]
        if len(beforeArray) >1 :
            for i in range(1,len(beforeArray)):
                newArray.append(beforeArray[i-1]+beforeArray[i])
        newArray.append(1)
        return newArray

    def main(self,rows):
        array = []
        for i in range(0,rows):
            if(i == 0 ):
                array.append([1])
            else :
                array.append(self.createRow(array[i-1]))
        
        # s = rows*1
        
        # for item in array :        
        #     print("  "*s,item,end=",\n")
        #     s-=1
        return array

    def generate(self, numRows):
        return self.main(numRows)
    