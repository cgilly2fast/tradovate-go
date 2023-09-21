package tradovate


type Dispatcher[T any] struct {
	ID				string
	StoreState		T
	StoreActions	[]Action
	Reducer 		func( T, Action) EventHandlerResults[T]
	Queue			ActionQueue
	Dispatching		bool
}

func (d Dispatcher[T]) getState() T {
	return d.StoreState
}

func (d Dispatcher[any]) getActions() []Action {
	return d.StoreActions
}

func (d *Dispatcher[T]) dispatch(action Action) {
	d.Queue.Add( action)
	
	if d.Dispatching {
		return
	}

	d.Dispatching = true
	for d.Queue.Len() > 0 {
		a := d.Queue.Remove()
		next := d.Reducer(d.StoreState, a)
		d.StoreState = next.state
		d.StoreActions = next.actions
	}
	d.Dispatching = false
	return
}

type ActionNode struct {
	Value 	Action
	Next	*ActionNode
}

type ActionQueue struct {
	Front 	*ActionNode
	Back 	*ActionNode
	Length 	int
}

func (q ActionQueue) Len() int {
	return q.Length
}

func (q *ActionQueue) Add(action Action) {
	q.Length = q.Length + 1
	q.Back.Next = &ActionNode{Value: action, Next: nil}
	
}

func (q ActionQueue) Remove() Action {
	if q.Length == 0 {
		panic("Queue is empty") 
	}
	q.Length = q.Length - 1
	data := q.Front.Value
	q.Front = q.Front.Next
	return data
}

