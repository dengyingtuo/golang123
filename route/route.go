package route

import (
	"github.com/kataras/iris"
	"github.com/shen100/golang123/config"
	"github.com/shen100/golang123/controller/common"
	"github.com/shen100/golang123/controller/auth"
	"github.com/shen100/golang123/controller/category"
	"github.com/shen100/golang123/controller/article"
	"github.com/shen100/golang123/controller/collect"
	"github.com/shen100/golang123/controller/comment"
	"github.com/shen100/golang123/controller/vote"
	"github.com/shen100/golang123/controller/user"
	"github.com/shen100/golang123/controller/message"
	"github.com/shen100/golang123/controller/crawler"
	"github.com/shen100/golang123/controller/baidu"
)

// Route 路由
func Route(app *iris.Application) {
	apiPrefix := config.ServerConfig.APIPrefix

	routes := app.Party(apiPrefix, common.SessShiftExpiration)
	{	
		routes.Post("/signin",                   user.Signin)
		routes.Post("/signup",                   user.Signup)
		routes.Post("/signout",                  user.Signout)
		routes.Post("/active/sendmail",          user.ActiveSendMail)
		routes.Post("/active/:id/:secret",       user.ActiveAccount)
		routes.Post("/reset",                    user.ResetPasswordMail)
		routes.Get("/reset/verify/:id/:secret",  user.VerifyResetPasswordLink)		
		routes.Post("/reset/:id/:secret",        user.ResetPassword)

		routes.Get("/user/info/public/:id",    user.PublicInfo)
		routes.Get("/user/info",               auth.SigninRequired,  
											   user.SecretInfo)
		routes.Get("/user/info/detail",        auth.SigninRequired, 
											   user.InfoDetail)
		routes.Post("/user/update/:field",     auth.ActiveRequired,       
										       user.UpdateInfo)
		routes.Post("/user/password/update",   auth.ActiveRequired,       
											   user.UpdatePassword)
		routes.Get("/user/score/top10",        user.Top10)
		routes.Get("/user/score/top100",       user.Top100)
		routes.Post("/user/career/add",        auth.ActiveRequired,
											   user.AddCareer)
		routes.Post("/user/school/add",        auth.ActiveRequired,
											   user.AddSchool)
		routes.Post("/user/career/delete/:id", auth.ActiveRequired,
											   user.DeleteCareer)
		routes.Post("/user/school/delete/:id", auth.ActiveRequired,
											   user.DeleteSchool)
		routes.Post("/user/updateavatar",      auth.ActiveRequired,
											   user.UpdateAvatar)
		

		routes.Post("/upload",               auth.ActiveRequired,          
											 common.UploadHandler)

		routes.Get("/message/unread",        auth.SigninRequired,  
											 message.Unread)
		routes.Get("/message/unread/count",  auth.SigninRequired,  
											 message.UnreadCount)

		routes.Get("/categories",              category.List)

		routes.Get("/articles",                article.List)
		routes.Get("/articles/user/:userID",   article.UserArticleList)
		routes.Get("/articles/maxcomment",     article.ListMaxComment)
		routes.Get("/articles/maxbrowse",      article.ListMaxBrowse)
		routes.Get("/article/{id:int min(1)}", article.Info)
		routes.Post("/article/create",         auth.ActiveRequired, 
										       article.Create)
		routes.Post("/article/update",         auth.ActiveRequired,    
											   article.Update)
		routes.Post("/article/delete/:id",     auth.ActiveRequired,
											   article.Delete)
		routes.Get("/articles/top",            article.Tops)
		routes.Post("/article/top/:id",        auth.EditorRequired,    
											   article.Top)
		routes.Post("/article/deltop/:id",     auth.EditorRequired,    
											   article.DeleteTop)

		routes.Post("/collect/folder/create",        auth.ActiveRequired,
											         collect.CreateFolder)									   
		routes.Post("/collect/create",               auth.ActiveRequired,
											         collect.CreateCollect)
		routes.Post("/collect/delete/:id",           auth.ActiveRequired,
											         collect.DeleteCollect)
		routes.Get("/collect/folders/:userID",       collect.Folders)
		routes.Get("/collect/folders/source",        auth.ActiveRequired,
												     collect.FoldersWithSource)
		routes.Get("/collects",                      collect.Collects)

		routes.Post("/comment/create",                auth.ActiveRequired,
											          comment.Create)
		routes.Post("/comment/delete/:id",            auth.ActiveRequired,
											          comment.Delete)
		routes.Post("/comment/update",                auth.ActiveRequired,
											          comment.Update)
		routes.Get("/comments/user/:userID",          comment.UserCommentList)
		routes.Get("/comments/:sourceName/:sourceID", comment.SourceComments)

		routes.Get("/votes",                 vote.List)
		routes.Get("/votes/maxbrowse",       vote.ListMaxBrowse)
		routes.Get("/votes/maxcomment",      vote.ListMaxComment)
		routes.Get("/votes/user/:userID",    vote.UserVoteList)
		routes.Post("/vote/create",          auth.EditorRequired,
											 vote.Create)
		routes.Post("/vote/update",          auth.EditorRequired,
											 vote.Update)
		routes.Post("/vote/delete/:id",      auth.EditorRequired,
											 vote.Delete)
		routes.Get("/vote/:id",              vote.Info)
		routes.Post("/vote/item/create",     auth.EditorRequired,
											 vote.CreateVoteItem)
		routes.Post("/vote/item/edit",       auth.EditorRequired,
											 vote.EditVoteItem)
		routes.Post("/vote/item/delete/:id", auth.EditorRequired,
											 vote.DeleteItem)
		routes.Post("/vote/uservote/:id",   auth.ActiveRequired,
											vote.UserVoteVoteItem)
    }

	adminRoutes := app.Party(apiPrefix + "/admin", common.SessShiftExpiration, auth.AdminRequired)
	{
		adminRoutes.Get("/categories",               category.List)
		adminRoutes.Post("/category/create",         category.Create)
		adminRoutes.Post("/category/update",         category.Update)

		adminRoutes.Get("/articles",                 article.AllList)
		adminRoutes.Post("/article/status/update",   article.UpdateStatus)

		adminRoutes.Get("/comments",                   comment.Comments)
		adminRoutes.Put("/comments/update/status/:id", comment.UpdateStatus)

		adminRoutes.Post("/crawl",         crawler.Crawl)
		adminRoutes.Post("/crawl/account", crawler.CreateAccount)
		adminRoutes.Get("/crawl/account",  crawler.CrawlAccount)

		adminRoutes.Post("/pushBaiduLink",  baidu.PushToBaidu)

		adminRoutes.Get("/users", user.AllList)
    }
}