package main

import (
	"os"
	"path/filepath"

	"github.com/chzyer/readline"

	"github.com/CS80-Team/Goolean/internal/engine/structuresFactory"

	"github.com/CS80-Team/Goolean/internal/engine"
	"github.com/CS80-Team/Goolean/internal/engine/tokenizer"
	"github.com/CS80-Team/Goolean/internal/textprocessing"
	"github.com/CS80-Team/gshell/pkg/gshell"
)

func main() {
	engine := engine.NewEngine(
		textprocessing.NewDefaultProcessor(
			textprocessing.NewNormalizer(),
			textprocessing.NewStemmer(),
			textprocessing.NewStopWordRemover(),
		),
		tokenizer.NewDelimiterManager(
			&map[rune]struct{}{
				' ': {},

				',':  {},
				'?':  {},
				'!':  {},
				'.':  {},
				';':  {},
				':':  {},
				'\\': {},

				'(': {},
				')': {},
				'[': {},
				']': {},
				'{': {},
				'}': {},

				'=': {},
				'+': {},
				'-': {},
				'*': {},
				'/': {},
				'%': {},
				'^': {},
			},
		),
		*engine.NewIndexManager(structuresFactory.NewOrderedSliceFactory[int]()),
	)

	engine.LoadDirectory(filepath.Join(filepath.Base("."), "dataset"))

	stdin, stdinW := readline.NewFillableStdin(os.Stdin)

	s := gshell.NewShell(
		stdin,
		stdinW,
		os.Stdout,
		os.Stdout,
		gshell.SHELL_PROMPT,
		".shell_history",
		gshell.NewLogger("shell.log"),
	)

	RegisterCommands(s, engine)

	s.Run("Welcome to the Goolean search engine shell, type `help` for list of commands\n")
}
