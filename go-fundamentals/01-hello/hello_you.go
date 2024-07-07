package main

func HelloName(name string) string {
	if name == "" {
		return engPrefixHello + "World" + engSuffixSymbol
	}
	return engPrefixHello + name + engSuffixSymbol
}
