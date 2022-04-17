package controllers

import (
	"github.com/revel/revel"
	"pet/app"
	"pet/app/models"
)

var isLoggedIn = false
var currentUser models.User

type App struct {
	*revel.Controller
}

func (c App) Index() revel.Result {
	return c.Render()
}

func (c App) Signup() revel.Result {
	return c.Render()
}

func (c App) SignupDone() revel.Result {
	var user models.User
	c.Params.Bind(&user, "User")
	c.Validation.Required(user.Name).Message("Numele este obligatoriu!")
	c.Validation.Required(user.Surname).Message("Prenumele este obligatoriu!")
	c.Validation.Required(user.Password).Message("Parola este obligatorie!")
	// c.Validation.MinSize(user.Password, 8).Message("Parola este prea scurta!")
	c.Validation.Required(user.Email).Message("Adresa de mail este obligatorie!")
	// c.Validation.Match(user.Email, regexp.MustCompile("^\\S+@\\S+\\.\\S+$")).Message("Adresa de mail este invalida!")
	c.Validation.Required(user.Phone).Message("Numarul de telefon este obligatoriu!")
	// c.Validation.Match(user.Phone, regexp.MustCompile("^[+]*[(]{0,1}[0-9]{1,4}[)]{0,1}[-\\s\\./0-9]*$")).Message("Numarul de telefon este invalid!")

	if !c.Validation.HasErrors() {
		app.DB.Create(&user)
		return c.Render(user)
	}
	c.Validation.Keep()
	c.FlashParams()
	return c.Redirect(App.Signup)
}

func (c App) Login() revel.Result {
	return c.Render()
}

func (c App) MyAccount() revel.Result {
	if isLoggedIn == true {
		name := currentUser.Name
		surname := currentUser.Surname
		email := currentUser.Email
		phone := currentUser.Phone
		return c.Render(name, surname, email, phone)
	}

	var tryUser models.User
	c.Params.Bind(&tryUser, "UserTry")
	c.Validation.Required(tryUser.Password).Message("Parola este obligatorie!")
	c.Validation.Required(tryUser.Email).Message("Adresa de mail este obligatorie!")

	if !c.Validation.HasErrors() {

		var userBuilder models.User

		app.DB.Where(&models.User{Email: tryUser.Email, Password: tryUser.Password}).First(&userBuilder)

		if userBuilder.ID == 0 {
			c.Validation.Error("Contul acesta nu exista! Verificati emailul si parola.")
		}

		if !c.Validation.HasErrors() {
			isLoggedIn = true
			currentUser = userBuilder
			name := currentUser.Name
			surname := currentUser.Surname
			email := currentUser.Email
			phone := currentUser.Phone
			return c.Render(name, surname, email, phone)
		}

		c.Validation.Keep()
		c.FlashParams()
		return c.Redirect(App.Login)
	}
	c.Validation.Keep()
	c.FlashParams()
	return c.Redirect(App.Login)
}

func (c App) Map() revel.Result {
	// app.DB.Where(&models.Animal{AnimalType: models.Dog}).First(&animal1)

	var animals []models.Animal
	app.DB.Find(&animals)
	type LocationHTML struct {
		X float64
		Y float64
	}
	var locations []LocationHTML
	var locationStrings []string
	for _, item := range animals {
		var loc LocationHTML
		loc.Y = item.FirstX
		loc.X = item.FirstY
		locations = append(locations, loc)
		if item.Location != "" {
			locationStrings = append(locationStrings, item.Location)
		}
	}
	return c.Render(locations, locationStrings)
}

func (c App) AnunturileTale() revel.Result {
	return c.Render()
}

func (c App) Formular() revel.Result {
	return c.Render()
}

func (c App) FormularSent() revel.Result {
	var lostAnimal models.Animal
	c.Params.Bind(&lostAnimal, "LostAnimal")
	app.DB.Create(&lostAnimal)
	lastLocation := lostAnimal.Location
	return c.Render(lastLocation)
}

func (c App) FormularGasit() revel.Result {
	return c.Render()
}

func (c App) Logout() revel.Result {
	isLoggedIn = false
	return c.Redirect(App.Index)
}

func (c App) AnimaleFaraStapan() revel.Result {
	return c.Render()
}
