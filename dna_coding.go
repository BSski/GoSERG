package main

var speed map[int]int = map[int]int{
	-5: 120,
	-4: 60,
	-3: 40,
	-2: 30,
	-1: 24,
	0:  20,
	1:  12,
	2:  10,
	3:  8,
	4:  6,
	5:  4,
	6:  2,
	7:  1,
}

var speedCost map[int]float64 = map[int]float64{
	0: 0.000,
	1: 0.005,
	2: 0.010,
	3: 0.015,
	4: 0.020,
	5: 0.025,
	6: 0.030,
	7: 0.035,
}

var bowelLength map[int]float64 = map[int]float64{
	0: 0.51,
	1: 0.59,
	2: 0.66,
	3: 0.73,
	4: 0.79,
	5: 0.86,
	6: 0.93,
	7: 1.00,
}

var bowelLengthCost map[int]float64 = map[int]float64{
	0: 0.00,
	1: 0.018,
	2: 0.036,
	3: 0.054,
	4: 0.072,
	5: 0.090,
	6: 0.108,
	7: 0.124,
}

var fatLimit map[int]float64 = map[int]float64{
	0: 1500,
	1: 2000,
	2: 2500,
	3: 3000,
	4: 3500,
	5: 4000,
	6: 4500,
	7: 5000,
}

var fatLimitCost map[int]float64 = map[int]float64{
	0: 0.00,
	1: 0.018,
	2: 0.036,
	3: 0.054,
	4: 0.072,
	5: 0.090,
	6: 0.108,
	7: 0.124,
}

var legsLength map[int]float64 = map[int]float64{
	0: 1.00,
	1: 0.965,
	2: 0.930,
	3: 0.895,
	4: 0.860,
	5: 0.825,
	6: 0.790,
	7: 0.755,
}

var legsLengthCost map[int]float64 = map[int]float64{
	0: 0.00,
	1: 0.018,
	2: 0.036,
	3: 0.054,
	4: 0.072,
	5: 0.090,
	6: 0.108,
	7: 0.124,
}
