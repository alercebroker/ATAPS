package alercedb

import (
	_ "github.com/jackc/pgx/v5/stdlib"
)

func (suite *AlerceSuite) TestCreateTables() {
	RestoreDatabase(suite.DB, suite)
	err := CreateTables(suite.DB)
	suite.Require().Nil(err)
	suite.assertTableExists("object")
	suite.assertTableExists("detection")
	suite.assertTableExists("non_detection")
}

func (suite *AlerceSuite) TestCreateObjectsTable() {
	RestoreDatabase(suite.DB, suite)
	err := createObjectTable(suite.DB)
	suite.Require().Nil(err)
	suite.assertTableExists("object")
}

func (suite *AlerceSuite) TestCreateObjectsTableIndex() {
	RestoreDatabase(suite.DB, suite)
	err := createObjectTable(suite.DB)
	suite.Require().Nil(err)
	query := `SELECT COUNT(*) FROM pg_indexes WHERE tablename = 'object'`
	var count int
	err = suite.DB.QueryRow(query).Scan(&count)
	suite.Require().Nil(err)
	suite.Require().Equal(5, count)
}

func (suite *AlerceSuite) TestCreateDetectionsTable() {
	RestoreDatabase(suite.DB, suite)
	err := createObjectTable(suite.DB)
	suite.Require().Nil(err)
	err = createDetectionsTable(suite.DB)
	suite.Require().Nil(err)
	suite.assertTableExists("detection")
}

func (suite *AlerceSuite) TestCreateDetectionsTableIndex() {
	RestoreDatabase(suite.DB, suite)
	err := createObjectTable(suite.DB)
	suite.Require().Nil(err)
	err = createDetectionsTable(suite.DB)
	suite.Require().Nil(err)
	query := `SELECT COUNT(*) FROM pg_indexes WHERE tablename = 'detection'`
	var count int
	err = suite.DB.QueryRow(query).Scan(&count)
	suite.Require().Nil(err)
	suite.Require().Equal(2, count)
}

func (suite *AlerceSuite) TestCreateNonDetectionsTable() {
	RestoreDatabase(suite.DB, suite)
	err := createObjectTable(suite.DB)
	suite.Require().Nil(err)
	err = createNonDetectionsTable(suite.DB)
	suite.Require().Nil(err)
	suite.assertTableExists("non_detection")
}

func (suite *AlerceSuite) TestCreateNonDetectionsTableIndex() {
	RestoreDatabase(suite.DB, suite)
	err := createObjectTable(suite.DB)
	suite.Require().Nil(err)
	err = createNonDetectionsTable(suite.DB)
	suite.Require().Nil(err)
	query := `SELECT COUNT(*) FROM pg_indexes WHERE tablename = 'non_detection'`
	var count int
	err = suite.DB.QueryRow(query).Scan(&count)
	suite.Require().Nil(err)
	suite.Require().Equal(2, count)
}

func (suite *AlerceSuite) TestCreateForcedPhotometryTable() {
	RestoreDatabase(suite.DB, suite)
	err := createObjectTable(suite.DB)
	suite.Require().Nil(err)
	err = createForcedPhotometryTable(suite.DB)
	suite.Require().Nil(err)
	suite.assertTableExists("forced_photometry")
}
func (suite *AlerceSuite) TestCreateForcedPhotometryTableIndex() {
	RestoreDatabase(suite.DB, suite)
	err := createObjectTable(suite.DB)
	suite.Require().Nil(err)
	err = createForcedPhotometryTable(suite.DB)
	suite.Require().Nil(err)
	query := `SELECT COUNT(*) FROM pg_indexes WHERE tablename = 'forced_photometry'`
	var count int
	err = suite.DB.QueryRow(query).Scan(&count)
	suite.Require().Nil(err)
	suite.Require().Equal(2, count)
}

func (suite *AlerceSuite) TestCreateFeatureTable() {
	RestoreDatabase(suite.DB, suite)
	err := createObjectTable(suite.DB)
	suite.Require().Nil(err)
	err = createFeaturesTable(suite.DB)
	suite.Require().Nil(err)
	suite.assertTableExists("feature")
}

func (suite *AlerceSuite) TestCreateFeatureTableIndex() {
	RestoreDatabase(suite.DB, suite)
	err := createObjectTable(suite.DB)
	suite.Require().Nil(err)
	err = createFeaturesTable(suite.DB)
	suite.Require().Nil(err)
	query := `SELECT COUNT(*) FROM pg_indexes WHERE tablename = 'feature'`
	var count int
	err = suite.DB.QueryRow(query).Scan(&count)
	suite.Require().Nil(err)
	suite.Require().Equal(2, count)
}

func (suite *AlerceSuite) TestCreateProbabilityTable() {
	RestoreDatabase(suite.DB, suite)
	err := createObjectTable(suite.DB)
	suite.Require().Nil(err)
	err = createProbabilitiesTable(suite.DB)
	suite.Require().Nil(err)
	suite.assertTableExists("probability")
}

func (suite *AlerceSuite) TestCreateProbabilityTableIndex() {
	RestoreDatabase(suite.DB, suite)
	err := createObjectTable(suite.DB)
	suite.Require().Nil(err)
	err = createProbabilitiesTable(suite.DB)
	suite.Require().Nil(err)
	query := `SELECT COUNT(*) FROM pg_indexes WHERE tablename = 'probability'`
	var count int
	err = suite.DB.QueryRow(query).Scan(&count)
	suite.Require().Nil(err)
	suite.Require().Equal(5, count)
}

func (suite *AlerceSuite) assertTableExists(tableName string) {
	query := `SELECT COUNT(*) FROM information_schema.tables WHERE table_name = $1;`
	var count int
	err := suite.DB.QueryRow(query, tableName).Scan(&count)
	suite.Require().Nil(err)
	suite.Require().Equal(1, count)
}
