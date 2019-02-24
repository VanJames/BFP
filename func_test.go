package main

import (
	"testing"
	"fmt"
	"sort"
)



/**
	1 2 3

	4 5 6

	(1,4) (1,5) (1,6) (2,4) (2,6) (3,4) (3,5)

	1 1 1
	1 0 1
	1 1 0


 */

const TABLESIZE = 3


type Edge struct {
	from int64
	to   int64
}

func (this *Edge)GetKey() int64 {
	return this.from*TABLESIZE +this.to
}


var allKey map[int64][]int64
var allPath [][]int64

func TestCuckoo(t *testing.T){
	//hash table 1
	hashTable1 := []int64{1,2,3}
	//hash table 2
	//hashTable2 := []int64{4,5,6}

	//所有边
	edges := [][]int64{{1,4},{1,5},{1,6},{2,4},{2,6},{3,4},{3,5}}
	//allkey每个key对应的所有关系

	allKey = make(map[int64][]int64,0)
	for _,edge := range edges{
		allKey[edge[0]] = append(allKey[edge[0]],edge[1])
		allKey[edge[1]] = append(allKey[edge[1]],edge[0])
	}
	fmt.Println(allKey)
	allPath = make([][]int64,0)
	//寻找路线 寻找环
	for _,k := range hashTable1{
		findCircle([]Edge{},k)
	}
	for _,p := range allPath{
		fmt.Println("环长度：",len(p)-1,p)
	}

}

func edgeExist(ed Edge,m []Edge) bool {
	for _,v := range m{
		if (v.from == ed.from && v.to == ed.to) ||
			(v.to == ed.from && v.from == ed.to) {
			return true
		}
	}
	return false
}

func findCircle(parents []Edge,k int64) {
	l := len(parents)
	//是否已经找到 环
	if l >= 4{
		//二分图 最小环路 4条边
		start := parents[0]
		end := parents[l-1]
		if start.from == end.to{
			keys := make([]int,0)
			for k,_ := range parents{
				keys = append(keys,k)
			}
			sort.Ints(keys)
			cpath := make([]int64,0)
			for _,index := range keys {
				if index == 0{
					cpath = append(cpath,parents[index].from,parents[index].to)
				} else{
					cpath = append(cpath,parents[index].to)
				}
				//cpath = append(cpath,parents[index].from,parents[index].to)

			}
			allPath = append(allPath,cpath)
		}
	}
	//是否有线路 边要大于1 因为回路
	if _,ok := allKey[k];ok && len(allKey[k]) > 1{
		for _,v := range allKey[k]{
			edge := Edge{}
			edge.from = k
			edge.to = v
			if edgeExist(edge,parents){
				continue
			}
			newparents := make([]Edge,0)
			newparents = append(newparents,parents...)
			newparents = append(newparents,edge)
			findCircle(newparents,v)
		}
	}
	return
}
