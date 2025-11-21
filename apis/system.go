package apis

import (
	"net/http"

	"github.com/labstack/gommon/log"
	"github.com/nanoteck137/pyrin"
	"github.com/nanoteck137/storebook"
	"github.com/nanoteck137/storebook/core"
)

type GetSystemInfo struct {
	Version string `json:"version"`
}

type Event struct {
	Type string `json:"type"`
	Data any    `json:"data"`
}

type EventData interface {
	GetEventType() string
}

// NOTE(patrik): Based on: https://gist.github.com/Ananto30/8af841f250e89c07e122e2a838698246
type Broker struct {
	Notifier chan EventData

	newClients     chan chan EventData
	closingClients chan chan EventData
	clients        map[chan EventData]bool
}

func NewServer() (broker *Broker) {
	// Instantiate a broker
	broker = &Broker{
		Notifier:       make(chan EventData, 1),
		newClients:     make(chan chan EventData),
		closingClients: make(chan chan EventData),
		clients:        make(map[chan EventData]bool),
	}

	// Set it running - listening and broadcasting events
	go broker.listen()

	return
}

func (broker *Broker) listen() {
	for {
		select {
		case s := <-broker.newClients:
			broker.clients[s] = true
			log.Debug("Client added", "numClients", len(broker.clients))
		case s := <-broker.closingClients:
			delete(broker.clients, s)
			log.Debug("Removed client", "numClients", len(broker.clients))
		case event := <-broker.Notifier:
			for clientMessageChan := range broker.clients {
				clientMessageChan <- event
			}
		}
	}
}

func (broker *Broker) EmitEvent(event EventData) {
	broker.Notifier <- event
}

// const (
// 	EventSyncing string = "syncing"
// 	EventReport  string = "report"
// )
//
// type SyncEvent struct {
// 	Syncing bool `json:"syncing"`
// }
//
// func (s SyncEvent) GetEventType() string {
// 	return EventSyncing
// }
//
// type ReportEvent struct {
// 	Report
// }
//
// func (s ReportEvent) GetEventType() string {
// 	return EventReport
// }

func InstallSystemHandlers(app core.App, group pyrin.Group) {
	group.Register(
		pyrin.ApiHandler{
			Name:         "GetSystemInfo",
			Path:         "/system/info",
			Method:       http.MethodGet,
			ResponseType: GetSystemInfo{},
			HandlerFunc: func(c pyrin.Context) (any, error) {
				return GetSystemInfo{
					Version: storebook.Version,
				}, nil
			},
		},

		// pyrin.NormalHandler{
		// 	Name:   "SseHandler",
		// 	Method: http.MethodGet,
		// 	Path:   "/system/library/sse",
		// 	HandlerFunc: func(c pyrin.Context) error {
		// 		r := c.Request()
		// 		w := c.Response()
		//
		// 		w.Header().Set("Content-Type", "text/event-stream")
		// 		w.Header().Set("Cache-Control", "no-cache")
		// 		w.Header().Set("Connection", "keep-alive")
		//
		// 		w.Header().Set("Access-Control-Allow-Origin", "*")
		//
		// 		rc := http.NewResponseController(w)
		//
		// 		eventChan := make(chan EventData)
		// 		syncHandler.broker.newClients <- eventChan
		//
		// 		defer func() {
		// 			syncHandler.broker.closingClients <- eventChan
		// 		}()
		//
		// 		sendEvent := func(eventData EventData) {
		// 			fmt.Fprintf(w, "data: ")
		//
		// 			event := Event{
		// 				Type: eventData.GetEventType(),
		// 				Data: eventData,
		// 			}
		//
		// 			encode := json.NewEncoder(w)
		// 			encode.Encode(event)
		//
		// 			fmt.Fprintf(w, "\n\n")
		// 			rc.Flush()
		// 		}
		//
		// 		sendEvent(SyncEvent{
		// 			Syncing: syncHandler.isSyncing.Load(),
		// 		})
		//
		// 		sendEvent(ReportEvent{
		// 			Report: syncHandler.GetReport(),
		// 		})
		//
		// 		for {
		// 			select {
		// 			case <-r.Context().Done():
		// 				syncHandler.broker.closingClients <- eventChan
		// 				return nil
		//
		// 			case event := <-eventChan:
		// 				sendEvent(event)
		// 			}
		// 		}
		// 	},
		// },
	)
}
