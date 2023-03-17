type Trie struct {
    word [26]*Trie
    isEnd bool
}


func Constructor() Trie {
    return Trie{}
}


func (this *Trie) Insert(word string)  {
    cur := this
    for i:= 0; i < len(word); i++{
        if cur.word[word[i]-'a'] != nil {
            cur = cur.word[word[i]-'a']
        } else {
            newTrie := Constructor()
            cur.word[word[i]-'a'] = &newTrie
            cur = cur.word[word[i]-'a']
        }
    }
    cur.isEnd = true
}


func (this *Trie) Search(word string) bool {
    cur := this
    for i:=0; i < len(word); i++ {
        if cur.word[word[i]-'a'] == nil {
            return false
        }
        cur = cur.word[word[i]-'a']
    }
    return cur.isEnd
}


func (this *Trie) StartsWith(prefix string) bool {
    cur := this
    for i:=0; i < len(prefix); i++ {
        if cur.word[prefix[i]-'a'] == nil {
            return false
        }
        cur = cur.word[prefix[i]-'a']
    }
    return true
}


/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */