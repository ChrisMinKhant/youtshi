package clientmanager

import "log"

type CategorizedLinkedClient struct {
	linkedClient                *LinkedClient
	nextCategorizedLinkedClient *CategorizedLinkedClient
}

func NewCategorizedLinkedClient() *CategorizedLinkedClient {
	return &CategorizedLinkedClient{
		linkedClient:                nil,
		nextCategorizedLinkedClient: nil,
	}
}

/*
*	Adds new linked list to the nested linked list
*	Time complexity of O(n)
 */
func (categorizedLinkedClient *CategorizedLinkedClient) AddLinkedClient(requestedLinkedClient *LinkedClient) {

	// Declares temporary nested linke list to operate on.
	tempCategorizedLinkedClient := categorizedLinkedClient

	// Check if the nested linked list to operate on is null.
	if tempCategorizedLinkedClient != nil {

		// If the linked list of the nested linked list is not null,
		// check if the linked list inside the nested linked list is null.
		if tempCategorizedLinkedClient.linkedClient != nil {

			// If the linked list inside the nested linked list is not null,
			// iterate through to the end of the nested linked list, where
			// there is nil in the next nested linked list.
			for tempCategorizedLinkedClient.nextCategorizedLinkedClient != nil {
				tempCategorizedLinkedClient = tempCategorizedLinkedClient.nextCategorizedLinkedClient
			}

			// Add new nested linked list with the linked list value
			// to the next nested linked list of the nested linked list,
			// once the iteration reaches the end.
			tempCategorizedLinkedClient.nextCategorizedLinkedClient = &CategorizedLinkedClient{
				linkedClient:                requestedLinkedClient,
				nextCategorizedLinkedClient: nil,
			}

			return
		}

		// If the linked list of the nested linked list is null,
		// add the linked list to it.
		tempCategorizedLinkedClient.linkedClient = requestedLinkedClient

		log.Printf("Fetched linked client of tempCategorizedLinkedClient ::: %v", tempCategorizedLinkedClient.linkedClient)
	}
}

/*
*	Iterates all the linked list of the nested linked list and
* 	prints out the values in the linked list by calling the FindAll() function
* 	of the linked list.
 */
func (categorizedLinkedClient *CategorizedLinkedClient) FetchClientByBusNumber(busNumber int) []*Client {

	log.Printf("Fetched reuqested bus number in categorized linked client ::: %v", busNumber)
	// declares tempoaray nested linked list to operate on.
	tempCategorizedLinkedClient := categorizedLinkedClient

	// Iterates through all the linked list of the nested linked list.
	for tempCategorizedLinkedClient.linkedClient != nil {

		log.Print("Fetched tempCategorizedLinkedClient status ::: [ Not-Nil ]")

		log.Printf("Fetched busNumber that already exists ::: %v", &tempCategorizedLinkedClient.linkedClient.nextLinkedClient.client.busNumber)

		log.Print("Fetched tempCategorizedLinkedClient status ::: [ Not-Nil ]")
		if tempCategorizedLinkedClient.linkedClient.nextLinkedClient.client.busNumber != busNumber {

			log.Printf("Fetched category found status ::: [ False ].")
			if tempCategorizedLinkedClient.nextCategorizedLinkedClient != nil {

				tempCategorizedLinkedClient = tempCategorizedLinkedClient.nextCategorizedLinkedClient

				continue
			}

		}

		log.Printf("Fetched category found status ::: [ True ].")
		// Then calls the FindAll() function of each linked list.
		return tempCategorizedLinkedClient.linkedClient.FetchClient()

	}

	log.Print("Fetched tempCategorizedLinkedClient status ::: [ Nil ]")
	return nil
}
