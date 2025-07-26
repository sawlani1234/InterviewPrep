package tree

type TrieNode struct {
    chars [26]*TrieNode
    isEnd bool 
}


func NewTrieNode() *TrieNode {
    return &TrieNode{}
}

func Add(root *TrieNode,word string) {
    
    for i := 0 ;i<len(word); i++ { 
         if root.chars[word[i]-'a'] == nil {
             root.chars[word[i]-'a'] = NewTrieNode()
         }
         root = root.chars[word[i]-'a']
    }
    root.isEnd = true 
}


func Search(root *TrieNode,word string) bool {
     
     for i := 0;i < len(word) && root != nil ; i++ {
         root = root.chars[word[i]-'a']
     }
     
     return root != nil && root.isEnd 
}