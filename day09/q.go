package day09

import (
	"fmt"
	"time"

	"aoc2024/util"
)

type Node struct {
	id, count  int
	next, prev *Node
}

func NewNode() *Node {
	return &Node{id: -1, count: 0}
}

const (
	invalidID = -1
)

func solve1(input string) string {
	buffer := make([]int, 0, len(input))
	nextID := 0
	for i, char := range input {
		repeatCount := int(char - '0')
		currID := invalidID
		if (i & 1) == 0 {
			currID = nextID
			nextID++
		}

		for j := 0; j < repeatCount; j++ {
			buffer = append(buffer, currID)
		}
	}

	result := 0
	left, right := 0, len(buffer)-1
	for left <= right {
		if buffer[left] != invalidID {
			result += left * buffer[left]
			left++
		}

		for left <= right && buffer[right] == invalidID {
			right--
		}

		if left <= right && buffer[left] == invalidID && buffer[right] != invalidID {
			result += left * buffer[right]
			left++
			right--
		}
	}

	return fmt.Sprintf("%d", result)
}

func solve2(input string) string {
	head := NewNode()
	tail := head
	nextID := 0

	for index, char := range input {
		repeatCount := int(char - '0')
		currID := invalidID
		if (index & 1) == 0 {
			currID = nextID
			nextID++
		}

		newNode := &Node{id: currID, count: repeatCount, prev: tail}
		tail.next = newNode
		tail = newNode
	}

	right := tail
	for right != head {
		if right.id == invalidID {
			right = right.prev
			continue
		}

		left := head.next
		for left != right {
			if left.id != invalidID {
				left = left.next
				continue
			}

			if left.count >= right.count {
				excessCount := left.count - right.count
				if excessCount > 0 {
					excessNode := &Node{id: invalidID, count: excessCount, prev: left, next: left.next}
					left.next.prev = excessNode
					left.next = excessNode
					left.count = right.count
				}

				left.id = right.id
				right.id = invalidID
				break
			}

			left = left.next
		}

		right = right.prev
	}

	result := 0
	position := 0
	curr := head.next
	for curr != nil {
		if curr.count == 0 {
			curr = curr.next
			continue
		}

		for i := 0; i < curr.count; i++ {
			if curr.id != invalidID {
				result += position * curr.id
			}
			position++
		}

		curr = curr.next
	}

	return fmt.Sprintf("%d", result)
}

func Run(day int, input []string) {
	for i, line := range input {
		if line == "" {
			util.Fatal("Invalid format on line %d: empty line", i)
		}
	}

	startTime := time.Now()
	util.Output(1, solve1(input[0]))
	elapsed := time.Since(startTime)
	util.TimeTaken(elapsed)

	util.Separator()

	startTime = time.Now()
	util.Output(2, solve2(input[0]))
	elapsed = time.Since(startTime)
	util.TimeTaken(elapsed)
}
