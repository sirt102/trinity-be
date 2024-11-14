package services

import (
	"time"
	"trinity-be/global"
	"trinity-be/internal/entities"
	"trinity-be/internal/entities/requests"
	"trinity-be/internal/repositories"
)

type SubscriptionService interface {
	UserRegisterSubscription(*requests.UserRegisterSubscriptionRequest) (*entities.Transaction, error)
	UserPaiedForSubscription(*requests.UserPaiedForSubscriptionRequest) error
}

type subscriptionService struct {
	subscriptionRepo repositories.SubscriptionRepository
	campaignRepo     repositories.CampaignRepository
	voucherRepo      repositories.VoucherRepository
	uvRepo           repositories.UserVoucherRepository
	transactionRepo  repositories.TransactionRepository
	usRepo           repositories.UserSubscriptionRepository
}

func NewSubscriptionService(subscriptionRepo repositories.SubscriptionRepository,
	campaignRepo repositories.CampaignRepository,
	voucherRepo repositories.VoucherRepository,
	uvRepo repositories.UserVoucherRepository,
	transactionRepo repositories.TransactionRepository,
	usRepo repositories.UserSubscriptionRepository,
) SubscriptionService {
	return &subscriptionService{
		subscriptionRepo: subscriptionRepo,
		campaignRepo:     campaignRepo,
		voucherRepo:      voucherRepo,
		uvRepo:           uvRepo,
		transactionRepo:  transactionRepo,
		usRepo:           usRepo,
	}
}

func (ss *subscriptionService) UserRegisterSubscription(req *requests.UserRegisterSubscriptionRequest) (*entities.Transaction, error) {
	// 1. Check any compaign is running
	listCampaigns, err := ss.campaignRepo.GetRunningCampaigns()
	if err != nil {
		global.LogError(err, "")

		return nil, err
	}

	// 2. Check condition of sale off
	// 3. Add voucher to user storage
	for _, campaign := range listCampaigns {
		if campaign.Available > 0 {
			vc, err := ss.voucherRepo.GetVoucherByCampaignID(campaign.CampaignID)
			if err != nil {
				global.LogError(err, "")
				continue
			}

			uv := &entities.UserVoucher{
				UserID:    req.UserID,
				VoucherID: vc.VoucherID,
			}

			err = ss.uvRepo.CreateNewUserVoucher(uv)
			if err != nil {
				global.LogError(err, "")

				continue
			}
		}

	}

	// 4. Init transaction with status is pending or processing.
	err = ss.transactionRepo.InsertNewTransaction(&entities.Transaction{
		UserID:         req.UserID,
		SubscriptionID: req.SubscriptionID,
		StatusPayment:  entities.PaymentStatusPending,
	})
	if err != nil {
		global.LogError(err, "")

		return nil, err
	}

	newTransaction, err := ss.transactionRepo.GetTransactionByUserAndSubID(req.UserID, req.SubscriptionID)
	if err != nil {
		global.LogError(err, "")

		return nil, err
	}

	return newTransaction, nil
}

func (ss *subscriptionService) UserPaiedForSubscription(req *requests.UserPaiedForSubscriptionRequest) error {
	// 0. Get Voucher
	vc, err := ss.uvRepo.GetUserVoucherByID(req.UserVoucherID)
	if err != nil {
		global.LogError(err, "")

		return err
	}

	// 1. Update voucher status to used
	err = ss.uvRepo.UsedUserVoucher(req.UserVoucherID)
	if err != nil {
		global.LogError(err, "")

		return err
	}

	// At this step, temporary accept that payment successfully. We will automatically update subscription status to completed.
	// After payment method is implemented, add transaction mechanic to handle this logic.

	// 2. Update user subscription status to completed
	expireDate := time.Now().AddDate(0, 1, 0)
	err = ss.usRepo.CreateNewUserSubscription(
		&entities.UserSubscription{
			UserID:         req.UserID,
			SubscriptionID: req.SubscriptionID,
			SubscribedAt:   time.Now(),
			ExpiryDate:     &expireDate,
		},
	)
	if err != nil {
		global.LogError(err, "")

		return err
	}

	// 3. Update transaction amount and voucher
	err = ss.transactionRepo.ApplyVoucherToTransaction(req.TransactionID, vc.VoucherID, vc.Voucher.DiscountRate)
	if err != nil {
		global.LogError(err, "")

		return err
	}

	// 4. Update transaction status to completed
	err = ss.transactionRepo.UpdatePaySuccessTransaction(req.TransactionID)
	if err != nil {
		global.LogError(err, "")

		return err
	}

	return nil
}
