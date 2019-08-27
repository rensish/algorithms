package maxSubArray

func main(nums []int) int {
    max := nums[0]
    cur := 0
    
    for i := 0 ; i < len(nums) ; i++ {
        if cur < 0 {
            cur = nums[i]
        } else {
            cur += nums[i]
        }
        
        if cur > max {
            max = cur
        }
    }
    
    return max
}