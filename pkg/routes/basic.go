package routes

import (
	"net/http"

	"github.com/ShariqT/csfood4/pkg/core"
	"github.com/ShariqT/csfood4/pkg/utils"
	"github.com/flosch/pongo2"
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
	data["markets"] = markets
	data["currentCart"] = core.NewCart(&core.Order{})
	// save to session
	return c.Render(http.StatusOK, "index.html", data)
}

// func AddOrderToCart(c echo.Context) error {
// 	currentCart := sess.Values["currentCart"]
// 	newItem = core.GetProductVariationById(c.FormValue("product_id"))
// 	currentCart.AddToCart(newItem)
// 	sess.Values["currentCart"] = currentCart
// 	data := make(map[string]string)
// 	data["status"] = "OK"
// 	return c.JSON(http.StatusOK, data)
// }

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
