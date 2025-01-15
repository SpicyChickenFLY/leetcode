package main

func josephus(n int, k int) int {
    if n == 1 {
        return 0
    }
    return (josephus(n - 1, k) + k) % n
}

func findTheWinner(n int, k int) int {
    // return josephus(n, k) + 1

    winner := 0
    for i := 1; i <= n ; i++ {
        winner = (winner + k) % i
    }
    return winner + 1
}
