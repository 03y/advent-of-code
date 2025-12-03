data = []
f = open('test_input.txt')
for line in f:
    row = []
    for char in line:
        if char != '\n':
            row.append(int(char))
    data.append(row)

part1, part2 = 0, 0

# gets the index of the highest value in the list
# the returned index cannot be in the filter list
def get_max_pos(arr: list[int], filter_arr: list[int]) -> int:
    max_value, pos = -1, 0

    for i in range(len(arr)):
        if i in filter_arr:
            continue
        elif arr[i] >= max_value:
            pos = i
            max_value = arr[i]

    return pos

for row in data:
    print(row)

    # add last index to filter, because we need 8..9 to give 89
    max_pos = get_max_pos(row, [len(row)-1])
    part1 += int(str(row[max_pos]) + str(max(row[max_pos+1:])))

    # part 2: we now need 12 digits
    max_positions, s = [], ''
    for i in range(12):
        if i == 0:
            next_max = get_max_pos(row, [len(row)-1])
        else:
            next_max = get_max_pos(row, max_positions)
        max_positions.append(next_max)

    for i in range(len(row)):
        if i in max_positions:
            print('\033[0;32m' + str(row[i]) + '\033[0m', end='')
        else:
            print(row[i], end='')
    print('')

    max_positions.sort()

    for x in max_positions:
        s += str(row[x])

    print(s + '\n')
    part2 += int(s)

print('''
987654321111
811111111119
434234234278
888911112111
''')

print('Part 1:', part1)
print('Part 2:', part2, part2 == 3121910778619)
