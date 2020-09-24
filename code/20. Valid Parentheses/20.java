class Solution {
    private static char getSymbol(char c) {
        switch(c) {
            case ')':
                return '(';
            case '}':
                return '{';
            case ']':
                return '[';
        }
        return '\0';
    }
    
    private static boolean isPrefix(char c) {
        String temp = "({[";
        for (int i = 0; i < 3; i++) {
            if (c == temp.charAt(i)) {
                return true;
            }
        }
        return false;
    }
    public boolean isValid(String s) {
        Stack<Character> stack = new Stack<>();
        
        for (int i = 0; i < s.length(); i++) {
            if (isPrefix(s.charAt(i))) {
                stack.push(s.charAt(i));
            } else {
                if (stack.empty()) return false;
                char c = stack.pop();
                
                if (getSymbol(s.charAt(i)) != c) {
                    return false;
                }
            }
        }
        
        return stack.empty() ? true : false;
    }
}