package service

import (
	"TPBDM/scraper/internal/entities"
	"TPBDM/scraper/internal/repositories"
	"TPBDM/scraper/internal/repositories/database"
	"TPBDM/scraper/test/mocks/repositories/database/measurements"
	"errors"
	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/stretchr/testify/suite"
	"testing"
)

type MeasurementsTestSuite struct {
	suite.Suite
	ctrl  *gomock.Controller
	mocks struct {
		dbMocks *dbmeasurementsmocks.MockRepository
	}
}

func (suite *MeasurementsTestSuite) SetupTest() {
	suite.ctrl = gomock.NewController(suite.T())
	suite.mocks.dbMocks = dbmeasurementsmocks.NewMockRepository(suite.ctrl)
}

func (suite *MeasurementsTestSuite) TearDownTest() {
	suite.ctrl.Finish()
}

func (suite *MeasurementsTestSuite) Test_GetMeasurementsList() {

	type args struct {
		query entities.MeasurementsQueryContract
	}

	cases := map[string]struct {
		setup   func()
		args    args
		want    entities.List[entities.Measurement]
		wantErr bool
	}{
		"successful get list": {
			setup: func() {
				suite.mocks.dbMocks.EXPECT().
					GetMeasurementsList(entities.MeasurementsQueryContract{Parameter: "param", City: "city", Country: "country", Location: "loc"}).
					Return([]entities.Measurement{{ID: uuid.UUID{1}}}, nil).
					Times(1)

			},
			args: args{
				query: entities.MeasurementsQueryContract{Parameter: "param", City: "city", Country: "country", Location: "loc"},
			},
			want: entities.List[entities.Measurement]{
				Meta: struct {
					Count int `json:"count"`
				}{Count: 1},
				Data: []entities.Measurement{{ID: uuid.UUID{1}}},
			},
			wantErr: false,
		},
		"failed get list": {
			setup: func() {
				suite.mocks.dbMocks.EXPECT().
					GetMeasurementsList(entities.MeasurementsQueryContract{Parameter: "param", City: "city", Country: "country", Location: "loc"}).
					Return(nil, errors.New("DB ERROR")).
					Times(1)

			},
			args: args{
				query: entities.MeasurementsQueryContract{Parameter: "param", City: "city", Country: "country", Location: "loc"},
			},
			want:    entities.List[entities.Measurement]{},
			wantErr: true,
		},
	}

	for name, cs := range cases {
		suite.Run(name, func() {
			cs.setup()

			s := New(
				repositories.Repository{
					DB: &database.Container{Measurements: suite.mocks.dbMocks},
				},
			)

			res, err := s.GetMeasurementsList(cs.args.query)

			if cs.wantErr {
				suite.Require().Error(err)
			} else {
				suite.Require().NoError(err)
				suite.Require().EqualValues(cs.want, res)
			}
		})
	}
}

func (suite *MeasurementsTestSuite) Test_GetMeasurementCitiesList() {

	type args struct {
		query entities.MeasurementsQueryContract
	}

	cases := map[string]struct {
		setup   func()
		args    args
		want    entities.List[string]
		wantErr bool
	}{
		"successful get list": {
			setup: func() {
				suite.mocks.dbMocks.EXPECT().
					GetMeasurementCitiesList(entities.MeasurementsQueryContract{Parameter: "param", City: "city", Country: "country", Location: "loc"}).
					Return([]string{"city"}, nil).
					Times(1)

			},
			args: args{
				query: entities.MeasurementsQueryContract{Parameter: "param", City: "city", Country: "country", Location: "loc"},
			},
			want: entities.List[string]{
				Meta: struct {
					Count int `json:"count"`
				}{Count: 1},
				Data: []string{"city"},
			},
			wantErr: false,
		},
		"failed get list": {
			setup: func() {
				suite.mocks.dbMocks.EXPECT().
					GetMeasurementCitiesList(entities.MeasurementsQueryContract{Parameter: "param", City: "city", Country: "country", Location: "loc"}).
					Return(nil, errors.New("DB ERROR")).
					Times(1)

			},
			args: args{
				query: entities.MeasurementsQueryContract{Parameter: "param", City: "city", Country: "country", Location: "loc"},
			},
			want:    entities.List[string]{},
			wantErr: true,
		},
	}

	for name, cs := range cases {
		suite.Run(name, func() {
			cs.setup()

			s := New(
				repositories.Repository{
					DB: &database.Container{Measurements: suite.mocks.dbMocks},
				},
			)

			res, err := s.GetMeasurementCitiesList(cs.args.query)

			if cs.wantErr {
				suite.Require().Error(err)
			} else {
				suite.Require().NoError(err)
				suite.Require().EqualValues(cs.want, res)
			}
		})
	}
}

func (suite *MeasurementsTestSuite) Test_GetMeasurementLocationsList() {

	type args struct {
		query entities.MeasurementsQueryContract
	}

	cases := map[string]struct {
		setup   func()
		args    args
		want    entities.List[string]
		wantErr bool
	}{
		"successful get list": {
			setup: func() {
				suite.mocks.dbMocks.EXPECT().
					GetMeasurementLocationsList(entities.MeasurementsQueryContract{Parameter: "param", City: "city", Country: "country", Location: "loc"}).
					Return([]string{"loc"}, nil).
					Times(1)

			},
			args: args{
				query: entities.MeasurementsQueryContract{Parameter: "param", City: "city", Country: "country", Location: "loc"},
			},
			want: entities.List[string]{
				Meta: struct {
					Count int `json:"count"`
				}{Count: 1},
				Data: []string{"loc"},
			},
			wantErr: false,
		},
		"failed get list": {
			setup: func() {
				suite.mocks.dbMocks.EXPECT().
					GetMeasurementLocationsList(entities.MeasurementsQueryContract{Parameter: "param", City: "city", Country: "country", Location: "loc"}).
					Return(nil, errors.New("DB ERROR")).
					Times(1)

			},
			args: args{
				query: entities.MeasurementsQueryContract{Parameter: "param", City: "city", Country: "country", Location: "loc"},
			},
			want:    entities.List[string]{},
			wantErr: true,
		},
	}

	for name, cs := range cases {
		suite.Run(name, func() {
			cs.setup()

			s := New(
				repositories.Repository{
					DB: &database.Container{Measurements: suite.mocks.dbMocks},
				},
			)

			res, err := s.GetMeasurementLocationsList(cs.args.query)

			if cs.wantErr {
				suite.Require().Error(err)
			} else {
				suite.Require().NoError(err)
				suite.Require().EqualValues(cs.want, res)
			}
		})
	}
}

func (suite *MeasurementsTestSuite) Test_GetMeasurementParameterList() {

	type args struct {
		query entities.MeasurementsQueryContract
	}

	cases := map[string]struct {
		setup   func()
		args    args
		want    entities.List[string]
		wantErr bool
	}{
		"successful get list": {
			setup: func() {
				suite.mocks.dbMocks.EXPECT().
					GetMeasurementParametersList(entities.MeasurementsQueryContract{Parameter: "param", City: "city", Country: "country", Location: "loc"}).
					Return([]string{"param"}, nil).
					Times(1)

			},
			args: args{
				query: entities.MeasurementsQueryContract{Parameter: "param", City: "city", Country: "country", Location: "loc"},
			},
			want: entities.List[string]{
				Meta: struct {
					Count int `json:"count"`
				}{Count: 1},
				Data: []string{"param"},
			},
			wantErr: false,
		},
		"failed get list": {
			setup: func() {
				suite.mocks.dbMocks.EXPECT().
					GetMeasurementParametersList(entities.MeasurementsQueryContract{Parameter: "param", City: "city", Country: "country", Location: "loc"}).
					Return(nil, errors.New("DB ERROR")).
					Times(1)

			},
			args: args{
				query: entities.MeasurementsQueryContract{Parameter: "param", City: "city", Country: "country", Location: "loc"},
			},
			want:    entities.List[string]{},
			wantErr: true,
		},
	}

	for name, cs := range cases {
		suite.Run(name, func() {
			cs.setup()

			s := New(
				repositories.Repository{
					DB: &database.Container{Measurements: suite.mocks.dbMocks},
				},
			)

			res, err := s.GetMeasurementParameterList(cs.args.query)

			if cs.wantErr {
				suite.Require().Error(err)
			} else {
				suite.Require().NoError(err)
				suite.Require().EqualValues(cs.want, res)
			}
		})
	}
}

func TestMeasurementsTestSuite(t *testing.T) {
	suite.Run(t, new(MeasurementsTestSuite))
}
