package merkleTree

import (
	"bytes"
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"errors"
	"hash"
	"math"
)

const (
	errEmptyData = "empty data"
)

// MerkleTree 默克尔树
type MerkleTree struct {
	Root        *Node
	Leaves      []*Node
	hashHandler func() hash.Hash
}

// Node 节点
type Node struct {
	Parent *Node
	Left   *Node
	Right  *Node
	leaf   bool
	single bool
	Hash   []byte
	Data   []byte
}

// GetRootHash 获取默克尔树根Hash
func (m *MerkleTree) GetRootHash() []byte {
	return m.Root.Hash
}

// NewMerkleTree 创建一个新的 默克尔树
func NewMerkleTree(hashType string, data [][]byte) (*MerkleTree, error) {
	mk := &MerkleTree{}
	mk.hashHandler = mk.buildHash(hashType)

	root, leaves, err := mk.buildMerkleTreeRoot(data)
	if err != nil {
		return nil, err
	}
	mk.Root = root
	mk.Leaves = leaves
	return mk, nil
}

// VerifyData 验证数据是否在默克尔树中
func (m *MerkleTree) VerifyData(data []byte) (bool, error) {
	dataHash := m.calHash(data)
	for _, leaf := range m.Leaves {
		if bytes.Compare(dataHash, leaf.Hash) == 0 {
			return true, nil
		}
	}
	return false, nil
}

// VerifyTree 验证默克尔树Hash
func (m *MerkleTree) VerifyTree() (bool, error) {
	calRoot, err := m.Root.verifyNode(m)
	if err != nil {
		return false, err
	}

	if bytes.Compare(calRoot, m.Root.Hash) == 0 {
		return true, nil
	}
	return false, err
}

// verifyNode 重新计算节点Hash
func (n *Node) verifyNode(mk *MerkleTree) ([]byte, error) {
	if n.leaf {
		return mk.calHash(n.Data), nil
	}
	leftNodeHash, err := n.Left.verifyNode(mk)
	if err != nil {
		return nil, err
	}
	rightNodeHash, err := n.Right.verifyNode(mk)
	if err != nil {
		return nil, err
	}

	return mk.calHash(append(leftNodeHash, rightNodeHash...)), nil
}

// buildMerkleTree 构建 默克尔树 节点
func (m *MerkleTree) buildMerkleTreeRoot(data [][]byte) (*Node, []*Node, error) {
	if len(data) <= 0 {
		return nil, nil, errors.New(errEmptyData)
	}
	leaf := m.buildMerkleTreeLeaf(data)
	root, err := m.buildMerkleTreeNode(leaf)
	return root, leaf, err
}

// buildMerkleTreeNode 构建 默克尔树 中间节点
func (m *MerkleTree) buildMerkleTreeNode(nodes []*Node) (*Node, error) {
	length := int(math.Ceil(float64(len(nodes)) / 2))

	var nodeSlice []*Node
	var single bool
	for i := 0; i < length; i++ {
		leftNode := nodes[i*2]
		var rightNode = new(Node)
		if i*2+1 < len(nodes) {
			rightNode = nodes[i*2+1]
		} else {
			single = true
		}

		node := &Node{
			Parent: nil,
			Left:   leftNode,
			Right:  rightNode,
			leaf:   false,
			single: single,
			Hash:   nil,
			Data:   nil,
		}
		// 将两个子节点Hash拼接后，计算自身Hash
		if !single {
			leftNode.Hash = append(leftNode.Hash, rightNode.Hash...)
		}

		node.Hash = m.calHash(leftNode.Hash)

		// 将当前节点设置为 两个子节点的父节点
		nodes[i*2].Parent = node
		nodes[i*2+1].Parent = node

		nodeSlice = append(nodeSlice, node)
	}

	if len(nodeSlice) > 1 {
		return m.buildMerkleTreeNode(nodeSlice)
	}
	return nodeSlice[0], nil
}

// buildMerkleTreeLeaf 构建 默克尔树 叶节点
func (m *MerkleTree) buildMerkleTreeLeaf(data [][]byte) []*Node {
	var leaf []*Node
	for _, item := range data {
		node := &Node{
			Parent: nil,
			Left:   nil,
			Right:  nil,
			leaf:   true,
			single: false,
			Hash:   m.calHash(item),
			Data:   item,
		}
		leaf = append(leaf, node)
	}

	return leaf
}

// buildHash 根据Hash类型，构建Hash
func (m *MerkleTree) buildHash(hashType string) func() hash.Hash {
	switch hashType {
	case "md5":
		return md5.New
	case "sha1":
		return sha1.New
	case "sha256":
		return sha256.New
	case "sha512":
		return sha512.New
	default:
		return sha1.New
	}
}

func (m *MerkleTree) calHash(data []byte) []byte {
	hashHandler := m.hashHandler()
	hashHandler.Write(data)
	return hashHandler.Sum(nil)
}
