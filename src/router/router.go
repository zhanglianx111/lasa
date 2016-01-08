package router

import (
	"handlers"

	//	"github.com/drone/routes"

	"github.com/go-martini/martini"
)

var Routers martini.Router

func init() {
	Routers = martini.NewRouter()
	/* method GET */
	/* jenkins */
	Routers.Get("/api/info", handlers.HandlerGetInfo)
	/* register */
	Routers.Get("/register", handlers.HandlerRegister)
	Routers.Post("/register", handlers.HandlerRegister)

	/* login */
	Routers.Get("/login", handlers.HandlerLogin)
	Routers.Post("/login", handlers.HandlerLogin)
	/* logout */
	Routers.Post("/logout", handlers.HandlerLogout)
	/* dashboard */
	Routers.Get("/dashboard", handlers.HandlerDashboard)
	/* job */
	Routers.Get("/", handlers.HandlerDefault)
	Routers.Get("/api/job/alljobs", handlers.HandlerGetAllJobs)
	Routers.Get("/api/job/ajob/:jobid", handlers.HandlerGetJob)
	Routers.Get("/api/job/isrunning/:jobid", handlers.HandlerJobRunning)
	Routers.Get("/api/job/:jobid/allbuilds", handlers.HandlerGetAllBuildIds)
	Routers.Get("/api/job/config/:jobid", handlers.HandlerJobConfig)
	Routers.Post("/api/job/config/:jobid", handlers.HandlerJobConfig)
	/* node */
	Routers.Get("/api/node/allnodes", handlers.HandlerGetAllNodes)
	Routers.Get("/api/node/anode/:nodeid", handlers.HandlerGetNode)
	Routers.Get("/api/node/nodeinfo/:nodeid", handlers.HandlerNodeInfo)
	Routers.Get("/isjnlpagent/:nodeid", handlers.HandlerIsJnlpAgent)
	Routers.Get("/api/node/isidle/:nodeid", handlers.HandlerIsIdle)
	Routers.Get("/api/node/isonline/:nodeid", handlers.HandlerIsOnline)
	Routers.Get("/api/node/log/:nodeid", handlers.HandlerGetLogText)
	/* queue */
	Routers.Get("/queue", handlers.HandlerGetQueue)
	Routers.Get("/queueurl", handlers.HandlerGetQueueUrl)
	/* view */
	Routers.Get("/api/view/allviews", handlers.HandlerGetAllViews)
	Routers.Get("/api/view/:viewid", handlers.HandlerGetView)
	/* plugin */
	Routers.Get("/api/plugins/:depth", handlers.HandlerGetPlugins)
	/* log */
	Routers.Get("/api/log/level", handlers.HandlerGetLogLevel)
	/* build */
	Routers.Get("/api/build/buildlog/:jobid", handlers.HandlerBuildConsoleOutput)
	Routers.Get("/api/build/result/:jobid", handlers.HandlerGetBuildResult)
	/* method POST */
	/* job */
	//Routers.Post("/api/job/create/:jobid", handlers.HandlerCreateJob)
	Routers.Get("/createjob", handlers.HandlerCreateJob)
	Routers.Post("/createjob", handlers.HandlerCreateJob)
	Routers.Post("/api/job/delete/:jobid", handlers.HandlerDeleteJob)
	Routers.Post("/api/job/disable/:jobid", handlers.HandlerDisableJob)
	Routers.Post("/api/job/enable/:jobid", handlers.HandlerEnableJob)
	Routers.Post("/api/job/rename/:oldjobid/:newjobid", handlers.HandlerRenameJob)
	Routers.Post("/api/job/build/:jobid", handlers.HandlerBuildJob)
	Routers.Post("/api/job/stopbuild/:jobid/:number", handlers.HandlerStopBuild)
	//Routers.Post("/api/job/copy/:from/:newname", handlers.HandlerCopyJob)
	/* view */
	Routers.Post("/api/view/create/:name/:type", handlers.HandlerCreateView)
	Routers.Post("/api/view/addjob/:name/:jobid", handlers.HandlerViewAddJob)
	Routers.Post("/api/view/deletejob/:name/:jobid", handlers.HandlerViewDeleteJob)
	/* node */
	Routers.Post("/api/node/create/:nodeid/:numexecutors/:description/:remotefs", handlers.HandlerAddNode)
	Routers.Post("/api/node/delete/:nodeid", handlers.HandlerDeleteNode)
	//Routers.Get("/toggletempoffline/:nodeid", handlers.HandlerNodeToggleTemporarilyOffline)
	Routers.Post("/api/node/online/:nodeid", handlers.HandlerSetOnline)
	Routers.Post("/api/node/offline/:nodeid", handlers.HandlerSetOffline)
	Routers.Post("/api/node/launch/:nodeid", handlers.HandlersLaunchNode)
	Routers.Post("/api/node/disconnect/:nodeid", handlers.HandlerDisconnect)
	/* log */
	Routers.Post("/api/log/set/:level", handlers.HandlerSetLogLevel)
}
