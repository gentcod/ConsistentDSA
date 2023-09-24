#Node class: Linked List node element
class ListNode:
   data = None
   next = None

   def __init__(self, data):
      self.data = data

   def __repr__(self) -> str:
      return f'Node data: {self.data}\n'


#Singly Linked List
class LinkedList:

   #Constructor
   def __init__(self):
      self.head = None

   #Check if Linked List is empty
   def is_empty(self):
      return self.head == None
   
   #Check size of Linked List
   def size(self):
      current = self.head
      count = 0

      while current:
         count += 1
         current = current.next

      return count
   
   #Add to linked list:
   """
   1. Create instance of ListNode class
   2. Set next reference to the previous Linked List head
   3. Set head to newly added ListNode
   """
   def add(self, data):
      node = ListNode(data)
      node.next = self.head
      self.head = node

   def search(self, key):
      current = self.head

      while current:
         if current.data == key:
            return current
         else:
            current = current.next
      return None

   def __repr__(self):
      nodes = []
      current = self.head

      while current:
         if current is self.head:
            nodes.append(f'[Head: {current.data}]')
         elif current.next is None:
            nodes.append(f'[Tail: {current.data}]')
         else:
            nodes.append(f'[{current.data}]')
         
         current = current.next
      return '-> '.join(nodes)
   

linked_list = LinkedList()
linked_list.add(1)
linked_list.add(2)
linked_list.add(3)

# print(linked_list.size())
print(linked_list)
print(linked_list.search(3))