package service

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"gitlab.com/projetos/orderbook/users-wallet-service/internal/mocks"
	"testing"
)

func Test_UserWallet(t *testing.T) {
	t.Run("Should Process Order Executed with success", func(t *testing.T) {
		//give
		userWalletDaoMock := mocks.NewIUserWalletDao(t)
		userWalletService := NewUserWalletService(userWalletDaoMock)

		userWalletDaoMock.EXPECT().GetUserWalletByUserIdAndProductType(mock.Anything, mock.Anything).Call.Return(&mocks.UserWalletMock, nil)
		userWalletDaoMock.EXPECT().SaveUserWallet(mock.Anything).Call.Return(nil)

		//when
		err := userWalletService.ProcessOrderExecuted(&mocks.OrderExecutedListenerMock)

		//then
		assert.Nil(t, err)
	})
}
