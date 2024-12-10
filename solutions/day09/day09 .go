package day09

import (
	"fmt"

	"github.com/dvan-sqsp/advent-of-code-2024/internal/solver"
	"github.com/dvan-sqsp/advent-of-code-2024/util"
)

type Day09 struct{}

var _ solver.Solver = (*Day09)(nil)

func New() *Day09 {
	return &Day09{}
}

func (d *Day09) Part1(lines []string) string {
	disk, _ := d.parseDisk(lines)

	disk.Compact()

	checkSum := d.getCheckSum(disk)

	return fmt.Sprintf("%d", checkSum)
}

func (d *Day09) Part2(lines []string) string {
	disk, files := d.parseDisk(lines)

	disk.CompactWholeFile(files)

	checkSum := d.getCheckSum(disk)

	return fmt.Sprintf("%d", checkSum)
}

func (d *Day09) parseDisk(lines []string) (Disk, []File) {
	diskMap, _ := util.ReadInts(lines[0], "")

	id := 0
	disk := NewDisk()
	files := make([]File, 0) // fileID is the idx and the value is the size

	startIDIdx := 0
	for idx, fileLength := range diskMap {
		if idx%2 == 0 {
			var freeSpace int
			if idx != len(diskMap)-1 {
				freeSpace = diskMap[idx+1]
			}
			for i := 0; i < fileLength; i++ {
				disk = append(disk, id)
			}
			for i := 0; i < freeSpace; i++ {
				disk = append(disk, -1) // -1 will represent free memory
			}
			// files represents a file
			// startIDIDx is the start of the ID such that we can use it to do the swaps easier
			files = append(files, NewFile(id, startIDIdx, startIDIdx+fileLength, fileLength))
			id++

			startIDIdx += fileLength + freeSpace
		}
	}
	return disk, files
}

func (d *Day09) getCheckSum(disk []int) int {
	checkSum := 0
	for blockPos, fileID := range disk {
		if fileID != -1 {
			checkSum += blockPos * fileID
		}
	}
	return checkSum
}
