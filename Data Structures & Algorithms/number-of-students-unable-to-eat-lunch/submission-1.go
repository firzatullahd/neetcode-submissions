func countStudents(students []int, sandwiches []int) int {
	output := 0
	for i := range sandwiches {
		if sandwiches[i] == students[0] {
			students = dequeue(students)
			continue
		}

		count := 0
		for sandwiches[i] != students[0] {
			students = requeue(students)

			if sandwiches[i] == students[0] {
				students = dequeue(students)
				break
			}

			if count == len(students) {
				return len(students)
			}
			count++
		}

	}

	return output
}

func requeue(students []int) []int{
	return append(students[1:], students[0])
}

func dequeue(students []int) []int {
	return students[1:]
}
