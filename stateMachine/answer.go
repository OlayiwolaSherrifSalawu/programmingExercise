package statemachine

func stripHtml(s string) string {
	tag := map[string]string{
		"<": ">",
	}
	store := ""
	tags := ""
	foundTag := false
	for i := 0; i < len(s); i++ {
		val, ok := tag[string(s[i])]
		if ok {
			tags = val

		}
		if ok {
			foundTag = true
			continue
		}
		if string(s[i]) == tags {
			foundTag = false
			continue
		}
		if !ok && !foundTag {
			store += string(s[i])
		}
		if !ok && foundTag {
			continue
		}
	}
	return store
}
