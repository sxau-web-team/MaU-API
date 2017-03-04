// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"mau-api/controllers"

	"github.com/astaxie/beego"
)

func init() {
	ns := beego.NewNamespace("/v1",

		beego.NSNamespace("/mau_article",
			beego.NSInclude(
				&controllers.MauArticleController{},
			),
		),

		beego.NSNamespace("/mau_article_tag",
			beego.NSInclude(
				&controllers.MauArticleTagController{},
			),
		),

		beego.NSNamespace("/mau_ask",
			beego.NSInclude(
				&controllers.MauAskController{},
			),
		),

		beego.NSNamespace("/mau_attention",
			beego.NSInclude(
				&controllers.MauAttentionController{},
			),
		),

		beego.NSNamespace("/mau_comment",
			beego.NSInclude(
				&controllers.MauCommentController{},
			),
		),

		beego.NSNamespace("/mau_friend",
			beego.NSInclude(
				&controllers.MauFriendController{},
			),
		),

		beego.NSNamespace("/mau_label",
			beego.NSInclude(
				&controllers.MauLabelController{},
			),
		),

		beego.NSNamespace("/mau_like",
			beego.NSInclude(
				&controllers.MauLikeController{},
			),
		),

		beego.NSNamespace("/mau_special",
			beego.NSInclude(
				&controllers.MauSpecialController{},
			),
		),

		beego.NSNamespace("/mau_subscription",
			beego.NSInclude(
				&controllers.MauSubscriptionController{},
			),
		),

		beego.NSNamespace("/mau_user",
			beego.NSInclude(
				&controllers.MauUserController{},
			),
			beego.NSRouter("/register", &controllers.MauUserController{}, "post:Register"),
		),
	)
	beego.AddNamespace(ns)
}
