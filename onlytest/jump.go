package main

package main

import "fmt"

func jump(nums []int) int {
	if len(nums) <= 1 {
		return 0
	}

	jumps := 0      // 记录跳跃次数
	farthest := 0    // 当前跳跃能到达的最远距离
	nextFarthest := 0 // 下一次跳跃能到达的最远距离==之所以维护这两个变量，是因为在到达下一跳的过程中，不断维护最大的下一跳量，这样每一步
	//都是最大的跨度，这样属于局部最优且无后效性，算是一种贪婪+动态规划

	for i := 0; i < len(nums)-1; i++ {
		// 更新下一次跳跃能到达的最远距离
		if i + nums[i] > nextFarthest {
			nextFarthest = i + nums[i]
		}
// nums = [2,3,1,1,4]
		// 如果已经到达当前跳跃的最远距离，则进行下一次跳跃
		if i == farthest {//初始化情况是一定能到达这里的，岂不是在index0出就已经开始算跳了？太不合理了，但是吧这个档做跳点，
		//那肯定是算一跳的了，也就是只算起点，这里一定会发生一跳，当越位时，（那就不算最后一次了），因为只有起跳点才算，理解为此点起跳
			//我觉得这是最巧妙的一步，把边界条件设置为起跳点就是jump++，那么这样的话，我们就可以给farthest初始化为0了，也就是当前可到点为0，
			//起跳，最后维护最大跳跨度
			jumps++
			farthest = nextFarthest

			// 如果已经到达最后一个位置，则提前返回
			if farthest >= len(nums)-1 {//这个是
				break
			}
		}
	}

	return jumps
}
//总结：贪婪+动规：我们维护一个当前最远，再沿着路径遍历到当前最远处时，必须维护一个下一个最远，
//起跳点的个数算跳数，这样可以吧当前最远和下一个最远设置为0

func jump2(nums []int) int {
	farthest:=0
	nextFarthest:=0
	jumps:=0
	for i := 0; i < len(nums); i++ {
		if farthest>=len(nums)-1{
			return jumps
		}
		nextFarthest=max(nextFarthest,i+nums[i])
		if(i==farthest) {
			jumps++
			farthest = nextFarthest
		}



	}
	return jumps

}
func canJump2(nums []int)bool{

	imax:=0
	for i := 0; i <=imax; i++ {
		m_each:=i+nums[i]
		imax=max(imax,m_each)
		if(imax>=len(nums)-1){
			return true
		}
	}
	return false

}