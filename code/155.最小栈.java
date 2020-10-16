import java.util.ArrayList;
import java.util.Stack;

/*
 * @lc app=leetcode.cn id=155 lang=java
 *
 * [155] 最小栈
 *
 * https://leetcode-cn.com/problems/min-stack/description/
 *
 * algorithms
 * Easy (55.13%)
 * Likes:    695
 * Dislikes: 0
 * Total Accepted:    169.4K
 * Total Submissions: 305.3K
 * Testcase Example:  '["MinStack","push","push","push","getMin","pop","top","getMin"]\n' +
  '[[],[-2],[0],[-3],[],[],[],[]]'
 *
 * 设计一个支持 push ，pop ，top 操作，并能在常数时间内检索到最小元素的栈。
 * 
 * 
 * push(x) —— 将元素 x 推入栈中。
 * pop() —— 删除栈顶的元素。
 * top() —— 获取栈顶元素。
 * getMin() —— 检索栈中的最小元素。
 * 
 * 
 * 
 * 
 * 示例:
 * 
 * 输入：
 * ["MinStack","push","push","push","getMin","pop","top","getMin"]
 * [[],[-2],[0],[-3],[],[],[],[]]
 * 
 * 输出：
 * [null,null,null,null,-3,null,0,-2]
 * 
 * 解释：
 * MinStack minStack = new MinStack();
 * minStack.push(-2);
 * minStack.push(0);
 * minStack.push(-3);
 * minStack.getMin();   --> 返回 -3.
 * minStack.pop();
 * minStack.top();      --> 返回 0.
 * minStack.getMin();   --> 返回 -2.
 * 
 * 
 * 
 * 
 * 提示：
 * 
 * 
 * pop、top 和 getMin 操作总是在 非空栈 上调用。
 * 
 * 
 */

// @lc code=start

// constant space version
// class MinStack {

//     Stack<Long> stack;
//     long min;
//     /** initialize your data structure here. */
//     public MinStack() {
//         stack = new Stack<Long>();
//     }
    
//     public void push(int x) {
//         if (stack.size() == 0) {
//             stack.push((long)0);
//             min = x;
//         } else {
//             long diff = x - min;
//             stack.push(diff);
//             if (diff < 0) {
//                 min = x;
//             }
//         }
//     }
    
//     public void pop() {
//         if (stack.size() > 0) {
//             if (stack.peek() < 0) {
//                 min = min - stack.peek();
//             }
//             stack.pop();
//         }
//     }
    
//     public int top() {
//         long tmp = stack.peek();
//         if (tmp < 0) {
//             return (int)min;
//         }
//         return (int)(tmp + min);
//     }
    
//     public int getMin() {
//         return (int) min;
//     }
// }
class MinStack {

    Stack<Integer> stack;
    Stack<Integer> minStack;
    int min;
    /** initialize your data structure here. */
    public MinStack() {
        stack = new Stack<Integer>();
        minStack = new Stack<Integer>();
    }
    
    public void push(int x) {
        stack.push(x);
        if (minStack.isEmpty()) {
            minStack.push(x);
            return;
        }

        int top = minStack.peek();
        if (x <= top) {
            minStack.push(x);
        }

    }
    
    public void pop() {
        int pop = stack.pop();

        int top = minStack.peek();

        if (pop == top) {
            minStack.pop();
        }
    }
    
    public int top() {
        return stack.peek();
    }
    
    public int getMin() {
        return minStack.peek();
    }

}

/**
 * Your MinStack object will be instantiated and called as such:
 * MinStack obj = new MinStack();
 * obj.push(x);
 * obj.pop();
 * int param_3 = obj.top();
 * int param_4 = obj.getMin();
 */
// @lc code=end

