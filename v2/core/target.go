package core

import (
	"errors"
	"fmt"
	"os"
	"sort"
	"text/tabwriter"

	log "github.com/sirupsen/logrus"

	"github.com/robgonnella/ardi/v2/rpc"
)

// Target represents a targeted arduino board
type Target struct {
	Board *rpc.Board
}

// NewTarget returns new target
func NewTarget(connectedBoards, allBoards []*rpc.Board, logger *log.Logger, fqbn string, onlyConnected bool) (*Target, error) {
	board, err := getTargetBoard(connectedBoards, allBoards, logger, fqbn, onlyConnected)
	if err != nil {
		return nil, err
	}
	return &Target{board}, nil
}

func getTargetBoard(connectedBoards, allBoards []*rpc.Board, logger *log.Logger, fqbn string, onlyConnected bool) (*rpc.Board, error) {
	if fqbn != "" {
		return &rpc.Board{FQBN: fqbn}, nil
	}

	if len(connectedBoards) == 0 {
		if onlyConnected {
			err := errors.New("No connected boards detected")
			logger.WithError(err).Error()
			return nil, err
		}
		printFQBNs(allBoards)
		logger.Errorf("You must supply fqbn to compile. You can find a list of board fqbns for installed platforms above.")
		return nil, errors.New("Must provide board fqbn")
	}

	if len(connectedBoards) == 1 {
		return connectedBoards[0], nil
	}

	if len(connectedBoards) > 1 {
		printFQBNs(connectedBoards)
		logger.Errorf("You must supply fqbn to compile. You can find a list of board fqbns for connected boards above.")
		return nil, errors.New("Must provide board fqbn")
	}

	return nil, errors.New("Error parsing target")
}

// private helpers
func printFQBNs(boardList []*rpc.Board) {
	sort.Slice(boardList, func(i, j int) bool {
		return boardList[i].Name < boardList[j].Name
	})

	printBoardsWithIndices(boardList)
}

func printBoardsWithIndices(boards []*rpc.Board) {
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 8, ' ', 0)
	defer w.Flush()
	fmt.Fprintln(w, "No.\tName\tFQBN")
	for i, b := range boards {
		fmt.Fprintf(w, "%d\t%s\t%s\n", i, b.Name, b.FQBN)
	}
}
