package main

import (
	"flag"
	"fmt"
	"net/http"
	"strings"
	//"github.com/gorilla/mux"
	"github.com/libgit2/git2go"
	"log"
)

func main() {
	fmt.Println("Hello gitnav world !")

	dir := flag.String("directory", "web/", "directory of web files")
	homeHandler := http.FileServer(http.Dir(*dir))

	http.Handle("/", homeHandler)
	http.HandleFunc("/repo/", indexHandler)
	http.ListenAndServe(":2340", nil)
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[5:])
	fmt.Println("Full path : ", r.URL)

	if r.URL.Path == "/favicon.ico" {
		fmt.Println("Ca va couper")
		return
	}

	fmt.Println("On va open le repo")
	repo, err := git.OpenRepository(r.URL.Path[5:])
	if err != nil {
		fmt.Println("Erreur sur ouverture : ", err)
		return
	}
	fmt.Println("On a open le repo")

	odb, err := repo.Odb()
	if err != nil {
		fmt.Println("Erreur sur Odb : ", err)
		log.Fatal(err)
	}

	fmt.Println("On va lancer le foreach")
	err = odb.ForEach(func(oid *git.Oid) error {

		fmt.Println("On va Lookup")
		obj, err := repo.Lookup(oid)
		if err != nil {
			log.Fatal("Lookup:", err)
		}

		switch obj := obj.(type) {
		default:
			fmt.Println("On a default")
		case *git.Blob:
			break
			fmt.Printf("==================================\n")
			fmt.Printf("obj:  %s\n", obj)
			fmt.Printf("Type: %s\n", obj.Type())
			fmt.Printf("Id:   %s\n", obj.Id())
			fmt.Printf("Size: %s\n", obj.Size())
		case *git.Commit:
			fmt.Printf("==================================\n")
			fmt.Printf("obj:  %s\n", obj)
			fmt.Printf("Type: %s\n", obj.Type())
			fmt.Printf("Id:   %s\n", obj.Id())
			author := obj.Author()
			fmt.Printf("    Author:\n        Name:  %s\n        Email: %s\n        Date:  %s\n", author.Name, author.Email, author.When)
			committer := obj.Committer()
			fmt.Printf("    Committer:\n        Name:  %s\n        Email: %s\n        Date:  %s\n", committer.Name, committer.Email, committer.When)
			fmt.Printf("    ParentCount: %s\n", obj.ParentCount())
			fmt.Printf("    TreeId:      %s\n", obj.TreeId())
			fmt.Printf("    Message:\n\n        %s\n\n", strings.Replace(obj.Message(), "\n", "\n        ", -1))
			//fmt.Printf("obj.Parent: %s\n",obj.Parent())
			//fmt.Printf("obj.ParentId: %s\n",obj.ParentId())
			//fmt.Printf("obj.Tree: %s\n",obj.Tree())
		case *git.Tree:
			break
			fmt.Printf("==================================\n")
			fmt.Printf("obj:  %s\n", obj)
			fmt.Printf("Type: %s\n", obj.Type())
			fmt.Printf("Id:   %s\n", obj.Id())
			fmt.Printf("    EntryCount: %s\n", obj.EntryCount())

		}
		return nil
	})

}

func panic(err error) {
	fmt.Println(err)
}
