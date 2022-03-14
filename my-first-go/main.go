package main

import (
	"context"
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	// "time"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

//this will create a schema and model

//schema in node
/*const employeeSchema = new mongoose.Schema({ //this is how we create a schema of a data entry that is going to go in the database
	name: String,
	address: String,
	position: String,
	salary: Number,
  }); */

type employeeSchema struct {
	ID          string       `json:"id"`
	Name        string       `json:"name"`
	Address     string       `json:"address"`
	Designation *Designation `json:"designation"`
	Salary      string       `json:"salary"`
	Email       string       `json:"email"`
	Phone       string       `json:"phone"`
}

/* ID string `json:"id"` first is the property that we are going to see in the database 2nd us the type and 3rd is the response
that we are going to recieve which is going to be json and then the key id remember in the json we write keys in "" so json:"id
means give me the response in json and key of that json is going to be "id"*/

//Designation schema struct
type Designation struct {
	Department string `json:"department"`
	Role       string `json:"role"`
}

//use the employeeSchema to create a var of slice:- its the indefinete length array
var employees []employeeSchema

func main() {
	//mongo db database connection
	clientURL := options.Client().ApplyURI("mongodb+srv://root:root@cluster0.kpteb.mongodb.net/myFirstDatabase?retryWrites=true&w=majority") //this how we establish the url for which later server will store data
	client, err := mongo.NewClient(clientURL) //create the instance of the db
	if err != nil {
	 log.Fatal(err)
	}
	err = client.Connect(context.Background())
	if err != nil {
	 log.Fatal(err)
	}
	// ctx, _ := context.WithTimeout(context.Background(), 15* time.Minute)
	 /*this will allow us to perform an action for certain time period and 
	if action is not performed suppose in this case in 15 sec then it will throw an error*/

	// database := client.Database("employeemodels") //creating a database
    // collection := database.Collection("employees") // collection to store data in that database

	// oneDoc := employeeSchema{
    //     ID:      "1",
	// 	Name:    "Harshal Upadhye",
	// 	Address: "Pune",
	// 	Designation: &Designation{ //this is the sub struct of the outer struct employeeSchema
	// 		Department: "Smart Building",
	// 		Role:       "Software Engineer",
	// 	},
	// 	Salary: "75,000",
	// 	Email:  "harshal.upadhye@siemens.com",
	// 	Phone:  "7498171447",
	// }
	
	// result, err := collection.InsertOne(ctx, oneDoc)
	// if err != nil {
	// 	log.Fatal("mongo.Connect() ERROR:", err)
	//   }
	// newID := result.InsertedID
	// log.Fatal(newID)

	//init Router
	router := mux.NewRouter() // exactly like const router = express.Router() but here := means (const router : any = mux.Router())
	headers := handlers.AllowedHeaders([]string{"content-type"}) //here content type application/json is the only thing we are allowing
	methods := handlers.AllowedMethods([]string{"GET","PUT","POST","DELETE"}) /* handler is the package that lets user decide what is allowed 
	in the header manually here we are defining the methods that are allowed to be performed in a string slice*/
	origins := handlers.AllowedOrigins([]string{"*"})
  
	
	
	//dummy data
	employees = append(employees, employeeSchema{
		ID:      "1",
		Name:    "Harshal Upadhye",
		Address: "Pune",
		Designation: &Designation{ //this is the sub struct of the outer struct employeeSchema
			Department: "Smart Building",
			Role:       "Software Engineer",
		},
		Salary: "75,000",
		Email:  "harshal.upadhye@siemens.com",
		Phone:  "7498171447",
	})
	employees = append(employees, employeeSchema{
		ID:      "2",
		Name:    "Abhishek Gaur",
		Address: "Pune",
		Designation: &Designation{ //this is the sub struct of the outer struct employeeSchema
			Department: "Smart Building",
			Role:       "Team Lead",
		},
		Salary: "1,00,000",
		Email:  "abhishek.gaur@siemens.com",
		Phone:  "1111111111",
	})
	employees = append(employees, employeeSchema{
		ID:      "3",
		Name:    "Ameya Pai",
		Address: "Karnataka",
		Designation: &Designation{ //this is the sub struct of the outer struct employeeSchema
			Department: "RB Tech",
			Role:       "Team Lead",
		},
		Salary: "3,00,000",
		Email:  "ameya.pai@rbtech.com",
		Phone:  "1111111111",
	})
	//router Handlers Endpoints
	//this will give the list of employees
	router.HandleFunc("/employees", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("content-type", "application/json") /*the w is response var and in that var we are setting up the header and telling application
		that we are receiving json object as a response*/
		json.NewEncoder(w).Encode(employees)

	}).Methods("GET") /*here in go we dont have app.get("/",(req, res)=>{}) or app.post or router.get or router.post so we have HandleFunc which takes route
	and function to be executed upon receiving that route where .Method tells method type plus here we have (req:Request, res:Response)=>{}
	where we were import {Request, Response} from 'express' but here in the place of req we have w and in the place of Request we have
	http.ResponseWriter and for res we have r and for Response we have *http.Request and * says handle multiple req*/

	//this will give single employee upon receiving id here we pass id like /{id} unlike node where we say /:id
	router.HandleFunc("/employees/{id}", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("content-type", "application/json")
		params := mux.Vars(r)
		for _, item := range employees {
			if item.ID == params["id"] {
				json.NewEncoder(w).Encode(item)
				return
			}
		}
	}).Methods("GET")

	//this one is to create an employee
	router.HandleFunc("/employees", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("content-type", "application/json")
		var employee employeeSchema                   /* create a var and assign the type as employeeSchema(our schema ie struct) same as per the var employees []employeeSchema there we were creating the slice array here we are creating a single employee*/
		_ = json.NewDecoder(r.Body).Decode(&employee) /* that _ automaticatlly point to the employee var, this say take the req.body decode it to
		json and then place it in the employee var but while placeing it in the var follow employeeSchema (.Decode(&employee)) where &employee
		points to the schema var employee employeeSchema*/
		employee.ID = strconv.Itoa(rand.Intn(10000000)) /* strconv.itoa converts int to string and rand.Intn picks random number between 0 to
		10000000*/
		employees = append(employees, employee) // now employees slice var stores employee single var
		json.NewEncoder(w).Encode(employee)     // send response



	}).Methods("POST")

	//this one is to update an employee
	router.HandleFunc("/employees/{id}", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("content-type", "application/json")
		params := mux.Vars(r)
		for index, item := range employees {
			if item.ID == params["id"] {
				employees = append(employees[:index], employees[index+1:]...)
				var employee employeeSchema
				_ = json.NewDecoder(r.Body).Decode(&employee)
				employee.ID = params["id"]
				employees = append(employees, employee)
				json.NewEncoder(w).Encode(employee)
				return

			}
		}

	}).Methods("PUT")

	//this one is to delete an employee
	router.HandleFunc("/employees/{id}", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("content-type", "application/json")
		params := mux.Vars(r)
		for index, item := range employees {
			if item.ID == params["id"] {
				employees = append(employees[:index], employees[index+1:]...) // exactly works like slice in js
				json.NewEncoder(w).Encode(employees)
				return
			}
		}

	}).Methods("DELETE")

	// this will spin the server

	log.Fatal(http.ListenAndServe(":8000", handlers.CORS(methods,origins,headers)(router) )) /* here we have log package which has fatal which will act as a try and catch block
	if it finds a problem it will through an error and the http.ListenAndServe is exactly like app.listen(3000,()=>{}) in node but in node it will
	take port number in int and a ananymous function as a callback to print server is running but here it will take string and the routers
	in node we use app.use(router) to let express know the routes*/
	//here we need to let our server know via handlers.CORS middleware that we are going to allow all this things

}
