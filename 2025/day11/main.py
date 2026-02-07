from igraph import *

def find_all_paths(graph, start, end):
    if start == end:
        return (1, 0, 0, 0)
    numPath = graph.vs[start]["numPath"]
    if numPath >= 0:
        return (numPath, graph.vs[start]["numPathDAC"], graph.vs[start]["numPathFFT"], graph.vs[start]["numPathBoth"])
    numPath = 0
    numPathDAC = 0
    numPathFFT = 0
    numPathBoth = 0
    neighbors = graph.neighbors(start, mode="out")
    dacVisited = graph.vs[start]["name"] == "dac"
    fftVisited = graph.vs[start]["name"] == "fft"
    for vertex in neighbors:
        nPath, nPathDAC, nPathFFT, nPathBoth = find_all_paths(graph, vertex, end)
        numPath = numPath + nPath
        numPathDAC = numPathDAC + nPathDAC
        numPathFFT = numPathFFT + nPathFFT
        numPathBoth = numPathBoth + nPathBoth
               
    if graph.vs[start]["name"] == "dac":
        numPathBoth = numPathFFT
        numPathDAC = numPath
    if graph.vs[start]["name"] == "fft":
        numPathBoth = numPathDAC
        numPathFFT = numPath
    graph.vs[start]["numPath"] = numPath
    graph.vs[start]["numPathDAC"] = numPathDAC
    graph.vs[start]["numPathFFT"] = numPathFFT
    graph.vs[start]["numPathBoth"] = numPathBoth
    return (numPath, numPathDAC, numPathFFT, numPathBoth)

# ****** MAIN ******
graph = Graph(directed=True)

with open("input.txt", 'r') as file:
    lines = file.readlines()

for line in lines:
    node_name, n = line.split(':')
    graph.add_vertex(node_name, numPath=-1)
graph.add_vertex("out", numPath=-1)

for line in lines:    
    node_name, n = line.split(':')
    neighbors = n.strip().split(' ')
    for neighbor in neighbors:
        graph.add_edge(node_name, neighbor)

(numPath, numPathDAC, numPathFFT, numPathBoth) = find_all_paths(graph, graph.vs.find("svr").index, graph.vs.find("out").index)
print("Answer: ", numPathBoth)
  
