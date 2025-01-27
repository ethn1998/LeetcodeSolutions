func checkIfPrerequisite(numCourses int, prerequisites [][]int, queries [][]int) []bool {
    /*
    What's with all the directed graph questions lately?
    A course can be a prerequisite for more than one (other) course.
    A course can have multiple (indirect) prerequisites
    Idea: "Sparse" matrix for storing prerequisites. ith entry stores all courses that course i is a prerequisite of
    */
    //Initialize "database" as a "sparse matrix"
    prereqMat := make([]map[int]bool,numCourses) //There are a total of numCourses courses
    for i := range(prereqMat){
        prereqMat[i] = make(map[int]bool)
    }
    for _,edge := range(prerequisites){
            prereqMat[edge[0]][edge[1]] = true
    }
    // BRUTE FORCE IMPLEMENTATION IS HIGHLY INEFFICIENT
    //Each course is a prerequisite for at most numCourses-1 courses.
    for i := 0; i < numCourses; i++ { //Loop numCourses times for redundancy. This increases computation time though.
        noNewPrereq := true
        for j,indReq := range(prereqMat){ //Check for indirect prerequisites
            for k := range(indReq){
                for key := range(prereqMat[k]){
                    if !prereqMat[j][key]{
                        noNewPrereq = false
                        prereqMat[j][key] = true
                    }
                }
            }
        }
        if noNewPrereq { //Premature exit if no more new edges are added
            break
        }
    }
    //*/
    

    fout := make([]bool,0) //Initialize output
    for _,q := range(queries){
        fout = append(fout,prereqMat[q[0]][q[1]]) //Lookup
    }
    return fout
}
