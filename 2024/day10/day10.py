from functools import cache

'''
returns the number of values we will end up with given the first

i was implementing the cache myself with a dict:
    - but saw that python3 has this built in with the @cache decorator

the key was to ignore the actual values, and just focus on the number of values

also: recursing the path of one stone at a time is better than tackling the whole list at once
'''
@cache
def calculate(n, depth):
    if depth == 0:
        return 1
    elif n == 0:
        return calculate(1, depth-1)
    elif len(str(n)) % 2 == 0:
        left = int(str(n)[:len(str(n)) // 2])
        right = int(str(n)[len(str(n)) // 2:])
        return calculate(int(left), depth-1) + calculate(int(right), depth-1)
    else:
        return calculate(n * 2024, depth-1)

with open('input.txt') as f:
    data = [int(x) for x in f.readline().split()]
    
    print('Part 1:', sum(calculate(n, 25) for n in data))
    print('Part 2:', sum(calculate(n, 75) for n in data))
