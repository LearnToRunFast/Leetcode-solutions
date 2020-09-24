
class Solution {
public:
    string convert(string s, int numRows) {
        if (numRows <= 1) {
            return s;
        }
        
        vector<string> rows(min(int(s.size()), numRows));
        int currRow = 0;
        bool down = false;
        
        for(char c : s) {
            rows[currRow] += c;
            
            //toggle down flag
            if (currRow == 0 || currRow == numRows - 1) {
                down = !down;
            }
            currRow += down ? 1 : -1;
        }
        
        string ans;
        for (string row: rows) {
            ans += row;
        }
        return ans;
    }
};