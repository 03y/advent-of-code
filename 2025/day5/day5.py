f = open('input.txt')

ranges = []
data = []

read_ranges = False

for line in f:
    if line == '\n':
        read_ranges = True
        continue
    if read_ranges:
        data.append(int(line))
    else:
        ranges.append((int(line.split('-')[0]), int(line.split('-')[1])))

part1 = 0

for x in data:
    for lo, hi in ranges:
        if x >= lo and x <= hi:
            part1 += 1
            break

print('Part 1:', part1)

sorted_ranges = sorted(ranges, key=lambda x: x[0])

part2 = 0
cur_lo, cur_hi = sorted_ranges[0][0], sorted_ranges[0][1]

for lo, hi in sorted_ranges:
    # overlap
    if lo <= cur_hi + 1:
        cur_hi = max(cur_hi, hi)
    # gap
    elif lo > cur_hi + 1:
        part2 += (cur_hi - cur_lo + 1)
        cur_lo, cur_hi = lo, hi

part2 += cur_hi - cur_lo + 1
print('Part 2:', part2)

