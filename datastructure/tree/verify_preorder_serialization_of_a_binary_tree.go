package tree

import (
	"regexp"
	"strings"
)

// 331. 验证二叉树的前序序列化
/**
 * 两个#，#消耗一个节点
 */
func isValidSerialization(preorder string) bool {
	re, _ := regexp.Compile("[0-9]+,#,#")
	for re.Match([]byte(preorder)) {
		preorder = re.ReplaceAllString(preorder, "#")
	}
	return preorder == "#"
}

/**
 * 根据节点入度和出度判断
 */
func isValidSerialization2(preorder string) bool {
	if preorder == "#" {
		return true
	}
	indegree, outdegree := 0, 0
	nodes := strings.Split(preorder, ",")

	for i := 0; i < len(nodes); i++ {
		if i == 0 { //根节点无入度
			if nodes[i] == "#" {
				return false
			}
			outdegree = 2
			continue
		}
		if nodes[i] == "#" { // nil节点一个入度
			indegree += 1
		} else { // 普通节点1个入度，2个出度
			indegree += 1
			outdegree += 2
		}
		// 如果未遍历完就出现 入度 >= 出度 说明不符合二叉树
		if i != len(nodes)-1 && indegree >= outdegree {
			return false
		}
	}
	return indegree == outdegree
}
