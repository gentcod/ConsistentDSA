#Node Class
class Node:
   data = None
   prev = None
   next = None

   def __init__(self, data):
      self.data = data

class LinkedList:
   def __init__(self):
      self.head = None

   def is_empty(self):
      return self.head == None
   
   def search(self, key):
      current = self.head

      while current:
         if current.data == key:
            return current
         else:
            current = current.next
      return None
   
   def add(self, data):
      node = Node(data)
      node.next = self.head
      node.prev = self.search(node.next)
      self.head = node

   
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

link = LinkedList()

link.add(1)
link.add(2)
link.add(3)

print(link)