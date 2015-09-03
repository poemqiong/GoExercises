package main

func main() {
    a := [...]float64{1,2,3,4,5}
    s := a[:]
    avg := Average(s)
    println(avg)
}

func Average(s []float64) (avg float64) {
    l := len(s)
    for i := 0; i < l; i++ {
        avg = avg + s[i]
    }
    if l > 0 {
        avg = avg / float64(l)
    }
    return
} 
