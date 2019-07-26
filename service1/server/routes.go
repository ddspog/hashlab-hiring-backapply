package server

import (
	"context"

	"github.com/ddspog/hashlab-hiring-backapply/service1/model/product"
	"github.com/ddspog/hashlab-hiring-backapply/service1/server/protocols"
	"github.com/ddspog/hashlab-hiring-backapply/service1/today"
)

// Server contains the "routes" of this service. That are
type Server struct{}

// New cretes the server, with no needed info.
func New() (s *Server) {
	s = &Server{}
	return
}

// VerifyDiscount checks the available discount of a product, for a
// specified user. It gives at max 10% if it's Black Friday, and 5%
// if it's only the user Birthday.
func (s *Server) VerifyDiscount(ctx context.Context, form *protocols.VerifyDiscountForm) (info *protocols.DiscountInfo, err error) {
	info = &protocols.DiscountInfo{Pct: 0, ValueInCents: 0}

	var p *product.Product
	if p, err = product.FetchByID(form.GetProdid()); err != nil {
		return
	}

	if today.IsBlackFriday() {
		info.Pct = 0.1
	} else {
		var ok bool
		if ok, err = today.IsUserBDay(form.GetUserid()); err != nil {
			return
		} else if ok {
			info.Pct = 0.05
		}
	}

	info.ValueInCents = int64(float32(p.PriceInCents) * info.Pct)
	return
}
