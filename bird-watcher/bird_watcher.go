package birdwatcher

// TotalBirdCount return the total bird count by summing
// the individual day's counts.
func TotalBirdCount(birdsPerDay []int) int {
	var total int

	for _, bird := range birdsPerDay {
		total += bird
	}

	return total
}

// BirdsInWeek returns the total bird count by summing
// only the items belonging to the given week.
func BirdsInWeek(birdsPerDay []int, week int) int {
	startPos := (week - 1) * 7

	return TotalBirdCount(birdsPerDay[startPos: startPos + 7])
}

// FixBirdCountLog returns the bird counts after correcting
// the bird counts for alternate days.
func FixBirdCountLog(birdsPerDay []int) []int {
	for i := 0; i * 2 < len(birdsPerDay); i++ {
		birdsPerDay[i * 2] += 1
	}

	return birdsPerDay
}
