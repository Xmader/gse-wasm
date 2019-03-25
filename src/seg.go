package gse

// Segment 文本中的一个分词
type Segment struct {
	// 分词在文本中的起始字节位置
	Start int

	// 分词在文本中的结束字节位置（不包括该位置）
	End int

	// 分词信息
	Token *Token
}

// GetStart 返回分词在文本中的起始字节位置
func (s *Segment) GetStart() int {
	return s.Start
}

// GetEnd 返回分词在文本中的结束字节位置（不包括该位置）
func (s *Segment) GetEnd() int {
	return s.End
}

// GetToken 返回分词信息
func (s *Segment) GetToken() *Token {
	return s.Token
}
