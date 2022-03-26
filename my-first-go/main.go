package main

import (
	"context"
	"encoding/json"

	// "fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"

	"go.mongodb.org/mongo-driver/bson"
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
	ID          string       `bson:"id"`
	Name        string       `bson:"name"`
	Address     string       `bson:"address"`
	Designation *Designation `bson:"designation"`
	Salary      string       `bson:"salary"`
	Email       string       `bson:"email"`
	Phone       string       `bson:"phone"`
}

/* ID string `json:"id"` first is the property that we are going to see in the database 2nd us the type and 3rd is the response
that we are going to recieve which is going to be json and then the key id remember in the json we write keys in "" so json:"id
means give me the response in json and key of that json is going to be "id"*/

//Designation schema struct
type Designation struct {
	Department string `bson:"department"`
	Role       string `bson:"role"`
}

//use the employeeSchema to create a var of slice:- its the indefinete length array
var employees []employeeSchema
var employeesRecord []bson.M

func main() {
	//mongo db database connection

	ctx, _ := context.WithTimeout(context.Background(), 15*time.Minute)
	/*this will allow us to perform an action for certain time period and
	if action is not performed suppose in this case in 15 sec then it will throw an error like connection timeout*/

	clientURL := options.Client().ApplyURI("mongodb+srv://root:root@cluster0.kpteb.mongodb.net/myFirstDatabase?retryWrites=true&w=majority") //this how we establish the url for which later server will store data

	client, err := mongo.Connect(ctx, clientURL) //this is actual connection ie mongoose.connect(url)
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(ctx, nil) //this will check the ping of the server and if server is not connected then it will throw an error
	if err != nil {
		log.Fatal(err)
	}
	database := client.Database("employeemodels")  //creating a database
	collection := database.Collection("employees") // collection to store data in that database

	//list all the data from db
	//result, err := collection.Find(ctx, bson.M{}) /*this is Find method to find data from database here we pass ctx our connection timeout and bson.M{} that mean all records*/
	// if err != nil {
	// 	log.Fatal(err)
	// }
	//this will push all the records at the same time
	// var employeeList []bson.M //create slice var with bson.M type as we are recieving bson object from db
	// result.All(ctx,&employeeList) //result has data and .All retreives everything and puts into employeeList
	// fmt.Println(employeeList)

	//this will show one by onee
	// defer result.Close(ctx)
	// for result.Next(ctx) {
	// 	var employeeList bson.M
	// 	result.Decode(&employeeList)
	// 	fmt.Println(employeeList)
	// }

	//insert one doc into db collection
	// oneDoc := employeeSchema{
	//     ID:      "1",
	// 	Name:    "Kalpana Panchal",
	// 	Address: "Pune",
	// 	Designation: &Designation{ //this is the sub struct of the outer struct employeeSchema
	// 		Department: "Smart Building",
	// 		Role:       "Sr Software Engineer",
	// 	},
	// 	Salary: "1,00,000",
	// 	Email:  "kalpana.panchal@siemens.com",
	// 	Phone:  "1234567899",
	// }

	// result, err := collection.InsertOne(ctx, oneDoc)
	// if err != nil {
	// 	log.Fatal("mongo.Connect() ERROR:", err)
	// }
	// fmt.Println(result)

	//find one by id
	//    var employeeList bson.M
	//    err = collection.FindOne(ctx, bson.M{"id":"8498081"}).Decode(&employeeList)
	//    	if err != nil {
	// 	log.Fatal("mongo.Connect() ERROR:", err)
	//    } else {
	// 	   fmt.Println(employeeList)
	//    }

	//find one and delete by id

	// var employeeList bson.M
	//    err = collection.FindOneAndDelete(ctx, bson.M{"id":"1"}).Decode(&employeeList)
	//    	if err != nil {
	// 	log.Fatal("mongo.Connect() ERROR:", err)
	//    } else {
	// 	   fmt.Println(employeeList)
	//    }

	//find one and update by id
	// var employeeList bson.M
	// err = collection.FindOneAndUpdate(ctx, bson.M{"id": "8498081"}, bson.D{
	// 	{"$set", bson.D{{"name", "Jishnu Jishnu"}}},
	// }).Decode(&employeeList)
	// if err != nil {
	// 	log.Fatal("mongo.Connect() ERROR:", err)
	// } else {
	// 	fmt.Println(employeeList)
	// }

	//init Router
	router := mux.NewRouter()                                                    // exactly like const router = express.Router() but here := means (const router : any = mux.Router())
	headers := handlers.AllowedHeaders([]string{"content-type"})                 //here content type application/json is the only thing we are allowing
	methods := handlers.AllowedMethods([]string{"GET", "PUT", "POST", "DELETE"}) /* handler is the package that lets user decide what is allowed
	in the header manually here we are defining the methods that are allowed to be performed in a string slice*/
	origins := handlers.AllowedOrigins([]string{"*"})

	//dummy data
	// employees = append(employees, employeeSchema{
	// 	ID:      "1",
	// 	Name:    "Harshal Upadhye",
	// 	Address: "Pune",
	// 	Designation: &Designation{ //this is the sub struct of the outer struct employeeSchema
	// 		Department: "Smart Building",
	// 		Role:       "Software Engineer",
	// 	},
	// 	Salary: "75,000",
	// 	Email:  "harshal.upadhye@siemens.com",
	// 	Phone:  "7498171447",
	// })
	// employees = append(employees, employeeSchema{
	// 	ID:      "2",
	// 	Name:    "Abhishek Gaur",
	// 	Address: "Pune",
	// 	Designation: &Designation{ //this is the sub struct of the outer struct employeeSchema
	// 		Department: "Smart Building",
	// 		Role:       "Team Lead",
	// 	},
	// 	Salary: "1,00,000",
	// 	Email:  "abhishek.gaur@siemens.com",
	// 	Phone:  "1111111111",
	// })
	// employees = append(employees, employeeSchema{
	// 	ID:      "3",
	// 	Name:    "Ameya Pai",
	// 	Address: "Karnataka",
	// 	Designation: &Designation{ //this is the sub struct of the outer struct employeeSchema
	// 		Department: "RB Tech",
	// 		Role:       "Team Lead",
	// 	},
	// 	Salary: "3,00,000",
	// 	Email:  "ameya.pai@rbtech.com",
	// 	Phone:  "1111111111",
	// })
	//router Handlers Endpoints
	//this will give the list of employees
	router.HandleFunc("/employees", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("content-type", "application/json") /*the w is response var and in that var we are setting up the header and telling application
		that we are receiving json object as a response*/
		result, err := collection.Find(ctx, bson.M{})
		if err != nil {
			log.Fatal(err)
		} else {
			result.All(ctx, &employeesRecord) // put all the received records from mongodb into our employeesRecord bson.M[]
		}
		json.NewEncoder(w).Encode(employeesRecord)

	}).Methods("GET") /*here in go we dont have app.get("/",(req, res)=>{}) or app.post or router.get or router.post so we have HandleFunc which takes route
	and function to be executed upon receiving that route where .Method tells method type plus here we have (req:Request, res:Response)=>{}
	where we were import {Request, Response} from 'express' but here in the place of req we have w and in the place of Request we have
	http.ResponseWriter and for res we have r and for Response we have *http.Request and * says handle multiple req*/

	//this will give single employee upon receiving id here we pass id like /{id} unlike node where we say /:id
	router.HandleFunc("/employees/{id}", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("content-type", "application/json")
		params := mux.Vars(r) //this willcatch all the params in the url after ? sign
		//this is local slice code
		// for _, item := range employees {
		// 	if item.ID == params["id"] {
		// 		json.NewEncoder(w).Encode(item)
		// 		return
		// 	}
		// }

		//this comes from db
		var singleEmployee bson.M
		err := collection.FindOne(ctx, bson.M{"id": params["id"]}).Decode(&singleEmployee)
		if err != nil {
			log.Fatal(err)
		} else {
			json.NewEncoder(w).Encode(singleEmployee)
		}
	}).Methods("GET")

	//this one is to create an employee
	router.HandleFunc("/employees", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("content-type", "application/json")
		//this comes for local slice var
		// var employee employeeSchema                   /* create a var and assign the type as employeeSchema(our schema ie struct) same as per the var employees []employeeSchema there we were creating the slice array here we are creating a single employee*/
		// _ = json.NewDecoder(r.Body).Decode(&employee) /* that _ automaticatlly point to the employee var, this say take the req.body decode it to
		// json and then place it in the employee var but while placeing it in the var follow employeeSchema (.Decode(&employee)) where &employee
		// points to the schema var employee employeeSchema*/
		// employee.ID = strconv.Itoa(rand.Intn(10000000)) /* strconv.itoa converts int to string and rand.Intn picks random number between 0 to
		// 10000000*/
		// employees = append(employees, employee) // now employees slice var stores employee single var
		// json.NewEncoder(w).Encode(employee)     // send response

		//this comes from db
		var singleEmployee employeeSchema
		_ = json.NewDecoder(r.Body).Decode(&singleEmployee)
		singleEmployee.ID = strconv.Itoa(rand.Intn(10000000))
		result, err := collection.InsertOne(ctx, singleEmployee)
		if err != nil {
			log.Fatal(err)
		} else {
			json.NewEncoder(w).Encode(result)
		}

	}).Methods("POST")

	//this one is to update an employee
	router.HandleFunc("/employees/{id}", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("content-type", "application/json")
		params := mux.Vars(r)
		//this comes from slice var
		// for index, item := range employees {
		// 	if item.ID == params["id"] {
		// 		employees = append(employees[:index], employees[index+1:]...)
		// 		var employee employeeSchema
		// 		_ = json.NewDecoder(r.Body).Decode(&employee)
		// 		employee.ID = params["id"]
		// 		employees = append(employees, employee)
		// 		json.NewEncoder(w).Encode(employee)
		// 		return

		// 	}
		// }

		//this comes from db
		var employee employeeSchema
		_ = json.NewDecoder(r.Body).Decode(&employee)
		employee.ID = params["id"]
		var employeeList bson.M
		err = collection.FindOneAndUpdate(ctx, bson.M{"id": params["id"]}, bson.D{
			{"$set", bson.M{
				"id":      params["id"],
				"name":    employee.Name,
				"address": employee.Address,
				"designation": bson.M{
					"department": employee.Designation.Department,
					"role": employee.Designation.Role,
				},
				"salary":  employee.Salary,
				"email":   employee.Email,
				"phone":   employee.Phone,
			}},
		}).Decode(&employeeList)
		// if err != nil {
		// 	log.Fatal(err)
		// }
		// if singleEmployee != nil{
		// 	result, err := collection.Find(ctx, bson.M{})
		// 	if err != nil {
		// 		log.Fatal(err)
		// 	} else {
		// 		result.All(ctx,&employeesRecord)
		// 	}
		// 	json.NewEncoder(w).Encode(employeesRecord)

		// }

	}).Methods("PUT")

	//this one is to delete an employee
	router.HandleFunc("/employees/{id}", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("content-type", "application/json")
		params := mux.Vars(r)
		//this comes from slice var
		// for index, item := range employees {
		// 	if item.ID == params["id"] {
		// 		employees = append(employees[:index], employees[index+1:]...) // exactly works like slice in js
		// 		json.NewEncoder(w).Encode(employees)
		// 		return
		// 	}
		// }

		//this comes from db
		var employee bson.M
		err := collection.FindOneAndDelete(ctx, bson.M{"id": params["id"]}).Decode(&employee)
		if err != nil {
			log.Fatal(err)
		}
		if employee != nil {
			result, err := collection.Find(ctx, bson.M{})
			if err != nil {
				log.Fatal(err)
			} else {
				result.All(ctx, &employeesRecord)
			}
			json.NewEncoder(w).Encode(employeesRecord)

		}

	}).Methods("DELETE")

	// this will spin the server

	log.Fatal(http.ListenAndServe(":8000", handlers.CORS(methods, origins, headers)(router))) /* here we have log package which has fatal which will act as a try and catch block
	if it finds a problem it will through an error and the http.ListenAndServe is exactly like app.listen(3000,()=>{}) in node but in node it will
	take port number in int and a ananymous function as a callback to print server is running but here it will take string and the routers
	in node we use app.use(router) to let express know the routes*/
	//here we need to let our server know via handlers.CORS middleware that we are going to allow all this things

}
