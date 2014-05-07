package search

var Book = map[string][]string{
	"": []string{"pd2-d4", "pe2-e4"},

	// C50 italian game
	"pe2-e4":                         []string{"pe7-e5", "pc7-c5", "pe7-e6"},
	"pe2-e4pe7-e5":                   []string{"ng1-f3"},
	"pe2-e4pe7-e5ng1-f3":             []string{"nb8-c6"},
	"pe2-e4pe7-e5ng1-f3nb8-c6":       []string{"bf1-c5", "bf1-b5"},
	"pe2-e4pe7-e5ng1-f3nb8-c6bf1-c4": []string{"ng8-f6", "bc8-c5"},

	// C68 ruy lopez exchange
	"pe2-e4pe7-e5ng1-f3nb8-c6bf1-b5":                         []string{"pa7-a6"},
	"pe2-e4pe7-e5ng1-f3nb8-c6bf1-b5pa7-a6":                   []string{"bb5-c6", "bb5-a4"},
	"pe2-e4pe7-e5ng1-f3nb8-c6bf1-b5pa7-a6bb5-c6":             []string{"pd7-c6"},
	"pe2-e4pe7-e5ng1-f3nb8-c6bf1-b5pa7-a6bb5-c6pd7-d6":       []string{"ke1-g1"},
	"pe2-e4pe7-e5ng1-f3nb8-c6bf1-b5pa7-a6bb5-c6pd7-d6nf3-e5": []string{"qd8-d4"},

	// D35 queens gambit declined
	"pd2-d4":                         []string{"pd7-d5", "ng8-f6"},
	"pd2-d4pd7-d5":                   []string{"pc2-c4"},
	"pd2-d4pd7-d5pc2-c4":             []string{"pe7-e6", "pc7-c6"},
	"pd2-d4pd7-d5pc2-c4pe7-e6":       []string{"nb1-c3", "ng1-f3"},
	"pd2-d4pd7-d5pc2-c4pe7-e6nb1-c3": []string{"ng8-f6", "pc7-c6"},
	"pd2-d4pd7-d5pc2-c4pe7-e6ng1-f3": []string{"ng8-f6", "pc7-c6"},

	// E20 nimzo-indian
	"pd2-d4ng8-f6": []string{"pc2-c4"},
}
