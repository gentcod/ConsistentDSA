def recursive_binary_search(list: list, target):
    if len(list) == 0:
        return False
    else:
        midpoint = (len(list))//2 #Divide and round down to nearest whole number

        if list[midpoint] == target:
            return True
        else:
            if list[midpoint] < target:
                return recursive_binary_search(list[midpoint+1:], target)
            else:
                return recursive_binary_search(list[:midpoint], target)
            

def verify(result):
    print(f"Target found: {result}")

# numbers = [1, 2, 3, 4, 5, 6, 7, 8]

# verify(recursive_binary_search(numbers, 6))