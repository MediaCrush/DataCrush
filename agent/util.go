package agent

func clean(result string) string {
	sz := len(result)
	return result[:sz-1]
}
