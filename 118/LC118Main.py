from LC118v0 import Solution1
from LC118v2 import Solution2
from datetime import datetime

def main():
    rowNo = 500

    y = Solution1()
    ts = datetime.now().timestamp()
    s = y.generate(rowNo)
    tf = datetime.now().timestamp()

    # print(s)    
    print("Solution 1 :",(tf-ts), " second")
    
    y = Solution2()
    ts = datetime.now().timestamp()
    s = y.generate(rowNo)
    tf = datetime.now().timestamp()

    print("\n")
    # print(s)    
    print("Solution 2 :",(tf-ts), " second")

if __name__ == '__main__':
    main()   
