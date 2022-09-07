#include <iostream>
#include <deque>
#include <map>
using namespace std;

// RunTime: 5.06%, MemoryUsage: 91.46%
class LRUCache
{
public:
    int capacity;
    deque<int> keyQueue;
    map<int, int> kvMap;

    LRUCache(int capacity)
    {
        this->capacity = capacity;
    }

    int get(int key)
    {
        int result = -1;
        for (deque<int>::iterator it = keyQueue.begin(); it != keyQueue.end(); it++)
        {
            if (*it == key)
            {
                result = kvMap.at(key);
                keyQueue.erase(it);
                keyQueue.push_back(key);
                break;   
            }
        }
        return result;
    }

    void put(int key, int value)
    {
        
        for (deque<int>::iterator it = keyQueue.begin(); it != keyQueue.end(); it++)
            if (*it == key)
            {
                kvMap.at(key) = value;
                keyQueue.erase(it);
                keyQueue.push_back(key);
                return;
            }
        if (capacity == keyQueue.size())
        {
            int deprecatedKey = keyQueue.at(0);
            keyQueue.pop_front();
            kvMap.erase(deprecatedKey);
        }
        keyQueue.push_back(key);
        kvMap[key] = value;
    }
};

/**
 * Your LRUCache object will be instantiated and called as such:
 * LRUCache* obj = new LRUCache(capacity);
 * int param_1 = obj->get(key);
 * obj->put(key,value);
 */

int main()
{
    try
    {
        LRUCache cache(10);
        cache.put(10, 13);
        cache.put(3, 17);
        cache.put(6, 11);
        cache.put(10, 5);
        cache.put(9, 10);
        cout << "\nget result: "
             << cache.get(13);
        cache.put(2, 19);
        cout << "\nget result: "
             << cache.get(2);
        cout << "\nget result: "
             << cache.get(3);
        cache.put(5, 25);
        cout << "\nget result: "
             << cache.get(8);
        cache.put(9, 22);
        cache.put(5, 5);
        cache.put(1, 30);
        cout << "\nget result: "
             << cache.get(11);
        cache.put(9, 12);
        cout << "\nget result: "
             << cache.get(7);
        cout << "\nget result: "
             << cache.get(5);
        cout << "\nget result: "
             << cache.get(8);
        cout << "\nget result: "
             << cache.get(9);
        cache.put(4, 30);
        cache.put(9, 3);
        cout << "\nget result: "
             << cache.get(9);
        cout << "\nget result: "
             << cache.get(10);
        cout << "\nget result: "
             << cache.get(10);
        cache.put(6, 14);
        cache.put(3, 1);
        cout << "\nget result: "
             << cache.get(3);
        cache.put(10, 11);
        cout << "\nget result: "
             << cache.get(8);
        cache.put(2, 14);
        cout << "\nget result: "
             << cache.get(1);
        cout << "\nget result: "
             << cache.get(5);
        cout << "\nget result: "
             << cache.get(4);
        cache.put(11, 4);
        cache.put(12, 24);
        cache.put(5, 18);
        cout << "\nget result: "
             << cache.get(13);
        cache.put(7, 23);
        cout << "\nget result: "
             << cache.get(8);
        cout << "\nget result: "
             << cache.get(12);
        cache.put(3, 27);
        cache.put(2, 12);
        cout << "\nget result: "
             << cache.get(5);
        cache.put(2, 9);
        cache.put(13, 4);
        cache.put(8, 18);
        cache.put(1, 7);
        cout << "\nget result: "
             << cache.get(6);
        cache.put(9, 29);
        cache.put(8, 21);
        cout << "\nget result: "
             << cache.get(5);
        cache.put(6, 30);
        cache.put(1, 12);
        cout << "\nget result: "
             << cache.get(10);
        cache.put(4, 15);
        cache.put(7, 22);
        cache.put(11, 26);
        cache.put(8, 17);
        cache.put(9, 29);
        cout << "\nget result: "
             << cache.get(5);
        cache.put(3, 4);
        cache.put(11, 30);
        cout << "\nget result: "
             << cache.get(12);
        cache.put(4, 29);
        cout << "\nget result: "
             << cache.get(3);
        cout << "\nget result: "
             << cache.get(9);
        cout << "\nget result: "
             << cache.get(6);
        cache.put(3, 4);
        cout << "\nget result: "
             << cache.get(1);
        cout << "\nget result: "
             << cache.get(10);
        cache.put(3, 29);
        cache.put(10, 28);
        cache.put(1, 20);
        cache.put(11, 13);
        cout << "\nget result: "
             << cache.get(3);
        cache.put(3, 12);
        cache.put(3, 8);
        cache.put(10, 9);
        cache.put(3, 26);
        cout << "\nget result: "
             << cache.get(8);
        cout << "\nget result: "
             << cache.get(7);
        cout << "\nget result: "
             << cache.get(5);
        cache.put(13, 17);
        cache.put(2, 27);
        cache.put(11, 15);
        cout << "\nget result: "
             << cache.get(12);
        cache.put(9, 19);
        cache.put(2, 15);
        cache.put(3, 16);
        cout << "\nget result: "
             << cache.get(1);
        cache.put(12, 17);
        cache.put(9, 1);
        cache.put(6, 19);
        cout << "\nget result: "
             << cache.get(4);
        cout << "\nget result: "
             << cache.get(5);
        cout << "\nget result: "
             << cache.get(5);
        cache.put(8, 1);
        cache.put(11, 7);
        cache.put(5, 2);
        cache.put(9, 28);
        cout << "\nget result: "
             << cache.get(1);
        cache.put(2, 2);
        cache.put(7, 4);
        cache.put(4, 22);
        cache.put(7, 24);
        cache.put(9, 26);
        cache.put(13, 28);
        cache.put(11, 26);
    }
    catch(const exception& e)
    {
        cout << e.what() << "\n";
    }


    return 0;
}
