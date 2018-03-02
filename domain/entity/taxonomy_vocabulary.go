package entity

import (
	"github.com/afranioce/goddd/domain"
)

type TaxonomyVocabulary struct {
	entityBase
	entityBlamed
	Name        string `gorm:"type:varchar(50);not null" sql:"index"`
	Description string `gorm:"type:varchar(1000);not null"`
	Status      domain.Status
}

func (entidade *TaxonomyVocabulary) ToDomain() EntityTransformer {
	return &taxonomyVocabularyDomain{
		domainBase: &domainBase{
			value: entidade,
		},
	}
}

type taxonomyVocabularyDomain struct {
	*domainBase
}

func NewTaxonomyVocabulary(name string, description string, author *userDomain) *taxonomyVocabularyDomain {
	return &taxonomyVocabularyDomain{
		domainBase: &domainBase{
			value: &TaxonomyVocabulary{
				Name:        name,
				Description: description,
				Status:      domain.StatusEnabled,
				entityBlamed: entityBlamed{
					CreatedBy:   *author.ToEntity().(*User),
					CreatedByID: author.Id(),
				},
			},
		},
	}
}

func (d *taxonomyVocabularyDomain) Id() uint {
	return d.value.(*TaxonomyVocabulary).ID
}

func (d *taxonomyVocabularyDomain) Name() string {
	return d.value.(*TaxonomyVocabulary).Name
}

func (d *taxonomyVocabularyDomain) Description() string {
	return d.value.(*TaxonomyVocabulary).Description
}

func (d *taxonomyVocabularyDomain) Author() EntityTransformer {
	return d.value.(*TaxonomyVocabulary).CreatedBy.ToDomain()
}

func (d *taxonomyVocabularyDomain) Editor() EntityTransformer {
	return d.value.(*TaxonomyVocabulary).ChangedBy.ToDomain()
}

func (d *taxonomyVocabularyDomain) Status() domain.Status {
	return d.value.(*TaxonomyVocabulary).Status
}