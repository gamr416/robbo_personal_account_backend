package main

import (
	"github.com/ory/dockertest/v3"
	"github.com/ory/dockertest/v3/docker"
	"github.com/skinnykaen/robbo_student_personal_account.git/package/config"
	"github.com/skinnykaen/robbo_student_personal_account.git/package/courses/gateway"
	"github.com/skinnykaen/robbo_student_personal_account.git/package/db_client"
	"github.com/skinnykaen/robbo_student_personal_account.git/package/models"
	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"testing"
	"time"
)

func initTestPostgresDb(testPostgresClient db_client.PostgresClient) (err error) {
	err = testPostgresClient.Migrate()
	return
}

func SetupTestPostgresClient() (testPostgresClient db_client.PostgresClient, err error) {
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold:             time.Second, // Slow SQL threshold
			LogLevel:                  logger.Info, // Log level
			IgnoreRecordNotFoundError: false,       // Ignore ErrRecordNotFound error for logger
			Colorful:                  true,        // Disable color
		},
	)
	db, err := gorm.Open(postgres.Open(viper.GetString("postgres.postgresDsn")), &gorm.Config{Logger: newLogger})
	testPostgresClient = db_client.PostgresClient{
		Db: db,
	}
	return
}

func TestMain(m *testing.M) {
	if err := config.InitForTests(); err != nil {
		log.Fatalf("%s", err.Error())
	}
	pool, poolErr := dockertest.NewPool("")
	if poolErr != nil {
		log.Fatalf("Could not construct pool: %s", poolErr)
	}

	if err := pool.Client.Ping(); err != nil {
		log.Fatalf("Could not connect to Docker: %s", err)
	}

	opts := dockertest.RunOptions{
		Name:         viper.GetString("docker.containerName"),
		Repository:   "postgres",
		Tag:          "13",
		Env:          viper.GetStringSlice("docker.environment"),
		ExposedPorts: []string{"5432"},
		PortBindings: map[docker.Port][]docker.PortBinding{
			"5432": {
				{HostIP: "localhost", HostPort: "5433"},
			},
		},
	}

	resource, err := pool.RunWithOptions(&opts,
		func(config *docker.HostConfig) {
			config.AutoRemove = true
			config.RestartPolicy = docker.RestartPolicy{
				Name: "no",
			}
		},
	)
	if err != nil {
		log.Fatalf("Could not start resource: %s", err)
	}

	log.Println("Connecting to test database on url: ", viper.GetString("postgres.postgresDsn"))
	var testPostgresClient db_client.PostgresClient
	if err = pool.Retry(func() error {
		testPostgresClient, err = SetupTestPostgresClient()
		if err != nil {
			log.Println("Test database not ready yet (it is booting up, wait for a few tries)...")
			return err
		}
		db, sqlErr := testPostgresClient.Db.DB()
		if sqlErr != nil {
			return sqlErr
		}
		return db.Ping()
	}); err != nil {
		log.Fatalf("Could not connect to docker: %s", err)
	}
	if err = initTestPostgresDb(testPostgresClient); err != nil {
		log.Fatalf("Could not init test database: %s", err)
	}
	//Run tests
	code := m.Run()

	if err = pool.Purge(resource); err != nil {
		log.Fatalf("Could not purge resource: %s", err)
	}
	os.Exit(code)
}

func TestHelloWorld(t *testing.T) {
	client, _ := SetupTestPostgresClient()
	gatewaycourse := gateway.SetupCoursesGateway(client)
	id, _ := gatewaycourse.CreateCourse(&models.CourseCore{
		ID:               "1",
		BlocksUrl:        "test",
		Effort:           "",
		EnrollmentStart:  time.Time{},
		EnrollmentEnd:    time.Time{},
		End:              time.Time{},
		Name:             "",
		Number:           "",
		Org:              "",
		ShortDescription: "",
		Start:            time.Time{},
		StartDisplay:     "",
		StartType:        "",
		Pacing:           "",
		MobileAvailable:  false,
		Hidden:           false,
		InvitationOnly:   false,
		CourseID:         "1",
		MediaID:          "1",
		Media: models.CourseApiMediaCollectionCore{
			ID: "1",
			BannerImage: models.AbsoluteMediaCore{
				ID:                         "1",
				Uri:                        "",
				UriAbsolute:                "",
				CourseApiMediaCollectionID: "1",
			},
			CourseImage: models.MediaCore{
				ID:                         "1",
				Uri:                        "",
				CourseApiMediaCollectionID: "",
			},
			CourseVideo: models.MediaCore{
				ID:                         "2",
				Uri:                        "",
				CourseApiMediaCollectionID: "",
			},
			Image: models.ImageCore{
				ID:                         "1",
				Raw:                        "",
				Small:                      "",
				Large:                      "",
				CourseApiMediaCollectionID: "",
			},
			CourseID: "1",
		},
	})
	log.Println(id)
}
