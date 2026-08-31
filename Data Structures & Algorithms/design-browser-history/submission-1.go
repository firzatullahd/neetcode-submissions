type BrowserHistory struct {
	homepage *History // head
	latest *History // tail
	currentPage *History
}

type History struct {
	url string
	next *History
	prev *History
}


func Constructor(homepage string) BrowserHistory {
	newHistory := &History{
		url: homepage,
	}
    return BrowserHistory{
		homepage: newHistory,
		currentPage: newHistory,
		latest: newHistory,
	}
}


func (this *BrowserHistory) Visit(url string)  {
	newHistory := &History{
		url: url,
		prev: this.currentPage,
	}
	this.currentPage.next = newHistory
	this.currentPage = newHistory
	this.latest = this.currentPage
}


func (this *BrowserHistory) Back(steps int) string {
	current := this.currentPage
	i := 0
	for current.prev != nil {
		if i == steps {
			this.currentPage = current
			return current.url
		}
		i++
		current = current.prev
	}
	
	this.currentPage = current
    return current.url
}


func (this *BrowserHistory) Forward(steps int) string {
	current := this.currentPage
	i := 0
	for current.next != nil {
		if i == steps {
			this.currentPage = current
			return current.url
		}
		i++
		current = current.next
	}
	
	this.currentPage = current	
    return current.url
}


/**
 * Your BrowserHistory object will be instantiated and called as such:
 * obj := Constructor(homepage);
 * obj.Visit(url);
 * param_2 := obj.Back(steps);
 * param_3 := obj.Forward(steps);
 */