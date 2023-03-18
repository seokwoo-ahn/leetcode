type BrowserHistory struct {
    history []string
    cur int
}


func Constructor(homepage string) BrowserHistory {
    return BrowserHistory{
        history: []string{homepage},
        cur: 0,
    }
}


func (this *BrowserHistory) Visit(url string)  {
    this.cur++
    this.history = this.history[:this.cur]
    this.history = append(this.history, url)
}


func (this *BrowserHistory) Back(steps int) string {
    temp := this.cur - steps
    if temp < 0 {
        this.cur = 0
        return this.history[this.cur]
    } else {
        this.cur = this.cur - steps
        return this.history[this.cur]
    }
}


func (this *BrowserHistory) Forward(steps int) string {
    temp := this.cur + steps
    if temp >= len(this.history) {
        this.cur = len(this.history) - 1
        return this.history[this.cur]
    } else {
        this.cur = this.cur + steps
        return this.history[this.cur]
    }
}


/**
 * Your BrowserHistory object will be instantiated and called as such:
 * obj := Constructor(homepage);
 * obj.Visit(url);
 * param_2 := obj.Back(steps);
 * param_3 := obj.Forward(steps);
 */