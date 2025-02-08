// OOP problem!

type NumberContainers struct { //This is just a hashmap with extra steps
    iv map[int]int //Given index, store value
    vi map[int]int //Given value, store min index
}


func Constructor() NumberContainers { //Initializer
    iv := make(map[int]int)
    vi := make(map[int]int)
    return NumberContainers{iv,vi}
}


func (this *NumberContainers) Change(index int, number int)  {
    //this.iv[index] = number //Naive implementation leads to TLE due to search during find function.
    
    //HINT: Update minimum value here instead. Only need to check few indices.
    v, okiv := this.iv[index] //Check if index already has something inside
    i0, okvi := this.vi[number] //Check minimum index, whether it exists
    if okiv { //If something else is already present inside the index
        if v != number && index == this.vi[v] { //If this is the minimum index of another value
            mode := true //Boolean for triggering search mode. true: min not found
            for key,val := range(this.iv) { //Look for next minimum value
                if mode {
                    if key != index && val == v {//Existence of another key storing value val
                        mode = false
                        this.vi[val] = key
                    }
                } else { //Look for potential smaller candidates
                    if key != index && val == v && key < this.vi[val]{
                        this.vi[val] = key
                    }
                }
            }
            if mode { //If value no longer exists in map
                delete(this.vi,v)
            }
        }
    }
    this.iv[index] = number

    if okvi { //If we have already seen this value before
        if index < i0 { //If we find a smaller value
            this.vi[number] = index
        }
    } else {
        this.vi[number] = index
    }
}


func (this *NumberContainers) Find(number int) int {
    
    i0, ok := this.vi[number]
    if ok {
        return i0
    }
    return -1
    
    /* Naive implementation leads to TLE due to looping over large map/array
    fout := -1
    mode := true //Search mode. True if number is not found
    for key, val := range(this.iv) {
        if mode {
            if val == number {
                fout = key
                mode = false
            }
        } else {
            if val == number && key < fout {
                fout = key
            }
        }
    }
    return fout
    //*/
}


/**
 * Your NumberContainers object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Change(index,number);
 * param_2 := obj.Find(number);
 */
