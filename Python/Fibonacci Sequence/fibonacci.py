def fibn(num):
    fibnum = [0, 1]
    fibnum_even = []
    fib_a = 0
    fib_b = 0
    fib_c = 0
    while fibnum[-1] < num:
        fib_a = fibnum[-2]
        fib_b = fibnum [-1]
        fib_c = (fib_b + fib_a)
        if fib_c > num: break
        fibnum.append(fib_c)
    for evn in fibnum:
        if evn % 2 == 0:
            fibnum_even.append(evn)

    print("Evem numbers in fibonacci sequence: ",fibnum_even)
    print("Sum of even numbers in fibonacci sequence: ", sum(fibnum_even))
    return fibnum
print("Fibonacci Sequence",fibn(100))