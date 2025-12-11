from graphviz import Digraph
from collections import defaultdict

def digraph_to_adjlist(dot):
    adj = defaultdict(list)

    for edge in dot.body:
        if '->' in edge:
            left, right = edge.split('->')
            tail = left.strip()
            head = right.strip().strip(';')
            adj[tail].append(head)

    return adj

# count paths using memoisation
def count_paths(graph, start, end, memo=None):
    if memo is None:
        memo = {}
    if start == end:
        return 1
    if start in memo:
        return memo[start]

    total = 0
    for neighbour in graph[start]:
        total += count_paths(graph, neighbour, end, memo)

    memo[start] = total
    return total

graph = Digraph(name='aoc')

f = open('input.txt')
for l in f:
    node_name = l[:l.index(':')]
    node_children = l[l.index(':')+1:].split()

    with graph.subgraph() as Devices:
        Devices.attr(name=node_name)
        Devices.node(node_name)

        for child in node_children:
            Devices.node(child)
            graph.edge(node_name, child)

# uncomment to render graph
#graph.render('aoc', view=True)

adjlist = digraph_to_adjlist(graph) # convert graphviz object to neighbour dict
print('Part 1:', count_paths(adjlist, 'you', 'out'))