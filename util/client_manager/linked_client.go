package clientmanager

import "log"

/*
* The ClientNode struct is basic
* node for organizing large websocket
* connection which is included in Client struct.
 */
type LinkedClient struct {
	client           *Client
	nextLinkedClient *LinkedClient
}

func NewLinkedClient() *LinkedClient {
	return &LinkedClient{
		client:           nil,
		nextLinkedClient: nil,
	}
}

func (linkedClient *LinkedClient) AddClient(requestedClient *Client) {

	log.Printf("Fetched requested client comming to the linked client ::: %v", requestedClient)
	// Check is there any next node after the head node
	if linkedClient.nextLinkedClient != nil {
		tempClient := linkedClient.nextLinkedClient

		// Iterate through till the next node is nil
		for tempClient.nextLinkedClient != nil {
			tempClient = tempClient.nextLinkedClient
		}

		// Iteration reaches the end of the list.
		// Therefore, add new node to the next node.
		tempClient.nextLinkedClient = &LinkedClient{
			client:           requestedClient,
			nextLinkedClient: nil,
		}

		log.Printf("Fetched next linked client status ::: [ Not-Nil ]")
		log.Printf("Fetched added client temporary linked client ::: %v", tempClient.nextLinkedClient)

		return
	}

	// There is no next node after head node.
	// Therefore, add new node to that next node.
	linkedClient.nextLinkedClient = &LinkedClient{
		client:           requestedClient,
		nextLinkedClient: nil,
	}

	log.Printf("Fetched next linked client status ::: [ Nil ]")
	log.Printf("Fetched added client linked client ::: %v", linkedClient.nextLinkedClient.client)
}

/*
*	Print out all the value in the nodes, which
* 	are linked in the list.
 */
func (linkedClient *LinkedClient) FetchClient() []*Client {

	log.Print("Fetched invoking FetchClient() of linked client status ::: [ Reached ]")
	fetchedClient := []*Client{}

	// There will be value only if there is a next node
	// after the head node.
	if linkedClient.nextLinkedClient != nil {
		tempLinkedClient := linkedClient.nextLinkedClient

		fetchedClient = append(fetchedClient, tempLinkedClient.client)

		for tempLinkedClient.nextLinkedClient != nil {

			tempLinkedClient = tempLinkedClient.nextLinkedClient
			fetchedClient = append(fetchedClient, tempLinkedClient.client)
		}
	}

	log.Printf("Fetched found client in linked client ::: %v", fetchedClient)

	return fetchedClient
}
