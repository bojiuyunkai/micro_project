package repositories

import (
	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"
	"go.uber.org/zap"
	"hz_project/internal/pkg/models"
)

type DetailsRepository interface {
	Get(ID uint64) (p *models.Detail, err error)
}

type MysqlDetailsRepository struct {
	logger *zap.Logger
	db     *gorm.DB
}

func NewMysqlDetailsRepository(logger *zap.Logger, db *gorm.DB) DetailsRepository {
	return &MysqlDetailsRepository{
		logger: logger.With(zap.String("type", "DetailsRepository")),
		db:     db,
	}
}

func (s *MysqlDetailsRepository) Get(ID uint64) (p *models.Detail, err error) {
	detailMgr,retErr:= models.DetailsMgr(s.db);
	if retErr!= nil {
		return nil, errors.Wrapf(retErr, "get product error[id=%d]", ID)
	}

	details,err:=detailMgr.GetFromID(ID);

	p = &details;
	return
}
