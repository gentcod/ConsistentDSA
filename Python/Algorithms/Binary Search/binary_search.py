def binary_search(list: list, target):
    first = 0
    last = len(list) - 1

    while first <= last:
        midpoint = (first + last) // 2

        if list[midpoint] == target:
            return True
        elif list[midpoint] < target:
            first = midpoint + 1
        elif list[midpoint] > target:
            last = midpoint - 1
        else: 
            return False


def verify(valid):
    if valid:
        print(f"Target found: {valid}")
    else:
        print("Target not found in list")


# numbers = [1, 2, 3, 4, 5, 6, 7, 8, 9, 10]

# verify(binary_search(numbers, 11))
