#Node class: Linked List node element
class Node:
   def __init__(self, data=None):
      self.data = data
      self.next = None

   def __repr__(self) -> str:
      return f'{self.data}\n'


#Singly Linked List
class LinkedList:
   #Constructor
   def __init__(self):
      self.head = None
      self.nodes = []


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
   def add(self, data):
      node = Node(data)
      node.next = self.head
      self.head = node
      self.nodes.insert(0, node)

   #Search for key
   def search(self, key):
      current = self.head

      while current:
         if current.data == key:
            return current
         else:
            current = current.next
      return None
   
   #Insert to linked list
   def insert(self, data, index):
      if index == 0:
         return self.add(data)
      
      if index > 0:
         new_node = Node(data)

         pos = index
         current = self.head

         while pos > 1:
            current = current.next
            pos -= 1

         prev_node = current
         next_node = current.next

         prev_node.next = new_node
         new_node.next = next_node
         self.nodes.insert(index, new_node)

   #remove from linked list
   def remove(self, key):
      current = self.head
      prev = None
      found = False
      
      while current and not found:
         if current.data == key and current is self.head:
            found = True
            self.head = current.next
         elif current.data == key:
            found = True
            prev.next = current.next
         else:
            prev = current
            current = current.next
      
      return current

   #Display linked list
   def __repr__(self):
      # nodes = []
      # current = self.head

      # while current:
      #    if current is self.head:
      #       nodes.append(f'[Head: {current.data}]')
      #    elif current.next is None:
      #       nodes.append(f'[Tail: {current.data}]')
      #    else:
      #       nodes.append(f'[{current.data}]')
         
      #    current = current.next
      # return '-> '.join(nodes)

      nodes = []
      for i in self.nodes:
         if i is self.head:
            nodes.append(f'[Head: {i.data}]')
         elif i.next is None:
            nodes.append(f'[Tail: {i.data}]')
         else:
            nodes.append(f'[{i.data}]')

      return '-> '.join(nodes)
   

# linkedist = LinkedList()
# linkedist.add(1)
# linkedist.add("Hello")
# linkedist.add(3)
# linkedist.add(4)
# print(linkedist.size())

# # linkedist.insert("There", 3)

# # print(linkedist)