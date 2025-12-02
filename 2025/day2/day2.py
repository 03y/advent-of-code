def valid_id(s: str) -> bool:
    # not valid if length 1
    if len(s) < 2:
        return False

    # not valid if leading zero
    if s[0] == '0':
        return False

    # chop string in half
    front = s[:len(s)//2]
    back  = s[len(s)//2:]

    # valid if palindrome
    return front == back

part1 = []

with open('input.txt', 'r') as f:
    for line in f:
        for id_range in line.split(','):
            # optimisation: 0-10 are all invalid IDs so we never check them
            lo = max(int(id_range.split('-')[0]), 11)
            hi = int(id_range.split('-')[1])

            # loop from lo to hi, inclusive
            for x in range(lo, hi+1):
                if valid_id(str(x)):
                    part1.append(x)

print('Part 1:', sum(part1))
