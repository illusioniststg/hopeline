package planning

import (
	"container/heap"
	"math"
	"net/http"

	"github.com/gin-gonic/gin"
)

// =====================
// DATA MODELS (ALL STRUCTS PREFIXED WITH S)
// =====================

type SLocation struct {
	Lat float64 `json:"lat"`
	Lon float64 `json:"lon"`
}

type SBoat struct {
	ID                   string    `json:"boat_id"`
	Location             SLocation `json:"location"`
	Capacity             int       `json:"capacity"`
	SpeedKmph            float64   `json:"speed_kmph"`
	FuelCapacity         float64   `json:"fuel_capacity"`
	FuelConsumptionPerKm float64   `json:"fuel_consumption_per_km"`
}

type SSurvivor struct {
	ID          string    `json:"survivor_id"`
	Location    SLocation `json:"location"`
	PeopleCount int       `json:"people_count"`
	Priority    string    `json:"priority"`
}

type SShelter struct {
	ID       string    `json:"shelter_id"`
	Location SLocation `json:"location"`
}

type SRouteRequest struct {
	Boat      SBoat       `json:"boat"`
	Survivors []SSurvivor `json:"survivors"`
	Shelter   SShelter    `json:"destination_shelter"`
}

// =====================
// GRAPH MODELS
// =====================

type SNode struct {
	ID  string
	Loc SLocation
}

type SEdge struct {
	To       string
	Distance float64
	Blocked  bool
}

type SGraph map[string][]SEdge

// =====================
// PRIORITY QUEUE
// =====================

type SPQItem struct {
	Node     string
	Priority float64
	Index    int
}

type SPriorityQueue []*SPQItem

func (pq SPriorityQueue) Len() int { return len(pq) }
func (pq SPriorityQueue) Less(i, j int) bool {
	return pq[i].Priority < pq[j].Priority
}
func (pq SPriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].Index = i
	pq[j].Index = j
}
func (pq *SPriorityQueue) Push(x interface{}) {
	item := x.(*SPQItem)
	item.Index = len(*pq)
	*pq = append(*pq, item)
}
func (pq *SPriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[0 : n-1]
	return item
}

// =====================
// HARD-CODED WATER GRAPH
// =====================

var nodes = map[string]SNode{
	"A": {"A", SLocation{18.5204, 73.8567}},
	"B": {"B", SLocation{18.5300, 73.8700}},
	"C": {"C", SLocation{18.5400, 73.8900}},
	"D": {"D", SLocation{18.5500, 73.9100}},
}

var waterGraph = SGraph{
	"A": {
		{"B", 3, false},
		{"C", 10, false},
	},
	"B": {
		{"C", 4, false},
		{"D", 12, false},
	},
	"C": {
		{"D", 5, false},
	},
}

// =====================
// UTILS
// =====================

func haversine(a, b SLocation) float64 {
	const R = 6371
	dLat := (b.Lat - a.Lat) * math.Pi / 180
	dLon := (b.Lon - a.Lon) * math.Pi / 180

	lat1 := a.Lat * math.Pi / 180
	lat2 := b.Lat * math.Pi / 180

	x := math.Sin(dLat/2)*math.Sin(dLat/2) +
		math.Sin(dLon/2)*math.Sin(dLon/2)*math.Cos(lat1)*math.Cos(lat2)

	return 2 * R * math.Asin(math.Sqrt(x))
}

// =====================
// DIJKSTRA
// =====================

func dijkstra(start, end string) ([]string, float64, bool) {
	dist := map[string]float64{}
	prev := map[string]string{}

	for k := range nodes {
		dist[k] = math.Inf(1)
	}

	dist[start] = 0

	pq := &SPriorityQueue{}
	heap.Init(pq)
	heap.Push(pq, &SPQItem{Node: start, Priority: 0})

	for pq.Len() > 0 {
		item := heap.Pop(pq).(*SPQItem)
		u := item.Node

		if u == end {
			break
		}

		for _, edge := range waterGraph[u] {
			if edge.Blocked {
				continue
			}
			alt := dist[u] + edge.Distance
			if alt < dist[edge.To] {
				dist[edge.To] = alt
				prev[edge.To] = u
				heap.Push(pq, &SPQItem{Node: edge.To, Priority: alt})
			}
		}
	}

	if math.IsInf(dist[end], 1) {
		return nil, 0, false
	}

	path := []string{}
	for at := end; at != ""; at = prev[at] {
		path = append([]string{at}, path...)
	}

	return path, dist[end], true
}

// =====================
// SURVIVOR PICKUP ORDER (GREEDY)
// =====================

func orderSurvivors(start SLocation, survivors []SSurvivor) []SSurvivor {
	result := []SSurvivor{}
	current := start
	remaining := survivors

	for len(remaining) > 0 {
		bestIdx := 0
		bestDist := math.Inf(1)

		for i, s := range remaining {
			d := haversine(current, s.Location)
			if d < bestDist {
				bestDist = d
				bestIdx = i
			}
		}

		result = append(result, remaining[bestIdx])
		current = remaining[bestIdx].Location
		remaining = append(remaining[:bestIdx], remaining[bestIdx+1:]...)
	}

	return result
}

// =====================
// API HANDLER
// =====================

func MultiPickupRoute(c *gin.Context) {
	var req SRouteRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	totalPeople := 0
	for _, s := range req.Survivors {
		totalPeople += s.PeopleCount
	}

	if totalPeople > req.Boat.Capacity {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "FAILED",
			"reason": "BOAT_CAPACITY_EXCEEDED",
		})
		return
	}

	ordered := orderSurvivors(req.Boat.Location, req.Survivors)

	totalDistance := 0.0
	fullPath := []string{}
	currentNode := "A"

	for range ordered {
		path, dist, ok := dijkstra(currentNode, "B")
		if !ok {
			c.JSON(http.StatusBadRequest, gin.H{
				"status": "FAILED",
				"reason": "NO_ROUTE_TO_SURVIVOR",
			})
			return
		}
		fullPath = append(fullPath, path...)
		totalDistance += dist
		currentNode = "B"
	}

	path, dist, ok := dijkstra(currentNode, "D")
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "FAILED",
			"reason": "NO_ROUTE_TO_SHELTER",
		})
		return
	}

	fullPath = append(fullPath, path...)
	totalDistance += dist

	fuelNeeded := totalDistance * req.Boat.FuelConsumptionPerKm
	if fuelNeeded > req.Boat.FuelCapacity {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "FAILED",
			"reason": "INSUFFICIENT_FUEL",
		})
		return
	}

	etaMinutes := (totalDistance / req.Boat.SpeedKmph) * 60

	c.JSON(http.StatusOK, gin.H{
		"status": "SUCCESS",
		"path":   fullPath,
		"metrics": gin.H{
			"distance_km":    totalDistance,
			"eta_minutes":    etaMinutes,
			"fuel_used":      fuelNeeded,
			"people_rescued": totalPeople,
		},
	})
}
