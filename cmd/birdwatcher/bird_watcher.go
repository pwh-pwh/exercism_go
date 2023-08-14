package birdwatcher

// TotalBirdCount return the total bird count by summing
// the individual day's counts.
func TotalBirdCount(birdsPerDay []int) int {
	count := 0
	for _, num := range birdsPerDay {
		count += num
	}
	return count
}

// BirdsInWeek returns the total bird count by summing
// only the items belonging to the given week.
func BirdsInWeek(birdsPerDay []int, week int) int {
	start := (week - 1) * 7
	count := 0
	end := start + 7
	for ; start < end && start < len(birdsPerDay); start++ {
		count += birdsPerDay[start]
	}
	return count
}

// FixBirdCountLog returns the bird counts after correcting
// the bird counts for alternate days.
func FixBirdCountLog(birdsPerDay []int) []int {
	for index, _ := range birdsPerDay {
		if index%2 == 0 {
			birdsPerDay[index] += 1
		}
	}
	return birdsPerDay
}
