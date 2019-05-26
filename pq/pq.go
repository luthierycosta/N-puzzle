package pq

type EmptyStackError struct {
	msg string
}

func (e *EmptyStackError) Error() string {
    return e.msg;
}

type Item struct {
	value float64;
}

func (this *Item) compare(other *Item) bool {
	return (this.value <= other.value);
}

type PriorityQueue struct { 
	items []Item;
	compare func(Item, Item) bool; // a <= b? true : false;
}

func (pq *PriorityQueue) len() int {
    return len((*pq).items);
}

func (pq *PriorityQueue) parent(i int) int {
	return ((i - 1) / 2);
}

func (pq *PriorityQueue) left(i int) int {
	return (2 * i + 1);
}

func (pq *PriorityQueue) right(i int) int {
	return (2 * i + 2);
}

func (pq *PriorityQueue) push(items ...Item) {
    for _, obj := range items {
		pq.items = append(pq.items, obj);
		var i int = pq.len() - 1;

		for i != 0 && !(pq.compare(pq.items[pq.parent(i)], pq.items[i])) {
			pq.items[pq.parent(i)], pq.items[i] = pq.items[i], pq.items[pq.parent(i)];
			i = pq.parent(i);
		}
    }
}

func (pq *PriorityQueue) top() (result Item, errorHandle error) {
	errorHandle = nil;
	if pq.len() == 0 {
		errorHandle = &EmptyStackError{"Pilha vázia"};
		return
	}
	result = pq.items[0];
	return;
}

func (pq *PriorityQueue) pop() (errorHandle error) {
	errorHandle = nil;
	if pq.len() == 0 {
		errorHandle = &EmptyStackError{"Pilha vázia"};
		return
	}
	pq.items[0], pq.items[pq.len() - 1] = pq.items[pq.len() - 1], pq.items[0];
	pq.items = pq.items[:(pq.len() - 1)];
	pq.fixHeapQueue(0);
	return;
}

func (pq *PriorityQueue) fixHeapQueue(i int) {
	var l int = pq.left(i);
	var r int = pq.right(i);
	var small int = i;
	if (l < pq.len()) && pq.compare(pq.items[l], pq.items[i]) {
		small = l;		
	}
	if (r < pq.len()) && pq.compare(pq.items[r], pq.items[small]) {
		small = r;
	}
	if (small != i) {
		pq.items[i], pq.items[small] = pq.items[small], pq.items[i];
		pq.fixHeapQueue(small);
	}
}