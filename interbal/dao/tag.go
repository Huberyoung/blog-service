package dao

import (
	"blog-service/interbal/model"
	"blog-service/pkg/app"
)

func (d *Dao) GetTag(id uint) (model.Tag, error) {
	tag := &model.Tag{Model: &model.Model{ID: id}}
	return tag.Get(d.engine)
}

func (d *Dao) CountTag(name string, state uint8) (int, error) {
	tag := &model.Tag{Name: name, State: state}
	return tag.Count(d.engine)
}

func (d *Dao) GetTagList(name string, state uint8, page, pageSize int) ([]*model.Tag, error) {
	tag := &model.Tag{Name: name, State: state}
	pageOffset := app.GetPageOffset(page, pageSize)
	return tag.List(d.engine, pageOffset, pageSize)
}

func (d *Dao) CreateTag(name string, state uint8, createdBy string) error {
	tag := &model.Tag{Model: &model.Model{CreatedBy: createdBy}, Name: name, State: state}
	return tag.Create(d.engine)
}

func (d *Dao) UpdateTag(id uint, name string, state uint8, modifiedBy string) error {
	tag := &model.Tag{
		Model: &model.Model{ID: id},
	}

	values := map[string]any{
		"state":       state,
		"modified_by": modifiedBy,
	}

	if name != "" {
		values["name"] = name
	}
	return tag.Update(d.engine, values)
}

func (d *Dao) DeleteTag(id uint) error {
	tag := &model.Tag{Model: &model.Model{ID: id}}
	return tag.Delete(d.engine)
}
