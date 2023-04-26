type SmallestInfiniteSet struct {
    m map[int]bool
    smallest int
}


func Constructor() SmallestInfiniteSet {
    return SmallestInfiniteSet{
        m: make(map[int]bool),
        smallest: 1,
    }
}


func (this *SmallestInfiniteSet) PopSmallest() int {
    res := this.smallest
    tmp := this.smallest
    this.m[tmp] = true
    
    for {
        tmp++
        if !this.m[tmp] {
            this.smallest = tmp
            break
        }
    }
    return res
}


func (this *SmallestInfiniteSet) AddBack(num int)  {
    if this.smallest > num {
        this.smallest = num
    }
    this.m[num] = false
}


/**
 * Your SmallestInfiniteSet object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.PopSmallest();
 * obj.AddBack(num);
 */

