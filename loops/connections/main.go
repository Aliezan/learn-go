package main

func countConnections(groupSize int) int {
	var totalConn int

	for i := 0; i < groupSize; i++ {
		totalConn += i
	}

	return totalConn

}
