package mysql
// From: https://github.com/megawubs/mysql_restore/blob/master/main.go
//
//package mysql
//
//import (
//	"flag"
//	"fmt"
//	"os"
//	"os/exec"
//	"database/sql"
//	"io/ioutil"
//	"bufio"
//	"log"
//)
//
//import (
//	"github.com/go-sql-driver/mysql"
//)
//
//func main() {
//	fileToImport := flag.String("file", "", "the file to import")
//	database := flag.String("database", "", "the database to use")
//	flag.Parse()
//	fmt.Println()
//	baseDir := os.Getenv("HOME") + "/Projects/mysql-dumps/"
//	if *fileToImport == "" {
//		fileToImport = getFileToImport(baseDir)
//	}
//	password := scan("mysql root password:")
//	if *database == "" {
//		database = getDatabase(password)
//	}
//
//	fmt.Printf("Going to import %s into %s\n", *fileToImport, *database)
//	command := fmt.Sprintf("mysql -u root -p'%s' -h 127.0.0.1 %s < %s", password, *database, *fileToImport)
//	cmd := exec.Command("sh", "-c", command)
//	cmd.Stderr = os.Stderr
//	cmd.Stdin = os.Stdin
//	cmd.Stdout = os.Stdout
//	err := cmd.Run()
//
//	if err != nil {
//		fmt.Errorf("error while runnin import command, %s", err)
//	}
//	fmt.Println("import done!")
//
//}
//
//func getDatabase(password string) (*string) {
//	dataSource := fmt.Sprintf("root:%s@tcp(127.0.0.1:3306)/", password)
//	db, err := sql.Open("mysql", dataSource)
//	defer db.Close()
//	if err != nil {
//		fmt.Errorf("error while setting up mysql connection %s", err)
//	}
//	err = db.Ping()
//	if err != nil {
//		fmt.Errorf("error while connecting to mysql")
//	}
//	dbs := databases(db)
//	for d, i := range dbs {
//		fmt.Printf("%d) %s\n", i, d)
//	}
//	createNew := len(dbs)
//	fmt.Println("-------")
//	fmt.Printf("%d) new database\n", createNew)
//	fmt.Printf("Choose database to use [0-%d]: ", createNew)
//	var index int
//	read("%d", &index)
//	var database string
//	if index != createNew {
//		for d, i := range dbs {
//			if i == index {
//				database = d
//				break
//			}
//		}
//	}
//	if index == createNew {
//		database = scan("The name of the database to import into: ")
//		createDatabase(db, database)
//	}
//	db.Close()
//	return &database
//}
//
//func getFileToImport(baseDir string) (*string) {
//	files, err := ioutil.ReadDir(baseDir)
//	if err != nil {
//		fmt.Errorf("someting went wrong %s", err)
//	}
//	for i, file := range files {
//		fmt.Printf("%d %s \n", i, file.Name())
//	}
//	fmt.Printf("Choose file to import [0:%d]: ", len(files)-1)
//	var index int
//	read("%d", &index)
//
//	name := baseDir + files[index].Name()
//	return &name
//}
//func createDatabase(db *sql.DB, d string) {
//
//	res, err := db.Exec("CREATE DATABASE " + d + " CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci")
//	if err != nil {
//		fmt.Errorf("error while creating new database, %s", err)
//	}
//	fmt.Print(res)
//	fmt.Printf("database '%s' created", d)
//}
//
//func read(format string, i interface{}) {
//	_, err := fmt.Scanf(format, i)
//	if err != nil {
//		fmt.Errorf("error while reading from stdin %s", err)
//	}
//}
//
//func scan(description string) string {
//	fmt.Print(description)
//	scanner := bufio.NewScanner(os.Stdin)
//	scanner.Scan() // use `for scanner.Scan()` to keep reading
//	return scanner.Text()
//}
//
//func databases(db *sql.DB) map[string]int {
//	rows, err := db.Query("SHOW DATABASES")
//	if err != nil {
//		fmt.Errorf("error while retreiving rows, %s", err)
//	}
//	defer rows.Close()
//	var name string
//	exclude := map[string]struct{}{
//		"information_schema": {},
//		"mysql":              {},
//		"performance_schema": {},
//		"sys":                {},
//	}
//	databases := make(map[string]int)
//	i := 0
//	for rows.Next() {
//		err := rows.Scan(&name)
//		if err != nil {
//			log.Fatal(err)
//		}
//		_, ok := exclude[name]
//		if ok == false {
//			databases[name] = i
//			i++
//		}
//
//	}
//	err = rows.Err()
//	if err != nil {
//		log.Fatal(err)
//	}
//	return databases
//}
