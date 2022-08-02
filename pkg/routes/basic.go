package routes

import (
	"net/http"
	"strconv"

	"github.com/ShariqT/csfood4/pkg/core"
	"github.com/ShariqT/csfood4/pkg/utils"
	"github.com/flosch/pongo2"
	echosession "github.com/go-session/echo-session"
	"github.com/labstack/echo"
)

func IndexHandler(c echo.Context) error {
	cc := c.(*utils.CommonstockContext)
	data := pongo2.Context{}
	db, err := cc.DB()
	if err != nil {
		data["error"] = err
		return c.Render(http.StatusInternalServerError, "./pkg/templates/error.html", data)
	}
	markets, err := core.GetMarkets(db)
	if err != nil {
		data["error"] = err
		return c.Render(http.StatusInternalServerError, "./pkg/templates/error.html", data)

	}
	data["markets"] = markets
	return c.Render(http.StatusOK, "./pkg/templates/index.html", data)
}

func ShowLogin(c echo.Context) error {
	return c.Render(http.StatusOK, "./pkg/templates/login.html", nil)
}

func LoginHandler(c echo.Context) error {
	cc := c.(*utils.CommonstockContext)
	data := pongo2.Context{}
	username := c.FormValue("username")
	password := c.FormValue("password")
	db, err := cc.DB()
	if err != nil {
		data["error"] = err
		return c.Render(http.StatusInternalServerError, "./pkg/templates/error.html", data)
	}
	user, err := core.AuthenticateUser(username, password, db)
	if err != nil {
		data["error"] = err
		return c.Render(http.StatusInternalServerError, "./pkg/templates/error.html", data)
	}
	if user == nil {
		data["error"] = "Wrong username/password combo"
		return c.Render(http.StatusInternalServerError, "./pkg/templates/error.html", data)
	}

	data["user"] = user
	markets, err := core.GetMarkets(db)
	if err != nil {
		data["error"] = err
		return c.Render(http.StatusInternalServerError, "./pkg/templates/error.html", data)
	}
	store := echosession.FromContext(c)
	store.Set("user", user)

	data["markets"] = markets
	cart := core.NewCart(&core.Order{})
	data["currentCart"] = cart
	store.Set("cart", cart)
	err = store.Save()
	if err != nil {
		c.Logger().Error(err)
	}
	// save to session
	return c.Render(http.StatusOK, "index.html", data)
}

func AddOrderToCart(c echo.Context) error {
	store := echosession.FromContext(c)
	cc := c.(*utils.CommonstockContext)
	data := make(map[string]string)
	currentCartRawValue, ok := store.Get("currentCart")
	if !ok {
		data["error"] = "Could not access customer cart"
		return c.JSON(500, data)
	}
	currentCart := currentCartRawValue.(*core.Cart)
	db, err := cc.DB()
	if err != nil {
		data["error"] = "Could not get access to database"
		return c.JSON(http.StatusInternalServerError, data)
	}

	newItem, err := core.GetProductVariationById(c.FormValue("product_id"), db)
	if err != nil {
		data["error"] = "Error accessing product variation."
		return c.JSON(http.StatusInternalServerError, data)
	}
	newQty, err := strconv.ParseFloat(c.FormValue("quantity"), 64)
	if err != nil {
		data["error"] = "Error parsing new quantity for item"
		return c.JSON(http.StatusInternalServerError, data)
	}
	currentCart.AddOrderItem(newItem, newQty)
	store.Set("currentCart", currentCart)
	err = store.Save()
	if err != nil {
		data["error"] = "Did not add item to cart. Error processing reqeust"
		return c.JSON(http.StatusInternalServerError, data)
	}
	data["status"] = "OK"
	return c.JSON(http.StatusOK, data)
}

// func EditdOrderToCart(c echo.Context) error {
// 	currentCart := sess.Values["currentCart"]
// 	newQty = c.FormValue("quantity")
// 	Item = core.models.GetProductVariationById(c.FormValue("product_id"))
// 	currentCart.SaleItems[Item] = newQty
// 	currentCart.Recalcuate()
// 	sess.Values["currentCart"] = currentCart
// 	data := make(map[string]string)
// 	data["status"] = "OK"
// 	return c.JSON(http.StatusOK, data)
// }

// func DeleteOrderToCart(c echo.Context) error {
// 	currentCart := sess.Values["currentCart"]
// 	Item = core.models.GetProductVariationById(c.FormValue("product_id"))
// 	currentCart.RemoveItem(Item)
// 	currentCart.RecalcuateSubTotal()
// 	sess.Values["currentCart"] = currentCart
// 	data := make(map[string]string)
// 	data["status"] = "OK"
// 	return c.JSON(http.StatusOK, data)
// }

// func ViewCheckout(c echo.Context) error {
// 	currentCart := sess.Values["currentCart"]
// 	fees := utils.CustomFees{}
// 	currentCart.RecalcuateSubTotal()
// 	currentCart.AddFees(fees)
// 	sess.Values["currentCart"] = currentCart
// 	data := pongo2.Context{}
// 	data["currentCart"] = currentCart
// 	return c.Render(http.StatusOK, "checkout.html", data)

// }

// func ProcessCheckout(c echo.Context) error {
// 	cc := c.(*CommonstockContext)
// 	currentCart := sess.Values["currentCart"]
// 	payProcessor := &utils.StripePaymentProcessor{}
// 	if txId, err := payProcessor.PayForOrder(currentCart); err != nil {
// 		return err
// 	}
// 	currentCart.Order.StripeTransactionId = txId
// 	currentCart.Order.Save()
// 	data := pongo2.Context{}
// 	data["processedCart"] := currentCart
// 	return c.Render(http.StatusOK, "receipt.html", data)
// }
