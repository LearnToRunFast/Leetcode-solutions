import java.util.ArrayList;
import java.util.Arrays;
import java.util.Comparator;
import java.util.List;
import java.util.PriorityQueue;

/*
 * @lc app=leetcode.cn id=37 lang=java
 *
 * [37] 解数独
 *
 * https://leetcode-cn.com/problems/sudoku-solver/description/
 *
 * algorithms
 * Hard (62.48%)
 * Likes:    526
 * Dislikes: 0
 * Total Accepted:    40.5K
 * Total Submissions: 64.6K
 * Testcase Example:  '[["5","3",".",".","7",".",".",".","."],["6",".",".","1","9","5",".",".","."],[".","9","8",".",".",".",".","6","."],["8",".",".",".","6",".",".",".","3"],["4",".",".","8",".","3",".",".","1"],["7",".",".",".","2",".",".",".","6"],[".","6",".",".",".",".","2","8","."],[".",".",".","4","1","9",".",".","5"],[".",".",".",".","8",".",".","7","9"]]'
 *
 * 编写一个程序，通过已填充的空格来解决数独问题。
 * 
 * 一个数独的解法需遵循如下规则：
 * 
 * 
 * 数字 1-9 在每一行只能出现一次。
 * 数字 1-9 在每一列只能出现一次。
 * 数字 1-9 在每一个以粗实线分隔的 3x3 宫内只能出现一次。
 * 
 * 
 * 空白格用 '.' 表示。
 * 
 * 
 * 
 * 一个数独。
 * 
 * 
 * 
 * 答案被标成红色。
 * 
 * Note:
 * 
 * 
 * 给定的数独序列只包含数字 1-9 和字符 '.' 。
 * 你可以假设给定的数独只有唯一解。
 * 给定数独永远是 9x9 形式的。
 * 
 * 
 */

// @lc code=start
// *************
// not optimaze version
// *************
// class Solution {
//     private boolean isSafe(char[][] board, int value, int x, int y, int n) {
//         // check horizontal and vertical
//         for(int i = 0; i < n; i++) {
//             if (board[i][y] == value || board[x][i] == value)  {
//                 return false;
//             }
//         }

//         // check square
//         int originX = x / 3 * 3;
//         int originY = y / 3 * 3;
//         for(int i = 0; i < 3; i++) {
//             for(int j = 0; j < 3; j++) {
//                 if (board[originX + i][originY + j] == value) {
//                     return false;
//                 }
//             }
//         }
//         return true;

//     }
//     private boolean solve(char[][] board, int n) {
//         int nextX = -1, nextY = -1;
//         for(int i = 0; i < n; i++) {
//             for(int j = 0; j < n; j++) {
//                 if (board[i][j] == '.') {
//                     nextX = i;
//                     nextY = j;
//                     break;
//                 }
//             }
//             if (nextX != -1) {
//                 break;
//             }
//         }
//         // no empty slot found.
//         if (nextX == -1) {
//             return true;
//         }

//         for(int i = 1; i <= 9; i++) {
//             char value = (char)('0' + i);
            
//             if (isSafe(board, value, nextX, nextY, n)) {
//                 board[nextX][nextY] = value;
//                 if (solve(board, n)) {
//                     return true;
//                 } else {
//                     board[nextX][nextY] = '.';
//                 }
//             }
//         }

//         return false;


//     }

//     public void solveSudoku(char[][] board) {

//         solve(board, 9);
//     }
// }


// // new version

// class Solution {

//     private void validate(char[][] board, int i , int j, boolean[] safeSquare) {
//         Arrays.fill(safeSquare, true); // fill square with true

//         for(int k = 0; k < 9; k++) {
//             char currentValue = board[i][k];
//             //validate horizontal
//             if (currentValue != '.') { // if it is not empty, set the corresponding index to false
//                 safeSquare[currentValue - '0'] = false;
//             }

//             // validate vertical
//             currentValue = board[k][j];
//             if (board[k][j] != '.') {
//                 safeSquare[board[k][j] - '0'] = false;
//             }

//             //validate square
//             int row = i / 3 * 3 + k / 3; // i / 3 + 3 is respective origin
//             int column = j / 3 * 3 + k % 3;

//             if (board[row][column] != '.') {
//                 safeSquare[board[row][column] - '0'] = false;
//             }

//         }
//     }
//     private boolean backtrack(char[][] board, int step) {
//         if (step == 81) {
//             return true;
//         }
//         int i = step / 9, j = step % 9;

//         if (board[i][j] != '.') return backtrack(board, step + 1);

//         boolean[] safeSquare = new boolean[10]; // only store 1-9
//         validate(board, i, j, safeSquare);

//         for(int k = 1; k < 10; k++) {
            
//             if (safeSquare[k]) {
//                 board[i][j] = (char) ('0' + k);
//                 if (backtrack(board, step + 1)) {
//                     return true;
//                 }
//             }
//         }
//         board[i][j] = '.'; // backtracking
//         return false;

//     }

//     public void solveSudoku(char[][] board) {

//         backtrack(board, 0);
//     }
// }



// bit set version

class Solution {

    // class newComparator implements Comparator<int[]> {
    //     private int getAvailableNum(int[] num) {
    //         int i = num[0];
    //         int j = num[1];
    //         return ~(row[i] | column[j] | block[i / 3][j / 3]) & 0x1ff;
    //     }

    //     public int compare(int[] x, int[] y) {
    //         return Integer.bitCount(getAvailableNum(x)) - Integer.bitCount(getAvailableNum(y));
    //     }
    // }

    private int[] row = new int[9];
    private int[] column = new int[9];
    private int[][] block = new int[3][3];
    // private PriorityQueue<int[]> spaces = new PriorityQueue<>(new newComparator());
    private List<int[]> spaces = new ArrayList<>();

    // private int[] getNext(char[][] board) {
    //     return spaces.poll();
    // }
    
    private int[] getNext(char[][] board) {
        int minCount = 10; // max is 9
        int index = 0;
        for(int i = 0; i < spaces.size(); i++) {
            int num = getAvailableNum(spaces.get(i));
            int count = Integer.bitCount(num);
            if (minCount > count) {
                minCount = count;
                index = i;
            }
        }
        int[] ans = spaces.get(index);
        spaces.remove(index);
        return ans;
    }

    private int getAvailableNum(int[] num) {
        int i = num[0];
        int j = num[1];
        return ~(row[i] | column[j] | block[i / 3][j / 3]) & 0x1ff;
    }

    private boolean backtrack(char[][] board, int step) {
        if (step == 0) return true;

        int[] next = getNext(board);
        int num = getAvailableNum(next);
        for(; num != 0; num &= num - 1) { // num &= num - 1 will remove the least-significant 1
            int rightMostBit = num & (-num); // extract the right most 1 bit
            int numOfZeros = Integer.numberOfTrailingZeros(rightMostBit);
            flip(next[0], next[1], numOfZeros);
            board[next[0]][next[1]] = (char)(numOfZeros + '1');
            if (backtrack(board, step - 1)) {
                return true;
            } else {
                flip(next[0], next[1], numOfZeros);
            }
        }
        spaces.add(next); // backtrack
        return false;
    }
    public void flip(int i, int j, int digit) {
        row[i] ^= (1 << digit);
        column[j] ^= (1 << digit);
        block[i / 3][j / 3] ^= (1 << digit);
    }

    public void solveSudoku(char[][] board) {

        for(int i = 0; i < board.length; i++) {
            for(int j = 0; j < board.length; j++) {
                if (board[i][j] == '.') {
                    spaces.add(new int[]{i, j});
                    continue;
                }
                int num = board[i][j] - '1';
                flip(i, j, num);
            }
        }
        int step = spaces.size();
        backtrack(board, step);
    }
}
// @lc code=end

