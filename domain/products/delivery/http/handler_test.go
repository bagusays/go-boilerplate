package http

import (
	"go-boilerplate/database/migration"
	"go-boilerplate/shared/config"
	"go-boilerplate/shared/database"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	productRepository "go-boilerplate/domain/products/repository"
	productService "go-boilerplate/domain/products/service"

	"github.com/labstack/echo"
	"github.com/stretchr/testify/suite"
	"gorm.io/gorm"
)

type testSuite struct {
	suite.Suite
	dbConn      *gorm.DB
	dbMigration *migration.Migration
}

func TestSuite(t *testing.T) {
	suite.Run(t, new(testSuite))
}

func (s *testSuite) SetupSuite() {
	os.Chdir("../../../../")
	config.LoadConfig()
	dbConn, err := database.OpenMysqlConn()
	s.Require().NoError(err)
	s.dbConn = dbConn
	s.dbMigration = migration.NewMigration()
}

func (s *testSuite) TearDownSuite() {

}

func (s *testSuite) SetupTest() {
	s.dbMigration.Up()
}

func (s *testSuite) TearDownTest() {
	s.dbMigration.Down()
}

func (s *testSuite) TestGetProducts() {
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	res := rec.Result()
	defer res.Body.Close()

	productRepo := productRepository.NewProductsRepository(s.dbConn)
	productService := productService.NewProductService(productRepo)
	h := NewProductHandler(e, productService)

	// Assertions
	if s.NoError(h.GetProducts(c)) {
		s.Equal(http.StatusOK, rec.Code)
	}
}
