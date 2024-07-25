class Node:
   def __init__(self, data):
      self.data = data
      self.left = None
      self.right = None



class BinaryTree:
   def __init__(self):
      self.root = None

   #Add to Tree
   def add(self, data):
      node = Node(data)
      self.root.left = self.root
      self.root = node
