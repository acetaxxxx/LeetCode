from LC118v0 import Solution1
from LC118v2 import Solution2
from datetime import datetime

def main():
    
    y = Solution1()
    ts = datetime.now().timestamp()
    s = y.generate(800)
    tf = datetime.now().timestamp()

    print(s)
    print("\n")
    print("Solution 1 :",(tf-ts)*100, " second")
    


if __name__ == '__main__':
    main()
    
