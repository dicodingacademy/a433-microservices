package main

func (app *application) getCount() int {
	return *app.counter
}

func (app *application) increment() {
	*app.counter++
}
