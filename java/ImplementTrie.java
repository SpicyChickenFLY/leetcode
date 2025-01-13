// Runtime: 79.16%, Memory Usage: 98.08%
class Trie {
    class Node {
        Node[] children;
        boolean end;
        Node() {
            children = new Node[26];
        }
    }
    
    Node root;
    public Trie() {
        root = new Node();
    }
    
    public void insert(String word) {
        Node temp = root;
        for(char letter : word.toCharArray()) {
            int index = letter - 'a';
            if(temp.children[index] == null) {
                temp.children[index] = new Node();
            }
            temp = temp.children[index];
        }
        temp.end = true;
    }
    
    protected Node getEndNode(String word){
        Node temp = root;
        for(char letter : word.toCharArray()) {
            int index = letter - 'a';
            if(temp.children[index] == null) {
                return null;
            }
            temp = temp.children[index];
        }
        return temp;
    }
    
    public boolean search(String word) {
        Node last = getEndNode(word);
        return last != null && last.end;
    }
    
    public boolean startsWith(String prefix) {
        return getEndNode(prefix) != null;
    }
}
