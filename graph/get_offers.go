package graph

import (
  "math"
  "log"
)

func CheckProfitable(edges map[int][]Offer) bool {
  product := 1.0
  for _, v := range edges {
    product = product * v[len(v)-1].Rate
  }
  return product > 1
}

func (graph *Graph) GetProfitableOffers() (map[int][]Offer) {

  // log.Println("XRP", graph.Graph["XRP"])
  // log.Println(graph.CreateSimpleGraph().Graph["XRP"])

  asset, predecessors := graph.CreateSimpleGraph().BellmanFord()
	if asset == "" {
    log.Println("No positive cycle")
    return nil
  }

	cycle := GetCycle(asset, predecessors)

  // Make a copy of the graph
  res := make(map[int][]Offer)
  quantities := make(map[int]float64)
  cycle_count := len(cycle)

  log.Println("cycle", cycle)

  for i, _ := range cycle {
    // Get best edge
    // log.Println("graph.Graph[cycle[i]]", graph.Graph[cycle[i]], len(graph.Graph), len(graph.Graph[cycle[i]]))
    edges := graph.Graph[cycle[i]][cycle[(i+1)%cycle_count]]

    if edges == nil || edges.List[0] == nil {
      panic("Should never happen")
    }

    edge := edges.List[0]
    // Update total quantities for that edge
    quantities[i] = edge.Volume
    // Remove used edge from graph
    graph.Graph[cycle[i]][cycle[(i+1)%cycle_count]].List = graph.Graph[cycle[i]][cycle[(i+1)%cycle_count]].List[1:]
    // Update selected edges
    res[i] = make([]Offer, 1000000)
    res[i][0] = *edge
  }

  if !CheckProfitable(res) {
    panic("Positive cycle doesn't exist.")
  }

  for true {
    minQuantity := math.MaxFloat64
    for _, v := range quantities {
      if (v < minQuantity) {
        minQuantity = v
      }
    }

    next_edges := make([]*Offer, 100)

    var bottleneck_edge *int

    for i, v := range quantities {
      if (v == minQuantity && len(graph.Graph[cycle[i]][cycle[(i+1)%cycle_count]].List) > 0) {
        *bottleneck_edge = i
      }
    }

    if bottleneck_edge == nil {
      return res
    } else {
      // Getting next edges
      copy(next_edges, graph.Graph[cycle[*bottleneck_edge]][cycle[(*bottleneck_edge+1)%cycle_count]].List)
      // Removing first one if existing
      graph.Graph[cycle[*bottleneck_edge]][cycle[(*bottleneck_edge+1)%cycle_count]].List = graph.Graph[cycle[*bottleneck_edge]][cycle[(*bottleneck_edge+1)%cycle_count]].List[1:]
    }

    next_edge := next_edges[capacityList-1]

    product := 1.0
    for i, edges := range res {
      if *bottleneck_edge == i {
          product = product * next_edge.Rate
      } else {
          product = product * edges[len(edges)-1].Rate
      }
    }

    if product > 1 {
      quantities[*bottleneck_edge] += next_edge.Volume
      res[*bottleneck_edge] = append(res[*bottleneck_edge], *next_edge)
    } else {
      return res
    }

  }

  return nil
}
