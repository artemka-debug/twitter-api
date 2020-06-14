package utils

func Filter(comments []map[string]interface{}, test func(int) bool) (ret []map[string]interface{}) {
	for i := range comments {
		if test(comments[i]["postId"].(int)) {
			ret = append(ret, comments[i])
		}
	}
	return
}
