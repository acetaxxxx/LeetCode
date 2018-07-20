from LC283_v1 import Solution

def main():
    temp = [0,1,0,3,12]
    expect = [1,2,3,0,0,0]
    s = Solution()
    result = s.moveZeroes(temp)
    print(result)
    print(expect)

if __name__ == '__main__':
    main()
    

