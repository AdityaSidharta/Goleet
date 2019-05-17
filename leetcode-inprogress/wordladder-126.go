package leetcode_inprogress

import (
	"fmt"
	"math/rand"
)

func shuffle(values []string) []string {
	dest := make([]string, len(values))
	perm := rand.Perm(len(values))
	for i, v := range perm {
		dest[v] = values[i]
	}
	return dest
}

func transWord(word1 string, word2 string) bool {
	nword1 := len(word1)
	nword2 := len(word2)
	if nword1 != nword2 {
		panic("length of word1 is different from word2")
	}
	nsame := 0
	for idx:=0; idx <nword1; idx++{
		if word1[idx] == word2[idx]{
			nsame = nsame + 1
		}
	}
	return nsame == nword1 - 1
}

func list2Maps(wordList []string) map[string][]string {
	maps := make(map[string][]string)
	for _, value := range wordList{
		maps[value] = make([]string, 0)
	}
	for i:=0; i < len(wordList); i++{
		for j :=i; j < len(wordList); j++ {
			if i != j{
				if transWord(wordList[i], wordList[j]){
					maps[wordList[i]] = append(maps[wordList[i]], wordList[j])
					maps[wordList[j]] = append(maps[wordList[j]], wordList[i])
				}
			}
		}
	}
	for idx :=0; idx < len(maps); idx ++ {
		maps[wordList[idx]] = shuffle(maps[wordList[idx]])
	}
	return maps
}


// TODO: Finish this function. This function fails to obtain all paths, it only returns one path
func dfs(curWord string, endWord string, maps map[string][]string, visited map[string]bool, path []string, paths [][]string) [][]string {
	visited[curWord] = true
	newpath := append(path, curWord)
	if curWord == endWord{
		paths = append(paths, newpath)
	} else {
		for _, neighWord := range maps[curWord]{
			if !visited[neighWord]{
				paths = dfs(neighWord, endWord, maps, visited, newpath, paths)
			}
		}
	}
	return paths
}



func FindLadders(beginWord string, endWord string, wordList []string){
	visited := make(map[string]bool)
	path := make([]string, 0)
	paths := [][]string{}
	for _, word := range wordList{
		visited[word] = false
	}
	fmt.Println(dfs("dog", "bob", list2Maps(wordList), visited, path, paths))
}