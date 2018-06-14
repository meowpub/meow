package models

import (
	"math"

	"github.com/bwmarrin/snowflake"
	"github.com/jinzhu/gorm"
	"github.com/meowpub/meow/config"
	"github.com/meowpub/meow/lib"
)

//go:generate mockgen -package=models -source=streamitem.go -destination=streamitem.mock.go

// StreamItem represents a single item in a string
type StreamItem struct {
	ItemID   snowflake.ID `json:"item_id"`
	StreamID snowflake.ID `json:"stream_id"`
	EntityID snowflake.ID `json:"entity_id"`
}

type Direction int

const (
	Before Direction = -1
	After  Direction = 1

	Beginning snowflake.ID = 0
	End       snowflake.ID = math.MaxInt64
)

type StreamItemStore interface {
	// GetItems returns count items from the specified stream, starting
	// after/before the specified item, in the specified direction. Pass "Beginning"/"End" to
	// start at the beginning or end of the stream
	GetItems(streamID snowflake.ID, startID snowflake.ID, direction Direction, count uint) ([]StreamItem, error)

	// GetItem returns a specific stream item
	GetItem(itemID snowflake.ID) (*StreamItem, error)

	// GetItemByEntityId returns a specific stream item by stream and entry ID
	GetItemByEntityID(streamID snowflake.ID, entityID snowflake.ID) (*StreamItem, error)

	// TryInsertItem attempts to insert the specified entity ID into the stream.
	// Returns:
	//   * If the entry was inserted
	//   * The stream item (regardless)
	TryInsertItem(stream snowflake.ID, entityId snowflake.ID) (bool, *StreamItem, error)
}

type streamItemStore struct {
	DB *gorm.DB
}

func NewStreamItemStore(db *gorm.DB) StreamItemStore {
	return &streamItemStore{db}
}

func (s *streamItemStore) GetItems(streamID snowflake.ID, startID snowflake.ID, direction Direction, count uint) ([]StreamItem, error) {
	var items []StreamItem

	q := s.DB.Where("stream_id = ?", streamID)

	if direction == Before {
		q = q.Where("item_id < ?", startID).Order("item_id DESC")
	} else if direction == After {
		q = q.Where("item_id > ?", startID).Order("item_id ASC")
	}
	q = q.Limit(count)
	err := q.Find(&items).Error
	if err != nil {
		return nil, err
	} else {
		return items, nil
	}
}

func (s *streamItemStore) GetItem(itemID snowflake.ID) (*StreamItem, error) {
	var item StreamItem

	return &item, s.DB.First(&item, StreamItem{ItemID: itemID}).Error
}

func (s *streamItemStore) GetItemByEntityID(streamID snowflake.ID, entityID snowflake.ID) (*StreamItem, error) {
	var item StreamItem

	return &item, s.DB.First(&item, StreamItem{StreamID: streamID, EntityID: entityID}).Error
}

func (s *streamItemStore) TryInsertItem(streamID snowflake.ID, entityID snowflake.ID) (bool, *StreamItem, error) {
	item, err := s.GetItemByEntityID(streamID, entityID)
	if err == nil {
		return false, item, nil
	} else if gorm.IsRecordNotFoundError(err) {
		flake, err := lib.GenSnowflake(config.NodeID())
		if err != nil {
			return false, nil, err
		}

		item := &StreamItem{
			StreamID: streamID,
			EntityID: entityID,
			ItemID:   flake,
		}
		err = s.DB.Create(item).Error

		return true, item, err
	} else {
		return false, nil, err
	}
}