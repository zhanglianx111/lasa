package router

import (
	"handlers"

	"github.com/drone/routes"
)

var Mux *routes.RouteMux

func init() {
	Mux = routes.New()
	/* method GET */
	/* jenkins */
	Mux.Get("/api/info", handlers.HandlerGetInfo)
	/* job */
	Mux.Get("/", handlers.HandlerDefault)
	Mux.Get("/api/job/alljobs", handlers.HandlerGetAllJobs)
	Mux.Get("/api/job/ajob/:jobid", handlers.HandlerGetJob)
	Mux.Get("/api/job/isrunning/:jobid", handlers.HandlerJobRunning)
	Mux.Get("/api/job/:jobid/allbuilds", handlers.HandlerGetAllBuildIds)
	Mux.Get("/api/job/config/:jobid", handlers.HandlerJobConfig)
	Mux.Get("/api/job/buildlog/:jobid/:number", handlers.HandlerBuildConsoleOutput)
	/* node */
	Mux.Get("/api/node/allnodes", handlers.HandlerGetAllNodes)
	Mux.Get("/api/node/anode/:nodeid", handlers.HandlerGetNode)
	Mux.Get("/api/node/nodeinfo/:nodeid", handlers.HandlerNodeInfo)
	Mux.Get("/isjnlpagent/:nodeid", handlers.HandlerIsJnlpAgent)
	Mux.Get("/api/node/isidle/:nodeid", handlers.HandlerIsIdle)
	Mux.Get("/api/node/isonline/:nodeid", handlers.HandlerIsOnline)
	Mux.Get("/api/node/log/:nodeid", handlers.HandlerGetLogText)
	/* queue */
	Mux.Get("/queue", handlers.HandlerGetQueue)
	Mux.Get("/queueurl", handlers.HandlerGetQueueUrl)
	/* view */
	Mux.Get("/api/view/allviews", handlers.HandlerGetAllViews)
	Mux.Get("/api/view/:viewid", handlers.HandlerGetView)
	/* plugin */
	Mux.Get("/api/plugins/:depth", handlers.HandlerGetPlugins)
	/* log */
	Mux.Get("/api/log/level", handlers.HandlerGetLogLevel)
	/* build */
	//Mux.Get("/api///")
	/* method POST */
	/* job */
	Mux.Post("/api/job/create/:jobid", handlers.HandlerCreateJob)
	Mux.Post("/api/job/delete/:jobid", handlers.HandlerDeleteJob)
	Mux.Post("/api/job/disable/:jobid", handlers.HandlerDisableJob)
	Mux.Post("/api/job/enable/:jobid", handlers.HandlerEnableJob)
	Mux.Post("/api/job/rename/:oldjobid/:newjobid", handlers.HandlerRenameJob)
	Mux.Post("/api/job/build/:jobid", handlers.HandlerBuildJob)
	Mux.Post("/api/job/stopbuild/:jobid/:number", handlers.HandlerStopBuild)
	//Mux.Post("/api/job/copy/:from/:newname", handlers.HandlerCopyJob)
	/* view */
	Mux.Post("/api/view/create/:name/:type", handlers.HandlerCreateView)
	Mux.Post("/api/view/addjob/:name/:jobid", handlers.HandlerViewAddJob)
	Mux.Post("/api/view/deletejob/:name/:jobid", handlers.HandlerViewDeleteJob)
	/* node */
	Mux.Post("/api/node/create/:nodeid/:numexecutors/:description/:remotefs", handlers.HandlerAddNode)
	Mux.Post("/api/node/delete/:nodeid", handlers.HandlerDeleteNode)
	//Mux.Get("/toggletempoffline/:nodeid", handlers.HandlerNodeToggleTemporarilyOffline)
	Mux.Post("/api/node/online/:nodeid", handlers.HandlerSetOnline)
	Mux.Post("/api/node/offline/:nodeid", handlers.HandlerSetOffline)
	Mux.Post("/api/node/launch/:nodeid", handlers.HandlersLaunchNode)
	Mux.Post("/api/node/disconnect/:nodeid", handlers.HandlerDisconnect)
	/* log */
	Mux.Post("/api/log/set/:level", handlers.HandlerSetLogLevel)
}
