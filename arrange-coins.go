func arrangeCoins(n int) int {
    staircaseCount := 0
    
    i := 1
    leftCoins := n
    for {
        if leftCoins == 0 || leftCoins - i < 0  {
            break
        }
        leftCoins = leftCoins - i
        i++
        staircaseCount++
    }
    
    return staircaseCount
}
