import java.util.*;

class MyStack {
    private Queue<Integer> q;

    /** Initialize your data structure here. */
    public MyStack() {
        q = new LinkedList<>();
    }

    /** Push element x onto stack. */
    public void push(int x) {
        q.add(x);
    }

    /** Removes the element on top of the stack and returns that element. */
    public int pop() {
        Queue<Integer> reverse = new LinkedList<>();
        int temp = q.poll();
        while (!q.isEmpty()) {
            reverse.add(temp);
            temp = q.poll();
        }
        q = reverse;
        return temp;
    }

    /** Get the top element. */
    public int top() {
        Queue<Integer> reverse = new LinkedList<>();
        int temp = 0;
        while (!q.isEmpty()) {
            temp = q.poll();
            reverse.add(temp);
        }
        q = reverse;
        return temp;
    }

    /** Returns whether the stack is empty. */
    public boolean empty() {
        return q.isEmpty();
    }
}

public class ImplementStackUsingQueues {
    public static void main(String[] args) {
        MyStack obj = new MyStack(); 
        obj.push(3); 
        int param_2 = obj.pop(); 
        int param_3 = obj.top();
        boolean param_4 = obj.empty();
    }
}
