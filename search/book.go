package search

var Book = map[string][]string{
	"": []string{"pd2-d4", "pe2-e4"},

	// C02 french advance
	"pe2-e4pe7-e6":                                           []string{"pd2-d4"},
	"pe2-e4pe7-e6pd2-d4":                                     []string{"pd7-d5"},
	"pe2-e4pe7-e6pd2-d4pd7-d5":                               []string{"pe4-e5", "nb1-c3"},
	"pe2-e4pe7-e6pd2-d4pd7-d5pe4-e5":                         []string{"pc7-c5"},
	"pe2-e4pe7-e6pd2-d4pd7-d5pe4-e5pc7-c5":                   []string{"pc2-c3"},
	"pe2-e4pe7-e6pd2-d4pd7-d5pe4-e5pc7-c5pc2-c3":             []string{"nb8-c6"},
	"pe2-e4pe7-e6pd2-d4pd7-d5pe4-e5pc7-c5pc2-c3nb8-c6":       []string{"ng1-f3"},
	"pe2-e4pe7-e6pd2-d4pd7-d5pe4-e5pc7-c5pc2-c3nb8-c6ng1-f3": []string{"bc8-d7", "qd8-b6"},

	// C18 french winawer
	"pe2-e4pe7-e6pd2-d4pd7-d5nb1-c3":             []string{"bf8-b5", "ng8-f6"},
	"pe2-e4pe7-e6pd2-d4pd7-d5nb1-c3bf8-b5":       []string{"pe4-e5"},
	"pe2-e4pe7-e6pd2-d4pd7-d5nb1-c3bf8-b5pe4-e5": []string{"pc7-c5"},
	"pe2-e4pe7-e6pd2-d4pd7-d5nb1-c3ng8-f6":       []string{"bc1-g5", "pe4-e5"},
	"pe2-e4pe7-e6pd2-d4pd7-d5nb1-c3ng8-f6bc1-g5": []string{"pd5-e4", "bf8-e7"},

	// C50 italian game
	"pe2-e4":                         []string{"pe7-e5", "pc7-c5", "pe7-e6"},
	"pe2-e4pe7-e5":                   []string{"ng1-f3"},
	"pe2-e4pe7-e5ng1-f3":             []string{"nb8-c6"},
	"pe2-e4pe7-e5ng1-f3nb8-c6":       []string{"bf1-c5", "bf1-b5"},
	"pe2-e4pe7-e5ng1-f3nb8-c6bf1-c4": []string{"ng8-f6", "bf8-c5"},

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
	"pd2-d4ng8-f6":                               []string{"pc2-c4"},
	"pd2-d4ng8-f6pc2-c4":                         []string{"pe7-e6"},
	"pd2-d4ng8-f6pc2-c4ng8-f6":                   []string{"nb1-c3", "ng1-f3"},
	"pd2-d4ng8-f6pc2-c4pe7-e6nb1-c3":             []string{"bf8-b4", "pc7-c5"},
	"pd2-d4ng8-f6pc2-c4pe7-e6nb1-c3bf8-b4":       []string{"ng1-f3", "qd1-c2"},
	"pd2-d4ng8-f6pc2-c4pe7-e6nb1-c3bf8-b4ng1-f3": []string{"pc7-c5", "pd7-d5"},
}
