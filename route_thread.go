package main

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"forum/data"
)

// GET /threads/new
// Show the new thread form page
func newThread(writer http.ResponseWriter, request *http.Request) {
	_, err := session(writer, request)
	if err != nil {
		http.Redirect(writer, request, "/login", 302)
	} else {
		generateHTML(writer, request, nil, "layout", "private.navbar", "new.thread")
	}
}

// POST /signup
// Create the user account
func createThread(writer http.ResponseWriter, request *http.Request) {

	sess, err := session(writer, request)
	if err != nil {
		http.Redirect(writer, request, "/login", 302)
	} else {
		err = request.ParseForm()
		if err != nil {
			danger(err, "Cannot parse form")
			error_message(writer, request, "500 Internal Server Error")
		}
		user, err := sess.User()
		if err != nil {
			danger(err, "Cannot get user from session")
			error_message(writer, request, "500 Internal Server Error")
		}
		var categories []string
		for i := 1; i <= 5; i++ {
			temp := request.PostFormValue("category" + strconv.Itoa(i))
			if temp != "" {
				categories = append(categories, temp)
			}
		}
		topic := request.PostFormValue("topic")
		if strings.Trim(topic, " ") == "" {
			error_message(writer, request, "Cannot create an empty post")
			return
		}
		thread, err := user.CreateThread(topic, categories)
		if err != nil {
			danger(err, "Cannot create thread")
			error_message(writer, request, "500 Internal Server Error")
		}
		if err := thread.AssignCategory(categories); err != nil {
			danger(err, "Cannot assign category")
			error_message(writer, request, "500 Internal Server Error")
		}
		http.Redirect(writer, request, "/", 302)
	}
}

// GET /thread/read
// Show the details of the thread, including the posts and the form to write a post
func readThread(writer http.ResponseWriter, request *http.Request) {
	vals := request.URL.Query()
	uuid := vals.Get("id")
	thread, err := data.ThreadByUUID(uuid)
	if err != nil {
		error_message(writer, request, "404 Page Not Found")
	} else {
		_, err := session(writer, request)
		if err != nil {
			generateHTML(writer, request,  &thread, "layout", "public.navbar", "public.thread")
		} else {
			generateHTML(writer, request, &thread, "layout", "private.navbar", "private.thread")
		}
	}
}

// POST /thread/post
// Create the post
func postThread(writer http.ResponseWriter, request *http.Request) {
	sess, err := session(writer, request)
	if err != nil {
		http.Redirect(writer, request, "/login", 302)
	} else {
		err = request.ParseForm()
		if err != nil {
			danger(err, "Cannot parse form")
			error_message(writer, request, "500 Internal Server Error")
		}
		user, err := sess.User()
		if err != nil {
			danger(err, "Cannot get user from session")
			error_message(writer, request, "500 Internal Server Error")
		}
		body := request.PostFormValue("body")
		if strings.Trim(body, " ") == "" {
			error_message(writer, request, "Cannot create an empty comment")
			return
		}
		uuid := request.PostFormValue("uuid")
		thread, err := data.ThreadByUUID(uuid)
		if err != nil {
			error_message(writer, request, "400 Bad request")
		}
		if _, err := user.CreatePost(thread, body); err != nil {
			danger(err, "Cannot create post")
			error_message(writer, request, "500 Internal Server Error")
		}
		url := fmt.Sprint("/thread/read?id=", uuid)
		http.Redirect(writer, request, url, 302)
	}
}

func addThreadLike(writer http.ResponseWriter, request *http.Request) {
	sess, err := session(writer, request)
	if err != nil {
		http.Redirect(writer, request, "/login", 302)
	} else {
		err = request.ParseForm()
		if err != nil {
			danger(err, "Cannot parse form")
			error_message(writer, request, "500 Internal Server Error")
		}
		user, err := sess.User()
		if err != nil {
			danger(err, "Cannot get user from session")
			error_message(writer, request, "500 Internal Server Error")
		}
		uuid := request.PostFormValue("uuid")
		thread, err := data.ThreadByUUID(uuid)
		if err != nil {
			error_message(writer, request, "400 Bad request")
		}
		if err := user.RateThread(thread); err != nil {
			danger(err, "Cannot rate thread")
			error_message(writer, request, "500 Internal Server Error")
		}
		url := request.PostFormValue("url")
		http.Redirect(writer, request, url, 302)
	}
}

func addThreadDislike(writer http.ResponseWriter, request *http.Request) {
	sess, err := session(writer, request)
	if err != nil {
		http.Redirect(writer, request, "/login", 302)
	} else {
		err = request.ParseForm()
		if err != nil {
			danger(err, "Cannot parse form")
			error_message(writer, request, "500 Internal Server Error")
		}
		user, err := sess.User()
		if err != nil {
			danger(err, "Cannot get user from session")
			error_message(writer, request, "500 Internal Server Error")
		}
		uuid := request.PostFormValue("uuid")
		thread, err := data.ThreadByUUID(uuid)
		if err != nil {
			error_message(writer, request, "400 Bad request")
		}
		if err := user.UnrateThread(thread); err != nil {
			danger(err, "Cannot rate thread")
			error_message(writer, request, "500 Internal Server Error")
		}
		url := request.PostFormValue("url")
		http.Redirect(writer, request, url, 302)
	}
}

func addPostLike(writer http.ResponseWriter, request *http.Request) {
	sess, err := session(writer, request)
	if err != nil {
		http.Redirect(writer, request, "/login", 302)
	} else {
		err = request.ParseForm()
		if err != nil {
			danger(err, "Cannot parse form")
			error_message(writer, request, "500 Internal Server Error")
		}
		user, err := sess.User()
		if err != nil {
			danger(err, "Cannot get user from session")
			error_message(writer, request, "500 Internal Server Error")
		}
		uuid := request.PostFormValue("uuid")
		post, err := data.PostByUUID(uuid)
		if err != nil {
			error_message(writer, request, "400 Bad request")
		}
		if err := user.RatePost(post); err != nil {
			danger(err, "Cannot rate post")
			error_message(writer, request, "500 Internal Server Error")
		}
		url := request.PostFormValue("url")
		http.Redirect(writer, request, url, 302)
	}
}

func addPostDislike(writer http.ResponseWriter, request *http.Request) {
	sess, err := session(writer, request)
	if err != nil {
		http.Redirect(writer, request, "/login", 302)
	} else {
		err = request.ParseForm()
		if err != nil {
			danger(err, "Cannot parse form")
			error_message(writer, request, "500 Internal Server Error")
		}
		user, err := sess.User()
		if err != nil {
			danger(err, "Cannot get user from session")
			error_message(writer, request, "500 Internal Server Error")
		}
		uuid := request.PostFormValue("uuid")
		post, err := data.PostByUUID(uuid)
		if err != nil {
			error_message(writer, request, "400 Bad request")
		}
		if err := user.UnratePost(post); err != nil {
			danger(err, "Cannot rate post")
			error_message(writer, request, "500 Internal Server Error")
		}
		url := request.PostFormValue("url")
		http.Redirect(writer, request, url, 302)
	}
}
