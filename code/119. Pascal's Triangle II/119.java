class Solution {
    public List<Integer> getRow(int rowIndex) {
        rowIndex++;
        List<Integer> cur = new ArrayList<>();
        
        int pre = 1;
        cur.add(1);
        
        for (int i = 1; i < rowIndex; i++) {
            for (int j = i - 1; j > 0; j--) {
                cur.set(j, cur.get(j - 1) + cur.get(j));
            }
            cur.add(1);
        }
        return cur;
    }
}
class Solution {
    public List<Integer> getRow(int rowIndex) {
        List<Integer> ans = new ArrayList<>();
        
        int N = rowIndex;
        long pre = 1;
        
        ans.add(1);
        
        for (long k = 1; k <= N; k++) {
            long curr = pre * (N - k + 1) / k;
            ans.add((int)curr);
            pre = curr;
        }
        return ans;
    }
}
