class Solution1:
    
    def __init__(self):
       self.__init__()
        
    def createRow(beforeArray):
        newArray = [1]
        if len(beforeArray) >1 :
            for i in range(1,len(beforeArray)):
                newArray.append(beforeArray[i-1]+beforeArray[i])
        newArray.append(1)
        return newArray

    def main(rows):
        array = [[1]]
        for item in range(rows-1):
            array.append(createRow(array[item]))
        
        s = rows*1
        
        for item in array :        
            print("  "*s,item,end=",\n")
            s-=1

     def generate(self, numRows):
         main(numRows)
    