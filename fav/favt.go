package fav

type FAVT int8

const (
	FAVT_BOARD  FAVT = 1
	FAVT_FOLDER FAVT = 2
	FAVT_LINE   FAVT = 3
)

func (t FAVT) String() string {
	switch t {
	case FAVT_BOARD:
		return "Board"
	case FAVT_FOLDER:
		return "Folder"
	case FAVT_LINE:
		return "Line"
	default:
		return "unknown"
	}
}

func (theType FAVT) IsValidFavType() bool {
	switch theType {
	case FAVT_BOARD:
		return true
	case FAVT_FOLDER:
		return true
	case FAVT_LINE:
		return true
	}
	return false
}

func (theType FAVT) GetTypeSize() uintptr {
	switch theType {
	case FAVT_BOARD:
		return SIZE_OF_FAV_BOARD
	case FAVT_FOLDER:
		return SIZE_OF_FAV_FOLDER
	case FAVT_LINE:
		return SIZE_OF_FAV_LINE
	}
	return 0
}

func (theType FAVT) GetFav4TypeSize() uintptr {
	switch theType {
	case FAVT_BOARD:
		return SIZE_OF_FAV4_BOARD
	case FAVT_FOLDER:
		return SIZE_OF_FAV_FOLDER
	case FAVT_LINE:
		return SIZE_OF_FAV_LINE
	}
	return 0
}
