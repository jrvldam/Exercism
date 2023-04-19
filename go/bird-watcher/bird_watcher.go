package birdwatcher

const (
	weekDays = 7
)

// TotalBirdCount return the total bird count by summing
// the individual day's counts.
func TotalBirdCount(birdsPerDay []int) int {
	total := 0
	for _, c := range birdsPerDay {
		total += c
	}

	return total
}

// BirdsInWeek returns the total bird count by summing
// only the items belonging to the given week.
func BirdsInWeek(birdsPerDay []int, week int) int {
	head := (week - 1) * weekDays
	tail := 0
	if head+weekDays == len(birdsPerDay) {
		tail = head + weekDays
	} else {
		tail = head + weekDays + 1
	}

	return TotalBirdCount(birdsPerDay[head:tail])
}

// FixBirdCountLog returns the bird counts after correcting
// the bird counts for alternate days.
func FixBirdCountLog(birdsPerDay []int) []int {
	for i := 0; i < len(birdsPerDay); i += 1 {
		if i%2 == 0 {
			birdsPerDay[i] = birdsPerDay[i] + 1
		}
	}

	return birdsPerDay
}
