// RunTime: 88.64%, MemoryUsage: 29.71%
class MinStack {

    class LinkNode {
        private int val;
        private int min;
        private LinkNode next;

        public LinkNode(int val, LinkNode next) {
            this.val = val;
            this.min = val;
            this.next = next;
        }
    }

    private LinkNode head;

    /** initialize your data structure here. */
    public MinStack() {
        head = null;
    }

    public void push(int x) {
        head = new LinkNode(x, head);
        if (head.next != null && x >= head.next.min)
            head.min = head.next.min;
    }

    public void pop() {
        head = head.next; 
    }

    public int top() {
        return head.val;
    }

    public int getMin() {
        return head.min;
    }
}



public class UseMinStack {
    public static void main(String[] args) {
        MinStack minStack = new MinStack();
        minStack.push(-2);
        minStack.push(0);
        minStack.push(-3);
        System.out.println(minStack.getMin());
        minStack.pop();
        System.out.println(minStack.top());
        System.out.println(minStack.getMin());
    }
}