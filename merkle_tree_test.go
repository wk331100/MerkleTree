package merkleTree

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

var (
	testData []byte
	testHash []byte
)

func TestMerkleTree(t *testing.T) {
	data := initData()
	mk, err := NewMerkleTree("md5", data)
	require.Nil(t, err)
	fmt.Sprintf("%x", mk.GetRootHash())
	printHash(mk)

	res, _ := mk.VerifyData(testData)
	require.True(t, res)

	res2, _ := mk.VerifyTree()
	require.True(t, res2)
}

func printHash(mk *MerkleTree) {
	if mk.Root.leaf {

		return
	}
	cyclePrint(mk.Root)
}

func cyclePrint(node *Node) {
	if node.leaf {
		fmt.Println(fmt.Sprintf("Leaf: hash[%x], data[%s], leaf[%v]\n", node.Hash, node.Data, node.leaf))
	} else {
		fmt.Println(fmt.Sprintf("Node: hash[%x], data[%s], leaf[%v]\n", node.Hash, node.Data, node.leaf))
	}

	if node.Left != nil {
		cyclePrint(node.Left)
	}
	if node.Right != nil {
		cyclePrint(node.Right)
	}
}

func TestBuildMerkleTreeLeaf(t *testing.T) {
	m := &MerkleTree{}
	data := initData()
	m.hashHandler = m.buildHash("sha256")
	leaf := m.buildMerkleTreeLeaf(data)
	fmt.Println(leaf)
}

func initData() [][]byte {
	var data [][]byte

	for i := 0; i < 4; i++ {
		str := fmt.Sprintf("test data %d", i)
		bz, _ := json.Marshal(str)
		if i == 3 {
			testData = bz
		}
		data = append(data, bz)
	}
	return data
}
