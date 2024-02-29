package app

import (
	"banking/domain"
	"banking/service"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
)

// CMD Application Configuration Options
// func sanityCheck() {
// 	if os.Getenv("SERVER_ADDRESS") == "" || os.Getenv("SERVER_PORT") == "" {
// 		log.Fatal("SERVER_ADDRESS and SERVER_PORT must be set")
// 	}
// }

func Start() {

	// sanityCheck()

	router := mux.NewRouter()

	//wiring
	//ch := CustomerHandlers{service.NewCustomerService(domain.NewCustomerRepositoryStub())}
	dbClient := getDbClient()
	customerRepositoryDB := domain.NewCustomerRepositoryDB(dbClient)
	accountRepositoryDB := domain.NewAccountRepositoryDB(dbClient)
	ch := CustomerHandlers{service.NewCustomerService(customerRepositoryDB)}
	ah := AccountHandler{service.NewAccountService(accountRepositoryDB)}
	//define routers
	router.HandleFunc("/customers", ch.getAllCustomers).Methods(http.MethodGet)
	router.HandleFunc("/customers/{customer_id:[0-9]+}", ch.getCustomer).Methods(http.MethodGet)
	router.HandleFunc("/customers/{customer_id:[0-9]+}/account", ah.NewAccount).Methods(http.MethodPost)

	//starting server
	// address := os.Getenv("SERVER_ADDRESS")
	// port := os.Getenv("SERVER_POST")
	// log.Fatal(http.ListenAndServe(fmt.Sprintf("%s:%s", address, port), router)) // CMD Application Configuration Options
	log.Fatal(http.ListenAndServe("localhost:8000", router)) // Standard Application Configuration Options
}

func getDbClient() *sqlx.DB {
	// dbUser := os.Getenv("DB_USER")
	// dbPasswd := os.Getenv("DB_PASSWD")
	// dbAddr := os.Getenv("DB_ADDR")
	// dbPort := os.Getenv("DB_PORT")
	// dbName := os.Getenv("DB_NAME")

	// root:my-secret-pw@tcp(localhost:3306)/banking // old way
	// export DB_USER=root
	// export DB_PASSWD=my-secret-pw
	// export DB_ADDR=localhost
	// export DB_PORT=3306
	// export DB_NAME=banking
	// SERVER_ADDRESS=localhost SERVER_PORT=8000 DB_USER=root DB_PASSWD=my-secret-pw DB_ADDR=localhost DB_PORT=3306 DB_NAME=banking

	// dataSource := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPasswd, dbAddr, dbPort, dbName)

	client, err := sqlx.Open("mysql", "root:my-secret-pw@tcp(localhost:3306)/banking")
	if err != nil {
		panic(err)
	}
	// See "Important settings" section.
	client.SetConnMaxLifetime(time.Minute * 3)
	client.SetMaxOpenConns(10)
	client.SetMaxIdleConns(10)
	return client
}
