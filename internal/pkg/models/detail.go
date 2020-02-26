package models
import (
	
	"github.com/jinzhu/gorm"
	"time"
	"errors"
)

type Detail struct {
	ID    uint64 `json:"id"`
	Name  string  `json:"name"`
	Price float32 `json:"price"`
	CreatedTime time.Time `json:"created_time"`
}


type _DetailsMgr struct {
	*_BaseMgr
}

// DetailsMgr open func
func DetailsMgr(db *gorm.DB)( *_DetailsMgr ,error){
	if db == nil {
		return nil,errors.New("DetailsMgr need init by db");
	}
	return &_DetailsMgr{_BaseMgr: &_BaseMgr{DB: db, isRelated: gloabIsRelated}},nil
}

// GetTableName get sql table name.获取数据库名字
func (obj *_DetailsMgr) GetTableName() string {
	return "details"
}

// Get 获取
func (obj *_DetailsMgr) Get() (result Detail, err error) {
	err = obj.DB.Table(obj.GetTableName()).Find(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_DetailsMgr) Gets() (results []*Detail, err error) {
	err = obj.DB.Table(obj.GetTableName()).Find(&results).Error

	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取
func (obj *_DetailsMgr) WithID(ID uint64) Option {
	return optionFunc(func(o *options) { o.query["id"] = ID })
}

// WithName name获取
func (obj *_DetailsMgr) WithName(Name string) Option {
	return optionFunc(func(o *options) { o.query["name"] = Name })
}

// WithPrice price获取
func (obj *_DetailsMgr) WithPrice(Price float64) Option {
	return optionFunc(func(o *options) { o.query["price"] = Price })
}

// WithCreatedTime created_time获取
func (obj *_DetailsMgr) WithCreatedTime(CreatedTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["created_time"] = CreatedTime })
}

// GetByOption 功能选项模式获取
func (obj *_DetailsMgr) GetByOption(opts ...Option) (result Detail, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.Table(obj.GetTableName()).Where(options.query).Find(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_DetailsMgr) GetByOptions(opts ...Option) (results []*Detail, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.Table(obj.GetTableName()).Where(options.query).Find(&results).Error

	return
}

//////////////////////////enume case ////////////////////////////////////////////

// GetFromID 通过id获取内容
func (obj *_DetailsMgr) GetFromID(ID uint64) (result Detail, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("id = ?", ID).Find(&result).Error

	return
}

// GetBatchFromID 批量唯一主键查找
func (obj *_DetailsMgr) GetBatchFromID(IDs []uint64) (results []*Detail, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("id IN (?)", IDs).Find(&results).Error

	return
}

// GetFromName 通过name获取内容
func (obj *_DetailsMgr) GetFromName(Name string) (results []*Detail, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("name = ?", Name).Find(&results).Error

	return
}

// GetBatchFromName 批量唯一主键查找
func (obj *_DetailsMgr) GetBatchFromName(Names []string) (results []*Detail, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("name IN (?)", Names).Find(&results).Error

	return
}

// GetFromPrice 通过price获取内容
func (obj *_DetailsMgr) GetFromPrice(Price float64) (results []*Detail, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("price = ?", Price).Find(&results).Error

	return
}

// GetBatchFromPrice 批量唯一主键查找
func (obj *_DetailsMgr) GetBatchFromPrice(Prices []float64) (results []*Detail, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("price IN (?)", Prices).Find(&results).Error

	return
}

// GetFromCreatedTime 通过created_time获取内容
func (obj *_DetailsMgr) GetFromCreatedTime(CreatedTime time.Time) (results []*Detail, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("created_time = ?", CreatedTime).Find(&results).Error

	return
}

// GetBatchFromCreatedTime 批量唯一主键查找
func (obj *_DetailsMgr) GetBatchFromCreatedTime(CreatedTimes []time.Time) (results []*Detail, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("created_time IN (?)", CreatedTimes).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primay or index 获取唯一内容
func (obj *_DetailsMgr) FetchByPrimaryKey(ID uint64) (result Detail, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("id = ?", ID).Find(&result).Error

	return
}
