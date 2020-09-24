/**
 * @param {string} s
 * @return {number}
 */
var lengthOfLongestSubstring = function(s) {
    let temp = {};
    let i = 0, ans = 0;
    for(let j=0; j < s.length; j++) {
        if (temp[s[j]] !== undefined) {
            i = temp[s[j]] > i ?  temp[s[j]] : i;    
        }
        ans = ans > (j - i + 1) ? ans : (j - i + 1) ;
        temp[s[j]] = j + 1;
    }
    return ans;
};