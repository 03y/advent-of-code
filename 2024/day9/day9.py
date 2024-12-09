input = '2333133121414131402'

class Block:
    def __init__(self, size, empty):
        self.size = size
        self.empty = empty

size, empty = 1, False
blocks = []

for i in range(len(input)):
    blocks.append(Block(int(input[i]), empty))
    empty = not empty

print(input)

id = 0
fs = ''
for i in range(len(blocks)):
    print(id, blocks[i].size)
    if blocks[i].empty == True:
        fs += '.' * blocks[i].size
    else:
        fs += str(id) * blocks[i].size
        id += 1

print(fs)

def get_last(fs):
    pos = len(fs)-1
    while pos > 0:
        if fs[pos] != '.':
            return pos
        pos -= 1
    return -1

def get_free(fs):
    for i in range(len(fs)):
        if fs[i] == '.':
            return i
    return -1

import re

def is_sorted(fs):
    return re.search('\d+\.+\d+', fs) == None

print(is_sorted(fs))

def replace(s, c, index):
    return s[:index] + c + s[index + 1:]

sorted = False
while not sorted:
    if is_sorted(fs):
        sorted = True
    else:
        n = get_last(fs)
        fs = replace(fs, fs[n], get_free(fs))
        fs = replace(fs, fs[n], n)
        fs = replace(fs, '.', n)

print(is_sorted(fs))
print(fs)

checksum = 0
n = 0
curr = ''
for i in range(len(fs)):
    if i == 0:
        curr = fs[i]
    if fs[i] != curr and curr != '.':
        print(curr, n)
        checksum += int(curr) * n
        curr = fs[i]
        n += 1

print(checksum)
