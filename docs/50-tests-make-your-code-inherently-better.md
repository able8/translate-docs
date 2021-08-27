# Tests Make Your Code Inherently Better

From: https://www.mitchdennett.com/tests-make-your-code-inherently-better/

I have been developing software for about 8 years now but my  journey into testing has only been within the last 2 years. In college  we had half a semester on testing. Not even a whole class on it. It was  definitely an after thought and not a main focus it should be. Writing  tests has made my code cleaner and more concise than any other endeavor. Let's take a quick look at why. 

In this short example we are going to look at a REST API that returns a list of Recipes. 

## The "Bad" Code

Let's take a look at some quote unquote bad code. 

```go
func (h *Handler) handleListRecipes(w http.ResponseWriter, r *http.Request) {
	pages, ok := r.URL.Query()["page"]

	if !ok || len(pages[0]) < 1 {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	page, err := strconv.Atoi(pages[0])

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	queryDocStmt := `SELECT recipe_id, title from recipe limit 50 offset $1`

	var offset int
	if page-1 < 0 {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	offset = (page - 1) * 50

	rows, err := h.DB.Query(queryDocStmt, offset)
	itemsList := make([]*better.Recipe, 0)

	if err != nil {
		log.Println(err)
	}
	defer rows.Close()
	for rows.Next() {
		var item better.Recipe

		if err := rows.Scan(&item.ID, &item.Title); err != nil {
			// Check for a scan error.
			// Query rows will be closed with defer.
			log.Println(err)
			continue
		}

		itemsList = append(itemsList, &item)

	}

	retJSON, err := json.Marshal(itemsList)
	fmt.Fprintf(w, string(retJSON))
}
```

This code is a  mess. It is not testable at all. When unit testing we don't want to  actually make a call to the database. We can assume the standard library has been tested exhaustively. 

So back to our code. There is no easy way to mock out our database calls so we can write a good test. 

## Refactoring Out Our Database Calls

So we need to make the database call mockable. This just means that we  need a way to replace the actual database call with something else  during our test so we don't actually make to request to the database. So let's pull it out and create a `Store` that is responsible for all interactions with the database. 

```go
type Store struct {
	DB *sql.DB
}

//ListRecipes will list all the recipes for a given page
func (d *DB) ListRecipes(page int) ([]*better.Recipe, error) {
	queryDocStmt := `SELECT recipe_id, title from recipe limit 50 offset $1`

	var offset int
	if page-1 < 0 {
		return nil, errors.New("Bad Request")
	}

	offset = (page - 1) * 50

	rows, err := d.DB.Query(queryDocStmt, offset)
	itemsList := make([]*better.Recipe, 0)

	if err != nil || rows == nil {
		log.Println(err)
		return nil, err
	}

	defer rows.Close()
	for rows.Next() {
		var item Recipe

		if err := rows.Scan(&item.ID, &item.Title); err != nil {
			// Check for a scan error.
			// Query rows will be closed with defer.
			log.Println(err)
			continue
		}

		itemsList = append(itemsList, &item)

	}

	return itemsList, nil

}
```

So this is a step in the right direction. With that in place we can go and edit our `handleListRecipes` func. 

```go
func (h *Handler) handleListRecipes(w http.ResponseWriter, r *http.Request) {
	page, err := strconv.Atoi(r.URL.Query().Get("page"))
	if err != nil {
		http.Error(w, "invalid page", http.StatusBadRequest)
		return
	}

	items, err := h.RecipeStore.ListRecipes(page)
	if err != nil {
		log.Print("http error", err)
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}

	if err := json.NewEncoder(w).Encode(items); err != nil {
		log.Print("http json encoding error", err)
	}
}
```

As you can see  this made our handler function way cleaner, more readable and more  concise. We are still not there yet though. In order to be able to mock  (replace) our `RecipeStore` in our test we need to create an  Interface. If the only thing you take away from this is to use  Interfaces more I will consider this a success. 

## Interfaces, Interfaces, Interfaces

So let's create our interface so we can mock out our `RecipeStore` when testing.

```go
type RecipeService interface {
	ListRecipes(page int) ([]*Recipe, error)
}
```

Super simple. But super powerful.

With that, just change our Handler to take a `RecipeService` interface instead of the concrete `RecipeStore`. And voila we can now get started on writing our test. 

```go
items, err := h.RecipeService.ListRecipes(page)
```

## Tests

So now that we have our code refactored we can start with our tests. It is now going to be super easy to replace our actual Store with a mocked  service that we can now control in our tests. No more database calls. 

First let's create a new `mock` package to hold all our mocks and create a `recipe_service.go` file in there with the following.

```go
package mock

import (
	better "github.com/mitchdennett/tests-make-your-code-inherently-better"
)

type MockRecipeService struct {
	ListRecipesFunc func(page int) ([]*better.Recipe, error)
}

func (s *MockRecipeService) ListRecipes(page int) ([]*better.Recipe, error) {
	return s.ListRecipesFunc(page)
}
```

This mock allows our test to inject a function into it to test different outcomes. Now we can write our first test. 

```go
func TestListRecipes(t *testing.T) {
	req, err := http.NewRequest("GET", "/recipes?page=1", nil)
	if err != nil {
		t.Fatal(err)
	}

	var store mock.MockRecipeService
	store.ListRecipesFunc = func(page int) ([]*better.Recipe, error) {
		return []*better.Recipe{{ID: 1, Title: "Pasta"}}, nil
	}
	handler := Handler{RecipeService: &store}

	rr := httptest.NewRecorder()

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}
}
```

## Why is this better?

It's better for multiple reasons. The obvious one is that we can now  correctly test our code which is obviously better. Secondly, our code is more concise, clear and readable. Finally, we have removed our  dependency with the database from all but one package. Should we ever  need to change our data store we can simply implement the `RecipeService` interface and nothing else. 

Now there is definitely more we can do to this code to improve it. But by  simply focusing on testing we inherently have to write better and  cleaner code. Check out this Github Repo with the final code https://github.com/mitchdennett/tests-make-your-code-inherently-better. 

# Join the Newsletter

Subscribe to get our latest content by email. We won't send you spam. Unsubscribe at any time.

