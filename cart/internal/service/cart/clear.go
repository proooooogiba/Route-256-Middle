package cart

import "context"

func (c *Service) Clear(ctx context.Context, userID int64) error {
	return c.repo.DeleteItemsByUserID(ctx, userID)
}
