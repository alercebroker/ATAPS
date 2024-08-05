package alercedb

func (suite *AlerceSuite) TestInsertSampleObjects() {
	RestoreDatabase(suite.DB, suite)
	err := CreateTables(suite.DB)
	suite.Require().Nil(err)
	_, err = InsertSampleObjects(suite.DB, 10)
	query := `SELECT COUNT(*) FROM object;`
	var count int
	err = suite.DB.QueryRow(query).Scan(&count)
	suite.Require().Nil(err)
	suite.Require().Equal(10, count)
}

func (suite *AlerceSuite) TestInsertSampleDetections() {
	RestoreDatabase(suite.DB, suite)
	err := CreateTables(suite.DB)
	suite.Require().Nil(err)
	oidPool, err := InsertSampleObjects(suite.DB, 10)
	suite.Require().Nil(err)
	err = InsertSampleDetections(suite.DB, 100, oidPool)
	suite.Require().Nil(err)
	query := `SELECT COUNT(*) FROM detection;`
	var count int
	err = suite.DB.QueryRow(query).Scan(&count)
	suite.Require().Nil(err)
	suite.Require().Equal(100, count)
}

func (suite *AlerceSuite) TestInsertSampleNonDetections() {
	RestoreDatabase(suite.DB, suite)
	err := CreateTables(suite.DB)
	suite.Require().Nil(err)
	oidPool, err := InsertSampleObjects(suite.DB, 10)
	suite.Require().Nil(err)
	err = InsertSampleNonDetections(suite.DB, 100, oidPool)
	suite.Require().Nil(err)
	query := `SELECT COUNT(*) FROM non_detection;`
	var count int
	err = suite.DB.QueryRow(query).Scan(&count)
	suite.Require().Nil(err)
	suite.Require().Equal(100, count)
}

func (suite *AlerceSuite) TestInsertSampleForcedPhotometry() {
	RestoreDatabase(suite.DB, suite)
	err := CreateTables(suite.DB)
	suite.Require().Nil(err)
	oidPool, err := InsertSampleObjects(suite.DB, 10)
	suite.Require().Nil(err)
	err = InsertSampleForcedPhotometry(suite.DB, 100, oidPool)
	suite.Require().Nil(err)
	query := `SELECT COUNT(*) FROM forced_photometry;`
	var count int
	err = suite.DB.QueryRow(query).Scan(&count)
	suite.Require().Nil(err)
	suite.Require().Equal(100, count)
}

func (suite *AlerceSuite) TestInsertSampleFeatures() {
	RestoreDatabase(suite.DB, suite)
	err := CreateTables(suite.DB)
	suite.Require().Nil(err)
	oidPool, err := InsertSampleObjects(suite.DB, 10)
	suite.Require().Nil(err)
	err = InsertSampleFeatures(suite.DB, 100, oidPool)
	suite.Require().Nil(err)
	query := `SELECT COUNT(*) FROM feature;`
	var count int
	err = suite.DB.QueryRow(query).Scan(&count)
	suite.Require().Nil(err)
	suite.Require().Equal(100, count)
}

func (suite *AlerceSuite) TestInsertSampleProbabilities() {
	RestoreDatabase(suite.DB, suite)
	err := CreateTables(suite.DB)
	suite.Require().Nil(err)
	oidPool, err := InsertSampleObjects(suite.DB, 10)
	suite.Require().Nil(err)
	classPool := []string{"class1", "class2", "class3"}
	err = InsertSampleProbabilities(suite.DB, oidPool, classPool, "classifier")
	suite.Require().Nil(err)
	query := `SELECT COUNT(*) FROM probability;`
	var count int
	err = suite.DB.QueryRow(query).Scan(&count)
	suite.Require().Nil(err)
	suite.Require().Equal(len(classPool)*len(oidPool), count)
}
