package iourt43

var KillWeapons = map[int]string{
	1: "Water",
	3: "Lava",
	5: "Telefrag",
	6: "Falling",
	7: "Suicide",
	9: "Trigger",
	10: "Change team",
	12: "Knife",
	13: "Thrown knife",
	14: "Beretta",
	15: "Desert Eagle",
	16: "Spas",
	17: "UMP45",
	18: "MP5K",
	19: "LR300",
	20: "G36",
	21: "PSG1",
	22: "HK69",
	23: "Bled",
	24: "Kicked",
	25: "HE Grenade",
	28: "SR8",
	30: "AK103",
	31: "Sploded",
	32: "Slapped",
	33: "Smitted",
	34: "Bombed",
	35: "Nuked",
	36: "Negev",
	37: "HK69 Hit",
	38: "M4",
	39: "Glock",
	40: "Colt 911",
	41: "Mac 11",
	42: "FR-F1",
	43: "Benelli",
	44: "P90",
	45: "Magnum",
	46: "TOD50",
	47: "Flag",
	48: "Goomba",
}

var Locations = map[int]string{
	1: "Head",
	2: "Helmet",
	3: "Torso",
	4: "Vest",
	5: "Left Arm",
	6: "Right arm",
	7: "Groin",
	8: "Butt",
	9: "Left upper leg",
	10: "Right upper leg",
	11: "Left lower leg",
	12: "Right lower leg",
	13: "Left foot",
	14: "Right foot",
}

var HitWeapons = map[int]string {
	1: "Knife",
	2: "Beretta",
	3: "Desert Eagle",
	4: "Spas",
	5: "MP5K",
	6: "UMP45",
	8: "LR300",
	9: "G36",
	10: "PSG-1",
	14: "SR8",
	15: "AK103",
	17: "Negev",
	19: "M4",
	20: "Glock",
	21: "Colt 911",
	22: "MAC 11",
	23: "FRF1",
	24: "Benelli",
	25: "P90",
	26: "Magnum",
	27: "TOD50",
	28: "Kicked",
	29: "Thrown knife",
}

var HitDamage = map[int][]int {
	1: []int{100, 60, 44, 35, 20, 20, 40, 37, 20, 20, 18, 18, 15, 15},
	2: []int{100, 40, 33, 22, 13, 13, 22, 22, 15, 15, 13, 13, 11, 11},
	3: []int{100, 66, 57, 38, 22, 22, 42, 38, 28, 28, 22, 22, 18, 18},
	4: []int{100, 80, 80, 40, 32, 32, 59, 59 ,40, 40, 40, 40, 40, 40},
	5: []int{ 50, 34, 30, 22, 13, 13, 22, 20, 15, 15, 13, 13, 11, 11},
	6: []int{100, 51, 44, 29, 17, 17, 31, 28, 20, 20, 17, 17, 14, 14},
	8: []int{100, 51, 44, 29, 17, 17, 31, 28, 20, 20, 17, 17, 14, 14},
	9: []int{100, 51, 44, 29, 17, 17, 29, 28, 20, 20, 17, 17, 14, 14},
	10: []int{100, 100, 97, 70, 36, 36, 75, 70, 41, 41, 36, 36, 29, 29},
	14: []int{100, 100, 100, 100, 50, 50, 100, 100, 60, 60, 50, 50, 40, 40},
	15: []int{100, 58, 51, 34, 19, 19, 39, 35, 22, 22, 19, 19, 15, 15},
	17: []int{ 50, 34, 30, 22, 11, 11, 23, 21, 13, 13, 11, 11,  9,  9},
	19: []int{100, 51, 44, 29, 17, 17, 31, 28, 20, 20, 17, 17, 14, 14},
	20: []int{100, 45, 29, 35, 15, 15, 29, 27, 20, 20, 15, 15, 11, 11},
	21: []int{100, 60, 40, 30, 15, 15, 32, 29, 22, 22, 15, 15, 11, 11},
	22: []int{ 50, 29, 20, 16, 13, 13, 16, 15, 15, 15, 13, 13, 11, 11},
	23: []int{100, 100, 96, 76, 40, 40, 76, 74, 50, 50, 40, 40, 30, 30},
	24: []int{100, 100, 90, 67, 32, 32, 60, 50, 35, 35, 30, 30, 20, 20},
	25: []int{ 50, 40, 33, 27, 16, 16, 27, 25, 17, 17, 15, 15, 12, 12},
	26: []int{100, 82, 66, 50, 33, 33, 57, 52, 40, 33, 25, 25},
	27: []int{100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100},
	28: []int{30, 20, 20, 20, 20, 20, 20, 20, 20, 20, 20, 20, 20, 20},
	29: []int{100, 60, 44, 35, 20, 20, 40, 37, 20, 20, 18, 18, 15, 15},
}