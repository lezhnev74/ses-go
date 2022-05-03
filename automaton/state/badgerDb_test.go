package state

import (
	"reflect"
	"testing"
)

func Test_MissingNodes(t *testing.T) {
	b := MakeBadgerDb("scope", "")
	// Normal find
	d := b.Find(1)
	if d != nil {
		t.Errorf("Returned not nil")
	}
	// Tree mode
	_, n := b.FindNode(1)
	if n != nil {
		t.Errorf("Returned not nil")
	}
}

func Test_Scoping(t *testing.T) {
	t.Run("scoping", func(t *testing.T) {
		b1, b2 := MakeBadgerDb("scope1", ""), MakeBadgerDb("scope2", "")

		// Test Read/Write Scoping
		key := uint64(1)
		data1, data2 := []byte("A"), []byte("B")
		b1.Save(key, data1)
		b2.Save(key, data2)
		recoveredData1, recoveredData2 := b1.Find(key), b2.Find(key)

		if !reflect.DeepEqual(recoveredData1, data1) || !reflect.DeepEqual(recoveredData2, data2) {
			t.Errorf("read/write scoping does not work")
		}

		// Test Sequence Scoping
		if b1.NextId() != b2.NextId() {
			t.Errorf("sequence scoping does not work")
		}
	})
}

func Test_Tree(t *testing.T) {
	db := MakeBadgerDb("scope1", "")

	// Node1<---Node2<---Node3
	db.SaveNode(1, 0, []byte("Node1"))
	db.SaveNode(2, 1, []byte("Node2"))
	db.SaveNode(3, 2, []byte("Node3"))

	// Find Node2
	n2_parent, n2_state := db.FindNode(2)
	if n2_parent != 1 || !reflect.DeepEqual(n2_state, []byte("Node2")) {
		t.Errorf("Find tree node does not work")
	}

	// Test unlink
	db.Unlink(3) // should remove node3

	// Now unlink a node with children
	db.Unlink(1) // <-- should not remove node1 as it has children
	if _, state := db.FindNode(1); state == nil {
		t.Errorf("node1 should not have been deleted")
	}

	// Now Unlink node2 and that should remove node1 and then node 1
	db.Unlink(2)
	if _, state := db.FindNode(2); state != nil {
		t.Errorf("node2 should have been deleted")
	}
	if _, state := db.FindNode(1); state != nil {
		t.Errorf("node1 should have been deleted")
	}
}
