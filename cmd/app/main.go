package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"log"
	"net/http"

	_ "github.com/lib/pq"
	"golang.org/x/crypto/bcrypt"
)

func registrationPageHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		// Если метод запроса POST, вызовите обработчик регистрации
		db, err := connectToDB()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer db.Close()
		registerUserHandler(w, r, db)
		return
	}

	data := map[string]interface{}{
		"Title": "Registration Page",
	}

	tmpl, err := template.ParseFiles("../../web/templates/registration.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := tmpl.Execute(w, data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func registerUserHandler(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	// Получите данные из формы регистрации, например:
	nickname := r.FormValue("nickname")
	email := r.FormValue("email")
	password := r.FormValue("password")

	// Проверьте, что все необходимые данные были заполнены
	if nickname == "" || email == "" || password == "" {
		http.Error(w, "Не заполнены все обязательные поля", http.StatusBadRequest)
		return
	}

	// Создайте объект User и заполните его данными
	user := User{
		Nickname:     nickname,
		Email:        email,
		PasswordHash: hashPassword(password), // Предполагается, что у вас есть функция для хеширования паролей
	}

	// Сохраните пользователя в базе данных
	err := saveUserToDB(db, user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Перенаправьте пользователя на страницу успешной регистрации или другую нужную вам страницу
	http.Redirect(w, r, "/success", http.StatusFound)
}

func saveUserToDB(db *sql.DB, user User) error {
	// Создайте SQL-запрос для вставки данных пользователя в таблицу "users"
	query := `INSERT INTO users (nickname, email, password_hash) VALUES ($1, $2, $3)`
	_, err := db.Exec(query, user.Nickname, user.Email, user.PasswordHash)
	if err != nil {
		return err
	}

	return nil
}

func hashPassword(password string) string {
	// Генерация соли и хэширование пароля
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		log.Fatal(err)
	}

	return string(hashedPassword)
}

func connectToDB() (*sql.DB, error) {
	connStr := "postgres://aizek:1234@localhost/quizzes?sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}
	return db, nil
}

func getUsersFromDB(db *sql.DB) ([]User, error) {
	rows, err := db.Query("SELECT * FROM users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []User

	for rows.Next() {
		var user User
		err := rows.Scan(&user.ID, &user.Nickname, &user.Email, &user.PasswordHash)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return users, nil
}

func successPageHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Registration Successful!")
}

type User struct {
	ID           int
	Nickname     string
	Email        string
	PasswordHash string
}

func main() {

	db, err := connectToDB()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Подключение к базе данных успешно!")

	users, err := getUsersFromDB(db)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Пользователи:")

	for _, user := range users {
		fmt.Println("ID:", user.ID, "| Username:", user.Nickname, "| E-Mail:", user.Email, "| Password_hash:", user.PasswordHash)
	}

	fs := http.FileServer(http.Dir("../../web/static/css"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	//http.HandleFunc("/", loginHandler)
	http.HandleFunc("/", registrationPageHandler)
	http.HandleFunc("/success", successPageHandler)

	log.Fatal(http.ListenAndServe(":5000", nil))
}
