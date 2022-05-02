package state

// Db - Common K-V Data store
type Db interface {
	NextId() uint64
	Save(id uint64, state []byte)
	Find(id uint64) []byte
	Delete(id uint64)
}

// DbLinked exposes Extra Functionality For Linked Nodes
// It maintains an index of Parent-children relationships of nodes, as well as a counter of links to the nodes
// The lifecycle of a node is controlled by 2 things: children and link flag
//	- children show that this data is used in a chain
//	- the link flag shows that this current data is used from the outside (somebody carries an id of this node)
type DbLinked interface {
	NextId() uint64
	FindNode(id uint64) (parent uint64, state []byte) // find state by known id (that is used for resolving the state by id)
	SaveNode(id, parent uint64, state []byte)         // create a new node and mark it as linked from the outside (the owner has created it)
	Unlink(id uint64)                                 // clear the link from the outside, (owner does not need it anymore, but children could still use it), should remove if no children
}
