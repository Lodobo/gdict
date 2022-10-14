package main

import (
	"database/sql"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"os/exec"

	"atomicgo.dev/keyboard/keys"
	_ "github.com/mattn/go-sqlite3"
	"github.com/pterm/pterm"
	"github.com/xyproto/unzip"
)

func main() {

	// Create TUI
	var options []string
	options = append(options,
		"English",           // en
		"Arabic",            // ar
		"Danish",            // da
		"German",            // de
		"Spanish",           // es
		"Finnish",           // fi
		"French",            // fr
		"Hindi",             // hi
		"Icelandic",         // is
		"Italian",           // it
		"Japanese",          // ja
		"Latin",             // la
		"Norwegian",         // no
		"Norwegian bokmål",  // nb
		"Norwegian nynorsk", // nn
		"Dutch",             // nl
		"Polish",            // pl
		"Portuguese",        // pt
		"Russian",           // ru
		"Northern sami",     // se
		"Swedish",           // sv
		"Urdu",              // ur
		"Telugu",            // te
		"Chinese",           // zh
	)

	printer := pterm.DefaultInteractiveMultiselect.WithMaxHeight(24).WithOptions(options).WithDefaultText("Select languages to install:")
	printer.Filter = false
	printer.KeyConfirm = keys.Enter
	printer.KeySelect = keys.Space
	selectedOptions, _ := printer.Show()

	// open database and create it if it doesn't exist
	home, err := os.UserHomeDir() // Path to home directory
	checkErr(err)
	err = os.MkdirAll(home+"/.local/share/gdict", os.ModePerm)
	checkErr(err)
	err = os.MkdirAll(home+"/.local/bin", os.ModePerm)
	checkErr(err)
	var dbpath string = home + "/.local/share/gdict/dictionary.db" // Path to db
	db, err := sql.Open("sqlite3", dbpath)
	if errors.Is(err, os.ErrNotExist) {
		file, err := os.Create(dbpath)
		checkErr(err)
		file.Close()
	}

	// Install selected languages
	for _, lang := range selectedOptions {
		switch lang {
		case "English":
			downloadFile("en.zip", en_url)
			unzipFile("en.zip")
			add_table("en.sql")
			createIndex("en", db)
		case "Arabic":
			downloadFile("ar.zip", ar_url)
			unzipFile("ar.zip")
			add_table("ar.sql")
			createIndex("ar", db)
		case "Danish":
			downloadFile("da.zip", da_url)
			unzipFile("da.zip")
			add_table("da.sql")
			createIndex("da", db)
		case "German":
			downloadFile("de.zip", de_url)
			unzipFile("de.zip")
			add_table("de.sql")
			createIndex("de", db)
		case "Spanish":
			downloadFile("es.zip", es_url)
			unzipFile("es.zip")
			add_table("es.sql")
			createIndex("es", db)
		case "Finnish":
			downloadFile("fi.zip", fi_url)
			unzipFile("fi.zip")
			add_table("fi.sql")
			createIndex("fi", db)
		case "French":
			downloadFile("fr.zip", fr_url)
			unzipFile("fr.zip")
			add_table("fr.sql")
			createIndex("fr", db)
		case "Hindi":
			downloadFile("hi.zip", hi_url)
			unzipFile("hi.zip")
			add_table("hi.sql")
			createIndex("hi", db)
		case "Icelandic":
			downloadFile("is.zip", is_url)
			unzipFile("is.zip")
			add_table("is.sql")
			createIndex("is", db)
		case "Italian":
			downloadFile("it.zip", it_url)
			unzipFile("it.zip")
			add_table("it.sql")
			createIndex("it", db)
		case "Japanese":
			downloadFile("ja.zip", ja_url)
			unzipFile("ja.zip")
			add_table("ja.sql")
			createIndex("ja", db)
		case "Latin":
			downloadFile("la.zip", la_url)
			unzipFile("la.zip")
			add_table("la.sql")
			createIndex("la", db)
		case "Norwegian":
			downloadFile("no.zip", no_url)
			unzipFile("no.zip")
			add_table("no.sql")
			createIndex("no", db)
		case "Norwegian bokmål":
			downloadFile("nb.zip", nb_url)
			unzipFile("nb.zip")
			add_table("nb.sql")
			createIndex("nb", db)
		case "Norwegian nynorsk":
			downloadFile("nn.zip", nn_url)
			unzipFile("nn.zip")
			add_table("nn.sql")
			createIndex("nn", db)
		case "Dutch":
			downloadFile("nl.zip", nl_url)
			unzipFile("nl.zip")
			add_table("nl.sql")
			createIndex("nl", db)
		case "Polish":
			downloadFile("pl.zip", pl_url)
			unzipFile("pl.zip")
			add_table("pl.sql")
			createIndex("pl", db)
		case "Portuguese":
			downloadFile("pt.zip", pt_url)
			unzipFile("pt.zip")
			add_table("pt.sql")
			createIndex("pt", db)
		case "Russian":
			downloadFile("ru.zip", ru_url)
			unzipFile("ru.zip")
			add_table("ru.sql")
			createIndex("ru", db)
		case "Northern sami":
			downloadFile("se.zip", se_url)
			unzipFile("se.zip")
			add_table("se.sql")
			createIndex("se", db)
		case "Swedish":
			downloadFile("sv.zip", sv_url)
			unzipFile("sv.zip")
			add_table("sv.sql")
			createIndex("sv", db)
		case "Urdu":
			downloadFile("ur.zip", ur_url)
			unzipFile("ur.zip")
			add_table("ur.sql")
			createIndex("ur", db)
		case "Telugu":
			downloadFile("te.zip", te_url)
			unzipFile("te.zip")
			add_table("te.sql")
			createIndex("te", db)
		case "Chinese":
			downloadFile("zh.zip", zh_url)
			unzipFile("zh.zip")
			add_table("zh.sql")
			createIndex("zh", db)
		}
	}

	fmt.Println("\nDone !")
}

// Parent directory: https://www.dropbox.com/sh/6do0czqrrrr4voe/AABRSXPM-xgy7bNLZW_m4foqa?dl=0
const en_url = "https://www.dropbox.com/s/bj24ksbx96twada/en.zip?raw=1"
const ar_url = "https://www.dropbox.com/s/rbjuuceuzamy0ax/ar.zip?raw=1"
const da_url = "https://www.dropbox.com/s/ki36u7zuj106olm/da.zip?raw=1"
const de_url = "https://www.dropbox.com/s/vmkjl7f06av6xcg/de.zip?raw=1"
const es_url = "https://www.dropbox.com/s/x93qf446eyqzcg7/es.zip?raw=1"
const fi_url = "https://www.dropbox.com/s/s6llmbfthn6kdkr/fi.zip?raw=1"
const fr_url = "https://www.dropbox.com/s/hha7dkz17iy1ndn/fr.zip?raw=1"
const hi_url = "https://www.dropbox.com/s/zsa2r9t4vergu1d/hi.zip?raw=1"
const is_url = "https://www.dropbox.com/s/mjt0ewe9k8bayfp/is.zip?raw=1"
const it_url = "https://www.dropbox.com/s/feks2kvdm6kgw6a/it.zip?raw=1"
const ja_url = "https://www.dropbox.com/s/dd1qzwfxyg53t9y/ja.zip?raw=1"
const la_url = "https://www.dropbox.com/s/sgbaphuj38665dw/la.zip?raw=1"
const nb_url = "https://www.dropbox.com/s/ke9suf2z0074apu/nb.zip?raw=1"
const nl_url = "https://www.dropbox.com/s/4fbgkfmxnxk4t9n/nl.zip?raw=1"
const nn_url = "https://www.dropbox.com/s/smxg28tzn85kesj/nn.zip?raw=1"
const no_url = "https://www.dropbox.com/s/zjj4rlk0sihm3qz/no.zip?raw=1"
const pl_url = "https://www.dropbox.com/s/esd3dix9syz0sc0/pl.zip?raw=1"
const pt_url = "https://www.dropbox.com/s/qu89ovnwz7ojpvd/pt.zip?raw=1"
const ru_url = "https://www.dropbox.com/s/wrqf4igf2rjysev/ru.zip?raw=1"
const se_url = "https://www.dropbox.com/s/9nups73rj9y16wt/se.zip?raw=1"
const sv_url = "https://www.dropbox.com/s/om16b5sf7uzy27h/sv.zip?raw=1"
const ur_url = "https://www.dropbox.com/s/eih6b0azqjs13nl/ur.zip?raw=1"
const te_url = "https://www.dropbox.com/s/67c456ujn7tu74z/te.zip?raw=1"
const zh_url = "https://www.dropbox.com/s/mm112s1s8smoja2/zh.zip?raw=1"

func downloadFile(file_name string, url string) {
	// Create the file
	out, err := os.Create(file_name)
	checkErr(err)
	defer out.Close()
	log.SetFlags(log.Ltime)
	log.Println("Downloading " + file_name + "...")
	// Get the data
	resp, err := http.Get(url)
	checkErr(err)
	defer resp.Body.Close()
	// Check server response
	if resp.StatusCode != http.StatusOK {
		err := fmt.Errorf("bad status: %s", resp.Status)
		checkErr(err)
	}
	// Writer the body to file
	_, err = io.Copy(out, resp.Body)
	checkErr(err)
}

func unzipFile(file_name string) {
	// Unzip file
	log.SetFlags(log.Ltime)
	log.Println("Unzipping " + file_name + "...")
	err := unzip.Extract(file_name, "./")
	checkErr(err)
	// zip file no longer necessary, delete it
	log.Println("Deleting " + file_name + "...")
	os.Remove(file_name)
}

func add_table(file_name string) {
	log.SetFlags(log.Ltime)
	log.Println("Reading " + file_name + "...")
	home, err := os.UserHomeDir()
	checkErr(err)
	exec.Command("sqlite3", home+"/.local/share/gdict/dictionary.db", "-init", file_name).Run()
	log.Println("Deleting " + file_name + "...")
	os.Remove(file_name)
}

func createIndex(lang string, db *sql.DB) {
	create_index_statement := "CREATE INDEX idx_" + lang + "_word ON [" + lang + "](word);"
	log.SetFlags(log.Ltime)
	log.Println("Creating index on table " + lang + "...")
	statement, err := db.Prepare(create_index_statement)
	checkErr(err)
	statement.Exec()
}

func checkErr(err error) {
	if err != nil {
		log.SetFlags(log.Ltime)
		log.Fatal(err.Error())
	}
}
