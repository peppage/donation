package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"donation/setting"
	"donation/view"

	"github.com/kataras/iris"
	"github.com/stripe/stripe-go"
	"github.com/stripe/stripe-go/charge"
)

func main() {
	setting.Initialize()

	stripe.Key = setting.SecretKey

	iris.Get("/", root)
	iris.Post("/", handlePay)

	iris.Listen(":8000")
}

func root(ctx *iris.Context) {
	p := &view.RootPage{}
	view.WritePageTemplate(ctx.GetRequestCtx(), p)
	ctx.HTML(http.StatusOK, "")
}

func handlePay(ctx *iris.Context) {
	fmt.Println(ctx.FormValueString("stripeToken"))
	fmt.Println(ctx.FormValueString("message"))
	a, err := strconv.ParseUint(ctx.FormValueString("amount"), 10, 64)
	if err != nil {
		ctx.EmitError(http.StatusBadRequest)
	}

	chargeParams := &stripe.ChargeParams{
		Amount:   a * 100,
		Currency: "usd",
		Desc:     "Charge for simple donation example",
	}
	chargeParams.SetSource(ctx.FormValueString("stripeToken"))

	_, err = charge.New(chargeParams)
	if err != nil {
		log.Printf("error " + err.Error())
		ctx.Redirect("/?success=f", http.StatusSeeOther)
		return
	}
	ctx.Redirect("/?success=t", http.StatusSeeOther)
}
