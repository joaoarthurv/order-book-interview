package service

import (
	"github.com/labstack/gommon/log"
	"gitlab.com/projetos/orderbook/users-wallet-service/internal/dao"
	"gitlab.com/projetos/orderbook/users-wallet-service/internal/model"
	"time"
)

type UserWalletService struct {
	userWalletDao dao.IUserWalletDao
}

type IUserWalletService interface {
	ProcessOrderExecuted(orderExecuted *model.OrderExecutedBody) error
	PutUserWallet(userWallet *model.UserWallet) error
	GetUserWalletByUserIdAndProductType(userId string, productType string) (*model.UserWallet, error)
}

func NewUserWalletService(userWalletDao dao.IUserWalletDao) *UserWalletService {
	return &UserWalletService{
		userWalletDao: userWalletDao,
	}
}

func (u *UserWalletService) ProcessOrderExecuted(orderExecuted *model.OrderExecutedBody) error {
	orderBuyer, err := u.GetUserWalletByUserIdAndProductType(orderExecuted.OrderBuyOwnerId, orderExecuted.OrderProductType)
	if err != nil {
		return err
	}

	err = u.updateBuyerInformation(orderBuyer, orderExecuted)
	if err != nil {
		log.Error("error to update buyer information")
		return err
	}

	orderSeller, err := u.GetUserWalletByUserIdAndProductType(orderExecuted.OrderSellOwnerId, orderExecuted.OrderProductType)
	if err != nil {
		return err
	}

	err = u.updateSellerInformation(orderSeller, orderExecuted)
	if err != nil {
		log.Error("error to update seller information")
		return err
	}

	return nil
}

func (u *UserWalletService) updateBuyerInformation(orderBuyer *model.UserWallet, orderExecuted *model.OrderExecutedBody) error {
	orderBuyer.Balance = orderBuyer.Balance - (orderExecuted.Price * float64(orderExecuted.OrderQuantity))
	orderBuyer.ProductQuantity = orderBuyer.ProductQuantity + orderExecuted.OrderQuantity
	orderBuyer.UpdatedAt = time.Now().UTC()
	return u.userWalletDao.SaveUserWallet(orderBuyer)
}

func (u *UserWalletService) updateSellerInformation(orderSeller *model.UserWallet, orderExecuted *model.OrderExecutedBody) error {
	orderSeller.Balance = orderSeller.Balance + (orderExecuted.Price * float64(orderExecuted.OrderQuantity))
	orderSeller.UpdatedAt = time.Now().UTC()
	return u.userWalletDao.SaveUserWallet(orderSeller)
}

func (u *UserWalletService) PutUserWallet(userWallet *model.UserWallet) error {
	err := u.userWalletDao.SaveUserWallet(userWallet)
	if err != nil {
		return err
	}
	return nil
}

func (u *UserWalletService) GetUserWalletByUserIdAndProductType(userId string, productType string) (*model.UserWallet, error) {
	userWallet, err := u.userWalletDao.GetUserWalletByUserIdAndProductType(userId, productType)
	if err != nil {
		return nil, err
	}
	return userWallet, nil
}

func (u *UserWalletService) CheckUserBalance(userWalletCheck *model.UserWalletCheck) (bool, error) {
	userWalletResponse, err := u.GetUserWalletByUserIdAndProductType(userWalletCheck.UserId, userWalletCheck.ProductType)
	if err != nil {
		return false, err
	}
	if userWalletResponse.Balance < userWalletCheck.Price {
		return false, nil
	}
	return true, nil
}

func (u *UserWalletService) CheckUserProductQuantity(userWalletCheck *model.UserWalletCheck) (bool, error) {
	userWalletResponse, err := u.GetUserWalletByUserIdAndProductType(userWalletCheck.UserId, userWalletCheck.ProductType)
	if err != nil {
		return false, err
	}
	if userWalletResponse.ProductQuantity < userWalletCheck.ProductQuantity {
		return false, nil
	}
	u.updateProductQuantity(userWalletResponse, userWalletCheck.ProductQuantity)
	return true, nil
}

func (u *UserWalletService) updateProductQuantity(userWalletResponse *model.UserWallet, quantity int) error {
	userWalletResponse.ProductQuantity = userWalletResponse.ProductQuantity - quantity
	err := u.userWalletDao.SaveUserWallet(userWalletResponse)
	if err != nil {
		return err
	}
	return nil
}
