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
def count_paths(graph, parent, goal, memo=None):
    if memo is None:
        memo = {}
    if parent == goal:
        return 1
    if parent in memo:
        return memo[parent]

    total = 0
    for child in graph[parent]:
        total += count_paths(graph, child, goal, memo)

    memo[parent] = total
    return total

graph = Digraph(name='Advent of Code 2025 day 11')

f = open('input.txt')
for l in f:
    node_name = l[:l.index(':')]
    node_children = l[l.index(':')+1:].split()

    with graph.subgraph() as Node:
        Node.attr(name=node_name)
        Node.node(node_name)

        for child in node_children:
            Node.node(child)
            graph.edge(node_name, child)

# uncomment to render graph
#graph.render('Advent of Code 2025 day 11', view=True)

nodes = digraph_to_adjlist(graph) # convert graphviz object to neighbour dict
print('Part 1:', count_paths(nodes, 'you', 'out'))
print('Part 2:', count_paths(nodes, 'svr', 'fft') *
                 count_paths(nodes, 'fft', 'dac') *
                 count_paths(nodes, 'dac', 'out'))