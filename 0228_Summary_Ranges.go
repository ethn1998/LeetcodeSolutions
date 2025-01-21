import "fmt"

func summaryRanges(nums []int) []string {
    var i,j int
    fout := make([]string,0)
    if len(nums) == 0{
        return fout
    }
    for j < len(nums)-1{
        if nums[j+1] != nums[j]+1{
            if i != j{
                fout = append(fout,fmt.Sprintf("%d->%d",nums[i],nums[j]))
            } else{
                fout = append(fout,fmt.Sprintf("%d",nums[j]))
            }
            i = j+1
        }
        j++
    }
    if i < j{
        fout = append(fout,fmt.Sprintf("%d->%d",nums[i],nums[len(nums)-1]))
    }else{
        fout = append(fout,fmt.Sprintf("%d",nums[len(nums)-1]))
    }
    return fout
}
