package search

var Book = map[string][]string{
	// Initial position
	"rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR b": []string{"pd2-d4", "pe2-e4"},

	// A10 English Opening
	"rnbqkbnr/pppppppp/8/8/2P5/8/PP1PPPPP/RNBQKBNR b": []string{"ng8-f6", "pe7-e5"},

	// A15 English Opening, Anglo-Indian Defense
	"rnbqkb1r/pppppppp/5n2/8/2P5/8/PP1PPPPP/RNBQKBNR w": []string{"ng1-f3", "pd2-d4"},

	// A40 Queen's Pawn
	"rnbqkbnr/pppppppp/8/8/3P4/8/PPP1PPPP/RNBQKBNR b": []string{"pd7-d5", "ng8-f6"},

	// A45 Queen's Pawn Game
	"rnbqkb1r/pppppppp/5n2/8/3P4/8/PPP1PPPP/RNBQKBNR w": []string{"pc2-c4"},

	// A50 Queen's Pawn Game
	"rnbqkb1r/pppppppp/5n2/8/2PP4/8/PP2PPPP/RNBQKBNR b": []string{"pe7-e6"},

	// B00 King's Pawn Opening
	"rnbqkbnr/pppppppp/8/8/4P3/8/PPPP1PPP/RNBQKBNR b": []string{"pe7-e5", "pc7-c5", "pe7-e6"},

	// B20 Sicilian Defence
	"rnbqkbnr/pp1ppppp/8/2p5/4P3/8/PPPP1PPP/RNBQKBNR w": []string{"ng1-f3", "nb1-c3"},

	// B27 Sicilian Defence
	"rnbqkbnr/pp1ppppp/8/2p5/4P3/5N2/PPPP1PPP/RNBQKB1R b": []string{"pd7-d6", "pe7-e6"},

	// B40 Sicilian Defence
	"rnbqkbnr/pp1p1ppp/4p3/2p5/4P3/5N2/PPPP1PPP/RNBQKB1R w": []string{"pd2-d4"},
	"rnbqkbnr/pp1p1ppp/4p3/2p5/3PP3/5N2/PPP2PPP/RNBQKB1R b": []string{"pc5-d4"},
	"rnbqkbnr/pp1p1ppp/4p3/8/3pP3/5N2/PPP2PPP/RNBQKB1R w":   []string{"nf3-d4"},
	"rnbqkbnr/pp1p1ppp/4p3/8/3pP3/5N2/PPP2PPP/RNBQKB1R w":   []string{"pa7-a6", "nb8-c6", "ng8-f6"},

	// B50 Sicilian
	"nbqkbnr/pp2pppp/3p4/2p5/4P3/5N2/PPPP1PPP/RNBQKB1R w":  []string{"pd2-d4"},
	"rnbqkbnr/pp2pppp/3p4/2p5/3PP3/5N2/PPP2PPP/RNBQKB1R b": []string{"pc5-d4"},
	"nbqkbnr/pp2pppp/3p4/8/3pP3/5N2/PPP2PPP/RNBQKB1R w":    []string{"nf3-d4"},

	// B54 Sicilian
	"rnbqkbnr/pp2pppp/3p4/8/3NP3/8/PPP2PPP/RNBQKB1R b":   []string{"ng8-f6"},
	"rnbqkb1r/pp2pppp/3p1n2/8/3NP3/8/PPP2PPP/RNBQKB1R w": []string{"nb1-c3"},

	// C00 French Defense
	"rnbqkbnr/pppp1ppp/4p3/8/4P3/8/PPPP1PPP/RNBQKBNR w":  []string{"pd2-d4"},
	"rnbqkbnr/pppp1ppp/4p3/8/3PP3/8/PPP2PPP/RNBQKBNR b":  []string{"pd7-d5"},
	"rnbqkbnr/ppp2ppp/4p3/3p4/3PP3/8/PPP2PPP/RNBQKBNR w": []string{"pe4-e5", "nb1-c3"},

	// C02 French, Advance Variation
	"rnbqkbnr/ppp2ppp/4p3/3pP3/3P4/8/PPP2PPP/RNBQKBNR b":      []string{"pc7-c5"},
	"rnbqkbnr/pp3ppp/4p3/2ppP3/3P4/8/PPP2PPP/RNBQKBNR w":      []string{"pc2-c3"},
	"rnbqkbnr/pp3ppp/4p3/2ppP3/3P4/2P5/PP3PPP/RNBQKBNR b":     []string{"nb8-c6"},
	"r1bqkbnr/pp3ppp/2n1p3/2ppP3/3P4/2P5/PP3PPP/RNBQKBNR w":   []string{"ng1-f3"},
	"r1bqkbnr/pp3ppp/2n1p3/2ppP3/3P4/2P2N2/PP3PPP/RNBQKB1R b": []string{"bc8-d7", "qd8-b6"},

	// C10 French, Paulsen Variation
	"rnbqkbnr/ppp2ppp/4p3/3p4/3PP3/2N5/PPP2PPP/R1BQKBNR b": []string{"bf8-b5", "ng8-f6"},

	// C11 French Defense
	"rnbqkb1r/ppp2ppp/4pn2/3p4/3PP3/2N5/PPP2PPP/R1BQKBNR w":  []string{"bc1-g5", "pe4-e5"},
	"rnbqkb1r/ppp2ppp/4pn2/3p2B1/3PP3/2N5/PPP2PPP/R2QKBNR b": []string{"pd5-e4", "bf8-e7"},

	// C15 French, Winawer, Nimzovich Variation
	"rnbqk1nr/ppp2ppp/4p3/3p4/1b1PP3/2N5/PPP2PPP/R1BQKBNR w": []string{"pe4-e5"},

	// C16 French, Winawer, Advance Variation
	"rnbqk1nr/ppp2ppp/4p3/3pP3/1b1P4/2N5/PPP2PPP/R1BQKBNR b": []string{"pc7-c5"},

	// C20 King's Pawn Game
	"rnbqkbnr/pppp1ppp/8/4p3/4P3/8/PPPP1PPP/RNBQKBNR w": []string{"ng1-f3"},

	// C40 King's Knight Opening
	"rnbqkbnr/pppp1ppp/8/4p3/4P3/5N2/PPPP1PPP/RNBQKB1R b": []string{"nb8-c6"},

	// C44 King's Pawn Game
	"r1bqkbnr/pppp1ppp/2n5/4p3/4P3/5N2/PPPP1PPP/RNBQKB1R w": []string{"bf1-c5", "bf1-b5"},

	// C50 King's Pawn Game
	"r1bqkbnr/pppp1ppp/2n5/4p3/2B1P3/5N2/PPPP1PPP/RNBQK2R b": []string{"ng8-f6", "bf8-c5"},

	// C60 Ruy Lopez
	"r1bqkbnr/pppp1ppp/2n5/1B2p3/4P3/5N2/PPPP1PPP/RNBQK2R b":  []string{"pa7-a6", "ng8-f6"},
	"r1bqkbnr/1ppp1ppp/p1n5/1B2p3/4P3/5N2/PPPP1PPP/RNBQK2R w": []string{"bb5-c6", "bb5-a4"},

	// C65 Ruy Lopez, Berlin Defence
	"r1bqkb1r/pppp1ppp/2n2n2/1B2p3/4P3/5N2/PPPP1PPP/RNBQK2R w": []string{"ke1-g1"},

	// C68 Ruy Lopez, Exchange Variation
	"r1bqkbnr/1ppp1ppp/p1B5/4p3/4P3/5N2/PPPP1PPP/RNBQK2R b": []string{"pd7-c6"},
	"r1bqkbnr/1pp2ppp/p1p5/4p3/4P3/5N2/PPPP1PPP/RNBQK2R w":  []string{"ke1-g1"},

	// C69 Ruy Lopez, Exchange Variation
	"r1bqkbnr/1pp2ppp/p1p5/4p3/4P3/5N2/PPPP1PPP/RNBQ1RK1 b": []string{"pf7-f6", "bc8-g4"},

	// D00 Queen's Pawn Game
	"rnbqkbnr/ppp1pppp/8/3p4/3P4/8/PPP1PPPP/RNBQKBNR w": []string{"pc2-c4"},

	// D06 Queen's Gambit
	"rnbqkbnr/ppp1pppp/8/3p4/2PP4/8/PP2PPPP/RNBQKBNR b": []string{"pe7-e6", "pc7-c6"},

	// D10 Queen's Gambit Declined Slav
	"rnbqkbnr/pp2pppp/2p5/3p4/2PP4/8/PP2PPPP/RNBQKBNR w":     []string{"nb1-c3", "ng1-f3"},
	"rnbqkbnr/pp2pppp/2p5/3p4/2PP4/2N5/PP2PPPP/R1BQKBNR b":   []string{"ng8-f6"},
	"rnbqkb1r/pp2pppp/2p2n2/3p4/2PP4/2N5/PP2PPPP/R1BQKBNR w": []string{"ng1-f3"},

	// D11 Queen's Gambit Declined Slav
	"rnbqkbnr/pp2pppp/2p5/3p4/2PP4/5N2/PP2PPPP/RNBQKB1R b":   []string{"ng8-f6"},
	"rnbqkb1r/pp2pppp/2p2n2/3p4/2PP4/5N2/PP2PPPP/RNBQKB1R w": []string{"nb1-c3"},

	// D15 Queen's Gambit Declined Slav
	"rnbqkb1r/pp2pppp/2p2n2/3p4/2PP4/2N2N2/PP2PPPP/R1BQKB1R b ": []string{"pe7-e6"},

	// D30 Queen's Gambit Declined
	"rnbqkbnr/ppp2ppp/4p3/3p4/2PP4/8/PP2PPPP/RNBQKBNR w":   []string{"nb1-c3", "ng1-f3"},
	"rnbqkbnr/ppp2ppp/4p3/3p4/2PP4/5N2/PP2PPPP/RNBQKB1R b": []string{"ng8-f6", "pc7-c6"},

	// D31 Queen's Gambit Declined
	"rnbqkbnr/ppp2ppp/4p3/3p4/2PP4/2N5/PP2PPPP/R1BQKBNR b": []string{"ng8-f6", "pc7-c6"},

	// E00 Queen's Pawn Game
	"rnbqkb1r/pppp1ppp/4pn2/8/2PP4/8/PP2PPPP/RNBQKBNR w":   []string{"nb1-c3", "ng1-f3"},
	"rnbqkb1r/pppp1ppp/4pn2/8/2PP4/2N5/PP2PPPP/R1BQKBNR b": []string{"bf8-b4", "pc7-c5"},

	// E20 Nimzo-Indian Defence
	"rnbqk2r/pppp1ppp/4pn2/8/1bPP4/2N5/PP2PPPP/R1BQKBNR w": []string{"ng1-f3", "qd1-c2"},

	// E21 Nimzo-Indian, Three Knights Variation
	"rnbqk2r/pppp1ppp/4pn2/8/1bPP4/2N2N2/PP2PPPP/R1BQKB1R b": []string{"pc7-c5", "pd7-d5"},
}
