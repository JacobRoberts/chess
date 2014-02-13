"Beware of bugs in the above code; I have only proved it correct, not tried it."
- Donald E. Knuth.

Golang Chess Engine and AI
=====

A gitignored web directory hosts JS, CSS, HTML, and images for the GUI

A list of required javascript libraries are:
* http://chessboardjs.com/
* https://github.com/jhlywa/chess.js

To run:
* Copy the repo into $GOPATH/src/chess
* $ go install chess/chess
* $ $GOPATH/bin/chess (or just $ chess if $GOPATH/bin is in your $PATH)

To test:
* $ go test chess/engine
* $ go test chess/negamax