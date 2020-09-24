class MinStack {
    private Stack<Integer> data;
    private Stack<Integer> helper;

    /** initialize your data structure here. */
    public MinStack() {
        data = new Stack();
        helper = new Stack();
    }
    
    public void push(int x) {
        data.add(x);
        if (helper.isEmpty() || helper.peek() >= x) {
            helper.add(x);
        } else {
            helper.add(helper.peek());
        }
    }
    
    public void pop() {
        if (data.isEmpty()) {
            throw new RuntimeException("The stack is empty");
        }
        data.pop();
        helper.pop();
    }
    
    public int top() {
        if (data.isEmpty()) {
            throw new RuntimeException("The stack is empty");
        }
        return data.peek();
    }
    
    public int getMin() {
        if (helper.isEmpty()) {
            throw new RuntimeException("The stack is empty");
        }
        return helper.peek();
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