import re

# part 1: 'optimised' bruteforce
def valid_palindrome(s: str) -> bool:
    # not valid if length 1
    if len(s) < 2:
        return False

    # not valid if leading zero
    if s[0] == '0':
        return False

    # optimisation: cant be valid if first and last chars dont match
    if s[0] != s[len(s)-1]:
        return False

    # valid if palindrome
    return s[:len(s)//2] == s[len(s)//2:]

# part 2: bruteforce with regex
def valid_pattern(s: str) -> bool:
    # not valid if length 1
    if len(s) < 2:
        return False

    # not valid if leading zero
    if s[0] == '0':
        return False

    # for each substring from the start, e.g. '1', '12', '123'
    for i in range(1, len(s)):
        substr = s[:i]

        # run regex for this substring
        matches = re.findall(substr, s)

        # if the number of matches we got is a factor of the length of the substring
        return len(s) == i * len(matches)

    return False

part1, part2 = 0, 0

f = open('input.txt', 'r')
line = f.readlines()[0]
f.close()

# further optimisation potential:
# use a thread for each range, or even each id to check

for id_range in line.split(','):
    # optimisation: 0-10 are all invalid IDs so we never check them
    lo = max(int(id_range.split('-')[0]), 11)
    hi = int(id_range.split('-')[1])

    # loop from lo to hi, inclusive
    for x in range(lo, hi+1):
        if valid_palindrome(str(x)):
            part1 += x
        if valid_pattern(str(x)):
            part2 += x

print('Part 1:', part1)
print('Part 2:', part2)

