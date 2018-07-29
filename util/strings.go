package util

func CharAt(s string, b byte) byte {
	sa:=[]rune(s)
	return byte(sa[b])
}

