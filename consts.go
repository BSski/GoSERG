package main

type consts struct {
	partialDnaRange [4]int
	vonNeumannPerms [24][4][2]int
	chartsInfo      map[int]map[string]interface{}
}

type chartData struct {
	Label  string
	X      int
	Y      int
	Offset int
}

var c = consts{
	partialDnaRange: [4]int{2, 3, 4, 5},
	vonNeumannPerms: [24][4][2]int{
		{{1, 0}, {-1, 0}, {0, 1}, {0, -1}},
		{{1, 0}, {-1, 0}, {0, -1}, {0, 1}},
		{{1, 0}, {0, 1}, {-1, 0}, {0, -1}},
		{{1, 0}, {0, 1}, {0, -1}, {-1, 0}},
		{{1, 0}, {0, -1}, {-1, 0}, {0, 1}},
		{{1, 0}, {0, -1}, {0, 1}, {-1, 0}},
		{{-1, 0}, {1, 0}, {0, 1}, {0, -1}},
		{{-1, 0}, {1, 0}, {0, -1}, {0, 1}},
		{{-1, 0}, {0, 1}, {1, 0}, {0, -1}},
		{{-1, 0}, {0, 1}, {0, -1}, {1, 0}},
		{{-1, 0}, {0, -1}, {1, 0}, {0, 1}},
		{{-1, 0}, {0, -1}, {0, 1}, {1, 0}},
		{{0, 1}, {1, 0}, {-1, 0}, {0, -1}},
		{{0, 1}, {1, 0}, {0, -1}, {-1, 0}},
		{{0, 1}, {-1, 0}, {1, 0}, {0, -1}},
		{{0, 1}, {-1, 0}, {0, -1}, {1, 0}},
		{{0, 1}, {0, -1}, {1, 0}, {-1, 0}},
		{{0, 1}, {0, -1}, {-1, 0}, {1, 0}},
		{{0, -1}, {1, 0}, {-1, 0}, {0, 1}},
		{{0, -1}, {1, 0}, {0, 1}, {-1, 0}},
		{{0, -1}, {-1, 0}, {1, 0}, {0, 1}},
		{{0, -1}, {-1, 0}, {0, 1}, {1, 0}},
		{{0, -1}, {0, 1}, {1, 0}, {-1, 0}},
		{{0, -1}, {0, 1}, {-1, 0}, {1, 0}},
	},
	chartsInfo: map[int]map[string]interface{}{
		0: {
			"type": "trait",
			"charts": []chartData{
				{"SPEED", 870, 75, 60},
				{"BOWEL LENGTH", 870, 225, 37},
				{"FAT LIMIT", 870, 375, 53},
				{"LEGS LENGTH", 870, 525, 44},
			},
		},
		1: {
			"title": "HERBIVORES",
			"type":  "distribution",
			"charts": []chartData{
				{"SPEED DISTRIBUTION", 874, 88, 25},
				{"BOWEL LENGTH DISTRIBUTION", 874, 238, -2},
				{"FAT LIMIT DISTRIBUTION", 874, 388, 16},
				{"LEGS LENGTH DISTRIBUTION", 874, 538, 5},
			},
		},
		2: {
			"title": "CARNIVORES",
			"type":  "distribution",
			"charts": []chartData{
				{"SPEED DISTRIBUTION", 874, 88, 25},
				{"BOWEL LENGTH DISTRIBUTION", 874, 238, -2},
				{"FAT LIMIT DISTRIBUTION", 874, 388, 16},
				{"LEGS LENGTH DISTRIBUTION", 874, 538, 5},
			},
		},
	},
}
