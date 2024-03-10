package clientmanager

import (
	"fmt"
	"log"
)

type ClientManager struct {
	categorizedLinkedClient *CategorizedLinkedClient
}

func NewClientManager() *ClientManager {
	return &ClientManager{
		categorizedLinkedClient: NewCategorizedLinkedClient(),
	}
}

/*
*	Adds single value to the node of the nested linked list,
*	according to their category.
*	The result nested linked list will be categorized by default.
 */
func (clientManager *ClientManager) AddClient(requestedClient *Client) {

	// Declares temporary nested linked list to operate on.
	tempCategorizedLinkedClient := clientManager.categorizedLinkedClient

	// Iterates through to the end of the nested linked list.
	// By then, we will get separate nested linked list to concern.
	for tempCategorizedLinkedClient != nil {

		tempLinkedClient := tempCategorizedLinkedClient.linkedClient

		// Check if the linked list of the nested linked list is nil.
		// This if-statement is to ensure that there is intialized linked list inside
		// the nested linked list, to avoid nil pointer exception.
		if tempLinkedClient != nil {

			log.Print("Fetched linked client status ::: [ Not-Nil ]")
			// Check if the next node right after the head node the linked list is nil.
			if tempLinkedClient.nextLinkedClient != nil {

				// Check if the value of the next node is equal to the value to be added.
				if tempLinkedClient.nextLinkedClient.client.busNumber == requestedClient.busNumber {

					fmt.Printf("Found a category for requested bus number : %v \n", requestedClient.busNumber)

					tempLinkedClient.AddClient(requestedClient)

					return
				}

				// We will checking only the value of the first node right after the head node.
				// Because, we can easily know that the linked list is the right category for the
				// value to be added, just by checking the first node.

			}

			// Swap the temporary nested linked list with the next nested linked list,
			// only if there is one.
			if tempCategorizedLinkedClient.nextCategorizedLinkedClient != nil {

				tempCategorizedLinkedClient = tempCategorizedLinkedClient.nextCategorizedLinkedClient

				continue
			}

		}

		log.Print("Fetched linked client status ::: [ Nil ]")
		// If there is no right category for the valued to be added,
		// we will added new nested linked with the value already in it.
		tempLinkedClient = NewLinkedClient()

		tempLinkedClient.AddClient(requestedClient)

		log.Printf("Fetched tempLinkedClient before add to tempCategorizedLinkedClient ::: %v", tempLinkedClient.nextLinkedClient)

		tempCategorizedLinkedClient.AddLinkedClient(tempLinkedClient)

		return
	}

	// This is the case where, there is no single nested linked list.
	tempLinkedClient := NewLinkedClient()

	tempLinkedClient.AddClient(requestedClient)

	tempCategorizedLinkedClient.AddLinkedClient(tempLinkedClient)

}

/*
*	Prints out all the value of the linked list of the nested linked list.
 */
func (clientManager *ClientManager) FetchClientByBusNumber(busNumber int) []*Client {

	log.Printf("Fetched requested bus number ::: %v", busNumber)
	// Check the nested linked list is not nil.
	// This is to avoid nil pointer exception.
	if clientManager.categorizedLinkedClient != nil {
		log.Print("Fetched categorized linked client status ::: [ not nil ].")
		return clientManager.categorizedLinkedClient.FetchClientByBusNumber(busNumber)
	}

	log.Print("Fetched categorized linked client status ::: [ nil ].")
	return nil
}
