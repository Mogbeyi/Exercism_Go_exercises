package annalyn

func CanFastAttack(knightIsAwake bool) bool {
	return !knightIsAwake
}

func CanSpy(knightIsAwake, archerIsAwake, prisonerIsAwake bool) bool {
	return knightIsAwake || archerIsAwake || prisonerIsAwake 
}

func CanSignalPrisoner(archerIsAwake, prisonerIsAwake bool) bool {
	return prisonerIsAwake && !archerIsAwake
}

// func canfreeprisoner(knightisawake, archerisawake, prisonerisawake, petdogispresent bool) bool {
// 	if prisonerisawake && (!archerisawake && !knightisawake) {
// 		return true
// 	} else if petdogispresent && !archerisawake {
// 		return true
// 	} else {
// 		return false
// 	}
// }


// CanFreePrisoner can be executed if prisoner is awake and the other 3 characters are asleep
// or if Annalyn's pet dog is with her and the archer is sleeping
func CanFreePrisoner(knightIsAwake, archerIsAwake, prisonerIsAwake, petdogispresent bool) bool {
	return (prisonerIsAwake && !archerIsAwake && !knightIsAwake && !petdogispresent) || (petdogispresent && !archerIsAwake)
}

