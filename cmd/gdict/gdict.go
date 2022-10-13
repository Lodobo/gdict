package main

import (
	"database/sql"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/charmbracelet/lipgloss"
	"github.com/pterm/pterm"
	"snai.pe/go-pager"

	_ "github.com/mattn/go-sqlite3"
)

var out *pager.Pager
var err error

func main() {
	home, err := os.UserHomeDir()
	checkErr(err)
	db, err := sql.Open("sqlite3", home+"/.local/share/gdict/dictionary.db")
	checkErr(err)
	defer db.Close()

	arg_w, arg_p, arg_l := arguments() // returns an object with type arg
	langs_installed := fetch_tables(db)

	valid_language := false

	for _, l := range langs_installed {
		if arg_l == l {
			valid_language = true
		}
	}
	if valid_language == false {
		fmt.Println(error_message.Render("Err:"), "Language not installed")
		fmt.Println("Installed languages:", langs_installed)
		os.Exit(3)
	}

	//validate part of speech and print valid ones on error

	var rows *sql.Rows
	if arg_w.given == true && arg_p.given == false {
		rows, err = db.Query(fmt.Sprintf("SELECT word, pos, sounds, etymology_text, senses FROM [%s] WHERE word = '%s'", arg_l, arg_w.value))
		checkErr(err) // make sure to handle error when table doesn't exist
		defer rows.Close()
	} else if arg_w.given == true && arg_p.given == true {
		rows, err = db.Query(fmt.Sprintf("SELECT word, pos, sounds, etymology_text, senses FROM [%s] WHERE word = '%s' AND pos = '%s'", arg_l, arg_w.value, arg_p.value))
		checkErr(err)
		defer rows.Close()
	} else {
		fmt.Println("\nExpected at least one argument:")
		fmt.Println("\t gdict", error_message.Render("-w word"))
		os.Exit(3)
	}

	// Open pager
	out, err = pager.Open()
	checkErr(err)
	defer out.Close()

	var word sql.NullString
	var pos sql.NullString
	var sounds sql.NullString
	var etymology_text sql.NullString
	var senses sql.NullString
	var count int
	for rows.Next() {
		count++
		err := rows.Scan(&word, &pos, &sounds, &etymology_text, &senses)
		checkErr(err)

		if word.Valid == true && pos.Valid == true {
			print_panel(word.String, pos.String)
		}
		if sounds.Valid == true {
			print_sounds(sounds.String)
		}
		if etymology_text.Valid == true {
			print_etymology(etymology_text.String)
		}
		if senses.Valid == true {
			print_senses(senses.String)
		}
	}
	if count == 0 {
		out.Write([]byte(error_message.Render("Word not in dictionary\n")))
	}
}

// Define styles
var error_message = lipgloss.NewStyle().
	Bold(true).
	Foreground(lipgloss.Color("1"))

var panel = lipgloss.NewStyle().
	Foreground(lipgloss.Color("6")).
	Bold(true).
	BorderStyle(lipgloss.RoundedBorder()).
	BorderForeground(lipgloss.Color("6")).
	BorderTop(true).
	BorderBottom(true).
	BorderLeft(true).
	BorderRight(true).
	PaddingLeft(2).
	PaddingRight(2).
	MarginTop(1).
	MarginBottom(1).
	MarginRight(4).
	MarginLeft(1)

var s_title = lipgloss.NewStyle().
	Bold(true).
	Foreground(lipgloss.Color("2")).
	PaddingBottom(1).
	MarginLeft(1)

var s_sounds = lipgloss.NewStyle().
	Foreground(lipgloss.Color("6")).
	PaddingBottom(1).
	PaddingLeft(4)

var s_tags = lipgloss.NewStyle().
	Foreground(lipgloss.Color("7")).
	PaddingBottom(1).
	PaddingLeft(0)

var s_etym = lipgloss.NewStyle().
	Foreground(lipgloss.Color("6")).
	MarginLeft(3).
	MarginRight(1).
	MarginBottom(1)

var s_gloss = lipgloss.NewStyle().
	Foreground(lipgloss.Color("6")).
	MarginLeft(3).
	MarginRight(1).
	MarginBottom(1)

var s_example = lipgloss.NewStyle().
	Foreground(lipgloss.Color("30")).
	MarginLeft(9).
	MarginRight(1).
	MarginBottom(1)

func checkErr(err error) {
	if err != nil {
		log.Fatal(err.Error())
	}
}

// Print functions
func print_panel(word string, pos string) {
	pos = strings.ToUpper(pos)
	var Panel1 = panel.Render(pos)
	var Panel2 = panel.Render(word)
	var panels = lipgloss.JoinHorizontal(lipgloss.Bottom, Panel1, Panel2)
	out.Write([]byte(panels + "\n"))
}
func print_sounds(sounds string) {
	var pron []Pronunciation
	json.Unmarshal([]byte(sounds), &pron)
	if len(pron) > 0 {
		out.Write([]byte(s_title.Render("# Pronunciation:")))
	}
	out.Write([]byte("\n"))
	for i := 0; i < len(pron); i++ {
		if len(pron[i].IPA) > 1 {
			if len(pron[i].Tags) > 0 {
				var tag string = pron[i].Tags[0]
				var Panel1 = s_sounds.Render(pron[i].IPA)
				var Panel2 = s_tags.Render(" (" + tag + ")")
				var panels = lipgloss.JoinHorizontal(lipgloss.Bottom, Panel1, Panel2)
				out.Write([]byte(panels + "\n"))

			} else {
				out.Write([]byte(s_sounds.Render(pron[i].IPA) + "\n"))
			}
		}
	}
}
func print_etymology(etymology_text string) {
	out.Write([]byte(s_title.Render("# Etymology:")))
	out.Write([]byte("\n"))
	var ety string = etymology_text
	ety = pterm.DefaultParagraph.WithMaxWidth(95).Sprint(ety)
	ety = s_etym.Render(ety)
	out.Write([]byte(ety + "\n"))
}
func print_senses(senses string) {
	var defs []Sense
	json.Unmarshal([]byte(senses), &defs)
	out.Write([]byte(s_title.Render("# Definitions:") + "\n"))
	for i := 0; i < len(defs); i++ {
		var glosses []string = defs[i].Glosses
		if len(glosses) > 1 {
			glosses = glosses[1:]
		}
		var gloss = strings.Join(glosses, "\n➜ ")
		gloss = pterm.DefaultParagraph.WithMaxWidth(87).Sprint(gloss)
		gloss = s_gloss.Render("• " + gloss)

		out.Write([]byte(gloss + "\n"))
		for y := 0; y < len(defs[i].Examples); y++ {
			var ex string = defs[i].Examples[y].Text
			ex = pterm.DefaultParagraph.WithMaxWidth(87).Sprint(ex)
			ex = s_example.Render("\"" + ex + "\"")
			out.Write([]byte(ex))
			out.Write([]byte("\n"))
		}

	}
}

// function to return command line arguments
func arguments() (arg, arg, string) {
	// valid languages :
	word := flag.String("w", "", "Search word in dictionary")
	pos := flag.String("p", "", "Specify part of speech")
	var lang *string = flag.String("l", "en", "Specify language")
	flag.Parse()

	var arg_w arg
	var arg_p arg
	var arg_l string = *lang
	if *word == "" {
		arg_w = arg{given: false, value: *word}
	} else {
		arg_w = arg{given: true, value: *word}
	}
	if *pos == "" {
		arg_p = arg{given: false, value: *pos}
	} else {
		arg_p = arg{given: true, value: *pos}
	}
	return arg_w, arg_p, arg_l
}
func fetch_tables(db *sql.DB) []string {
	// check if table exists
	stmt, err := db.Prepare("SELECT count(*) FROM sqlite_master WHERE type='table' AND name= ? ;")
	checkErr(err)
	var langs_installed []string
	var result int
	var lang_list = [22]string{"en", "ar", "da", "de", "es", "fi", "fr", "is", "it", "ja", "la", "no", "nb", "nn", "nl", "pl", "pt", "ru", "se", "sv", "ur", "zh"}
	for _, l := range lang_list {
		stmt.QueryRow(l).Scan(&result)
		if result == 1 {
			langs_installed = append(langs_installed, l)
		}
	}
	return langs_installed
}

// Types for decoding JSON
type Pronunciation struct {
	IPA     string   `json:"ipa"`
	Tags    []string `json:"tags"`
	Audio   string   `json:"audio"`
	Text    string   `json:"text"`
	Ogg_url string   `json:"ogg_url"`
	Mp3_url string   `json:"mp3_url"`
}
type Sense struct {
	Glosses  []string `json:"glosses"`
	Tags     []string `json:"tags"`
	Examples []struct {
		Text string `json:"text"`
		Ref  string `json:"ref"`
	} `json:"examples"`
	Categories []struct {
		Name   string `json:"name"`
		Kind   string `json:"kind"`
		Source string `json:"source"`
	} `json:"categories"`
}

// Type for command line argument results
type arg struct {
	given bool
	value string
}
