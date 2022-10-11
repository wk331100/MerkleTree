# MerkleTree

一款使用Golang实现的默克尔树

```
func TestMerkleTree(t *testing.T) {
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
