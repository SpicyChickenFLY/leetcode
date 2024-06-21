package problems

func maxSatisfied(customers []int, grumpy []int, minutes int) int {
    current := 0
    for i := 0 ; i < len(customers); i++ {
        if i < minutes || grumpy[i] == 0 {
            current += customers[i]
        }
    }
    result := current

    for start := 0; start + minutes < len(customers); start++ {
        current -= (grumpy[start]) * customers[start]
        current += grumpy[start + minutes] * customers[start + minutes]
        if current > result {
            result = current
        }
    }

    return result
}
