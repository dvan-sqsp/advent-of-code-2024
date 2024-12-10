package day09

type Disk []int

func NewDisk() Disk {
	return make([]int, 0)
}

func (d *Disk) Compact() {
	// start from end of disk
	for i := len(*d) - 1; i >= 0; i-- {
		// free memory, skip it
		if (*d)[i] == -1 {
			continue
		}

		// find first free space from left
		for j := 0; j < i; j++ {
			if (*d)[j] == -1 {
				//swap
				(*d)[j] = (*d)[i]
				(*d)[i] = -1
				break
			}
		}
	}
}

func (d *Disk) CompactWholeFile(files []File) {
	// start from last file ID
	for fileID := len(files) - 1; fileID >= 0; fileID-- {
		if fileID == 0 {
			break
		}
		// the file size we're trying to find space for
		file := files[fileID]

		// now we need to find first chunk of free memory
		freeMemoryStart := -1
		freeMemoryLength := 0

		// start from beginning of disk
		for i := 0; i < len(*d); i++ {
			// we found free space!
			if (*d)[i] == -1 {
				if freeMemoryStart == -1 {
					freeMemoryStart = i
				}

				// stop computing if we're overlapping
				if freeMemoryStart > file.StartIdx {
					break
				}

				freeMemoryLength++
				if freeMemoryLength == file.Length { // file size fits within memory chunk and we haven't start to overlap
					fileIDsToSwap := (*d)[file.StartIdx:file.EndIdx]
					freeMemory := (*d)[freeMemoryStart : freeMemoryStart+freeMemoryLength]

					// copy a new disk
					// [0, 0, -1, -1, 9, 9 -1]
					newDisk := make([]int, 0)
					newDisk = append(newDisk, *d...)

					// copy over from the start of the disk to where the free memory is
					// newDisk = [0, 0, -1, -1, 9, 9, -1]
					// [0, 0] - newDisk[:freeMemoryStart]- copying all elements until the free memory
					// [0, 0, 9, 9] - append(fileIDsToSwap, ...) - append the fileIDs we're swapping as well as all other elements after it
					// [0, 0, 9, 9, 9, 9, -1] - newDisk[freeMemoryStart+freeMemoryLength:]... - everything else afterwards
					// then we'll append the fileIDs that we're swapping then append the rest of the disk
					newDisk = append(newDisk[:freeMemoryStart], append(fileIDsToSwap, newDisk[freeMemoryStart+freeMemoryLength:]...)...)

					// same thing but now we're swapping the free memory
					// newDisk = [0, 0, 9, 9, 9, 9, -1]
					// [0, 0, 9, 9] - newDisk[:file.StartIdx] - copying all elements until the the start ID IDX
					// [0, 0, 9, 9, -1, -1] - append(freeMemory, ...) - append the free memory we're swapping
					// [0, 0, 9, 9, 9, 9, -1] - newDisk[freeMemoryStart+freeMemoryLength:]... - everything else afterwards
					newDisk = append(newDisk[:file.StartIdx], append(freeMemory, newDisk[file.EndIdx:]...)...)

					// overwrite disk
					*d = newDisk
					break
				}
			} else {
				// reset since we did not find enough free space
				freeMemoryStart = -1
				freeMemoryLength = 0
			}
		}
	}
}
