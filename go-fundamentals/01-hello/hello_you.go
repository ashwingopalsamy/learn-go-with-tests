package main

func HelloName(name, language string) string {
	switch language {
	case langSpanish:
		return spanPrefixHello + name + suffixSymbol

	case langFrench:
		return frenPrefixHello + name + suffixSymbol

	case langEnglish:
		return engPrefixHello + name + suffixSymbol

	default:
		return engPrefixHello + "World" + suffixSymbol
	}
}
