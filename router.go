package main

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-playground/validator/v10"
	"log"
	"net/http"
	"webapp/auth"
	"webapp/config"
	"webapp/handler"
	dailyRecord "webapp/handler/daily_record"
	"webapp/handler/user"
	userDefinedRecord "webapp/handler/user_defined_record"
	"webapp/lib"
	"webapp/service"
	"webapp/store"
)

func NewApp() (http.Handler, error) {
	cfg, err := config.New()
	if err != nil {
		log.Fatalf("Error loading config: %s", err.Error())
	}
	log.Printf("Running on go app environment: %s\n", cfg.GoEnv)

	db, err := store.InitDB(cfg)
	if err != nil {
		log.Fatalf("Error loading db: %s", err.Error())
	}

	ac, err := auth.NewNoopClient("test-user-1", "test-user-email")
	if err != nil {
		log.Fatalf("Error loading firebase: %s", err.Error())
	}
	// todo: uncomment this when you want to use firebase auth
	//ac, err := auth.NewFirebaseAdminClient(
	//	context.Background(),
	//	cfg.GoEnv,
	//	cfg.GoogleProjectId,
	//	cfg.FirebaseCredentialJSON)
	//if err != nil {
	//	log.Fatalf("Error loading firebase: %s", err.Error())
	//}
	//

	recaptchaCli, err := lib.NewRecaptchaCli(cfg.RecaptchaSecretKey)
	//recaptchaCli, err := lib.NewNoopRecaptchaCli(cfg.RecaptchaSecretKey)
	if err != nil {
		log.Fatalf("Error loading recaptcha: %s", err.Error())
	}

	v := validator.New()

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(handler.CorsMiddleware(cfg.GoEnv))

	r.MethodFunc("GET", "/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`{"status": "ok"}`))
	})

	userRepository := store.NewUserRepository(db, ac)
	userService := service.NewUserService(userRepository, recaptchaCli)

	getUserHandler := user.NewGetUserHandler(userService)
	registerUserHandler := user.NewRegisterUserHandler(userService, v)
	updateUserHandler := user.NewUpdateUserHandler(userService)
	r.Route("/user", func(inner chi.Router) {
		inner.Use(handler.AuthMiddleware(ac))
		inner.Get("/", getUserHandler.ServeHTTP)
		inner.Post("/", registerUserHandler.ServeHTTP)
		inner.Put("/", updateUserHandler.ServeHTTP)
	})
	dailyRecordRepository := store.NewDailyRecordRepository(db, ac)
	dailyRecordService := service.NewDailyRecordService(dailyRecordRepository)

	getDailyRecordHandler := dailyRecord.NewGetDailyRecordHandler(dailyRecordService)
	listDailyRecordHandler := dailyRecord.NewListDailyRecordHandler(dailyRecordService)
	registerDailyRecordHandler := dailyRecord.NewRegisterDailyRecordHandler(dailyRecordService, v)
	updateDailyRecordHandler := dailyRecord.NewUpdateDailyRecordHandler(dailyRecordService, v)
	deleteDailyRecordHandler := dailyRecord.NewDeleteDailyRecordHandler(dailyRecordService, v)
	r.Route("/daily_record", func(inner chi.Router) {
		inner.Use(handler.AuthMiddleware(ac))
		inner.Get("/", listDailyRecordHandler.ServeHTTP)
		inner.Get("/{id}", getDailyRecordHandler.ServeHTTP)
		inner.Post("/", registerDailyRecordHandler.ServeHTTP)
		inner.Put("/", updateDailyRecordHandler.ServeHTTP)
		inner.Delete("/", deleteDailyRecordHandler.ServeHTTP)
	})

	userDefinedRecordRepository := store.NewUserDefinedRecordRepository(db, ac)
	userDefinedRecordService := service.NewUserDefinedRecordService(userDefinedRecordRepository)

	getUserDefinedRecordHandler := userDefinedRecord.NewGetUserDefinedRecordHandler(userDefinedRecordService)
	listUserDefinedRecordHandler := userDefinedRecord.NewListUserDefinedRecordHandler(userDefinedRecordService)
	registerUserDefinedRecordHandler := userDefinedRecord.NewRegisterUserDefinedRecordHandler(userDefinedRecordService, v)
	deleteUserDefinedRecordHandler := userDefinedRecord.NewDeleteUserDefinedRecordHandler(userDefinedRecordService, v)
	r.Route("/user_defined_record", func(inner chi.Router) {
		inner.Use(handler.AuthMiddleware(ac))
		inner.Get("/{id}", getUserDefinedRecordHandler.ServeHTTP)
		inner.Get("/", listUserDefinedRecordHandler.ServeHTTP)
		inner.Post("/", registerUserDefinedRecordHandler.ServeHTTP)
		inner.Delete("/", deleteUserDefinedRecordHandler.ServeHTTP)
	})

	return r, nil
}
