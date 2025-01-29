class Solution(object):
    def findRedundantConnection(self, edges):
        """
        :type edges: List[List[int]]
        :rtype: List[int]
        """
        """
        NOTE: 
        DEFN: A tree is a (1) connected (2) acyclic (3) undirected graph.
        If there are multiple answers, return the final valid answer in input edges.
        (3) is trivially satisfied.
        (1) Eliminate all answers that link only one edge.
        IDEA: Some kind of variant of Kahn's algorithm targetting nodes with only one edge,
              and remove the edge connecting said node from edges in place. Then return final answer remaining.
        """
        edgeCount = {} #Dictionary for storing number of edges of each node
        for edge in edges:
            for node in edge:
                if node in edgeCount:
                    edgeCount[node] += 1
                else:
                    edgeCount[node] = 1
        #Remove all essential edges
        Rep = True
        while Rep:
            tNode = -1
            Rep = False
            #Locate node with only one edge.
            for node in edgeCount:
                if edgeCount[node] == 1:
                    tNode = node
                    Rep = True 
            #Trim edge
            for i,edge in enumerate(edges):
                if edge[0] == tNode:
                    edgeCount[edge[1]] -= 1
                    edgeCount[tNode] -= 1
                    edges.pop(i)
                    break
                if edge[1] == tNode:
                    edgeCount[edge[0]] -= 1
                    edgeCount[tNode] -= 1
                    edges.pop(i)
                    break
        return edges[-1]

        
