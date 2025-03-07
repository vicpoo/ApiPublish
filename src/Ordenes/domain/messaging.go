// domain/messaging.go
package domain

type MessagingService interface {
	PublishOrderCreated(order []byte) error
}
