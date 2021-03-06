package talent

import (
	"context"
	"fmt"

	"github.com/huypher/kit/container"

	"github.com/huypher/talent-acquisition-management/domain"
)

type talentUsecase struct {
	talentRepository domain.TalentRepository
}

func NewTalentUsecase(userRepository domain.TalentRepository) *talentUsecase {
	return &talentUsecase{
		talentRepository: userRepository,
	}
}

func (u *talentUsecase) GetByID(ctx context.Context, id int) (*domain.Talent, error) {
	return u.talentRepository.GetByID(ctx, id)
}

func (u *talentUsecase) GetList(ctx context.Context, filter container.Map, pageID, perPage int) ([]domain.Talent, error) {
	return u.talentRepository.GetList(ctx, filter, pageID, perPage)
}

func (u *talentUsecase) AddTalent(ctx context.Context, talent domain.Talent) error {
	return u.talentRepository.Create(ctx, talent)
}

func (u *talentUsecase) UpdateTalent(ctx context.Context, talentID int, params container.Map) error {
	talent, err := u.talentRepository.GetByID(ctx, talentID)
	if err != nil {
		return err
	}
	if talent == nil {
		return NewErrTalentNotFound(fmt.Sprintf("not found talent id=%d", talentID))
	}

	return u.talentRepository.Update(ctx, talentID, params)
}
