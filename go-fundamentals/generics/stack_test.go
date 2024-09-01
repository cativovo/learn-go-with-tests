package generics

import "testing"

func TestStack(t *testing.T) {
	t.Run("integer stack", func(t *testing.T) {
		myStackOfInts := new(Stack[int])

		// check stack is empty
		AssertTrue(t, myStackOfInts.IsEmpty())

		// add a thing, then check it's not empty
		myStackOfInts.Push(123)
		AssertFalse(t, myStackOfInts.IsEmpty())

		// add another thing, pop it back again
		myStackOfInts.Push(456)
		AssertEqual(t, myStackOfInts.Length(), 2)
		value, ok := myStackOfInts.Pop()
		AssertEqual(t, value, 456)
		AssertTrue(t, ok)
		AssertEqual(t, myStackOfInts.Length(), 1)

		value, ok = myStackOfInts.Pop()
		AssertEqual(t, value, 123)
		AssertTrue(t, ok)
		AssertTrue(t, myStackOfInts.IsEmpty())
		_, ok = myStackOfInts.Pop()
		AssertFalse(t, ok)

		// can get the numbers we put in as numbers, not untyped interface{}
		myStackOfInts.Push(1)
		myStackOfInts.Push(2)
		firstNum, _ := myStackOfInts.Pop()
		secondNum, _ := myStackOfInts.Pop()
		AssertEqual(t, firstNum+secondNum, 3)
	})
}
