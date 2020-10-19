/*
 * @lc app=leetcode.cn id=208 lang=java
 *
 * [208] 实现 Trie (前缀树)
 *
 * https://leetcode-cn.com/problems/implement-trie-prefix-tree/description/
 *
 * algorithms
 * Medium (68.55%)
 * Likes:    436
 * Dislikes: 0
 * Total Accepted:    58.8K
 * Total Submissions: 85.1K
 * Testcase Example:  '["Trie","insert","search","search","startsWith","insert","search"]\n' +
  '[[],["apple"],["apple"],["app"],["app"],["app"],["app"]]'
 *
 * 实现一个 Trie (前缀树)，包含 insert, search, 和 startsWith 这三个操作。
 * 
 * 示例:
 * 
 * Trie trie = new Trie();
 * 
 * trie.insert("apple");
 * trie.search("apple");   // 返回 true
 * trie.search("app");     // 返回 false
 * trie.startsWith("app"); // 返回 true
 * trie.insert("app");   
 * trie.search("app");     // 返回 true
 * 
 * 说明:
 * 
 * 
 * 你可以假设所有的输入都是由小写字母 a-z 构成的。
 * 保证所有输入均为非空字符串。
 * 
 * 
 */

// @lc code=start
class Trie {
    private final int R = 26; // 26 letter
    private Trie next[];
    private boolean _isEnd;
    private Trie root;

    /** Initialize your data structure here. */
    public Trie() {
        root = this;
        next = new Trie[26];    
    }
    private int letterToDigit(char c) {
        return c - 'a';
    }

    /** Inserts a word into the trie. */
    public void insert(String word) {
        Trie tmp = root;
        for(char c : word.toCharArray()) {
            int idx = letterToDigit(c);
            if (tmp.next[idx] == null) {
                tmp.next[idx] = new Trie();
            }
            tmp = tmp.next[idx];
        }
        tmp._isEnd = true;
    }
    private boolean isEnd() {
        return _isEnd;
    }

    private Trie getLastTrie(String word) {
        Trie tmp = root;
        for(char c : word.toCharArray()) {
            Trie nextTrie = tmp.next[letterToDigit(c)];
            if (nextTrie == null) {
                return null;
            }
            tmp = nextTrie;
        }
        return tmp;
    }

    /** Returns if the word is in the trie. */
    public boolean search(String word) {
        Trie lastTrie = getLastTrie(word);
        return lastTrie != null && lastTrie.isEnd();
    }
    
    /** Returns if there is any word in the trie that starts with the given prefix. */
    public boolean startsWith(String prefix) {
        return getLastTrie(prefix) != null;
    }
}

/**
 * Your Trie object will be instantiated and called as such:
 * Trie obj = new Trie();
 * obj.insert(word);
 * boolean param_2 = obj.search(word);
 * boolean param_3 = obj.startsWith(prefix);
 */
// @lc code=end

