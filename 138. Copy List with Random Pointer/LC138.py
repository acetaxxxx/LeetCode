"""
# Definition for a Node.
class Node:
    def __init__(self, x: int, next: 'Node' = None, random: 'Node' = None):
        self.val = int(x)
        self.next = next
        self.random = random
"""


class Node:
    def __init__(self, x: int, next: 'Node' = None, random: 'Node' = None):
        self.val = int(x)
        self.next = next
        self.random = random


class Solution:
    def copyRandomList(self, head: 'Node') -> 'Node':
        if head == None:
            return None

        copyIterator = head

        # DeepCopy Main Line
        while(copyIterator != None):
            newNode = Node(copyIterator.val, copyIterator.next)
            copyIterator.next = newNode
            copyIterator = newNode.next

        # CopyRandomNode
        copyIterator = head
        while(copyIterator != None):
            duplicateNode = copyIterator.next
            originRandomNode = copyIterator.random
            duplicateNode.random = None if originRandomNode is None else originRandomNode.next
            copyIterator = copyIterator.next.next

        # Destruction Linked List
        newHead = head.next

        newListIterator = newHead
        while(newListIterator is not None):
            newListIterator.next = newListIterator.next.next
            newListIterator = newListIterator.next
        return newHead
        
        
            