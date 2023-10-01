#Node class: Linked List node element
class Node:
   def __init__(self, data=None):
      self.data = data
      self.next = None

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
   def add(self, data):
      node = Node(data)
      node.next = self.head
      self.head = node

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
linked_list.add(4)

print(linked_list)

linked_list.insert("Hello", 2)

print(linked_list)
print(linked_list.size())
