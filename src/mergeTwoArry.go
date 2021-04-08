package leetcode

func Merge(A,B []int,m,n,int)  {
	sortedArry := make([]int, 0,m+n)
	pa,pb := 0,0
	for{
		if pa==m{
			sortedArry = append(sortedArry,B[pb:]...)
			break
		}
		if pb==n{
			sortedArry = append(sortedArry,A[pa:]...)
			break
		}
		if A[pa]<B[pb]{
			sortedArry = append(sortedArry,A[pa])
			pa++
		}
		else{
			sortedArry = append(sortedArry,B[pb])
			pb++
		}
	}
	copy(A,sortedArry)
}