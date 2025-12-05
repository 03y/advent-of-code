f = open('test_input.txt')

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
print(sorted_ranges)

part2 = 0
merged_range = sorted_ranges[0]
cur_lo, cur_hi = 0, 0

for lo, hi in sorted_ranges:
    if cur_lo == 0 and cur_hi == 0:
        cur_lo, cur_hi = merged_range[0], merged_range[1]
        continue

    print('cur_lo:', cur_lo)
    print('cur_hi:', cur_hi)
    print('lo:', lo)

    # overlap
    if lo <= cur_hi + 1:
        print('overlap')
        merged_range = (cur_hi, hi)
    # gap
    elif lo > cur_hi + 1:
        print('gap')
        part2 += (cur_hi - cur_lo + 1)
        merged_range = (lo, hi)

part2 += sorted_ranges[-1][1] - sorted_ranges[-1][0]

print('Part 2:', part2)
