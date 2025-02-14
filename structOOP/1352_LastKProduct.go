//Challenge: O(1) time complexity instead of O(k).

type ProductOfNumbers struct {
    stream []int //IDEA: Make this a stream of products rather than a stream of inputs.
}


func Constructor() ProductOfNumbers {
    stream := make([]int,0)
    return ProductOfNumbers{stream}
}


func (this *ProductOfNumbers) Add(num int)  {
    //this.stream = append(this.stream,num)
    if num == 0 {
        this.stream = make([]int,0) //Reset
    } else if len(this.stream) == 0 { //First nonzero
        this.stream = append(this.stream,num) 
    } else {
        this.stream = append(this.stream,num*this.stream[len(this.stream)-1])
    }
}


func (this *ProductOfNumbers) GetProduct(k int) int {
    /* This runs in O(k) time
    fout := 1
    n := len(this.stream) - 1
    for i := 0; i < k; i++ {
        fout *= this.stream[n-i]
    }
    */
    //Lookup should be O(1)
    if k > len(this.stream) { //This reaches the previous zero.
        return 0
    } else if k == len(this.stream) { //Product of all numbers
        return this.stream[len(this.stream)-1]
    }
    p := this.stream[len(this.stream)-1] //Numerator
    q := this.stream[len(this.stream)-k-1] //Denominator
    return p/q
}


/**
 * Your ProductOfNumbers object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Add(num);
 * param_2 := obj.GetProduct(k);
 */
