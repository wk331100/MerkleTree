# MerkleTree

一款简单的基于Golang实现的默克尔树

## 作用
可以通过默克尔树根Hash或者节点Hash，快速验证数据的一致性。

## 使用方式
```
go get github.com/wk331100/MerkleTree
```

## 外部调用方法
- NewMerkleTree()	// 新建一个默克尔树
- mk.VerifyTree		// 验证一个默克尔树是否完整
- mk.VerifyData		// 验证数据是否在树中
- mk.GetRootHash	// 获取默克尔树根Hash

## 测试示例
```
func TestMerkleTree(t *testing.T) {
  // data为[][]byte类型数据
	data := initData()
  // 实例化一个默克尔树
	mk, err := NewMerkleTree("md5", data)
	require.Nil(t, err)
	fmt.Sprintf("%x", mk.GetRootHash())
	printHash(mk)

  // 验证数据是否在默克尔树中
	res, _ := mk.VerifyData(testData)
	require.True(t, res)

  // 验证默克尔树是否完整
	res2, _ := mk.VerifyTree()
	require.True(t, res2)
}
```
