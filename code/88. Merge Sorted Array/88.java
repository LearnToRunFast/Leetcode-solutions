public List<String> letterCombinations(String digits) {
    LinkedList<String> ans = new LinkedList<String>();
    if(digits.isEmpty()) return ans;
    String[] mapping = new String[] {"0", "1", "abc", "def", "ghi", "jkl", "mno", "pqrs", "tuv", "wxyz"};
    ans.add("");
    for(int i =0; i<digits.length();i++){
        int x = Character.getNumericValue(digits.charAt(i));
        while(ans.peek().length()==i){
            String t = ans.remove();
            for(char s : mapping[x].toCharArray())
                ans.add(t+s);
        }
    }
    return ans;
}

作者：xxxzzz-2
链接：https://leetcode-cn.com/problems/letter-combinations-of-a-phone-number/solution/javafei-chang-jing-qiao-de-dui-lie-jie-fa-by-xxxzz/
来源：力扣（LeetCode）
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。