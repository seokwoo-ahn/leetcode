type WordDictionary struct {
    letter [26]*WordDictionary
    end bool
}


func Constructor() WordDictionary {
    return WordDictionary{}
}


func (this *WordDictionary) AddWord(word string)  {
    cur := this
    for i := 0; i < len(word); i++ {
        if cur.letter[word[i]- 'a'] == nil {
             temp := Constructor()
             cur.letter[word[i]- 'a'] = &temp   
        }
        cur = cur.letter[word[i] - 'a']
    }
    cur.end = true
}


func (this *WordDictionary) Search(word string) bool {
    cur := this
    for i := 0; i < len(word); i++ {
        if word[i] == '.' {
            for j := 0; j < 26; j++ {
                if cur.letter[j] != nil {
                    if cur.letter[j].Search(word[i+1:]) {
                        return true
                    }
                }
            }
            return false
        }
        
        if cur.letter[word[i]-'a'] == nil {
            return false
        }
        cur = cur.letter[word[i] - 'a']
    }
    
    return cur.end
}


/**
 * Your WordDictionary object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddWord(word);
 * param_2 := obj.Search(word);
 */