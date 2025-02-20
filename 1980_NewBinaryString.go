import "fmt"

func findDifferentBinaryString(nums []string) string {
    var strout string
    for i,str := range(nums){
        char := string(str[i])
        if char == "0" {
            strout = fmt.Sprintf("%s%d",strout,1)
        } else {
            strout = fmt.Sprintf("%s%d",strout,0)
        }
    }
    return strout
}
