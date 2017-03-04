package routers

import (
	"github.com/astaxie/beego"
)

func init() {

	beego.GlobalControllerRouter["mau-api/controllers:MauArticleController"] = append(beego.GlobalControllerRouter["mau-api/controllers:MauArticleController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["mau-api/controllers:MauArticleController"] = append(beego.GlobalControllerRouter["mau-api/controllers:MauArticleController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["mau-api/controllers:MauArticleController"] = append(beego.GlobalControllerRouter["mau-api/controllers:MauArticleController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["mau-api/controllers:MauArticleController"] = append(beego.GlobalControllerRouter["mau-api/controllers:MauArticleController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["mau-api/controllers:MauArticleController"] = append(beego.GlobalControllerRouter["mau-api/controllers:MauArticleController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["mau-api/controllers:MauArticleTagController"] = append(beego.GlobalControllerRouter["mau-api/controllers:MauArticleTagController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["mau-api/controllers:MauArticleTagController"] = append(beego.GlobalControllerRouter["mau-api/controllers:MauArticleTagController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["mau-api/controllers:MauArticleTagController"] = append(beego.GlobalControllerRouter["mau-api/controllers:MauArticleTagController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["mau-api/controllers:MauArticleTagController"] = append(beego.GlobalControllerRouter["mau-api/controllers:MauArticleTagController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["mau-api/controllers:MauArticleTagController"] = append(beego.GlobalControllerRouter["mau-api/controllers:MauArticleTagController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["mau-api/controllers:MauAskController"] = append(beego.GlobalControllerRouter["mau-api/controllers:MauAskController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["mau-api/controllers:MauAskController"] = append(beego.GlobalControllerRouter["mau-api/controllers:MauAskController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["mau-api/controllers:MauAskController"] = append(beego.GlobalControllerRouter["mau-api/controllers:MauAskController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["mau-api/controllers:MauAskController"] = append(beego.GlobalControllerRouter["mau-api/controllers:MauAskController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["mau-api/controllers:MauAskController"] = append(beego.GlobalControllerRouter["mau-api/controllers:MauAskController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["mau-api/controllers:MauAttentionController"] = append(beego.GlobalControllerRouter["mau-api/controllers:MauAttentionController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["mau-api/controllers:MauAttentionController"] = append(beego.GlobalControllerRouter["mau-api/controllers:MauAttentionController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["mau-api/controllers:MauAttentionController"] = append(beego.GlobalControllerRouter["mau-api/controllers:MauAttentionController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["mau-api/controllers:MauAttentionController"] = append(beego.GlobalControllerRouter["mau-api/controllers:MauAttentionController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["mau-api/controllers:MauAttentionController"] = append(beego.GlobalControllerRouter["mau-api/controllers:MauAttentionController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["mau-api/controllers:MauCommentController"] = append(beego.GlobalControllerRouter["mau-api/controllers:MauCommentController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["mau-api/controllers:MauCommentController"] = append(beego.GlobalControllerRouter["mau-api/controllers:MauCommentController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["mau-api/controllers:MauCommentController"] = append(beego.GlobalControllerRouter["mau-api/controllers:MauCommentController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["mau-api/controllers:MauCommentController"] = append(beego.GlobalControllerRouter["mau-api/controllers:MauCommentController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["mau-api/controllers:MauCommentController"] = append(beego.GlobalControllerRouter["mau-api/controllers:MauCommentController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["mau-api/controllers:MauFriendController"] = append(beego.GlobalControllerRouter["mau-api/controllers:MauFriendController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["mau-api/controllers:MauFriendController"] = append(beego.GlobalControllerRouter["mau-api/controllers:MauFriendController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["mau-api/controllers:MauFriendController"] = append(beego.GlobalControllerRouter["mau-api/controllers:MauFriendController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["mau-api/controllers:MauFriendController"] = append(beego.GlobalControllerRouter["mau-api/controllers:MauFriendController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["mau-api/controllers:MauFriendController"] = append(beego.GlobalControllerRouter["mau-api/controllers:MauFriendController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["mau-api/controllers:MauLabelController"] = append(beego.GlobalControllerRouter["mau-api/controllers:MauLabelController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["mau-api/controllers:MauLabelController"] = append(beego.GlobalControllerRouter["mau-api/controllers:MauLabelController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["mau-api/controllers:MauLabelController"] = append(beego.GlobalControllerRouter["mau-api/controllers:MauLabelController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["mau-api/controllers:MauLabelController"] = append(beego.GlobalControllerRouter["mau-api/controllers:MauLabelController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["mau-api/controllers:MauLabelController"] = append(beego.GlobalControllerRouter["mau-api/controllers:MauLabelController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["mau-api/controllers:MauLikeController"] = append(beego.GlobalControllerRouter["mau-api/controllers:MauLikeController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["mau-api/controllers:MauLikeController"] = append(beego.GlobalControllerRouter["mau-api/controllers:MauLikeController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["mau-api/controllers:MauLikeController"] = append(beego.GlobalControllerRouter["mau-api/controllers:MauLikeController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["mau-api/controllers:MauLikeController"] = append(beego.GlobalControllerRouter["mau-api/controllers:MauLikeController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["mau-api/controllers:MauLikeController"] = append(beego.GlobalControllerRouter["mau-api/controllers:MauLikeController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["mau-api/controllers:MauSpecialController"] = append(beego.GlobalControllerRouter["mau-api/controllers:MauSpecialController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["mau-api/controllers:MauSpecialController"] = append(beego.GlobalControllerRouter["mau-api/controllers:MauSpecialController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["mau-api/controllers:MauSpecialController"] = append(beego.GlobalControllerRouter["mau-api/controllers:MauSpecialController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["mau-api/controllers:MauSpecialController"] = append(beego.GlobalControllerRouter["mau-api/controllers:MauSpecialController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["mau-api/controllers:MauSpecialController"] = append(beego.GlobalControllerRouter["mau-api/controllers:MauSpecialController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["mau-api/controllers:MauSubscriptionController"] = append(beego.GlobalControllerRouter["mau-api/controllers:MauSubscriptionController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["mau-api/controllers:MauSubscriptionController"] = append(beego.GlobalControllerRouter["mau-api/controllers:MauSubscriptionController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["mau-api/controllers:MauSubscriptionController"] = append(beego.GlobalControllerRouter["mau-api/controllers:MauSubscriptionController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["mau-api/controllers:MauSubscriptionController"] = append(beego.GlobalControllerRouter["mau-api/controllers:MauSubscriptionController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["mau-api/controllers:MauSubscriptionController"] = append(beego.GlobalControllerRouter["mau-api/controllers:MauSubscriptionController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["mau-api/controllers:MauUserController"] = append(beego.GlobalControllerRouter["mau-api/controllers:MauUserController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["mau-api/controllers:MauUserController"] = append(beego.GlobalControllerRouter["mau-api/controllers:MauUserController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["mau-api/controllers:MauUserController"] = append(beego.GlobalControllerRouter["mau-api/controllers:MauUserController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["mau-api/controllers:MauUserController"] = append(beego.GlobalControllerRouter["mau-api/controllers:MauUserController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["mau-api/controllers:MauUserController"] = append(beego.GlobalControllerRouter["mau-api/controllers:MauUserController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

}
