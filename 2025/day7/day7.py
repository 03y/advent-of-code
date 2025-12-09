from collections import defaultdict

data = []
f = open('input.txt')
for l in f:
    temp = []
    for c in l:
        if c != '\n':
            temp.append(c)
    data.append(temp)

# i called it 'seed' as that makes more sense in my head
seeds = {data[0].index('S'): 1}
part1, part2  = 0, 0
for i in range(1, len(data)): # skip row 0
    new_seeds = defaultdict(int) # dict of all ints

    for pos, n in seeds.items():
        if data[i][pos] == '.':
            new_seeds[pos] += n
        else:
            part1 += 1
            # a seed means we add two more lines below
            # we do += because if we count the same line twice, thats still one line
            # for part 2 we need to know how many times a new line was created
            new_seeds[pos - 1] += n
            new_seeds[pos + 1] += n

    seeds = new_seeds

print('Part 1:', part1)
print('Part 2:', sum(seeds.values())) # we add up all the positions the lines moved to