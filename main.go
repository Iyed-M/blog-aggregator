package main

func main() {
	srv := newSrv()
	srv.addRoutes()

	if err := srv.app.Listen(":" + srv.apiCfg.PORT); err != nil {
		panic(err)
	}
}
