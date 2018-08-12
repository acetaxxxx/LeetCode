import LC121_v1 
def main():
    
    inputArray = [7,1,5,3,6,4]
    
    act = LC121_v1.Solution()
    expect = 6
    actual = act.maxProfit(inputArray)
    print("expect :",expect," actual : ",actual)

    inputArray = [7,6,4,3,1]
    expect = 0
    actual = act.maxProfit(inputArray)
    print("expect :",expect," actual : ",actual)

if __name__ == '__main__':
    main()