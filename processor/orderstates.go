package processor

const (
	// as per PaymentProcessor.sol enum State {Null, Created, Paid, Finalized, Refunding, Refunded, Cancelled}

	// OrderStateNull = 0, order's state in Ethereum storage
	OrderStateNull = iota
	// OrderStateCreated = 1, order's state in Ethereum storage
	OrderStateCreated
	// OrderStatePaid = 2, order's state in Ethereum storage
	OrderStatePaid
	// OrderStateFinalized = 3, order's state in Ethereum storage
	OrderStateFinalized
	// OrderStateRefunding = 4, order's state in Ethereum storage
	OrderStateRefunding
	// OrderStateRefunded = 5, order's state in Ethereum storage
	OrderStateRefunded
	// OrderStateCanceled = 6, order's state in Ethereum storage
	OrderStateCanceled
)
