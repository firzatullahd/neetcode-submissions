func isAnagram(s string, t string) bool {
    if len(s) != len(t) {
        return false
    }
    
    mapStr1 := make(map[rune]int)
    mapStr2 := make(map[rune]int)
    for _, v := range s {
        mapStr1[v]++
    }

     for _, v := range t {
        mapStr2[v]++
    }

    for k,v := range mapStr1 {
        v2, ok := mapStr2[k]
        if !ok ||  v2 != v {
            return false
        }
    }

    return true
}
