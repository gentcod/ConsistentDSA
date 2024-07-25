from singly_linked_list import LinkedList

list1 = LinkedList()
list1.add(1)
list1.add(2)
list1.add(3)

list2 = LinkedList()
list2.add(1)
list2.add(2)
list2.add(4)

def mergedSortedList(list1: LinkedList, list2: LinkedList):
   if not list1: return list2
   elif not list2: return list1

   bound = max(len(list1.nodes), len(list2.nodes))

   for i in range(bound):
     if list1.nodes[i].data <= list2.nodes[i].data:
          list1.nodes[i].next = mergedSortedList(list1.nodes[i].next, list2)
          return list1
     else:
          list2.next = mergedSortedList(list1, list2.next)
          return list2

print(mergedSortedList(list1, list2))

def merge(list1: LinkedList, list2):
    for i in list1.nodes:
        print(i.next)

# print(merge(list1, list2))