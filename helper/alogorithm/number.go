package algorithm

const (
	Positive = "Positive Order"
	Reverse  = "Reverse order"
)

func BubbleSort(arr []int, Order string) []int {
	for i := len(arr) - 1; i > 0; i-- {
		for j := 0; j < i; j++ {
			switch Order {
			case Positive:
				if arr[j] < arr[j+1] {
					arr[j], arr[j+1] = arr[j+1], arr[j]
				}
			case Reverse:
				if arr[j] > arr[j+1] {
					arr[j], arr[j+1] = arr[j+1], arr[j]
				}
			}
		}
	}
	return arr
}
